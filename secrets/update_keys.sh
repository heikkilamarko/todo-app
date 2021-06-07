#!/bin/bash

ENV_DIR=$1

for f in $ENV_DIR/*.env; do
    sops updatekeys -y $f
done;
