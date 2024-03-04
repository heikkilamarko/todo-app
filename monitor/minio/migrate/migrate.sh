#!/bin/sh
set -e

# fail fast
# curl -s $MINIO_URL > /dev/null

set +o history

mc alias set minio $MINIO_URL $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD

if ! mc admin user info minio "$MINIO_CONSOLE_USER" | grep -q consoleAdmin; then
    mc mb --ignore-existing minio/loki
    mc admin user add minio $MINIO_CONSOLE_USER $MINIO_CONSOLE_PASSWORD
    mc admin policy create minio consoleAdmin consoleAdmin.json
    mc admin policy attach minio consoleAdmin --user $MINIO_CONSOLE_USER
fi

set -o history
