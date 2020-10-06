---
platform: linux

image_resource:
  type: docker-image
  source:
    repository: onsdigital/dp-concourse-tools-nancy
    tag: latest

inputs:
  - name: dp-hello-world-controller
    path: dp-hello-world-controller

run:
  path: dp-hello-world-controller/ci/scripts/audit.sh 