#!/bin/bash -eux

export cwd=$(pwd)

pushd $cwd/dp-hello-world-event
  make audit
popd 