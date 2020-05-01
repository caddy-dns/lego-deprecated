DNS Providers for Caddy (deprecated)
====================================

This one module gives Caddy the ability to solve the ACME DNS challenge with over 75 DNS providers.


## ⚠️ This module is deprecated

This module wraps DNS providers that are implemented by [go-acme/lego](https://github.com/go-acme/lego) which uses an old API that is no longer supported by Caddy. As such, this module is a temporary shim until a sufficient number of providers are ported to the [new `libdns` interfaces](https://github.com/libdns/libdns).

You can use this module to get up and running quickly with your provider of choice, but instead of using this module long-term, please consider [contributing to a libdns package](https://github.com/libdns/libdns/wiki/Implementing-providers) for your provider instead.

The `libdns` implementations offer better performance, lighter dependencies, easier maintainability with growth, and more flexible configuration.


## Instructions

First, [make sure Caddy is built with this module installed](https://github.com/caddyserver/caddy/#with-version-information-andor-plugins).

Then [find the documentation for your DNS provider](https://go-acme.github.io/lego/dns/).

Next, configure [the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```
{
	"module": "acme",
	"dns": {
		"provider": {
			"name": "lego_deprecated",
			"provider_name": "<provider_code>"
		}
	}
}
```

and replace `<provider_code>` with the name of your provider, as given in the docs linked above.

Your provider's credentials and other settings are configured via environment variables, which are also described in the docs linked above.

You can also use this with the Caddyfile:

```
tls {
	dns lego_deprecated <provider_code>
}
```

and replace `<provider_code>` with the name of your provider.



## Compatibility note

Unlike other modules in the caddy-dns repositories, this one can _only_ be used in the ACME issuer module for solving the DNS challenge. Even though it shares the more general `dns.providers` namespace with other provider modules, using this module in any other place in your config will result in errors.
