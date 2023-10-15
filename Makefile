#Test

test:
	@go test -race -cover ./...
.PHONY: test

#Lint

lint:
	@golangci-lint run --config=.golangci.yml ./...
.PHONY: lint

#Build

build:
	@goreleaser release --rm-dist --snapshot
.PHONY: build

#Dev

dev-up:
	@docker-compose up -d
.PHONY: dev-up

dev-down:
	@docker-compose down
.PHONY: dev-down
