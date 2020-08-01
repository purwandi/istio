# Instal Istio

TLDR, istio is open platform to create service mesh in your microservice.

## About

About this project is to simulate to installing istio using your own existing 
grafana and prometheus. 

## Why?

It's very hard to manually install grafana and prometheus because the istio website
lack documentation about to installing and configuring individual addons component

## Installing

**Dump Istio Profile**

```sh
istioctl profile dump demo > istio-dump.yaml
```

**Edit yaml file**

Edit `istio-dump.yaml` to disable grafana and prometheus component

```yaml
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
spec:
  addonComponents:
    grafana:
      enabled: false
      k8s:
        replicaCount: 1
    istiocoredns:
      enabled: false
    kiali:
      enabled: true
      k8s:
        replicaCount: 1
    prometheus:
      enabled: false
      k8s:
        replicaCount: 1
    tracing:
      enabled: true
```



Before

istioctl profile dump demo

istioctl manifest apply -f istio-dump.yaml \
  --set values.kiali.prometheusAddr="http://prometheus-server" \
  --set values.kiali.dashboard.grafanaURL="http://grafana"

istioctl manifest apply -f istio-dump.yaml \
  --set values.kiali.dashboard.grafanaURL="http://grafana"

istioctl manifest apply -f istio-dump.yaml \
  --set values.kiali.prometheusAddr="http://prometheus-server.default" \
  --set values.kiali.dashboard.grafanaURL="http://grafana.default"


istioctl manifest generate -f istio-dump.yaml | kubectl delete -f -

https://istiobyexample.dev/prometheus/#:~:text=Istio%20By%20Example-,Bring%20Your%20Own%20Prometheus,and%20your%20Envoy%2Dinjected%20workloads.


helm fetch stable/prometheus --version 9.1.2 --untar 

helm upgrade prometheus ./helm-charts/prometheus --values ./helm-charts/prometheus/values-istio.yaml --install --namespace istio-system --debug 