#!/bin/sh
set -e

mc alias set minio $MINIO_URL $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb --ignore-existing minio/loki
