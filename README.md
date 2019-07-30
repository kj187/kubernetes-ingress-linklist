# Kubernetes Ingress Linklist

UI with a list of all available Kubernetes ingresses per namespace

## Example UI

![Example UI](assets/documentation/example-ui.png "Example UI")

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

