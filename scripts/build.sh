#!/bin/bash

BUILD_TIME=$(date +"%Y-%m-%d.%H:%M:%S")
COMMIT=$(git log -1 --pretty=format:%H || echo 'N/A')
GO_VERSION=$(go version | awk '{ print $3 }')
TAG=$(git tag || echo 'N/A')
VERSION=$(date +"%Y-%m-%d.%H%M%S")

LDFLAGS="-s -w"
LDFLAGS="${LDFLAGS} -X main.BuildTime=$BUILD_TIME"
LDFLAGS="${LDFLAGS} -X main.CommitHash=$COMMIT"
LDFLAGS="${LDFLAGS} -X main.GoVersion=$GO_VERSION"
LDFLAGS="${LDFLAGS} -X main.GitTag=$TAG"
LDFLAGS="${LDFLAGS} -X main.Version=$VERSION"

go build -o bin/jmimg -compiler gc -ldflags "${LDFLAGS}" cmd/jmimg/*.go
echo "BUILT AT $BUILD_TIME"
