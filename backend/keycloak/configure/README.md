# Configuring Keycloak

[Keycloak Admin REST API](https://www.keycloak.org/docs-api/17.0/rest-api)

## Export realm

```bash
> ./export_realm.sh todo-app.json
```

## Import realm

```bash
> ./import_realm.sh realms/todo-app.json
```

## Import users

```bash
> ./import_user.sh users/demouser.json
> ./import_user.sh users/demoviewer.json
```
