#!/bin/bash -eux

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59.1

pushd dp-hello-world-controller
  make lint
popd
