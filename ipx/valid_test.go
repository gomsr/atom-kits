package ipx

import "testing"

func TestIsValidIPv4(t *testing.T) {

	println(IsValidIP("172.16.0.0"))
	println(IsValidIP("38.207.137.78"))
	println(IsValidIP("10.0.0.1"))
	println(IsValidIP("192.168.1.1"))
	println(IsValidIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	println(IsValidIP("::1"))
}
