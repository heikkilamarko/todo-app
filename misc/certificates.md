# TLS Certificates

## Certbot

[Docs](https://certbot.eff.org/docs/install.html)

```bash
sudo docker run -it --rm --name certbot \
  -v "/etc/letsencrypt:/etc/letsencrypt" \
  -v "/var/lib/letsencrypt:/var/lib/letsencrypt" \
  -p 80:80 \
  certbot/certbot certonly
```

## Kestrel

[Docs](https://docs.microsoft.com/en-us/aspnet/core/fundamentals/servers/kestrel/endpoints)

### `appsettings.json`

```json
{
  "Kestrel": {
    "Endpoints": {
      "Https": {
        "Url": "https://+:443",
        "Certificate": {
          "Path": "<path to .pem/.crt file>",
          "KeyPath": "<path to .key file>",
          "Password": "<certificate password>"
        }
      }
    }
  }
}
```

## Keycloak

[Docs](https://hub.docker.com/r/jboss/keycloak/)

Section: **Setting up TLS(SSL)**

## Bypassing Certificate Check

⚠️ You shouldn't do this in production environments.

### .NET

```csharp
// using System.Net.Http;

services
  .AddHttpClient<MyHttpClient>()
  .ConfigurePrimaryHttpMessageHandler(
      () => new HttpClientHandler
      {
          ServerCertificateCustomValidationCallback = (_, _, _, _) => true
      }
  )
```

```csharp
// using System.Net.Http;

services.AddAuthentication(JwtBearerDefaults.AuthenticationScheme).AddJwtBearer(o =>
{
    // ...
    o.BackchannelHttpHandler = new HttpClientHandler
    {
        ServerCertificateCustomValidationCallback = (_, _, _, _) => true
    };
    // ...
}
```
