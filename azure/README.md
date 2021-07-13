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
  HostName <domain_name_label>.westeurope.cloudapp.azure.com
  Port 22
  User azureuser
  IdentityFile <keyfile>
  IdentitiesOnly yes
  ControlMaster auto
  ControlPath ~/.ssh/control-%C
  ControlPersist 10m
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

> ./generate_certs.sh <domain_name_label>.westeurope.cloudapp.azure.com
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
