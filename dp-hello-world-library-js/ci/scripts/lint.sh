#!/bin/bash -eux

pushd dp-hello-world-library-js
  npm ci --silent
  make lint
popd
