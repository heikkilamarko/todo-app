#!/bin/bash
set -euo pipefail

cd "$(dirname "$0")"

export VM_ADMIN_USERNAME=$(terraform -chdir=infra output -raw vm_admin_username)

envsubst < vm_install_script.sh | ssh todo-app 'sudo bash -s' -
