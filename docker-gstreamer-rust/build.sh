#!/usr/bin/env bash
set -euo pipefail

UBUNTU_VERSION=20.04
RUST_VERSION=1.50.0
IMAGE_NAME=gstreamer-rust
IMAGE_TAG="${UBUNTU_VERSION}-${RUST_VERSION}"
IMAGE_NAME_FULL="${IMAGE_NAME}:${IMAGE_TAG}"
GITHUB_IMAGE_NAME_FULL="docker.pkg.github.com/$GITHUB_USERNAME/platune/$IMAGE_NAME_FULL"

docker build . -t $IMAGE_NAME_FULL \
--build-arg UBUNTU_VERSION=${UBUNTU_VERSION} --build-arg RUST_VERSION=${RUST_VERSION}

docker tag $IMAGE_NAME_FULL $GITHUB_IMAGE_NAME_FULL
echo "Tagged $GITHUB_IMAGE_NAME_FULL"
[[ $# -eq 1 && "$1" == "-p" ]] && docker push $GITHUB_IMAGE_NAME_FULL
