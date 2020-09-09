#!/bin/bash -eux

pushd dp-hello-world-contoller
  make build
  cp build/dp-hello-world-contoller Dockerfile.concourse ../build
popd
