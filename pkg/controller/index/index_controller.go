/*
Copyright 2019 kaitoy.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package index

import (
	"context"
	"fmt"
	"time"

	elasticsearchdbv1beta1 "github.com/kaitoy/elasticsearch-db-operator/pkg/apis/elasticsearchdb/v1beta1"
	utilsstrings "github.com/kaitoy/elasticsearch-db-operator/pkg/utils"
	resty "gopkg.in/resty.v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const controllerName = "elasticsearchdb-index-operator"

var log = logf.Log.WithName(controllerName)

// Add creates a new Index Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	resty.SetHeaders(map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	})
	// resty.SetLogger(log)
	resty.SetTimeout(20 * time.Second)
	resty.SetRetryCount(3)
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileIndex{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Index
	err = c.Watch(&source.Kind{Type: &elasticsearchdbv1beta1.Index{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileIndex{}

// ReconcileIndex reconciles a Index object
type ReconcileIndex struct {
	client.Client
	scheme *runtime.Scheme
}

const finalizerName = "index.finalizers.elasticsearchdb.kaitoy.github.com"

// Reconcile reads that state of the cluster for a Index object and makes changes based on the state read
// and what is in the Index.Spec
// +kubebuilder:rbac:groups=elasticsearchdb.kaitoy.github.com,resources=indices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=elasticsearchdb.kaitoy.github.com,resources=indices/status,verbs=get;update;patch
func (r *ReconcileIndex) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instanceName := request.NamespacedName.String()
	log.Info("Reconciling an index: " + instanceName)

	instance := &elasticsearchdbv1beta1.Index{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info(instanceName + " was not found.")
			return reconcile.Result{}, nil
		}

		log.Error(err, "Failed to get "+instanceName)
		return reconcile.Result{}, err
	}

	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		err := reconcileCreationOrUpdate(instance, r)
		return reconcile.Result{}, err
	}
	err = reconcileDeletion(instance, r)
	return reconcile.Result{}, err
}

func reconcileCreationOrUpdate(index *elasticsearchdbv1beta1.Index, r *ReconcileIndex) error {
	instanceName := types.NamespacedName{Name: index.Name, Namespace: index.Namespace}.String()
	indexURL, err := index.GetURL()
	if err != nil {
		log.Error(err, "Failed to get URL of "+instanceName)
		return err
	}

	if !utilsstrings.ContainsString(index.ObjectMeta.Finalizers, finalizerName) {
		log.Info(fmt.Sprintf("Adding a finalizer to %s.", instanceName))
		index.ObjectMeta.Finalizers = append(index.ObjectMeta.Finalizers, finalizerName)
	}

	response, err := resty.R().
		Get(indexURL.String())
	if err != nil {
		log.Error(err, "Failed to GET "+indexURL.String())
		return err
	}

	if len(index.Status.Conditions) > 0 {
		lastCond := index.Status.Conditions[len(index.Status.Conditions)-1]
		if lastCond.StatusCode != response.StatusCode() || lastCond.Status != response.Status() {
			appendCondition(index, response)
		}
	} else {
		appendCondition(index, response)
	}

	if response.IsSuccess() {
		log.Info(fmt.Sprintf("Index %s exists.", indexURL.String()))
		if err := r.Update(context.Background(), index); err != nil {
			log.Error(err, "Failed to update "+instanceName)
			return err
		}
		return nil
	} else if response.StatusCode() == 404 {
		response, err = resty.R().
			SetBody(index.Spec).
			Put(indexURL.String())
		if err != nil {
			log.Error(err, "Failed to PUT "+indexURL.String())
			if err := r.Update(context.Background(), index); err != nil {
				log.Error(err, "Failed to update "+instanceName)
				return err
			}
			return err
		}
		appendCondition(index, response)

		if response.IsError() {
			err = fmt.Errorf(
				"Got an error response '%s' for %s %s",
				response.Status(),
				response.Request.Method,
				response.Request.URL,
			)
			log.Error(err, "An error occurred during creating an index "+indexURL.String())
			if err := r.Update(context.Background(), index); err != nil {
				log.Error(err, "Failed to update "+instanceName)
				return err
			}
			return err
		}

		log.Info(fmt.Sprintf("Succeeded to create an index %s.", indexURL.String()))
	} else {
		err = fmt.Errorf(
			"Got an error response '%s' for %s %s",
			response.Status(),
			response.Request.Method,
			response.Request.URL,
		)
		log.Error(err, "An error occurred during getting an index "+indexURL.String())
		if err := r.Update(context.Background(), index); err != nil {
			log.Error(err, "Failed to update "+instanceName)
			return err
		}
		return err
	}

	if err := r.Update(context.Background(), index); err != nil {
		log.Error(err, "Failed to update "+instanceName)
		return err
	}

	return nil
}

func reconcileDeletion(index *elasticsearchdbv1beta1.Index, r *ReconcileIndex) error {
	instanceName := types.NamespacedName{Name: index.Name, Namespace: index.Namespace}.String()
	indexURL, err := index.GetURL()
	if err != nil {
		log.Error(err, "Failed to get URL of "+instanceName)
		return err
	}

	if utilsstrings.ContainsString(index.ObjectMeta.Finalizers, finalizerName) {
		log.Info(fmt.Sprintf("Deleting an index %s.", indexURL.String()))
		response, err := resty.R().
			Delete(indexURL.String())
		if err != nil {
			log.Error(err, "Failed to DELETE "+indexURL.String())
			return err
		}
		if response.IsError() {
			err = fmt.Errorf(
				"Got an error response '%s' for %s %s",
				response.Status(),
				response.Request.Method,
				response.Request.URL,
			)
			log.Error(err, "An error occurred during deleting an index "+indexURL.String())
			return err
		}

		log.Info(fmt.Sprintf("Succeeded to delete an index %s.", indexURL.String()))

		index.ObjectMeta.Finalizers = utilsstrings.RemoveString(index.ObjectMeta.Finalizers, finalizerName)
		if err := r.Update(context.Background(), index); err != nil {
			return err
		}
	}

	return nil
}

func appendCondition(index *elasticsearchdbv1beta1.Index, response *resty.Response) {
	index.Status.Conditions = append(
		index.Status.Conditions,
		elasticsearchdbv1beta1.IndexCondition{
			StatusCode:         response.StatusCode(),
			Status:             response.Status(),
			LastProbeTime:      metav1.NewTime(response.Request.Time),
			LastTransitionTime: metav1.NewTime(response.ReceivedAt()),
		},
	)
}
