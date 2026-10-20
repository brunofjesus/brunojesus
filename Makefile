TMP_DIR := tmp
BIN     := $(TMP_DIR)/main

.PHONY: dev generate build run fmt tidy test clean

## run-watch: run with live reload 
run-watch:
	go tool wgo run ./...

## temp-watch: regenerate templ files on change
templ-watch:
	go tool templ generate --watch

## dev: run with live reload (regenerates templ, rebuilds and restarts on change)
dev:
	make -j2 templ-watch run-watch

## generate: regenerate Go code from .templ files
generate:
	go tool templ generate

## build: compile the binary into tmp/main
build:
	@mkdir -p $(TMP_DIR)
	go build -o $(BIN) .

## run: run the server once, without watching
run:
	go generate
	go run .

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
