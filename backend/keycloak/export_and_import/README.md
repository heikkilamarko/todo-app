# Export Realm

```bash
# Run the following command inside the keycloak container
/opt/jboss/keycloak/bin/standalone.sh \
-Djboss.socket.binding.port-offset=100 \
-Dkeycloak.migration.action=export \
-Dkeycloak.migration.provider=singleFile \
-Dkeycloak.migration.realmName=todo-app \
-Dkeycloak.migration.usersExportStrategy=REALM_FILE \
-Dkeycloak.migration.file=/tmp/todo-app.json
```

# Import Realm

```bash
> ./import_realm.sh ../docker/todo-app.json
```
