GO              ?= GO15VENDOREXPERIMENT=1 go
GOPATH          := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))
PROMU           ?= $(GOPATH)/bin/promu
GOLINTER        ?= $(GOPATH)/bin/gometalinter
pkgs            = $(shell $(GO) list ./... | grep -v /vendor/)
TARGET          ?= logstash_exporter

PREFIX          ?= $(shell pwd)
BIN_DIR         ?= $(shell pwd)

all: clean format vet gometalinter build test

test:
	@echo ">> running tests"
	@$(GO) test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

gometalinter: $(GOLINTER)
	@echo ">> linting code"
	@$(GOLINTER) --install --update > /dev/null
	@$(GOLINTER) --config=./.gometalinter.json ./...

build: build_linux_amd64 build_linux_arm64 build_darwin_amd64 build_darwin_arm64

build_linux_amd64: $(PROMU)
	@echo ">> building binaries for amd64"
	@mkdir -p $(PREFIX)/bin/linux/amd64
	@GOOS=linux GOARCH=amd64 $(PROMU) build --prefix $(PREFIX)/bin/linux/amd64 logstash_exporter

build_linux_arm64: $(PROMU)
	@echo ">> building binaries for arm64"
	@mkdir -p $(PREFIX)/bin/linux/arm64
	@GOOS=linux GOARCH=arm64 $(PROMU) build --prefix $(PREFIX)/bin/linux/arm64 logstash_exporter

build_darwin_amd64: $(PROMU)
	@echo ">> building binaries for amd64"
	@mkdir -p $(PREFIX)/bin/darwin/amd64
	@GOOS=darwin GOARCH=amd64 $(PROMU) build --prefix $(PREFIX)/bin/darwin/amd64 logstash_exporter

build_darwin_arm64: $(PROMU)
	@echo ">> building binaries for arm64"
	@mkdir -p $(PREFIX)/bin/darwin/arm64
	@GOOS=darwin GOARCH=arm64 $(PROMU) build --prefix $(PREFIX)/bin/darwin/arm64 logstash_exporter

clean:
	@echo ">> Cleaning up"
	@find . -type f -name '*~' -exec rm -fv {} \;
	@rm -fv $(TARGET)

$(GOPATH)/bin/promu promu:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get -u github.com/prometheus/promu

$(GOPATH)/bin/gometalinter lint:
	@GOOS=$(shell uname -s | tr A-Z a-z) \
		GOARCH=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m))) \
		$(GO) get -u github.com/alecthomas/gometalinter

.PHONY: all format vet build test promu clean $(GOPATH)/bin/promu $(GOPATH)/bin/gometalinter lint
