package ipx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const ipInfoKeyUrl = "https://ipinfo.io/%s?token=%s"

func QueryByIPInfo(ip string, key string) (*IPInfo, error) {
	//ip = ExpandIPv6(ip)
	url := fmt.Sprintf(ipInfoKeyUrl, ip, key)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to make request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return nil, err
	}

	ipInfo := &IPInfoS{}
	if err := json.Unmarshal(body, ipInfo); err != nil {
		return nil, err
	} else {
		lat, lng := tryParseLoc(ipInfo.Loc)
		return &IPInfo{
			IP:          ipInfo.IP,
			CountryCode: ipInfo.Country,
			CountryName: ipInfo.Country,
			RegionName:  ipInfo.Region,
			CityName:    ipInfo.City,
			Latitude:    lat,
			Longitude:   lng,
			ZipCode:     ipInfo.Postal,
			TimeZone:    ipInfo.Timezone,
			Asn:         ipInfo.Org,
		}, nil
	}
}

func tryParseLoc(loc string) (float64, float64) {
	// Split the string by the comma
	parts := strings.Split(loc, ",")
	if len(parts) != 2 {
		return 0, 0
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, 0
	}

	lng, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, 0
	}

	return lat, lng
}

type IPInfoS struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}
