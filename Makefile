# vi: set ft=make ts=2 sw=2 sts=0 noet:

SHELL := /bin/bash
COVERPROFILE := coverage.out
TARGET := dbreport
CMD := $(wildcard cmd/*)

.PHONY: default
default: help

# http://postd.cc/auto-documented-makefile/
.PHONY: help help-common
help: help-common

help-common:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m %-30s\033[0m %s\n", $$1, $$2}'

test: ## go test
	go test -test.v ./...
	golangci-lint run ./...

cover: ## coverage
	go test -coverprofile=$(COVERPROFILE) -test.v ./...
	go tool cover -func=$(COVERPROFILE)

fmt: ## go fmt
	gofmt -l -s -d -w .

