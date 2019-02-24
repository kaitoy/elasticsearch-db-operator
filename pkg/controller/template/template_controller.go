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
	"fmt"
	"net/url"
	"time"

	elasticsearchdbv1beta1 "github.com/kaitoy/elasticsearch-db-operator/pkg/apis/elasticsearchdb/v1beta1"
	strings "github.com/kaitoy/elasticsearch-db-operator/pkg/controller/utils"
	resty "gopkg.in/resty.v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
	return &ReconcileTemplate{Client: mgr.GetClient(), scheme: mgr.GetScheme()}
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

	// TODO(user): Modify this to be the types you create
	// Uncomment watch a Deployment created by Template - change this for objects you create
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &elasticsearchdbv1beta1.Template{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileTemplate{}

// ReconcileTemplate reconciles a Template object
type ReconcileTemplate struct {
	client.Client
	scheme *runtime.Scheme
}

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

	endpoint, err := url.Parse(instance.URL.ElasticsearchEndpoint)
	if err != nil {
		log.Error(err, "Invalid url: "+instance.URL.ElasticsearchEndpoint)
		return reconcile.Result{}, err
	}
	endpoint.Path = "/_template/" + instance.URL.Template
	log.Info("Operating a template: " + endpoint.String())

	const finalizerName = "template.finalizers.elasticsearchdb.kaitoy.github.com"
	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		if !strings.ContainsString(instance.ObjectMeta.Finalizers, finalizerName) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, finalizerName)
		}

		response, err := resty.R().
			Get(endpoint.String())
		if err != nil {
			log.Error(err, "Failed to GET "+endpoint.String())
			return reconcile.Result{}, err
		}

		if len(instance.Status.Conditions) > 0 {
			lastCond := instance.Status.Conditions[len(instance.Status.Conditions)-1]
			if lastCond.StatusCode != response.StatusCode() || lastCond.Status != response.Status() {
				instance.Status.Conditions = append(
					instance.Status.Conditions,
					elasticsearchdbv1beta1.TemplateCondition{
						StatusCode:         response.StatusCode(),
						Status:             response.Status(),
						LastProbeTime:      metav1.NewTime(response.Request.Time),
						LastTransitionTime: metav1.NewTime(response.ReceivedAt()),
					},
				)
			} else {
				log.Info(fmt.Sprintf("Nothing to do for %s.", instanceName))
				return reconcile.Result{}, nil
			}
		}

		if response.IsSuccess() {
			log.Info(fmt.Sprintf("Template %s exists.", endpoint.String()))
			if err := r.Update(context.Background(), instance); err != nil {
				log.Error(err, "Failed to update "+instanceName)
				return reconcile.Result{}, err
			}
			return reconcile.Result{}, nil
		} else if response.StatusCode() == 404 {
			response, err = resty.R().
				SetBody(instance.Spec).
				Put(endpoint.String())
			if err != nil {
				log.Error(err, "Failed to PUT"+endpoint.String())
				if err := r.Update(context.Background(), instance); err != nil {
					log.Error(err, "Failed to update "+instanceName)
					return reconcile.Result{}, err
				}
				return reconcile.Result{}, err
			}
			instance.Status.Conditions = append(
				instance.Status.Conditions,
				elasticsearchdbv1beta1.TemplateCondition{
					StatusCode:         response.StatusCode(),
					Status:             response.Status(),
					LastProbeTime:      metav1.NewTime(response.Request.Time),
					LastTransitionTime: metav1.NewTime(response.ReceivedAt()),
				},
			)

			if response.IsError() {
				err = fmt.Errorf(
					"Got an error response '%s' for %s %s",
					response.Status(),
					response.Request.Method,
					response.Request.URL,
				)
				log.Error(err, "An error occurred during creating a template "+endpoint.String())
				if err := r.Update(context.Background(), instance); err != nil {
					log.Error(err, "Failed to update "+instanceName)
					return reconcile.Result{}, err
				}
				return reconcile.Result{}, err
			}

			log.Info(fmt.Sprintf("Scceeded to create a template %s.", endpoint.String()))
		} else {
			err = fmt.Errorf(
				"Got an error response '%s' for %s %s",
				response.Status(),
				response.Request.Method,
				response.Request.URL,
			)
			log.Error(err, "An error occurred during getting a template "+endpoint.String())
			if err := r.Update(context.Background(), instance); err != nil {
				log.Error(err, "Failed to update "+instanceName)
				return reconcile.Result{}, err
			}
			return reconcile.Result{}, err
		}

		if err := r.Update(context.Background(), instance); err != nil {
			log.Error(err, "Failed to update "+instanceName)
			return reconcile.Result{}, err
		}
	} else {
		// This instance is being deleted.
		if strings.ContainsString(instance.ObjectMeta.Finalizers, finalizerName) {
			response, err := resty.R().
				Delete(endpoint.String())
			if err != nil {
				log.Error(err, "Failed to Delete"+endpoint.String())
				return reconcile.Result{}, err
			}
			if response.IsError() {
				err = fmt.Errorf(
					"Got an error response '%s' for %s %s",
					response.Status(),
					response.Request.Method,
					response.Request.URL,
				)
				log.Error(err, "An error occurred during deleting a template "+endpoint.String())
				return reconcile.Result{}, err
			}

			log.Info(fmt.Sprintf("Scceeded to delete a template %s.", endpoint.String()))

			instance.ObjectMeta.Finalizers = strings.RemoveString(instance.ObjectMeta.Finalizers, finalizerName)
			if err := r.Update(context.Background(), instance); err != nil {
				return reconcile.Result{}, err
			}
		}
	}

	// TODO(user): Change this to be the object type created by your controller
	// Define the desired Deployment object
	// deploy := &appsv1.Deployment{
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      instance.Name + "-deployment",
	// 		Namespace: instance.Namespace,
	// 	},
	// 	Spec: appsv1.DeploymentSpec{
	// 		Selector: &metav1.LabelSelector{
	// 			MatchLabels: map[string]string{"deployment": instance.Name + "-deployment"},
	// 		},
	// 		Template: corev1.PodTemplateSpec{
	// 			ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"deployment": instance.Name + "-deployment"}},
	// 			Spec: corev1.PodSpec{
	// 				Containers: []corev1.Container{
	// 					{
	// 						Name:  "nginx",
	// 						Image: "nginx",
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// if err := controllerutil.SetControllerReference(instance, deploy, r.scheme); err != nil {
	// 	return reconcile.Result{}, err
	// }

	// TODO(user): Change this for the object type created by your controller
	// Check if the Deployment already exists
	// found := &appsv1.Deployment{}
	// err = r.Get(context.TODO(), types.NamespacedName{Name: deploy.Name, Namespace: deploy.Namespace}, found)
	// if err != nil && errors.IsNotFound(err) {
	// 	log.Info("Creating Deployment", "namespace", deploy.Namespace, "name", deploy.Name)
	// 	err = r.Create(context.TODO(), deploy)
	// 	if err != nil {
	// 		return reconcile.Result{}, err
	// 	}
	// } else if err != nil {
	// 	return reconcile.Result{}, err
	// }

	// TODO(user): Change this for the object type created by your controller
	// Update the found object and write the result back if there are any changes
	// if !reflect.DeepEqual(deploy.Spec, found.Spec) {
	// 	found.Spec = deploy.Spec
	// 	log.Info("Updating Deployment", "namespace", deploy.Namespace, "name", deploy.Name)
	// 	err = r.Update(context.TODO(), found)
	// 	if err != nil {
	// 		return reconcile.Result{}, err
	// 	}
	// }
	return reconcile.Result{}, nil
}