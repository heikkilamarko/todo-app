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

## Tips

With SOPS, editing will happen in whatever `$EDITOR` is set to, or, if it's not set, in `vim`. SOPS will wait for the editor to exit, and then try to reencrypt the file. Here is how to configure SOPS to use `Visual Studio Code` as the editor:

```bash
# Add this to your ~/.zshrc or ~/.bashrc
export EDITOR="code --wait"
```
