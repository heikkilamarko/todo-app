#!/bin/bash

ENV_DIR=$1

for f in $ENV_DIR/*; do
    sops -e -i $f
done;
