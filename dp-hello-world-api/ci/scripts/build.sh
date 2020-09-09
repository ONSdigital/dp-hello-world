#!/bin/bash -eux

pushd dp-hello-world-api
  make build
  cp build/dp-hello-world-api Dockerfile.concourse ../build
popd
