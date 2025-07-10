.PHONY: build clean test run

BIN_DIR = ./bin
MAIN_PKG = ./cmd/arbbot
BIN_NAME = arbbot

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) $(MAIN_PKG)

clean:
	rm -rf $(BIN_DIR)

test:
	go test -v ./...

run:
	go run $(MAIN_PKG)

deps:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

all: clean build test