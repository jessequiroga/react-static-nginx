#!/bin/sh
set -e
sed -i -e "s~__APP_SERVER_URI__~${APP_SERVER_URI}~g" /usr/share/nginx/html/index.html
sed -i -e "s~__APP_STAGE__~${APP_STAGE}~g" /usr/share/nginx/html/index.html

exec "$@"