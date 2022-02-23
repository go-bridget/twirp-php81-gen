.PHONY: all build test clean

export CGO_ENABLED := 0
export PATH := $(PWD)/build:$(PATH)

all:
	@drone exec --trusted

build:
	go build -o build/ ./cmd/...

codegen:
	@drone exec --trusted --include=codegen

clean:
	go fmt ./...
	go mod tidy
