#!/bin/sh

go get -d ./...
go run -race ./cmd/main.go

eval "$@"