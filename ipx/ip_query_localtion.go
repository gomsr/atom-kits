package ipx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const freeUrl = "https://api.ip2location.io?ip=%s"
const keyUrl = "https://api.ip2location.io?key=%s&ip=%s"

func QueryFree(ip string) (*IPInfo, error) {
	url := fmt.Sprintf(freeUrl, ip)
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

	ipInfo := &IPInfo{}
	if err := json.Unmarshal(body, ipInfo); err != nil {
		return nil, err
	} else {
		return ipInfo, nil
	}
}

func QueryWithKey(ip string, key string) (*IPInfo, error) {
	url := fmt.Sprintf(keyUrl, key, ip)
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

	ipInfo := &IPInfo{}
	if err := json.Unmarshal(body, ipInfo); err != nil {
		return nil, err
	} else {
		return ipInfo, nil
	}
}
