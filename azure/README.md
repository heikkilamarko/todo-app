# Hosting Todo App in Azure VM

## Configure SSH

```bash
# Generate SSH key pair

ssh-keygen -t rsa -b 4096 -f <keyfile>

# Outputs:

<keyfile>        # private key file
<keyfile>.pub    # public key file
```

```bash
# ~/.ssh/config

Host todo-app
  HostName <azure_vm_hostname>
  Port 22
  User azureuser
  IdentityFile <keyfile>
  IdentitiesOnly yes
  ControlMaster auto
  ControlPath ~/.ssh/control-%C
  ControlPersist 10m
```

## Create and Set Docker Context

```bash
docker context create todo-app --docker "host=ssh://todo-app"
docker context use todo-app
```

## Run Terraform

```bash
# In 'infra' directory

# Before running the below commands, set
# values ​​to variables in 'terraform.tfvars'

terraform init
terraform apply

# Test SSH connection
ssh todo-app
```

## Decrypt Secrets

```bash
# In repository root directory

config/decrypt_secrets.sh config/prod env
```

## Configure TLS Settings

Set the following environment variables in the `/env/caddy.env` file:

```dotenv
CADDY_TLS_DOMAIN=
CADDY_TLS_EMAIL=
CADDY_TLS_GODADDY_TOKEN=
```

## Deploy

```bash
# In repository root directory

docker compose -f docker-compose.yml -f docker-compose.prod.yml build
docker stack deploy -c docker-compose.yml -c docker-compose.prod.yml todo-app
```

## Configure Keycloak

Import the Keycloak realm. See [Configuring Keycloak](../backend/keycloak/) for details.
