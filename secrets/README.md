# Secrets Management

This repository uses [SOPS](https://github.com/mozilla/sops) with [age](https://github.com/mozilla/sops#encrypting-using-age) for managing secrets.

[age docs](https://age-encryption.org/)

⚠️ `keys.txt` (contains one or more secret keys) is included in this repository just for demo purposes.

## Decrypting Secrets

1. Install [SOPS](https://github.com/mozilla/sops).
2. Copy the `keys.txt` file to the [correct location](https://github.com/mozilla/sops#22encrypting-using-age) on your computer.
3. Create `/env` directory into the root of the repository. Secrets will be decrypted to this directory.

```bash
> mkdir ../env
```

4. Decrypt the secrets by running the following command.

```bash
> ./decrypt_env.sh env.enc ../env
```

## Tips

With SOPS, editing will happen in whatever `$EDITOR` is set to, or, if it's not set, in `vim`. SOPS will wait for the editor to exit, and then try to reencrypt the file. Here is how to configure SOPS to use `Visual Studio Code` as the editor:

```bash
# Add this to your ~/.zshrc or ~/.bashrc
export EDITOR="code --wait"
```
