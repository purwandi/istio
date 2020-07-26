#!bin/sh

set -e

docker build -t docker.pkg.github.com/purwandi/istio/frontend:latest ./frontend
docker build -t docker.pkg.github.com/purwandi/istio/product:latest ./product
docker build -t docker.pkg.github.com/purwandi/istio/review:latest ./review

docker push docker.pkg.github.com/purwandi/istio/frontend
docker push docker.pkg.github.com/purwandi/istio/product
docker push docker.pkg.github.com/purwandi/istio/review