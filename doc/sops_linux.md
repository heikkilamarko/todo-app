# Install Mozilla SOPS on Linux

## Install

```bash
> sudo wget https://github.com/mozilla/sops/releases/download/VERSION/sops-VERSION.linux -O /usr/local/bin/sops
> sudo chmod +x /usr/local/bin/sops
```

Substitute `VERSION` with the version of SOPS you want to use. For example: `v3.7.1`

## Configure age

Copy the age key file (`keys.txt`) into `~/.config/sops/age` directory.
