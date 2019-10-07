.PHONY: build

ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

GOCMD = go
GOBUILD = $(GOCMD) build

BINARY_PATH = "./build"

build:
	@CGO_ENABLED=0 $(GOBUILD) -o "$(BINARY_PATH)/scg/scg" -v ./cmd/scg/*.go
	@cp -r $(ROOT_DIR)/config $(BINARY_PATH)/scg/
	echo "api build done"

start:
	$(BINARY_PATH)/scg/scg start

