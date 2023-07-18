# -- SH Adjustments
null      :=
SPACE     := $(null) $(null)
export GO111MODULE=on
export GOFLAGS=-mod=vendor

# -- Code
PROJECT_PATH = $(subst $(notdir $(SPACE)),/,$(CURDIR))
# Get dir        $(notdir $(subst $(SPACE),,$(CURDIR)))
# Get parent dir $(subst $(notdir $(CURDIR)),,$(CURDIR))
GO = go
BIN	= $(PROJECT_PATH)bin
FILES = $(wildcard *.go)

#-- Management
MODULE   = $(shell env GO111MODULE=$(GO111MODULE) $(GO) list -m)
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=$(GO111MODULE)  $(GO) list ./...))
TESTPKGS = $(shell env GO111MODULE=$(GO111MODULE)  $(GO) list -f \
			'{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' \
			$(PKGS))
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m➡\033[0m")

# -- No params call
.PHONY: all
all: help

# Tools
$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) Building $(PACKAGE)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(PACKAGE) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret
GOLINT = $(BIN)/golint
$(BIN)/golint: PACKAGE=golang.org/x/lint/golint

.PHONY: build
build: $(BIN) ; $(info $(M) Building executable...) @ ## Build program binary
	$Q CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build \
		-tags release \
		-ldflags '-X $(MODULE)/cmd.Version=$(VERSION) -X $(MODULE)/cmd.BuildDate=$(DATE)' \
		-o $(BIN)/$(basename $(MODULE)) ./cmd/app/$(FILES)

.PHONY: run
run: ; $(info $(M) Running dev build (on the fly) ...) @ ## Run intermediate builds
	$Q $(GO) run ./cmd/app/$(FILES)

.PHONY: run-build
run-build: build | ; $(info $(M) Running dev build (compiled) ...) @ ## Run dev build
	$Q $(BIN)/$(basename $(MODULE))

.PHONY: install
install: ; $(info $(M) Installing dependencies...)	@ ## Install project dependencies
	$Q $(GO) mod vendor

# -- Testing
.PHONY: test-unit
test-unit: ; $(info $(M) Running unit tests ...)	@ ## Run unit tests
	$Q go test $(go list ./... | grep -v /vendor/ | grep -v /cmd/) -cover  ./...

.PHONY: test-full
test-full: ; $(info $(M) Running full tests ...)	@ ## Run all test wit coverage tests
	$Q go test -v -cover -covermode=atomic ./...

.PHONY:test-cover
test-cover: ; $(info $(M) Running coverage check ...)	@ ## Run coverage check
	$Q go test $(go list ./... | grep -v /vendor/ | grep -v /cmd/) -short -coverprofile=bin/cov.out ./...

.PHONY: proto-gen
proto-gen: ; $(info $(M) Generate protobuf sources ...)	@ ## Generate proto sources
	$Q protoc -I=./proto --go_opt=paths=source_relative --go_out=plugins=grpc:./proto ./proto/*.proto

# -- Misc
.PHONY: sys-check
sys-check: ; $(info $(M) Go environment checking...)	@ ## Check node version
	$Q go version

.PHONY: version
version: ; $(info $(M) Version: $(VERSION))	@ ## Current version

help:
	@echo "\nGo Clean Arch\n----------------"
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

