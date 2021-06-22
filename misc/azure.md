# Host Todo App in Azure VM

## Create VM

[Quickstart: Create a Linux virtual machine in the Azure portal](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-portal)

Download the SSH key file (`ssh-key.pem`) when asked.

## Setup Networking

Allow:

- HTTP (80, 8080)
- SSH (22)

## Setup DNS Name

```text
Public IP address > Configuration > DNS name label (VM_DNS_NAME) > Save
```

## Connect to VM

```bash
> chmod 700 /path/to/ssh-key.pem
> ssh -i /path/to/ssh-key.pem azureuser@VM_DNS_NAME
```

## Update Packages

```bash
> sudo apt update
> sudo apt upgrade
```

## Install Docker

1. [Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
2. [Post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/)
3. [Install Docker Compose](https://docs.docker.com/compose/install/)

## Install Mozilla SOPS

```bash
> sudo wget https://github.com/mozilla/sops/releases/download/VERSION/sops-VERSION.linux -O /usr/local/bin/sops
> sudo chmod +x /usr/local/bin/sops
```

In the above command, `VERSION` is `v3.7.1` or higher.

## Clone App Repository

```bash
> cd ~
> git clone https://github.com/heikkilamarko/todo-app.git
```

## Configure App

1. Copy the `age` key file.

```bash
> mkdir -p ~/.config/sops/age
> cp ~/todo-app/secrets/keys.txt ~/.config/sops/age/
```

2. Prepare the `env` folder.

```bash
> mkdir ~/todo-app/env
> ~/todo-app/secrets/decrypt_env.sh ~/todo-app/secrets/env.enc ~/todo-app/env
```

3. Update configuration files.

```bash
> vim ~/todo-app/docker-compose.yml
# [EDIT] services/todo-app/ports: 80:80

> vim ~/todo-app/backend/keycloak/docker/todo-app.json
# [EDIT] Replace localhost with VM_DNS_NAME

> vim ~/todo-app/env/todo-app.env
# [EDIT] Replace localhost with VM_DNS_NAME

> vim ~/todo-app/env/api-gateway.env
# [EDIT] Replace localhost with VM_DNS_NAME

> vim ~/todo-app/env/grafana.env
# [EDIT] Replace localhost with VM_DNS_NAME
```

## Run

```bash
> cd ~/todo-app
> docker-compose up --build -d
```
