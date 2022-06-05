# Hosting Todo App in Azure VM

## Configure SSH

```bash
# Generate SSH key pair

> ssh-keygen -t rsa -b 4096 -f <keyfile>

# Outputs:

- <keyfile>        # private key file
- <keyfile>.pub    # public key file
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
> docker context create todo-app --docker "host=ssh://todo-app"
> docker context use todo-app
```

## Run Terraform

```bash
# In 'infra' directory

# Before running the below commands, set
# values ​​to variables in 'terraform.tfvars'

> terraform init
> terraform apply

# Test SSH connection
> ssh todo-app
```

## Generate Certificates

```bash
# In 'azure' directory

> ./generate_certs.sh <domain> <email_for_account_notifications>
```

[Certbot docs](https://certbot.eff.org/docs/install.html)

## Decrypt Secrets

```bash
# In repository root directory

> mkdir env
> config/decrypt_secrets.sh config/prod env
```

## Deploy

```bash
# In repository root directory

> docker compose -f docker-compose.yml -f docker-compose.prod.yml build
> docker stack deploy -c docker-compose.yml -c docker-compose.prod.yml todo-app
```

## Configure Keycloak

Import the Keycloak realm. See [Configuring Keycloak](../backend/keycloak/configure/) for details.
