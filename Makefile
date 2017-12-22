export GOPATH := $(GOPATH):$(shell pwd)/vendor:$(shell pwd)
GOBIN ?= $(shell pwd)/bin

NS ?= zion
VERSION ?= latest
IMAGE_NAME_GO ?= health-checker-build
IMAGE_NAME ?= health-checker

DB_HOST ?= localhost
DB_NAME ?= checker
DB_USER ?= postgres

MIGRATIONS_PATH = src/checker/storage/postgres/migrations

.PHONY:install-deps
install-deps:
	$(info * Install dependencies)
	cd src && glide install
	cd src && glide rebuild

.PHONY: build-local
build-local: migrate
	$(info * Building service )
	@go build -o bin/service cmd/service/main.go

.PHONY: build-local-cli
build-local-cli: migrate
	$(info * Building cli )
	@go build -o bin/cli cmd/cli/main.go

.PHONY: run-local
run-local: migrate build-local
	bin/service -database $(DB_NAME) -host $(DB_HOST) -user $(DB_USER) -interval 10 -source src/checker/testdata/url_list.txt

.PHONY: run-detail-stats
run-detail-stats: migrate build-local-cli
	bin/cli -database $(DB_NAME) -host $(DB_HOST) -user $(DB_USER) -url "https://google.com" -start 01-01-2017 -end 31-12-2017

.PHONY: run-base-stats
run-base-stats: migrate build-local-cli
	bin/cli -database $(DB_NAME) -host $(DB_HOST) -user $(DB_USER) -start 01-01-2017 -end 31-12-2017

.PHONY: migrate
migrate:
	$(info * Creating migration for go-bindata)
	@bin/go-bindata -pkg migrations -ignore "\w+\.go" -prefix "$(MIGRATIONS_PATH)" -o $(MIGRATIONS_PATH)/bindata.go $(MIGRATIONS_PATH)


## Docker build rules

.PHONY: build-go
build-go: build/Dockerfile.build
	$(info * Building docker build container $(NS)/$(IMAGE_NAME_GO))
	@docker build -f build/Dockerfile.build -t $(NS)/$(IMAGE_NAME_GO):$(VERSION) .

.PHONY: build
build: build-go build/Dockerfile
	$(info * Building $(VERSION) version of $(NS)/$(IMAGE_NAME))
	@docker run --rm -v "$(PWD)":/app -w /app $(NS)/$(IMAGE_NAME_GO):$(VERSION) make install-deps build-local build-local-cli
	@docker build -f build/Dockerfile -t $(NS)/$(IMAGE_NAME):$(VERSION) .

.PHONY: push
push:
	docker push $(NS)/$(IMAGE_NAME):$(VERSION)

.PHONY: shell
shell: build-go
	docker run --rm -it -v "$(PWD)":/app -w /app $(NS)/$(IMAGE_NAME_GO):$(VERSION) /bin/bash

.PHONY: release
release: build push
