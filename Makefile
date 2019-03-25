.PHONY: help test build vendor

export GO111MODULE=on
GO=go
VERSION=0.0.1

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.DEFAULT_GOAL:= build

build:
	@cd cmd/goecma && ${GO} build -ldflags '-X main.Version=${VERSION}' 
 
