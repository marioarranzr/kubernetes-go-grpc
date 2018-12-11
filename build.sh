#!/bin/bash
set -eo pipefail

# Generate proto files
protoc -I ./pb --go_out=plugins=grpc:./pb ./pb/*.proto

# Just in case it is necessary to remove the docker images
docker rmi local/api -f
docker rmi local/gcd -f

# Build docker images
docker build -t local/gcd -f Dockerfile.gcd .
docker build -t local/api -f Dockerfile.api .

# Just in case it is necessary to remove the kubernetes services
kubectl delete service gcd-service
kubectl delete service api-service

# Just in case it is necessary to remove the kubernetes deployments
kubectl delete deployment api-deployment
kubectl delete deployment gcd-deployment

# Apply configuration to kubernetes
kubectl apply -f api.yaml
kubectl apply -f gcd.yaml
