#!/bin/bash -eux

pushd dp-hello-world-library-js
  npm ci --silent
  make build

  # copy build to the location expected by the CI
  cp -r build package.json package-lock.json ../build
popd
