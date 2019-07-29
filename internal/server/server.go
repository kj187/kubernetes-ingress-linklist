package server

import (
	"log"
	"net/http"

	"github.com/kj187/kubernetes-ingress-linklist/internal/routes/index"

	"github.com/gorilla/mux"
	"github.com/kj187/kubernetes-ingress-linklist/internal/routes/check"
	"github.com/kj187/kubernetes-ingress-linklist/internal/routes/links"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", index.IndexHandler).Methods("GET")
	r.HandleFunc("/health", check.HealthCheckHandler).Methods("GET")
	r.HandleFunc("/favicon.ico", NullHandler)
	r.HandleFunc("/{namespace}", links.LinksHandler).Methods("GET")
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("./web"))))
	log.Fatal(http.ListenAndServe(":8080", r))
}

func NullHandler(w http.ResponseWriter, r *http.Request) {}
