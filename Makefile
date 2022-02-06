.PHONY: all build test clean

PATH := $(PWD)/build:$(PATH)

all:
	@drone exec

build:
	go build -o build/ ./cmd/...

test:
	GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/...@v1.0.0-rc12
	buf --version
	cd example && buf mod update
	buf generate --template example/buf.gen.yaml --path example
	buf generate --template example/buf.gen.yaml --path example/src/upload.proto
	buf generate --template example/buf.gen.yaml --path example/src/google_timestamp.proto

clean:
	go fmt ./...
	go mod download
	go mod tidy
