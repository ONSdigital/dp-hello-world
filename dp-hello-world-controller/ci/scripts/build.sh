#!/bin/bash -eux

pushd dp-hello-world-controller
  make build
  cp build/dp-hello-world-controller Dockerfile.concourse ../build
popd
