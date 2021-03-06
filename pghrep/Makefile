.DEFAULT_GOAL = all

BINARY = pghrep
GOARCH = amd64

VERSION?=0.1
BUILD_TIME?=$(shell date -u '+%Y%m%d-%H%M')
COMMIT?=no #$(shell git rev-parse HEAD)
BRANCH?=no #$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=dmius
BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}
#BUILD_DIR=${GOPATH}/${BINARY}

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -w \
	-X main.version=${VERSION} \
	-X main.commit=${COMMIT} \
	-X main.branch=${BRANCH}\
	-X main.buildTime=${BUILD_TIME}"

PLUGINS_SRC := $(wildcard plugins/*.go)

# Build the project
all: clean install vet main

install:
	go get golang.org/x/net/html
	go get github.com/dustin/go-humanize
	go get gopkg.in/yaml.v2

main:
	GOARCH=${GOARCH} go build ${LDFLAGS} -o bin/${BINARY} ./src/

test:
	go test ./src/
	go test ./src/fmtutils/
	go test ./src/checkup/*/

vet:
	go vet ./src/...

fmt:
	go fmt $$(go list ./... | grep -v /vendor/)

clean:
	-rm -f bin/*

run:
	go run ${LDFLAGS} ./src/*

.PHONY: all main test vet fmt clean run

