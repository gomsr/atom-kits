package ipx

import "strings"

const (
	IPLocation = "IPLocation"
	IPInfos    = "IPInfos"
)

func Query(ip string, key string, vendors ...string) (*IPInfo, error) {
	var vendor string
	if len(vendors) == 0 {
		vendor = IPInfos
	} else {
		vendor = vendors[0]
	}

	if strings.EqualFold(vendor, IPLocation) {
		return QueryWithKey(ip, key)
	}

	return QueryByIPInfo(ip, key)
}
