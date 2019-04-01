# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	go build -o build/slasher cmd/main.go
    
.PHONY: test
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -race -cover -v ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: generate_mocks
generate_mocks:
	@go get github.com/golang/mock/gomock
	@go install github.com/golang/mock/mockgen
	mockgen -package slasher -destination ./mock_client.go github.com/taeho-io/idl/gen/go/slasher SlasherClient

.PHONY: clean_mocks
clean_mocks:
	find . -name "mock_*.go" -type f -delete
	rm -rf mocks
