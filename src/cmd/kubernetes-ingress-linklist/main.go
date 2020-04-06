package main

import (
	"github.com/kj187/kubernetes-ingress-linklist/src/internal/app"
	"github.com/kj187/kubernetes-ingress-linklist/src/internal/server"
)

func main() {
	app.Init()
	server.StartServer()
}
