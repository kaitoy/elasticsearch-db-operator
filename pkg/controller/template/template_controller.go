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

package template

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/url"
	"time"

	"github.com/google/go-cmp/cmp"
	elasticsearchdbv1beta1 "github.com/kaitoy/elasticsearch-db-operator/pkg/apis/elasticsearchdb/v1beta1"
	utilsstrings "github.com/kaitoy/elasticsearch-db-operator/pkg/utils"
	resty "gopkg.in/resty.v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const controllerName = "elasticsearchdb-template-operator"

var log = logf.Log.WithName(controllerName)

// Add creates a new Template Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
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
	return &ReconcileTemplate{
		Client:                    mgr.GetClient(),
		scheme:                    mgr.GetScheme(),
		indicesWatchingCancellers: make(map[string]context.CancelFunc),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New(controllerName, mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to Template
	err = c.Watch(&source.Kind{Type: &elasticsearchdbv1beta1.Template{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTemplate{}

// ReconcileTemplate reconciles a Template object
type ReconcileTemplate struct {
	client.Client
	scheme                    *runtime.Scheme
	indicesWatchingCancellers map[string]context.CancelFunc
}

const finalizerName = "template.finalizers.elasticsearchdb.kaitoy.github.com"

// Reconcile reads that state of the cluster for a Template object and makes changes based on the state read
// and what is in the Template.Spec
// +kubebuilder:rbac:groups=elasticsearchdb.kaitoy.github.com,resources=indices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=elasticsearchdb.kaitoy.github.com,resources=indices/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=elasticsearchdb.kaitoy.github.com,resources=templates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=elasticsearchdb.kaitoy.github.com,resources=templates/status,verbs=get;update;patch
func (r *ReconcileTemplate) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	instanceName := request.NamespacedName.String()
	log.Info("Reconciling a template: " + instanceName)

	instance := &elasticsearchdbv1beta1.Template{}
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

func reconcileCreationOrUpdate(template *elasticsearchdbv1beta1.Template, r *ReconcileTemplate) error {
	instanceName := types.NamespacedName{Name: template.Name, Namespace: template.Namespace}.String()
	templateURL, err := template.GetURL()
	if err != nil {
		log.Error(err, "Failed to get URL of "+instanceName)
		return err
	}

	if !utilsstrings.ContainsString(template.ObjectMeta.Finalizers, finalizerName) {
		log.Info(fmt.Sprintf("Adding a finalizer to %s.", instanceName))
		template.ObjectMeta.Finalizers = append(template.ObjectMeta.Finalizers, finalizerName)
	}
	if _, ok := r.indicesWatchingCancellers[instanceName]; !ok {
		log.Info(fmt.Sprintf("Starting to watch indices for a template %s.", instanceName))
		ctx, cancel := context.WithCancel(context.Background())
		r.indicesWatchingCancellers[instanceName] = cancel
		go watchIndices(ctx, template, r)
	}

	response, err := resty.R().
		Get(templateURL.String())
	if err != nil {
		log.Error(err, "Failed to GET "+templateURL.String())
		return err
	}

	if len(template.Status.Conditions) > 0 {
		lastCond := template.Status.Conditions[len(template.Status.Conditions)-1]
		if lastCond.StatusCode != response.StatusCode() || lastCond.Status != response.Status() {
			appendCondition(template, response)
		}
	} else {
		appendCondition(template, response)
	}

	if response.IsSuccess() {
		log.Info(fmt.Sprintf("Template %s exists.", templateURL.String()))
		if err := r.Update(context.Background(), template); err != nil {
			log.Error(err, "Failed to update "+instanceName)
			return err
		}
		return nil
	} else if response.StatusCode() == 404 {
		response, err = resty.R().
			SetBody(template.Spec).
			Put(templateURL.String())
		if err != nil {
			log.Error(err, "Failed to PUT "+templateURL.String())
			if err := r.Update(context.Background(), template); err != nil {
				log.Error(err, "Failed to update "+instanceName)
				return err
			}
			return err
		}
		appendCondition(template, response)

		if response.IsError() {
			err = fmt.Errorf(
				"Got an error response '%s' for %s %s",
				response.Status(),
				response.Request.Method,
				response.Request.URL,
			)
			log.Error(err, "An error occurred during creating a template "+templateURL.String())
			if err := r.Update(context.Background(), template); err != nil {
				log.Error(err, "Failed to update "+instanceName)
				return err
			}
			return err
		}

		log.Info(fmt.Sprintf("Succeeded to create a template %s.", templateURL.String()))
	} else {
		err = fmt.Errorf(
			"Got an error response '%s' for %s %s",
			response.Status(),
			response.Request.Method,
			response.Request.URL,
		)
		log.Error(err, "An error occurred during getting a template "+templateURL.String())
		if err := r.Update(context.Background(), template); err != nil {
			log.Error(err, "Failed to update "+instanceName)
			return err
		}
		return err
	}

	if err := r.Update(context.Background(), template); err != nil {
		log.Error(err, "Failed to update "+instanceName)
		return err
	}

	return nil
}

func watchIndices(ctx context.Context, template *elasticsearchdbv1beta1.Template, r *ReconcileTemplate) {
	instanceName := types.NamespacedName{Name: template.Name, Namespace: template.Namespace}.String()
	indicesURL, err := template.GetIndicesURL()
	if err != nil {
		log.Error(err, fmt.Sprintf("Failed to get indices URL of %s. Can't start to watch the indices.", instanceName))
		return
	}

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			pollIndices(indicesURL, template, r)
		}
	}
}

func pollIndices(indicesURL *url.URL, template *elasticsearchdbv1beta1.Template, r *ReconcileTemplate) error {
	response, err := resty.R().
		SetResult(&map[string]elasticsearchdbv1beta1.IndexSpec{}).
		Get(indicesURL.String())
	if err != nil {
		log.Error(err, "Failed to GET "+indicesURL.String())
		return err
	}

	if response.IsError() {
		err = fmt.Errorf(
			"Got an error response '%s' for %s %s",
			response.Status(),
			response.Request.Method,
			response.Request.URL,
		)
		log.Error(err, "An error occurred during getting indices "+indicesURL.String())
		return err
	}

	if len(*response.Result().(*map[string]elasticsearchdbv1beta1.IndexSpec)) == 0 {
		return nil
	}

	indexSpecs, ok := response.Result().(*map[string]elasticsearchdbv1beta1.IndexSpec)
	if !ok {
		err = fmt.Errorf(
			"Failed to cast %v to IndexSpec",
			response.Result(),
		)
		log.Error(err, "An error occurred during casting a json to IndexSpec ")
		return err
	}

	for indexName, spec := range *indexSpecs {
		createOrUpdateIndex(template, r, indexName, &spec)
	}

	return nil
}

func createOrUpdateIndex(
	template *elasticsearchdbv1beta1.Template,
	r *ReconcileTemplate, indexName string,
	indexSpec *elasticsearchdbv1beta1.IndexSpec,
) error {
	index := &elasticsearchdbv1beta1.Index{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-%x", template.Name, sha1.Sum([]byte(indexName))),
			Namespace: template.Namespace,
		},
		URL: elasticsearchdbv1beta1.IndexURL{
			ElasticsearchEndpoint: template.URL.ElasticsearchEndpoint,
			Index:                 indexName,
		},
		Spec: *indexSpec,
		Status: elasticsearchdbv1beta1.IndexStatus{
			Closed:     false,
			Conditions: []elasticsearchdbv1beta1.IndexCondition{},
		},
	}
	if err := controllerutil.SetControllerReference(template, index, r.scheme); err != nil {
		log.Error(err, fmt.Sprintf("An error occurred during setting %s/%s as the controller of %s.", template.Namespace, template.Name, index.Name))
		return err
	}

	found := &elasticsearchdbv1beta1.Index{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: index.Name, Namespace: index.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		log.Info(fmt.Sprintf("Creating an Index: %s/%s", index.Namespace, index.Name))
		if err := r.Create(context.TODO(), index); err != nil {
			log.Error(err, fmt.Sprintf("Failed to create an Index: %s/%s", index.Namespace, index.Name))
			return err
		}

		return nil
	} else if err != nil {
		log.Error(err, fmt.Sprintf("Failed to get an Index: %s/%s", index.Namespace, index.Name))
		return err
	}

	// Update the found object and write the result back if there are any changes
	if !cmp.Equal(
		index.Spec, found.Spec,
		cmp.Comparer(intOrStringComparer),
	) {
		log.Info(fmt.Sprintf(
			"Updating an Index %s/%s to eliminate a difference from actual one: %s",
			index.Namespace,
			index.Name,
			cmp.Diff(
				index.Spec, found.Spec,
				cmp.Comparer(intOrStringComparer),
			),
		))
		found.Spec = index.Spec
		if err := r.Update(context.TODO(), found); err != nil {
			log.Error(err, fmt.Sprintf("Failed to update an Index: %s/%s", index.Namespace, index.Name))
			return err
		}
	}

	return nil
}

func intOrStringComparer(a *intstr.IntOrString, b *intstr.IntOrString) bool {
	if a == nil {
		return b == nil
	}
	if b == nil {
		return false
	}
	return a.IntValue() == b.IntValue()
}

func reconcileDeletion(template *elasticsearchdbv1beta1.Template, r *ReconcileTemplate) error {
	instanceName := types.NamespacedName{Name: template.Name, Namespace: template.Namespace}.String()
	templateURL, err := template.GetURL()
	if err != nil {
		log.Error(err, "Failed to get URL of "+instanceName)
		return err
	}

	cancel, ok := r.indicesWatchingCancellers[instanceName]
	if ok {
		log.Info(fmt.Sprintf("Stopping watching indices for a template %s.", instanceName))
		cancel()
		delete(r.indicesWatchingCancellers, instanceName)
	}

	if utilsstrings.ContainsString(template.ObjectMeta.Finalizers, finalizerName) {
		log.Info(fmt.Sprintf("Deleting a template %s.", templateURL.String()))
		response, err := resty.R().
			Delete(templateURL.String())
		if err != nil {
			log.Error(err, "Failed to DELETE "+templateURL.String())
			return err
		}
		if response.IsError() {
			err = fmt.Errorf(
				"Got an error response '%s' for %s %s",
				response.Status(),
				response.Request.Method,
				response.Request.URL,
			)
			log.Error(err, "An error occurred during deleting a template "+templateURL.String())
			return err
		}

		log.Info(fmt.Sprintf("Succeeded to delete a template %s.", templateURL.String()))

		template.ObjectMeta.Finalizers = utilsstrings.RemoveString(template.ObjectMeta.Finalizers, finalizerName)
		if err := r.Update(context.Background(), template); err != nil {
			return err
		}
	}

	return nil
}

func appendCondition(templ *elasticsearchdbv1beta1.Template, response *resty.Response) {
	templ.Status.Conditions = append(
		templ.Status.Conditions,
		elasticsearchdbv1beta1.TemplateCondition{
			StatusCode:         response.StatusCode(),
			Status:             response.Status(),
			LastProbeTime:      metav1.NewTime(response.Request.Time),
			LastTransitionTime: metav1.NewTime(response.ReceivedAt()),
		},
	)
}
