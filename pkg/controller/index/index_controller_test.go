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
	"testing"
	"time"

	elasticsearchdbv1beta1 "github.com/kaitoy/elasticsearch-db-operator/pkg/apis/elasticsearchdb/v1beta1"
	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var c client.Client

var expectedRequest = reconcile.Request{NamespacedName: types.NamespacedName{Name: "foo", Namespace: "default"}}
var depKey = types.NamespacedName{Name: "foo-deployment", Namespace: "default"}

const timeout = time.Second * 5

func TestReconcile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	instance := &elasticsearchdbv1beta1.Index{
		ObjectMeta: metav1.ObjectMeta{Name: "foo", Namespace: "default"},
		URL: elasticsearchdbv1beta1.IndexURL{
			ElasticsearchEndpoint: "http://elasticsearch:9200",
			Index:                 "foo",
		},
		Spec: elasticsearchdbv1beta1.IndexSpec{
			Settings: &elasticsearchdbv1beta1.Settings{
				Index: elasticsearchdbv1beta1.IndexSettings{
					NumberOfShards:   &intstr.IntOrString{Type: 0, IntVal: 10, StrVal: ""},
					NumberOfReplicas: &intstr.IntOrString{Type: 0, IntVal: 2, StrVal: ""},
				},
			},
			Mappings: map[string]elasticsearchdbv1beta1.Mapping{
				"_doc": elasticsearchdbv1beta1.Mapping{
					Source: &elasticsearchdbv1beta1.Source{
						Enabled: true,
					},
					Properties: map[string]elasticsearchdbv1beta1.Property{
						"age": elasticsearchdbv1beta1.Property{
							Type: "integer",
						},
						"name": elasticsearchdbv1beta1.Property{
							Properties: map[string]elasticsearchdbv1beta1.Property{
								"first": elasticsearchdbv1beta1.Property{
									Type:  "keyword",
									Boost: 2.5,
								},
								"last": elasticsearchdbv1beta1.Property{
									Type: "keyword",
								},
							},
						},
					},
				},
			},
		},
	}

	// Setup the Manager and Controller.  Wrap the Controller Reconcile function so it writes each request to a
	// channel when it is finished.
	mgr, err := manager.New(cfg, manager.Options{})
	g.Expect(err).NotTo(gomega.HaveOccurred())
	c = mgr.GetClient()

	recFn, requests := SetupTestReconcile(newReconciler(mgr))
	g.Expect(add(mgr, recFn)).NotTo(gomega.HaveOccurred())

	stopMgr, mgrStopped := StartTestManager(mgr, g)

	defer func() {
		close(stopMgr)
		mgrStopped.Wait()
	}()

	// Create the Index object and expect the Reconcile and Deployment to be created
	err = c.Create(context.TODO(), instance)
	// The instance object may not be a valid object because it might be missing some required fields.
	// Please modify the instance object by adding required fields and then remove the following if statement.
	if apierrors.IsInvalid(err) {
		t.Logf("failed to create object, got an invalid object error: %v", err)
		return
	}
	g.Expect(err).NotTo(gomega.HaveOccurred())
	defer c.Delete(context.TODO(), instance)
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))

	deploy := &appsv1.Deployment{}
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, deploy) }, timeout).
		Should(gomega.Succeed())

	// Delete the Deployment and expect Reconcile to be called for Deployment deletion
	g.Expect(c.Delete(context.TODO(), deploy)).NotTo(gomega.HaveOccurred())
	g.Eventually(requests, timeout).Should(gomega.Receive(gomega.Equal(expectedRequest)))
	g.Eventually(func() error { return c.Get(context.TODO(), depKey, deploy) }, timeout).
		Should(gomega.Succeed())

	// Manually delete Deployment since GC isn't enabled in the test control plane
	g.Expect(c.Delete(context.TODO(), deploy)).To(gomega.Succeed())

}
