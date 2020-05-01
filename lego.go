package legodeprecated

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/go-acme/lego/v3/challenge"
	"github.com/go-acme/lego/v3/providers/dns"
)

func init() {
	caddy.RegisterModule(LegoDeprecated{})
}

// LegoDeprecated is a shim module that allows any and all of the
// DNS providers in go-acme/lego to be used with Caddy. They must
// be configured via environment variables, they do not support
// cancellation in the case of frequent config changes.
//
// Even though this module is in the dns.providers namespace, it
// is only a special case for solving ACME challenges, intended to
// replace the modules that used to be in the now-defunct tls.dns
// namespace. Using it in other places of the Caddy config will
// result in errors.
//
// This module will eventually go away in favor of the modules that
// make use of the libdns APIs: https://github.com/libdns
type LegoDeprecated struct {
	ProviderName string `json:"provider_name,omitempty"`

	prov challenge.Provider
}

// CaddyModule returns the Caddy module information.
func (LegoDeprecated) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.lego_deprecated",
		New: func() caddy.Module { return new(LegoDeprecated) },
	}
}

// Provision initializes the underlying DNS provider.
func (ld *LegoDeprecated) Provision(ctx caddy.Context) error {
	prov, err := dns.NewDNSChallengeProviderByName(ld.ProviderName)
	if err != nil {
		return err
	}
	ld.prov = prov
	return nil
}

// Present wraps the challenge.Provider interface.
func (ld LegoDeprecated) Present(domain, token, keyAuth string) error {
	return ld.prov.Present(domain, token, keyAuth)
}

// CleanUp wraps the challenge.Provider interface.
func (ld LegoDeprecated) CleanUp(domain, token, keyAuth string) error {
	return ld.prov.CleanUp(domain, token, keyAuth)
}

// Interface guard
var _ challenge.Provider = (*LegoDeprecated)(nil)
