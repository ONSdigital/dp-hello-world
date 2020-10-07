---
platform: linux

image_resource:
  type: docker-image
  source:
    repository: onsdigital/dp-concourse-tools-nancy
    tag: latest

inputs:
  - name: dp-hello-world-api
    path: dp-hello-world-api

run:
  path: dp-hello-world-api/ci/scripts/audit.sh 