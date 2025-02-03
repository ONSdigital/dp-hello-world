#!/bin/bash -eux

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4

pushd dp-hello-world-api
  make lint
popd
