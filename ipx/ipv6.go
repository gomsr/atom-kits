package ipx

import (
	"fmt"
	"net"
	"strings"
)

func ExpandIPv6Ptr(ip *string) string {
	if ip == nil {
		return ""
	}

	return ExpandIPv6(*ip)
}

// ExpandIPv6 from 2a09:bac2:a919:8c::e:2d2 to 2a09:bac2:a919:008c:0000:0000:000e:02d2
func ExpandIPv6(ip string) string {

	// 1. ip is ipv4
	if !IsIPv6(ip) {
		return ip
	}

	// 2. ip is full ipv6
	if len(ip) == 39 {
		return ip
	}

	// 判断是否为IPv6
	ipv6 := net.ParseIP(ip).To16()
	if ipv6 == nil || ipv6.To4() != nil {
		return "Not a valid IPv6 address"
	}

	// 按 2 字节（16 位）分隔，并转换为 4 位十六进制数
	parts := make([]string, 8)
	for i := 0; i < len(ipv6); i += 2 {
		parts[i/2] = fmt.Sprintf("%04x", int(ipv6[i])<<8|int(ipv6[i+1]))
	}

	// 使用冒号连接分片并返回结果
	return strings.Join(parts, ":")
}

func IsIPv6(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// 检查是否为 IPv4，如果不是则为 IPv6
	return parsedIP.To4() == nil
}
