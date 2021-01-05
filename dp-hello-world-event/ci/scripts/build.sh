#!/bin/bash -eux

pushd dp-hello-world-event
  make build
  cp build/dp-hello-world-event Dockerfile.concourse ../build
popd
