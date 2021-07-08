#!/bin/bash

ENV_DIR=$1

for f in $ENV_DIR/*; do
    sops updatekeys -y $f
done;
