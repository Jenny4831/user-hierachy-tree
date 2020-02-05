# VARIABLES
PACKAGE="github.com/Jenny4831/project-d"
BINARY_NAME="userapp"
TEST_TIME_OUT='50s'
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)
# PROJECTNAME := "userapp"

default: usage

clean: ## Trash binary files
	@echo "--> cleaning..."
	@go clean || (echo "Unable to clean project" && exit 1)
	@rm -rf $(GOPATH)/bin/$(BINARY_NAME) 2> /dev/null
	@echo "Clean OK"

test: ## Run all tests
	@echo "--> testing..."
	@go clean -testcache
	@go test -timeout ${TEST_TIME_OUT} -v $(PACKAGE)/...

install: clean ## Compile sources and build binary
	@echo "--> building..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(BINARY_NAME) $(GOFILES)
	@echo "Install OK"

build: ## Run your application
	@echo "--> installing..."
	@go install $(PACKAGE)/src || (echo "Compilation error" && exit 1)
	@echo "Install OK"

run: install ## Run your application
	@echo "--> running application..."
	@$(GOBASE)/bin/$(BINARY_NAME)

usage: ## List available targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all test clean
