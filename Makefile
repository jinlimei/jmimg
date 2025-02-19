.PHONY: all build dep clean

all: build

dep:
	@go mod download
	@go mod verify

build: dep
	@/bin/bash ./scripts/build.sh

install: dep
	@/bin/bash ./scripts/install.sh

clean:
	@rm -f bin/jmimg
