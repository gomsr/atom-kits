package ipx

import "regexp"

const (
	Browser        = "Unknown"
	BrowserVersion = "Unknown"
	Os             = "Unknown"
	Device         = "Unknown"
)

// ParseAgent Function to extract detailed information from User-Agent
func ParseAgent(userAgent string) (browser string, browserVersion string, os string, device string) {

	// Regular expressions for different browsers, OS, and devices
	browserRegex := regexp.MustCompile(`(?i)(firefox|msie|chrome|safari|opera|edg|edge|trident)\/?([0-9.]*)`)
	osRegex := regexp.MustCompile(`(?i)(windows nt|mac os x|android|iphone os|linux)`)
	deviceRegex := regexp.MustCompile(`(?i)(mobile|tablet|ipad|ipod|windows|android|windows phone|kindle|iphone|mac|X11)`)

	// Find browser information
	browserMatches := browserRegex.FindStringSubmatch(userAgent)
	if len(browserMatches) > 1 {
		browser = browserMatches[1]
		if len(browserMatches) > 2 {
			browserVersion = browserMatches[2]
		}
	}

	// Find OS information
	osMatch := osRegex.FindString(userAgent)
	if osMatch != "" {
		os = osMatch
	}

	// Find device information
	deviceMatch := deviceRegex.FindString(userAgent)
	if deviceMatch != "" {
		device = deviceMatch
	}

	return browser, browserVersion, os, device
}
