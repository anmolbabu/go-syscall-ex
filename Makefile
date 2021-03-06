PROJECT := github.com/anmolbabu/tcp-server
GITCOMMIT := $(shell git rev-parse --short HEAD 2>/dev/null)
PKGS := $(shell go list  ./... | grep -v $(PROJECT)/vendor)
BUILD_FLAGS := -ldflags="-w -X $(PROJECT)/cmd.GITCOMMIT=$(GITCOMMIT)"
FILES := tcp-server dist

default: bin

.PHONY: bin
bin:
	go build ${BUILD_FLAGS} -o go-syscall-ex main.go

.PHONY: test
test:
	go test -race $(PKGS)