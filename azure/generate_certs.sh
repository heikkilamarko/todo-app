#!/bin/bash

docker context use todo-app

docker run -it --rm --name certbot \
  -v "/etc/letsencrypt:/etc/letsencrypt" \
  -v "/var/lib/letsencrypt:/var/lib/letsencrypt" \
  -p 80:80 \
  certbot/certbot certonly \
  --manual \
  --preferred-challenges=dns \
  --agree-tos \
  --no-eff-email \
  --cert-name todo-app \
  -d $1 \
  -m $2

ssh todo-app 'sudo cp -Lr /etc/letsencrypt/live/todo-app/ ~/certs && sudo chown -R azureuser ~/certs'

scp todo-app:~/certs/privkey.pem ../secrets/prod/private.key
scp todo-app:~/certs/fullchain.pem ../secrets/prod/public.crt

ssh todo-app 'sudo rm -rf ~/certs'

pushd ../secrets > /dev/null

sops -e -i prod/private.key
sops -e -i prod/public.crt

popd > /dev/null
