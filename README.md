DNS Providers for Caddy (deprecated)
====================================

This one module gives Caddy the ability to solve the ACME DNS challenge with over 75 DNS providers.


## ⚠️ This module is deprecated

This module wraps DNS providers that are implemented by [go-acme/lego](https://github.com/go-acme/lego) which uses an old API that is no longer supported by Caddy. As such, this module is a temporary shim until a sufficient number of providers are ported to the [new `libdns` interfaces](https://github.com/libdns/libdns).

You can use this module to get up and running quickly with your provider of choice, but instead of using this module long-term, please consider [contributing to a libdns package](https://github.com/libdns/libdns/wiki/Implementing-providers) for your provider instead.

The `libdns` implementations offer better performance, lighter dependencies, easier maintainability with growth, and more flexible configuration.


## Instructions

1. Make sure you either [download](https://caddyserver.com/download?package=github.com%2Fcaddy-dns%2Flego-deprecated) or [build](https://github.com/caddyserver/caddy/#with-version-information-andor-plugins) (i.e. with [xcaddy](https://github.com/caddyserver/xcaddy)) Caddy with this module installed:
   - Download: <https://caddyserver.com/download?package=github.com%2Fcaddy-dns%2Flego-deprecated>
   - Build:
     ```bash
     xcaddy build --with github.com/caddy-dns/lego-deprecated
     ```
2. Find your **DNS Provider** and **provider code**, in the [lego DNS documentation](https://go-acme.github.io/lego/dns/)
3. You will need to set the lego provider's **credentials** and **other ENVs**, such as `CLOUDFLARE_API_KEY=xxxxxxxx`, in your environment configuration
4. Configure the ACME issuer in your [Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) or `Caddyfile` like so:
   ```json
   {
   	"module": "acme",
   	"challenges": {
   		"dns": {
   			"provider": {
   				"name": "lego_deprecated",
   				"provider_name": "<provider_code>"
   			}
   		}
   	}
   }
   ```
   or
   ```caddy
   tls {
   	dns lego_deprecated <provider_code>
   }
   ```
5. (don't forget to replace `<provider_code>` with the name of [your provider](https://go-acme.github.io/lego/dns/), such as `cloudflare` or `dnsimple`)

## Compatibility note

Unlike other modules in the caddy-dns repositories, this one can _only_ be used in the ACME issuer module for solving the DNS challenge. Even though it shares the more general `dns.providers` namespace with other provider modules, using this module in any other place in your config will result in errors.
