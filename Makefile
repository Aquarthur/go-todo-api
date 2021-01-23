SHELL := /bin/bash

COVERAGE_OUT := "test-coverage.out"
EXECUTABLE_NAME := "main"

.PHONY: all build run clean nistall uninstall fmt vet lint test

all: fmt vet lint test run

build:
	@echo "Cleaning dependencies..."
	@go mod tidy
	@go build -o $(EXECUTABLE_NAME) .

run: build
	@echo ""
	@echo "Starting $(EXECUTABLE_NAME)..."
	@./$(EXECUTABLE_NAME)

clean:
	@rm -f $(EXECUTABLE_NAME)
	@rm -f $(COVERAGE_OUT)

fmt:
	@echo "Running go fmt..."
	@go fmt $$(go list ./... | grep -v ./vendor/)

vet:
	@echo "Running go fmt..."
	@go vet $$(go list ./... | grep -v ./vendor/)

lint:
	@echo "Running go fmt..."
	@golint $$(go list ./... | grep -v ./vendor/)

test:
	@echo "Running tests..."
	@go test -coverprofile=$(COVERAGE_OUT) $$(go list ./... | grep -v ./vendor/)

coverage: test
	@echo "Showing test coverage..."
	@go tool cover -html=./$(COVERAGE_OUT)
