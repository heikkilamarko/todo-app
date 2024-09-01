# Configuration and Secrets Management

This repository uses [SOPS](https://getsops.io/) with [age](https://github.com/mozilla/sops#22encrypting-using-age) for managing secrets.

[age docs](https://age-encryption.org/)

⚠️ `keys.txt` (contains one or more secret keys) is included in this repository just for demo purposes.

## Installing and Configuring SOPS

1. Install [SOPS](https://getsops.io/).
2. Copy the `keys.txt` file to the [correct location](https://github.com/mozilla/sops#22encrypting-using-age) on your computer.

## Encrypting Secrets

```bash
# Encrypt a file
sops -e <file> > <encrypted_file>

# Encrypt a file (in-place)
sops -e -i <file>

# Encrypt all files (in-place) in a directory
./encrypt_secrets.sh <directory>
```

## Decrypting Todo App Secrets

Decrypt secrets into the `../env` directory.

```bash
# dev environment
./decrypt_secrets.sh dev ../env

# prod environment
./decrypt_secrets.sh prod ../env
```

## Tips

### Configuring SOPS to Use Visual Studio Code

With SOPS, editing will happen in whatever `$EDITOR` is set to, or, if it's not set, in `vim`. SOPS will wait for the editor to exit, and then try to reencrypt the file. Here is how to configure SOPS to use `Visual Studio Code` as the editor:

```bash
# Add this to your ~/.zshrc or ~/.bashrc
export EDITOR="code --wait"
```

### Installing SOPS on Linux

```bash
sudo wget https://github.com/mozilla/sops/releases/download/VERSION/sops-VERSION.linux -O /usr/local/bin/sops
sudo chmod +x /usr/local/bin/sops
```

Substitute `VERSION` with the version of SOPS you want to use. For example: `v3.9.0`
