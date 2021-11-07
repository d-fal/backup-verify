
GOPATH:=$(shell go env GOPATH)

APP_NAME="bverify"



.PHONY: build
build:
	go build -o $(APP_NAME) main.go

.PHONY: test
test:
	go test -v `go list ./...` -cover

