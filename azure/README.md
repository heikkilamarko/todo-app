# Hosting Todo App in Azure VM

## Generate SSH Key Pair

```bash
ssh-keygen -t rsa -b 4096 -f <keyfile>

# Outputs:

<keyfile>        # private key file
<keyfile>.pub    # public key file
```

## Create Azure Resources

Before running the below commands, set values ​​to variables in `infra/terraform.tfvars`.

```bash
terraform -chdir=infra init
```

```bash
terraform -chdir=infra apply
```

## Configure SSH

Add the following configuration to your `~/.ssh/config` file:

```bash
Host todo-app
  HostName <AZURE_VM_PUBLIC_IP>
  Port 22
  User azureuser
  IdentityFile <keyfile>
  IdentitiesOnly yes
  ControlMaster auto
  ControlPath ~/.ssh/control-%C
  ControlPersist 10m
```

## Check SSH Access to the VM

```bash
ssh todo-app
```

## Install Docker on the VM

```bash
./vm_install.sh
```

## Decrypt Secrets

To decrypt the secrets, run the following command in the repository root directory:

```bash
config/decrypt_secrets.sh config/prod env
```

## Configure TLS Settings

Set the following environment variables in the `/env/caddy.env` file:

```dotenv
CADDY_DOMAIN=
CADDY_TLS_EMAIL=
CADDY_CLOUDFLARE_API_TOKEN=
```

## Configure DNS

Create DNS A records for the domain $CADDY_DOMAIN pointing to the public IP address of the VM.

## Create and Set Docker Context

```bash
docker context create todo-app --docker "host=ssh://todo-app"
```

```bash
docker context use todo-app
```

## Deploy

In the repository root directory, run the following command to build and deploy the application:

```bash
docker compose -f docker-compose.yml -f docker-compose.prod.yml up --build -d
```

## Destroy Azure Resources

```bash
terraform -chdir=infra destroy
```
