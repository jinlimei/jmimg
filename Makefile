.PHONY: all build dep clean

all: build

dep:
	@go mod download
	@go mod verify

build: dep
	@/bin/bash ./scripts/build.sh

clean:
	@rm -f bin/jmimg
