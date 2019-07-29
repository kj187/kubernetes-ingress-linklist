
PACKAGES        = $(shell go list ./... | grep -v /vendor/)
GO_FILES        = $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
GO111          ?= on
CGO_ENABLED    ?= 0

default: info

info:
	@go version

setup: info ## Initial setup
	@GO111MODULE=$(GO111) go get k8s.io/client-go@v12.0.0
	@GO111MODULE=$(GO111) go mod vendor

lint: info ## Lint the files
	GO111MODULE=on go get -u golang.org/x/lint/golint
	golint -set_exit_status ${PACKAGES}

unit_test: info ## Run unittests
	GO111MODULE=$(GO111) go test -short ${PACKAGES}

go_build_linux: info ## Build the binary file for linux/amd64
	GO111MODULE=$(GO111) go mod vendor
	GO111MODULE=$(GO111) CGO_ENABLED=$(CGO_ENABLED) GOOS=linux GOARCH=amd64 go build -o bin/kubernetes-ingress-linklist -mod vendor cmd/kubernetes-ingress-linklist/main.go

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: info setup lint help