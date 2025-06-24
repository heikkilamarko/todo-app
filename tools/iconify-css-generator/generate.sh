#!/bin/bash
set -euo pipefail

cd "$(dirname "$0")"

ICONS=$(jq -r '.bi | join(",")' icons.json)
OUTPUT_DIR="../../frontend/src/styles"

curl https://api.iconify.design/bi.css?icons=$ICONS -o "$OUTPUT_DIR/_icons.scss"

# TODO: Temporary workaround for https://github.com/twbs/icons/issues/913
cat <<EOF >> "$OUTPUT_DIR/_icons.scss"

.icon--bi {
    vertical-align: -0.125em;
}
EOF
