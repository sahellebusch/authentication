#!/usr/bin/env make
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=30
ENV ?= dev

 # Go parameters	
GOCMD=go	
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
BINARY_NAME=main
SQLITEDB=test.db
MAIN=authentication.go

.PHONY: test clean

## Builds package
build: 
	$(GOBUILD) $(MAIN)

## Run the tests
test: 
	$(GOTEST) ./...

## Clean dev environment
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(SQLITEDB)

## Formats all files
fmt:
	$(GOFMT) ./...

## Runs authentication server
run:
	$(GOBUILD) main.go
	./$(BINARY_NAME)

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
