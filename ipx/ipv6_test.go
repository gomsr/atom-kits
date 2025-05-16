package ipx

import (
	"fmt"
	"testing"
)

func TestIpv6(t *testing.T) {
	ip := "2a09:bac2:a919:8c::e:2d2"
	fmt.Println(ExpandIPv6(ip))

	ipv4 := "152.42.179.218"
	fmt.Println(ExpandIPv6(ipv4))
}
