GIT_COMMIT   ?= $(shell git rev-parse HEAD)
GIT_TAG      ?= $(shell git tag --points-at HEAD)
BRANCH       ?= $(shell git rev-parse --abbrev-ref HEAD)

GO           := go
PYTHON3      := python3

BINARY_NAME  := ebiten-ecs
PACKAGE      := ebiten-ecs
USER         := jtbonhomme 
VERSION=1.12.4

BUILD_CONSTS = -X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT) -X main.name=$(BINARY_NAME)
BUILD_OPTS   = -ldflags="$(BUILD_CONSTS) -s -w" -gcflags="-trimpath=$(GOPATH)/src"

version:
ifneq ($(GIT_TAG),)
	$(eval VERSION := $(GIT_TAG))
	$(eval VERSION_FILE := $(GIT_TAG))
else
	$(eval VERSION := $(subst /,-,$(BRANCH)))
	$(eval VERSION_FILE := $(GIT_COMMIT)-SNAPSHOT)
endif
	@test -n "$(VERSION)"
	$(info Version is $(VERSION)/$(VERSION_FILE) on sha1 $(GIT_COMMIT))

update-pkg-cache: version
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/$(USER)/$(PACKAGE)@v$(VERSION)

run-example:
	go run ./example/main.go
