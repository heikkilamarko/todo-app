# Secrets Management

This repository uses [SOPS](https://github.com/mozilla/sops) with [age](https://github.com/mozilla/sops#encrypting-using-age) for managing secrets.

[age docs](https://age-encryption.org/)

⚠️ `keys.txt` is included in this repository just for demo purposes.

## Decrypting Secrets

```bash
# Create the output directory
> mkdir ../env
# Decrypt the secrets
> ./decrypt_env.sh env.enc ../env
```
