.PHONY: all build test clean google

PATH := $(PWD)/build:$(PATH)

all:
	@drone exec

build:
	go build -o build/ ./cmd/...

test:
	GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/...@v1.0.0-rc12
	buf --version
	buf mod update
	buf generate --template buf.gen.yaml --path example

clean:
	go fmt ./...
	go mod download
	go mod tidy

google:
	git clone https://github.com/googleapis/googleapis google
