package server

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kj187/kubernetes-ingress-linklist/pkg/kubernetes"
)

// StartServer ...
func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/favicon.ico", NullHandler)
	r.HandleFunc("/{namespace}", LinksHandler).Methods("GET")
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}

// NullHandler ...
func NullHandler(w http.ResponseWriter, r *http.Request) {}

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "//"+r.Host+"/default", http.StatusPermanentRedirect)
}

// HealthCheckHandler ...
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}

// LinksHandler ...
func LinksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := "default"
	if vars["namespace"] != "" {
		namespace = vars["namespace"]
	}

	k8s := kubernetes.Kubernetes{Namespace: namespace}
	namespaces, err := k8s.GetNamespaces()
	CheckAndWriteHTTPError(err, w)
	ingresses, err := k8s.GetIngresses()
	CheckAndWriteHTTPError(err, w)

	tpl, err := template.New("").Funcs(GetFuncMap()).ParseFiles(
		"internal/frontend/templates/namespace.gohtml", "internal/frontend/layouts/base.gohtml",
	)
	CheckAndWriteHTTPError(err, w)
	CheckAndWriteHTTPError(tpl.ExecuteTemplate(w, "layout", struct {
		CurrentNamespace string
		Namespaces       kubernetes.NamespacesOutput
		Ingresses        kubernetes.IngressesOutput
	}{
		namespace,
		namespaces,
		ingresses,
	}), w)
}

// CheckAndWriteHTTPError ...
func CheckAndWriteHTTPError(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 500)
	}
}

// GetFuncMap ...
func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"example": func(title string) string { return "example" },
	}
}
