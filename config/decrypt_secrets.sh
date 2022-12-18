#!/bin/bash

INPUT_DIR=$1
OUTPUT_DIR=$2

rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

for f in $INPUT_DIR/*; do
    sops -d $f > $OUTPUT_DIR/$(basename $f)
done;
