#!/bin/bash

/bin/sh -c
/usr/bin/mc config host rm local;
/usr/bin/mc config host add --quiet --api s3v4 local $MINIO_URL $MINIO_ACCESS_KEY $MINIO_SECRET_KEY;
/usr/bin/mc mb --insecure --ignore-existing --quiet local/loki/;
