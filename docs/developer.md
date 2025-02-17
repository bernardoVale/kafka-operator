# Developer Guide

## How to run Kafka-operator in your cluster with your changes

The Kafka operator is built on the [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) project.

To build the operator and run tests:

1. Run `make`

If you make changes and would like to try your own version, create your own image:

1. `make docker-build IMG={YOUR_USERNAME}/kafka-operator:v0.0.1`
2. `make docker-push IMG={YOUR_USERNAME}/kafka-operator:v0.0.1`
3. `make deploy IMG={YOUR_USERNAME}/kafka-operator:v0.0.1`

Watch the operator's logs with:

`kubectl logs -f -n kafka kafka-operator-controller-manager-0 -c manager`

Alternatively, run the operator on your machine:

1. `export $KUBECONFIG`
2. `make install`
3. `make run`

Create CR and let the operator set up Kafka in your cluster (you can change the `spec` of `Kafka` for your needs in the yaml file):

> Remember you need Zookeeper server to run Kafka

`kubectl create -n kafka -f config/samples/banzaicloud_v1alpha1_kafkacluster.yaml`

#### Limitations on minikube

Minikube does not have a load balancer implementation, thus our envoy service will not get an external IP and the operator will get stuck at this point.

A possible solution to overcome this problem is to use https://github.com/elsonrodriguez/minikube-lb-patch. The operator will be able to proceed if you run the following command:
```go
kubectl run minikube-lb-patch --replicas=1 --image=elsonrodriguez/minikube-lb-patch:0.1 --namespace=kube-system
``` 
