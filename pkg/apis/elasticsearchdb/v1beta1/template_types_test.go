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

package v1beta1

import (
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestStorageTemplate(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &Template{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		URL: TemplateURL{
			ElasticsearchEndpoint: "http://elasticsearch:9200",
			Template:              "foo",
		},
		Spec: TemplateSpec{
			IndexPatterns: []string{
				"user_*",
			},
			Settings: &Settings{
				Index: IndexSettings{
					NumberOfShards:   &intstr.IntOrString{Type: 0, IntVal: 10, StrVal: ""},
					NumberOfReplicas: &intstr.IntOrString{Type: 0, IntVal: 2, StrVal: ""},
				},
			},
			Order:   10,
			Version: 5,
			Mappings: map[string]Mapping{
				"_doc": Mapping{
					Source: &Source{
						Enabled: true,
					},
					Properties: map[string]Property{
						"age": Property{
							Type: "integer",
						},
						"name": Property{
							Properties: map[string]Property{
								"first": Property{
									Type:  "keyword",
									Boost: 2.5,
								},
								"last": Property{
									Type: "keyword",
								},
							},
						},
					},
				},
			},
		},
	}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &Template{}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}
