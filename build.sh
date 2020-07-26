#!bin/sh

set -e

docker build -t registry.gitlab.com/purwandi/istio/frontend:latest ./frontend
docker build -t registry.gitlab.com/purwandi/istio/product:latest ./product
docker build -t registry.gitlab.com/purwandi/istio/review:latest ./review

docker push registry.gitlab.com/purwandi/istio/frontend
docker push registry.gitlab.com/purwandi/istio/product
docker push registry.gitlab.com/purwandi/istio/review