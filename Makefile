.PHONY: build, build-api, start-web

ROOT_DIR = $(PWD)
API_DIR = $(ROOT_DIR)/api
WEB_DIR = $(ROOT_DIR)/web

build: build-api build-web

build-api:
	$(MAKE) -C api

build-web:
	$(MAKE) -C web

start-web: build-web
	$(MAKE) -C web start

start-api: build-api
	$(MAKE) -C api start
