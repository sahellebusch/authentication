#!/usr/bin/env make
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=30
ENV ?= dev

SQLITEDB=test.db
MAIN=authentication.go
BINARY=authentication

.PHONY: test clean

## Builds package
build: lint vet fmt tidy
	go build authentication.go

## Cleans up go modules
tidy:
	go mod tidy
## Run the tests
test: 
	GO_ENV=test go test -v ./controllers

## Clean dev environment
clean:
	go clean
	rm -f $(BINARY)
	rm -f $(SQLITEDB)

## Formats all files
fmt:
	gofmt -w ./

## Runs govet on code
vet:
	go vet

##n Lints the code
lint:
	golint -set_exit_status

## Runs authentication server
run:
	go build $(MAIN)
	chmod u+x $(BINARY)
	./$(BINARY)

help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9\%]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
		} \
		{ lastLine = $$0 }' $(MAKEFILE_LIST)
