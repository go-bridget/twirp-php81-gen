.PHONY: all build codegen lint test clean

export CGO_ENABLED := 0
export PATH := $(PWD)/build:$(PATH)

all:
	@drone exec --trusted

build:
	go build -o build/ ./cmd/...

codegen:
	@drone exec --trusted --include=codegen

lint:
	@drone exec --trusted --include=lint

test:
	go test -v ./...

clean:
	go fmt ./...
	go mod tidy
