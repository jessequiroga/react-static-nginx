#!/bin/bash
set -e
set -x

REVISION=$1

source ./scripts/env.sh

docker build --build-arg NGINX_IMAGE=nginx:$NGINX_VERSION-alpine \
  -t registry.gitlab.com/stackworx-public/react-static-nginx:${NGINX_VERSION}-${REVISION} .
