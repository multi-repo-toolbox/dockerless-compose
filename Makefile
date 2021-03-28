export REPO?=github.com/multi-repo-tools/dockerless

GIT_HASH:=$(shell git log -1 --pretty=format:"%h")
GIT_HEAD:=$(shell git rev-parse HEAD)
GIT_BRANCH:=$(shell git describe --contains --always --all | grep -v HEAD | head -1 | cut -d'/' -f3)
BUILD_AT:=$(shell date +"%Y-%m-%dT%H:%M:%S%z")
GIT_TAG:=$(shell git show-ref --tags -d | grep $(GIT_HASH) | sort -Vr | head -1 | cut -d'/' -f3)
DEFAULT:=master
ifdef GIT_TAG
  VERSION=$(GIT_TAG)
else
  ifdef GIT_BRANCH
    VERSION=$(GIT_BRANCH)
  else
    VERSION=$(DEFAULT)
  endif
endif
BUILD_FLAGS:="-s \
-X $(REPO)/command/version.AppName=$(REPO)\
-X $(REPO)/command/version.GitHash=$(GIT_HASH)\
-X $(REPO)/command/version.Version=$(VERSION)\
-X $(REPO)/command/version.BuildAt=$(BUILD_AT)\
"

build:
	go build -ldflags $(BUILD_FLAGS) -o dockerless cmd/dockerless/*.go

build-race:
	go build -race -ldflags $(BUILD_FLAGS) -o dockerless cmd/dockerless/*.go

run:
	go run -ldflags $(BUILD_FLAGS) cmd/dockerless/*.go

test:
	go test -v ./...

.PHONY: build
