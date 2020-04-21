include .env
export

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

OS := $(shell uname)
HOST_IP=$(shell ipconfig getifaddr en0)
PROJECT_NAME := $(shell basename "$(PWD)")

STDERR := /tmp/.$(PROJECT_NAME)-stderr.txt
PID := .$(PROJECT_NAME).pid

GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

DB_MIGRATION_PATH=db/migrations

.PHONY: help
all: help
help:
	@echo "---------------------------------------------"
	@echo "List of available targets:"
	@echo "  start                    - Start the Go application."
	@echo "  stop                     - Stop the Go application."
	@echo "  compile                  - Compile Go binary."
	@echo "  build                    - Build Go binary."
	@echo "  clean                    - Clean up project."
	@echo "  install                  - Install dependencies."
	@echo "  help                     - Shows this dialog."
	@exit 0

.PHONY: install
install: dep_tidy

start:
	bash -c "trap 'make stop' EXIT; $(MAKE) compile start-server watch run='make compile start-server'"

stop: stop-server

run:
	@echo "$(OK_COLOR)==> $(PROJECT_NAME) is available at $(APP_URL) (DEV) $(NO_COLOR)"
	@go run main.go

start-server: stop-server
	@echo "$(OK_COLOR)==> $(PROJECT_NAME) is available at $(APP_URL) $(NO_COLOR)"
	@$(GOBIN)/$(PROJECT_NAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

compile: dep_get build

build:
	@echo "$(OK_COLOR)==> Building binary... $(NO_COLOR)"
	@GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECT_NAME) $(GOFILES)

dep_tidy:
	@echo "$(OK_COLOR)==> Checking if there is any missing dependencies... $(NO_COLOR)"
	@go mod tidy

dep_get:
	@echo "$(OK_COLOR)==> Get dependency... $(NO_COLOR)"
	@GOBIN=$(GOBIN) go get

clean:
	@ echo "$(OK_COLOR)==> Clean up build... $(NO_COLOR)"
	@ go clean -modcache
	@ rm -rf $(GOBIN)

migrate_up:
	@echo "$(OK_COLOR)==> Migrating DB... $(NO_COLOR)"
	@migrate -source file://${DB_MIGRATION_PATH} -database ${DB_DRIVE}://${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} up

migrate_down:
	@echo "$(OK_COLOR)==> Migrating DB... $(NO_COLOR)"
	@migrate -source file://${DB_MIGRATION_PATH} -database ${DB_DRIVE}://${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} down

watch:
	@GOBIN=$(GOBIN) yolo -i . -e vendor -e bin -c "$(run)"
