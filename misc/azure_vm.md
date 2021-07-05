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

## Configure SSH

```bash
> chmod 700 /path/to/ssh-key.pem
```

```bash
# ~/.ssh/config

Host todo-app
  HostName VM_DNS_NAME
  Port 22
  User azureuser
  IdentityFile /path/to/ssh-key.pem
  IdentitiesOnly yes
  ControlMaster     auto
  ControlPath       ~/.ssh/control-%C
  ControlPersist    10m
```

## Connect to VM

```bash
> ssh todo-app
```

### Update Packages

```bash
> sudo apt update
> sudo apt upgrade
```

### Install Docker

1. [Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
2. [Post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/)
3. [Install Docker Compose](https://docs.docker.com/compose/install/)

### Prepare Certificates

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

### Exit from VM

```bash
> exit
```

## Create and Populate `env` Directory

```bash
> mkdir ~/todo-app/env
> ~/todo-app/secrets/decrypt_env.sh ~/todo-app/secrets/env.prod.enc ~/todo-app/env
```

## Create and Set Docker Context

```bash
> docker context create todo-app --docker "host=ssh://todo-app"
> docker context use todo-app
```

## Run

```bash
> cd ~/todo-app
> docker-compose -f docker-compose.yml -f docker-compose.prod.yml up --build -d
```
