Elasticsearch DB Operator
=========================

[![Go Report](https://goreportcard.com/badge/kaitoy/elasticsearch-db-operator)](https://goreportcard.com/report/kaitoy/elasticsearch-db-operator)

Elasticsearch DB Operator is a [Kubernetes operator](https://github.com/operator-framework/awesome-operators) built with [Kubebuilder](https://github.com/kubernetes-sigs/kubebuilder), that operates [Elasticsearch](https://www.elastic.co/products/elasticsearch) database (i.e. indices and settings).

Deploy to Kubernetes cluster
-----------------------------------

You need [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) to deploy Elasticsearch DB Operator with it.

1. `git clone https://github.com/kaitoy/elasticsearch-db-operator.git`
2. `cd elasticsearch-db-operator`
3. `kubectl apply -f config/crds/elasticsearchdb_v1beta1_index.yaml`
4. `kubectl apply -f config/crds/elasticsearchdb_v1beta1_template.yaml`
5. `kubectl apply -f elasticsearch-db-operator.yaml`

Usage
-----

### Operate an index

1. Write an index manifest.

    e.g.)

    ```yaml
    apiVersion: elasticsearchdb.kaitoy.github.com/v1beta1
    kind: Index
    metadata:
      labels:
        controller-tools.k8s.io: "1.0"
      name: index-sample
    url:
      elasticsearchEndpoint: http://elasticsearch:9200
      index: user
    spec:
      settings:
        index:
          number_of_shards: 5
          number_of_replicas: 1
      mappings:
        _doc:
          _source:
            enabled: true
          properties:
            age:
              type: integer
            name:
              properties:
                first:
                  type: keyword
                  boost: 2.5
                last:
                  type: keyword
    ```

2. Apply the manifest to create the index.
3. Delete the manifest to delete the index.

### Operate a template

1. Write a template manifest.

    e.g.)

    ```yaml
    apiVersion: elasticsearchdb.kaitoy.github.com/v1beta1
    kind: Template
    metadata:
      labels:
        controller-tools.k8s.io: "1.0"
      name: template-sample
    url:
      elasticsearchEndpoint: http://elasticsearch:9200
      template: user_template
    spec:
      index_patterns:
        - user-*
      settings:
        index:
          number_of_shards: 5
          number_of_replicas: 1
      order: 10
      version: 3
      mappings:
        _doc:
          _source:
            enabled: true
          properties:
            age:
              type: integer
            name:
              properties:
                first:
                  type: keyword
                  boost: 2.5
                last:
                  type: keyword
    ```

2. Apply the manifest to create the template.
3. Delete the manifest to delete the template.

Development
-----------

You need [Docker](https://www.docker.com/) to build Elasticsearch DB Operator container image, and [kustomize](https://github.com/kubernetes-sigs/kustomize) to generate a Kubernetes manifest file.

1. Build container image
    1. `git clone https://github.com/kaitoy/elasticsearch-db-operator.git`
    2. `cd elasticsearch-db-operator`
    3. `docker build .`
2. Generate Kubernetes manifest file
    1. `kustomize build config/default > elasticsearch-db-operator.yaml`
