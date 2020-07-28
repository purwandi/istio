#!bin/sh

set -e

SERVICE=$1
HELM_BIN=~/Code/bin/onelabs/helm 

kubectl -n istio-system get configmap istio-sidecar-injector -o=jsonpath='{.data.config}' > ./helm-charts/istio/inject-config.yaml
kubectl -n istio-system get configmap istio-sidecar-injector -o=jsonpath='{.data.values}' > ./helm-charts/istio/inject-values.yaml
kubectl -n istio-system get configmap istio -o=jsonpath='{.data.mesh}' > ./helm-charts/istio/mesh-config.yaml

rm -rf ./helm-charts/$SERVICE-service
$HELM_BIN template ./helm-charts/$SERVICE --values ./helm-charts/$SERVICE/values.yaml --output-dir ./helm-charts

# Istio injections
for chart in ./helm-charts/$SERVICE-service/templates/*.yaml; do
  istioctl kube-inject \
    --injectConfigFile ./helm-charts/istio/inject-config.yaml \
    --meshConfigFile ./helm-charts/istio/mesh-config.yaml \
    --valuesFile ./helm-charts/istio/inject-values.yaml \
    --filename $chart > $chart.bak
  
   rm -rf $chart
   mv $chart.bak $chart
done

cp ./helm-charts/$SERVICE/*.yaml ./helm-charts/$SERVICE-service/
$HELM_BIN upgrade $SERVICE-service ./helm-charts/$SERVICE-service/ --install 