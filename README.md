# Microservices proof of concept using Go, gRPC and Kubernetes

## Running it locally

```
./runLocally.sh
```

It runs the `gcd` service in port 3000 and the server `api` that responds to our request in port 3001.
Try making a request to `localhost:3001/gcd/294/44`

## Running it in Minikube

Firstly, we need to install docker, minikube and protoc.

Switch to Docker daemon inside Minikube VM:
```
$ eval $(minikube docker-env)
```

Build images and deploy to Kubernetes cluster:
```
$ ./build.sh
```

Try it out.
```
$ curl $(minikube service api-service --url)/gcd/294/462
```
