#!/bin/sh
set -e

. ./scripts/env.sh

export REVISION=${CI_COMMIT_SHA:-dev}

./scripts/build.sh $REVISION

IMAGE=registry.gitlab.com/stackworx-public/react-static-nginx:${NGINX_VERSION}-${REVISION}

# Build test image
docker build --file Dockerfile.testing --build-arg IMAGE=$IMAGE -t testing .

DOCKER_RUN_NAME=react-static-nginx-testing-${REVISION}

# Start Container, delete if already running
docker rm -f $DOCKER_RUN_NAME || true
docker run --name $DOCKER_RUN_NAME --detach --publish 8080:80 --rm testing

docker ps

set +e
yarn install --frozen-lockfile --silent
yarn test
TEST_RESULT=$?
set -e

docker rm -f $DOCKER_RUN_NAME

if [[ "$TEST_RESULT" -ne 0 ]]; then
  exit "$TEST_RESULT"
fi