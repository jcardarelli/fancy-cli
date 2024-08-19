.DEFAULT_GOAL := build

fmt:
	go fmt
.PHONY:fmt

lint: fmt
	golangci-lint run
.PHONY:lint

vet: fmt
	go vet
.PHONY:vet

build: vet
	go build -o fancy
.PHONY:build

local-github-actions:
	act --container-architecture linux/amd64 --remote-name upstream
