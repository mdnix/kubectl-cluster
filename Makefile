GO := go
TARGET=kubectl-cluster
ARTIFACTS := _out
CURRENT_DIR := $(shell pwd)
VERSION=0.0.1-dev
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_HASH=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date "+%Y-%m-%d %H:%M:%S %Z")

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

compile: ## Compile the cli for linux-amd64 and darwin-amd64
	# 64-Bit Linux
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags "-w -s -X 'kubectl-cluster/cmd/version.BuildDate=$(BUILD_DATE)' -X 'kubectl-cluster/cmd/version.GitBranch=$(GIT_BRANCH)' -X 'kubectl-cluster/cmd/version.GitHash=$(GIT_HASH)'  -X 'kubectl-cluster/cmd/version.Version=$(VERSION)'" -o $(ARTIFACTS)/linux/$(TARGET)

	# 64-Bit Dardwin
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags "-w -s -X 'kubectl-cluster/cmd/version.BuildDate=$(BUILD_DATE)' -X 'kubectl-cluster/cmd/version.GitBranch=$(GIT_BRANCH)' -X 'kubectl-cluster/cmd/version.GitHash=$(GIT_HASH)'  -X 'kubectl-cluster/cmd/version.Version=$(VERSION)'" -o $(ARTIFACTS)/darwin/$(TARGET)

all: compile

.PHONY: clean
clean: ## Cleans up all artifacts
	@-rm -rf $(ARTIFACTS)
