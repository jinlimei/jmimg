PKG := "github.com/jinlimei/jmimg"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
GOPATH := $(shell go env GOPATH)
CC := clang-15
CXX := clang-15

export DCWD=$(shell pwd)

.PHONY: all build dep clean

all: build

dep:
	@go mod download
	@go mod verify

build: dep
	@go build -o bin/jmimg -compiler gc -ldflags="-s -w" cmd/jmimg/*.go

clean:
	@rm -f bin/jmimg