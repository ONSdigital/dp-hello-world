#!/bin/bash -eux

pushd dp-hello-world-controller
  make test-component
popd
