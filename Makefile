GO_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | sort)

docker.start:
	docker compose --profile all up -d --quiet-pull

docker.stop:
	docker compose --profile all down

test:
	$(MAKE) docker.start
	$(MAKE) test.ci
	$(MAKE) docker.stop

test.ci:
	set -e; for dir in $(GO_MOD_DIRS); do \
	  echo "go test in $${dir}"; \
	  (cd "$${dir}" && \
	    go mod tidy && \
	    go vet && \
	    set CGO_ENABLED=1 go test -coverprofile=coverage.txt -covermode=atomic ./... -race); \
	done

bench:
	go test ./... -test.run=NONE -test.bench=. -test.benchmem

.PHONY: all test bench fmt

build:
	go build .

fmt:
	gofumpt -w ./
	goimports -w  -local github.com/redis/go-redis ./

go_mod_tidy:
	set -e; for dir in $(GO_MOD_DIRS); do \
	  echo "go mod tidy in $${dir}"; \
	  (cd "$${dir}" && \
	    go get -u ./... && \
	    go mod tidy); \
	done

ENV:=dev
GROUP_NAME:=viebiz
PROJECT_NAME:=redis

DOCKER_BUILD_BIN := docker
COMPOSE_BIN := ENV=$(ENV) GROUP_NAME=$(GROUP_NAME) PROJECT_NAME=$(PROJECT_NAME) docker compose
COMPOSE_TOOL_RUN := $(COMPOSE_BIN) run --rm --service-ports tool

gen-mocks:
	@$(COMPOSE_TOOL_RUN) sh -c "mockery"
