# Makefile for building and testing a Go project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=myapp

all: test build

build:
    $(GOBUILD) -o $(BINARY_NAME) -v

test:
    $(GOTEST) -v ./...

clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)

run:
    $(GOBUILD) -o $(BINARY_NAME) -v ./...
    ./$(BINARY_NAME)

deps:
    $(GOGET) -v ./...

.PHONY: all build test clean run deps
