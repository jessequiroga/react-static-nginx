#!/bin/sh
#
# Build docker in docker image with node and yarn to speed up build times
#
docker build -t registry.gitlab.com/stackworx-public/react-static-nginx/dnd_nodejs:latest -f Dockerfile.build .
docker push registry.gitlab.com/stackworx-public/react-static-nginx/dnd_nodejs:latest