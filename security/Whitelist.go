package security

import (
	"github.com/adrian-lorenz/nox-vault/globals"
	"github.com/adrian-lorenz/nox-vault/tools"
	"slices"
)

func CheckWhitelistsPub(ip string) bool {
	if ip == "" {
		return false
	}
	if len(globals.Whitelist) > 0 {
		if slices.Contains(globals.Whitelist, ip) {
			return true
		}
	}
	if len(globals.WhitelistDNS) > 0 {
		for _, w := range globals.WhitelistDNS {
			dnsIp, err := tools.GetDnsIP(w)
			if err != nil {
				return false
			}
			if dnsIp == ip {
				return true
			}
		}
	}
	return false
}

func CheckWhitelistsInt(ip string) bool {
	if ip == "" {
		return false
	}
	if len(globals.SystemWhitelist) > 0 {
		if slices.Contains(globals.SystemWhitelist, ip) {
			return true
		}
	}
	if len(globals.SystemWhitelistDNS) > 0 {
		for _, w := range globals.SystemWhitelistDNS {
			dnsIp, err := tools.GetDnsIP(w)
			if err != nil {
				return false
			}
			if dnsIp == ip {
				return true
			}
		}
	}
	return false
}
