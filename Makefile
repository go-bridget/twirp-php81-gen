.PHONY: all build test clean

PATH := $(PWD)/build:$(PATH)

all:
	@drone exec

build:
	go build -o build/ ./cmd/...

test:
	# use go run so we do not have install buf command
	# go get -u github.com/bufbuild/buf/cmd/...@v1.0.0-rc10
	GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/...@v1.0.0-rc10
	buf generate --template example/buf.gen.yaml --path example

clean:
	go fmt ./...
	go mod download
	go mod tidy
