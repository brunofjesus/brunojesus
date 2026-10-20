TMP_DIR := tmp
BIN     := $(TMP_DIR)/main

.PHONY: dev generate build run fmt tidy test clean

## generate: regenerate Go code from .templ files
generate:
	go generate .
	go tool templ generate

## build: compile the binary into tmp/main
build:
	@mkdir -p $(TMP_DIR)
	go generate .
	go build -o $(BIN) ./cmd/server

## run: run the server once, without watching
run:
	go generate .
	go run ./cmd/server

## fmt: format templ files and Go code
fmt:
	go tool templ fmt .
	go fmt ./...

## tidy: sync go.mod / go.sum
tidy:
	go mod tidy

## test: run the test suite
test:
	go test ./...

## clean: remove build artifacts
clean:
	rm -rf $(TMP_DIR)
