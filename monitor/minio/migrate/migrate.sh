#!/bin/bash

/bin/sh -c
/usr/bin/mc config host rm local;
/usr/bin/mc config host add --quiet --api s3v4 local $MINIO_URL $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD;
/usr/bin/mc mb --insecure --ignore-existing --quiet local/loki/;
