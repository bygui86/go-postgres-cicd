
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

start-postgres :		## Start PostgreSQL container
	docker run -d --rm --name postgres \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		postgres:13.1-alpine

stop-postgres :		## Stop PostgreSQL container
	docker container stop postgres


## application

build :		## Build application
	go build

test :		## Run all tests
	go clean -testcache ./...
	go test -v -count=5 -race -cover -coverprofile=coverage.out ./...
	go test -v -count=1 -race -cover -coverprofile=coverage.out -tags=integration ./...

unit-test :		## Run unit tests only
	go clean -testcache ./...
	go test -v -count=5 -race -cover -coverprofile=coverage.out ./...

integr-test :		## Run integration tests only
	go clean -testcache ./...
	go test -v -count=1 -race -cover -coverprofile=coverage.out -tags=integration ./...

coverage-cli :		## Show coverage results in cli
	go tool cover -func=coverage.out

coverage-browser :		## Show coverage results in browser
	go tool cover -html=coverage.out

run :		## Run application from source code
	godotenv -f local.env go run main.go


## container

__check-container-tag :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

container-build : __check-container-tag		## Build container
	docker build . -t bygui86/go-postgres-cicd:$(CONTAINER_TAG) --no-cache

container-push : __check-container-tag		## Push container to Docker hub
	docker push bygui86/go-postgres-cicd:$(CONTAINER_TAG)


## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
