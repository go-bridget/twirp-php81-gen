FROM golang:1.17-alpine

RUN apk --no-cache add make && GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/...@v1.0.0

RUN buf --version
