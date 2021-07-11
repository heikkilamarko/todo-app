#!/bin/sh

set +o history

mc alias set minio $MINIO_URL $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb --ignore-existing minio/loki

mc admin user add minio $MINIO_CONSOLE_USER $MINIO_CONSOLE_PASSWORD
mc admin policy add minio consoleAdmin consoleAdmin.json
mc admin policy set minio consoleAdmin user=$MINIO_CONSOLE_USER

set -o history
