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
