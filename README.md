# Kubernetes Ingress Linklist


[![pipeline](https://github.com/kj187/kubernetes-ingress-linklist/workflows/pipeline/badge.svg?branch=master)](https://github.com/kj187/kubernetes-ingress-linklist/actions?query=workflow%3Apipeline)

UI with a list of all available Kubernetes ingresses per namespace

## Example UI

![Example UI](assets/documentation/example-ui.png "Example UI")

## Usage

### Deploy using Helm Chart

TODO

### Deploying to Minikube

TODO 

```
kubectl apply -f deploy/
kubectl port-forward service/kuberetes-ingress-linklist 8080:80
```


## Custom pages

It is also possible to add custom pages via settings. 
To do that, just create a new ConfigMap and mount this into the `kubernetes-ingress-linklist`.

The mount path must be `/kubernetes-ingress-linklist` and the filename must be `settings.yaml` 

**Example: `/kubernetes-ingress-linklist/settings.yaml`**
```
customPages:
  demo1:
    title: Demo 1
    links:
    - title: Demo Link 1
      link: http://kj187.de
    - title: Demo Link 2
      link: http://www.aoe.com
    - title: Demo Link 3
      link: http://www.google.com
  demo2:
    title: Demo 2
    links:
    - title: Demo 2 Link 1
      link: http://kj187.de
```


![Example custom page](assets/documentation/custom-page.png "Example custom page")

## Index page redirect

You can set a temporary redirect via settings. The default behaviour is a redirect to the `default` namespace.

```
redirectRootPageTo: /default
```

You can change this by overriding the value. You can add a namespace or a custom page (keep in mind to prefix the path with `custom/`). 
Example: 

```
redirectRootPageTo: /custom/demo1
customPages:
  demo1:
    title: Demo 1
    links:
    - title: Demo Link 1
      link: http://kj187.de
```