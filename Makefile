PROJECT_NAME:=bc_prototype
FILE_HASH := $(shell git rev-parse HEAD)
GOLANGCI_LINT := $(shell command -v golangci-lint 2> /dev/null)
SOLC_DOCKER_IMAGE=ethereum/solc:stable

init_repo: ## create necessary configs
	cp configs/sample.common.env configs/common.env
	cp configs/sample.app_conf.yml configs/app_conf.yml
	cp configs/sample.app_conf_docker.yml configs/app_conf_docker.yml
	find . -type f -name "*.go" -exec sed -i 's/go_project_template/${PROJECT_NAME}/g' {} +
	find . -type f -name "*.mod" -exec sed -i 's/go_project_template/${PROJECT_NAME}/g' {} +
	go mod tidy && go mod vendor
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -local github.com/$(PROJECT_NAME) -w .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install-lint: ## Installs golangci-lint tool which a go linter
ifndef GOLANGCI_LINT
	${info golangci-lint not found, installing golangci-lint@latest}
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
endif

gogen: ## generate code
	${info generate code...}
	go generate ./internal...

test: ## Runs tests
	${info Running tests...}
	go test -v -race ./... -cover -coverprofile cover.out
	go tool cover -func cover.out | grep total

bench: ## Runs benchmarks
	${info Running benchmarks...}
	go test -bench=. -benchmem ./... -run=^#

lint: install-lint ## Runs linters
	@echo "-- linter running"
	golangci-lint run -c .golangci.yaml ./internal...
	golangci-lint run -c .golangci.yaml ./cmd...

stop: ## Stops the local environment
	${info Stopping containers...}
	docker container ls -q --filter name=${PROJECT_NAME} ; true
	${info Dropping containers...}
	docker rm -f -v $(shell docker container ls -q --filter name=${PROJECT_NAME}) ; true

dev_up: stop ## Runs local environment
	${info Running docker-compose up...}
	GIT_HASH=${FILE_HASH} docker compose -p ${PROJECT_NAME} up --build dbPostgres

build: ## Builds binary
	@echo "-- building binary"
	go build -o ./bin/binary ./cmd

run: ## Runs binary local with environment in docker
	${info Run app containered}
	GIT_HASH=${FILE_HASH} docker compose -p ${PROJECT_NAME} up --build

migrate_new: ## Create new migration
	migrate create -ext sql -dir migrations -seq data

migrate:
	migrate -path migrations/ -database "postgres://aHAjeK:AOifjwelmc8dw@127.0.0.1:5249/sybill?sslmode=disable" up

compile_contract:
	docker run --rm -v $(PWD)/internal/contracts:/contracts $(SOLC_DOCKER_IMAGE) \
            --bin --abi /contracts/CircleStorage.sol -o /contracts \
            --overwrite

	#docker run --rm -v $(PWD)/internal/service/circle:/abi ethereum/client-go abigen --abi /abi/CircleContract.abi --pkg service --out /abi/circle_contract.go
	abigen --bin=$(PWD)/internal/contracts/CircleStorage.bin --abi $(PWD)/internal/contracts/CircleStorage.abi --pkg contracts --out $(PWD)/internal/contracts/circle_contract.go

.PHONY: help install-lint test gogen lint stop dev_up build run init_repo migrate_new
.DEFAULT_GOAL := help