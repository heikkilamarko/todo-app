#!/bin/bash

/bin/sh -c
/usr/bin/mc config host rm local;
/usr/bin/mc config host add --quiet --api s3v4 local http://minio:9000 $MINIO_ACCESS_KEY $MINIO_SECRET_KEY;
/usr/bin/mc mb --ignore-existing --quiet local/loki/;
