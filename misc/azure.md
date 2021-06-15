# Host Todo App in Azure VM

## 1. Create a Linux VM

[Quickstart: Create a Linux virtual machine in the Azure portal](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/quick-create-portal)

- Linux (ubuntu 20.04)
- Standard E2s v3 (2 vcpus, 16 GiB memory)

Download the SSH key file (`ssh-key.pem`) when asked.

## 2. Setup VM Networking

Add inbound port rule:

```text
todo-app | 8000,8002,8080,3000,9000 | TCP | Any | Any | Allow
```

## 3. Connect to VM

```bash
> chmod 700 /path/to/ssh-key.pem
> ssh -i /path/to/ssh-key.pem azureuser@<VM_PUBLIC_IP>
```

## 4. Update VM

```bash
> sudo apt update
> sudo apt upgrade
```

## 5. Install Docker

1. [Install Docker Engine on Ubuntu](https://docs.docker.com/engine/install/ubuntu/)
2. [Post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/)
3. [Install Docker Compose](https://docs.docker.com/compose/install/)

## 6. Install Mozilla SOPS

```bash
> sudo wget https://github.com/mozilla/sops/releases/download/<VERSION>/sops-<VERSION>.linux -O /usr/local/bin/sops
> sudo chmod +x /usr/local/bin/sops
```

In the above command, `<VERSION>` is `v3.7.1` or higher.

## 7. Clone Todo App Repository

```bash
> cd ~
> git clone https://github.com/heikkilamarko/todo-app.git
```

## 8. Preparations before Running the App

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
> vim ~/todo-app/backend/keycloak/docker/todo-app.json
# [EDIT] Replace localhost with VM_PUBLIC_IP
# [EDIT] Set "sslRequired": "none"

> vim ~/todo-app/env/todo-app.env
# [EDIT] Replace localhost with VM_PUBLIC_IP

> vim ~/todo-app/env/api-gateway.env
# [EDIT] Replace localhost with VM_PUBLIC_IP
```

## 9. Run the App

```bash
> cd ~/todo-app
> docker-compose up --build -d
```
