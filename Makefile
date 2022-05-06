SHELL := /bin/bash

MAJOR := 0
MINOR := 0
MICRO := 0
NEXT_MICRO := 2
CURRENT_VERSION_MICRO := $(MAJOR).$(MINOR).$(MICRO)

DATE                = $(shell date +'%d.%m.%Y')
TIME                = $(shell date +'%H:%M:%S')

KERNEL=$(shell if [ "$$(uname -s)" == "Linux" ]; then echo linux; fi)
ARCH=$(shell if [ "$$(uname -m)" == "x86_64" ]; then echo amd64; fi)

.PHONY: build fmt vet test clean install acctest local-dev-install vendor docs



all: build docs

vendor:
	mkdir -p vendor_talos
	git clone --depth=1 https://github.com/siderolabs/talos.git vendor_talos/talos
	mv talos vendor_talos/talos
	go mod vendor

fmt:
	@echo " -> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -path ./talos_vendor -prune -o -name '*.go' -print) | grep '^'

vet:
	@echo " -> vetting code"
	@go vet ./...

test:
	@echo " -> testing code"
	@go test -v ./...

build: docs
	@echo " -> Building"
	goreleaser build --rm-dist --single-target --snapshot
	@echo "Built terraform-provider-talos"

docs:
	tfplugindocs

install: build
	cp dist/provider-terraform-talos_linux_amd64_v1/provider-terraform-talos_* $$GOPATH/bin/terraform-provider-talos

local-dev-install: build
	find examples -name '.terraform.lock.hcl' -delete
	@echo "$(CURRENT_VERSION_MICRO)"
	@echo "$(KERNEL)"
	@echo "$(ARCH)"
	mkdir -p ~/.terraform.d/plugins/localhost/j-lgs/talos/$(MAJOR).$(MINOR).$(NEXT_MICRO)/$(KERNEL)_$(ARCH)/
	cp dist/provider-terraform-talos_linux_amd64_v1/provider-terraform-talos_* ~/.terraform.d/plugins/localhost/j-lgs/talos/$(MAJOR).$(MINOR).$(NEXT_MICRO)/$(KERNEL)_$(ARCH)/terraform-provider-talos
