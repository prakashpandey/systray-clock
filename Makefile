GOPATH=$(shell go env GOPATH)
BINARY_NAME=x-clock

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

.PHONY: install 
install:
	go build -o $(GOPATH)/$(BINARY_NAME)

.PHONY: all 
all: clean fmt test run

-include User.mk