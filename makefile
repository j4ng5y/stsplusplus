all: test build

.PHONY: test
test:
	@go test -coverprofile coverage.out ./... -race

.PHONY: build
build:
	@go build -o bin/sts++ cmd/sts/main.go