# Host Todo App in Azure VM

## Create VM

[Quickstart: Create a Linux virtual machine in the Azure portal](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-portal)

Download the SSH key file (`ssh-key.pem`) when asked.

## Setup Networking

`Allow`:

- HTTP(S): `443`, `8443`, `9443`
- SSH: `22`

## Setup DNS Name

```text
Public IP address > Configuration > DNS name label (VM_DNS_NAME) > Save
```

Example: `todo-app.westeurope.cloudapp.azure.com`

## Configure SSH

```bash
> chmod 400 /path/to/ssh-key.pem
```

```bash
# ~/.ssh/config

Host todo-app
  HostName VM_DNS_NAME
  Port 22
  User azureuser
  IdentityFile /path/to/ssh-key.pem
  IdentitiesOnly yes
  ControlMaster auto
  ControlPath ~/.ssh/control-%C
  ControlPersist 10m
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

### Initialize Docker Swarm

```bash
> docker swarm init
```

### Generate Certificates

1. Generate certificates (see [here](certificates.md))

   - Domain name: `VM_DNS_NAME`

2. Copy, rename, and [encrypt](../secrets/) the certificates from VM into `/secrets/prod` directory.
   ```bash
   /secrets/prod/
   ├── ...
   ├── private.key # encrypted
   └── public.crt  # encrypted
   ```

### Exit from VM

```bash
> exit
```

## Create and Populate `env` Directory

```bash
# In repository root directory

> mkdir env
> secrets/decrypt_env.sh secrets/prod env
```

## Create and Set Docker Context

```bash
> docker context create todo-app --docker "host=ssh://todo-app"
> docker context use todo-app
```

## Run

```bash
# In repository root directory

> docker compose -f docker-compose.yml -f docker-compose.prod.yml build
> docker stack deploy -c docker-compose.yml -c docker-compose.prod.yml todo-app
```
