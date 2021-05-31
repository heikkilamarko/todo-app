#!/bin/bash

INPUT_DIR=$1
OUTPUT_DIR=$2

for f in $INPUT_DIR/*.env; do
    sops -d $f > $OUTPUT_DIR/$(basename $f)
done;
