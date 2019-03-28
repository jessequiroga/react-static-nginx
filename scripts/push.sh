#!/bin/sh
set -e

REVISION=$1

source ./scripts/env.sh

docker push registry.gitlab.com/stackworx-public/react-static-nginx:${NGINX_VERSION}-${REVISION}