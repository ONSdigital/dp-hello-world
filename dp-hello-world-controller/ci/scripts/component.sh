#!/bin/bash -eux

export SITE_DOMAIN=localhost

pushd dp-hello-world-controller
  make test-component
popd