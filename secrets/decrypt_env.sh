#!/bin/bash

INPUT_DIR=$1
OUTPUT_DIR=$2

for f in $INPUT_DIR/*; do
    sops -d $f > $OUTPUT_DIR/$(basename $f)
done;
