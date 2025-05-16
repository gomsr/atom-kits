package ipx

import (
	"net"
	"strings"
)

// IsValidIP checks if the provided string is a valid IPv4 address.
func IsValidIP(ip string) bool {

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	return !parsedIP.IsPrivate() && !strings.EqualFold("::1", ip)
}
