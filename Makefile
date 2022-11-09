.PHONY: all
all: build

.PHONY: build
build: cloud-burster

cloud-burster:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ ./cmd

.PHONY: unit
unit:
	go test -v -race -covermode=atomic -tags=unit -timeout=30s ./...

.PHONY: coverage
coverage:
	go test -v -race -covermode=atomic -tags=unit -timeout=30s -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

.PHONY: integration
integration:
	go test -v -race -covermode=atomic -tags=integration -timeout=30s ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: mocks
mocks:
	mockery --all
