package emailx

import "strings"

// ── Provider config ───────────────────────────────────────────

const (
	GmailSuffix = "@gmail.com"
	GmailType   = "google"

	NetSuffix = "@163.com"
	NetType   = "163"

	QqSuffix = "@qq.com"
	QqType   = "qq"

	ICloudSuffix = "@icloud.com"
	ICloudType   = "icloud"
)

// ZohoFunc returns Zoho SMTP server configuration.
// Use with WithServerFunc: emailx.Send(cfg, emailx.WithServerFunc(emailx.ZohoFunc))
var ZohoFunc = func() (string, int, bool) { return "smtp.zoho.com", 587, false }
var ICloudFunc = func() (string, int, bool) { return "smtp.mail.me.com", 587, false }
var QqFunc = func() (string, int, bool) { return "smtp.qq.com", 587, false }
var NetFunc = func() (string, int, bool) { return "smtp.163.com", 465, true }
var GmailFunc = func() (string, int, bool) { return "smtp.gmail.com", 587, false }

// resolveServer determines the SMTP server to use.
// Precedence: ServerFunc > auto-detect from suffix.
func resolveServer(from string, defVal ...func() (string, int, bool)) (host string, port int, isSSL bool) {

	switch {
	case strings.HasSuffix(from, GmailSuffix):
		return GmailFunc()
	case strings.HasSuffix(from, NetSuffix):
		return NetFunc()
	case strings.HasSuffix(from, QqSuffix):
		return QqFunc()
	case strings.HasSuffix(from, ICloudSuffix):
		return ICloudFunc()
	}
	if len(defVal) > 0 {
		return defVal[0]()
	}

	return "", 0, false
}
