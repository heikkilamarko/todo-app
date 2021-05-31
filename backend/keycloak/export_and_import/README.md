# Export and Import

[Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/index.html#_export_import)

## Export

To export the `todo-app` realm into a JSON file:

1. Run the following command inside the `keycloak` container.

```bash
> /opt/jboss/keycloak/bin/standalone.sh \
-Djboss.socket.binding.port-offset=100 \
-Dkeycloak.migration.action=export \
-Dkeycloak.migration.provider=singleFile \
-Dkeycloak.migration.realmName=todo-app \
-Dkeycloak.migration.usersExportStrategy=REALM_FILE \
-Dkeycloak.migration.file=/tmp/todo-app.json
```

2. The exported realm file: `/tmp/todo-app.json`

## Import

To import the `todo-app` realm from the JSON file, run the following command:

```bash
> ./import_realm.sh ../docker/todo-app.json
```
