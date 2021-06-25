# Host Todo App in Azure VM

## Create VM

[Quickstart: Create a Linux virtual machine in the Azure portal](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-portal)

Download the SSH key file (`ssh-key.pem`) when asked.

## Setup Networking

`Allow`:

- HTTP(S): `443`, `8443`, `9000`
- SSH: `22`

## Setup DNS Name

```text
Public IP address > Configuration > DNS name label (VM_DNS_NAME) > Save
```

Example: `todo-app.westeurope.cloudapp.azure.com`

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

## Prepare Certificates

1. [Generate certificates](certificates.md)
   - Domain name: `VM_DNS_NAME`
2. Copy and rename the certificates into `/etc/todo-app/certs/`

```text
/etc/todo-app/certs/
├── private.key
├── public.crt
├── tls.crt      # renamed public.crt
└── tls.key      # renamed private.key
```

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

## Configure Mozilla SOPS

Copy the `age` key file.

```bash
> mkdir -p ~/.config/sops/age
> cp ~/todo-app/secrets/keys.txt ~/.config/sops/age/
```

## Create and Populate `env` Directory

```bash
> mkdir ~/todo-app/env
> ~/todo-app/secrets/decrypt_env.sh ~/todo-app/secrets/env.prod.enc ~/todo-app/env
```

## Run

```bash
> cd ~/todo-app
> docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build -d
```
