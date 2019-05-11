GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
BINARY_BASE_NAME=systray-clock
BINARY_NAME=$(BINARY_BASE_NAME)-$(GOOS)-$(GOARCH)

.ONESHELL:
export SHELL:=/bin/bash

.PHONY: clean
clean:
	go clean
	rm -f $(BINARY_NAME)
	
.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test: 
	go test -v ./...

.PHONY: build
build:
	go build -o $(BINARY_NAME)

.PHONY: run 
run: build
	./$(BINARY_NAME)

.PHONY: all 
all: clean fmt test run

-include User.mk