CONFIG_FILE ?= ./config/local.yml

.PHONY: default
default: help

.PHONY: run
run: ## run the API server
	go build -o bin/firecracker-api cmd/api/main.go && sudo ./bin/firecracker-api
	#go run cmd/api/main.go
