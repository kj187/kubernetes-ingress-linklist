package main

import (
	"github.com/kj187/kubernetes-ingress-linklist/internal/app"
	"github.com/kj187/kubernetes-ingress-linklist/internal/server"
)

func main() {
	app.Init()
	server.StartServer()
}
