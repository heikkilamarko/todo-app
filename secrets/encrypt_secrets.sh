#!/bin/bash

SECRETS_DIR=$1

for f in $SECRETS_DIR/*; do
    sops -e -i $f
done;
