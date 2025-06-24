#!/bin/bash
set -euo pipefail

cd "$(dirname "$0")"

npm i -silent
node index.js -c config.yaml
