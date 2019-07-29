package links

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kj187/kubernetes-ingress-linklist/internal/routes"
	"github.com/kj187/kubernetes-ingress-linklist/pkg/kubernetes"
)

// LinksHandler ...
func LinksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := "default"
	if vars["namespace"] != "" {
		namespace = vars["namespace"]
	}

	k8s := kubernetes.Kubernetes{Namespace: namespace}
	namespaces, err := k8s.GetNamespaces()
	routes.CheckAndWriteHTTPError(err, w)
	ingresses, err := k8s.GetIngresses()
	routes.CheckAndWriteHTTPError(err, w)

	tpl, err := template.New("").Funcs(routes.GetFuncMap()).ParseFiles(
		"internal/frontend/templates/namespace.gohtml", "internal/frontend/layouts/base.gohtml",
	)
	routes.CheckAndWriteHTTPError(err, w)
	routes.CheckAndWriteHTTPError(tpl.ExecuteTemplate(w, "layout", struct {
		CurrentNamespace string
		Namespaces       kubernetes.NamespacesOutput
		Ingresses        kubernetes.IngressesOutput
	}{
		namespace,
		namespaces,
		ingresses,
	}), w)
}
