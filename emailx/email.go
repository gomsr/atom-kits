package emailx

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"regexp"
	"strings"

	"github.com/jordan-wright/email"
)

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 发件人  你自己要发邮件的邮箱
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称    发件人昵称 通常为自己的邮箱
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`       // 是否SSL   是否开启SSL
}

// Deprecated: DoSend
// @function: send
// @description: Email发送方法
// @param: subject string, body string
// @return: error
func SendEmail(config Email, to []string, subject string, body string) error {
	from := config.From
	nickname := config.Nickname
	secret := config.Secret
	host := config.Host
	port := config.Port
	isSSL := config.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

const (
	GmailSuffix = "@gmail.com"
	GmailType   = "google"
	GmailHost   = "smtp.gmail.com"
	GmailPort   = 587
	GmailIsSSL  = false
)

const (
	NetSuffix = "@163.com"
	NetType   = "163"
	NetHost   = "smtp.163.com"
	NetPort   = 465
	NetIsSSL  = true
)

const (
	QqSuffix = "@qq.com"
	QqType   = "qq"
	QqHost   = "smtp.qq.com"
	QqPort   = 587
	QqIsSSL  = false
)

// DoSend 发送邮件
func DoSend(to, subject, body, from, nickname, secret string) (err error) {
	if len(from) == 0 {
		return errors.New("函数配置的发件人不能为空")
	}
	if len(to) == 0 {
		return errors.New("函数配置的收件人不能为空")
	}

	// parse host, ssl, port info
	host, port, isSSL := "", 0, false
	if strings.HasSuffix(from, GmailSuffix) {
		host = GmailHost
		port = GmailPort
		isSSL = GmailIsSSL
	} else if strings.HasSuffix(from, NetSuffix) {
		host = NetHost
		port = NetPort
		isSSL = NetIsSSL
	} else if strings.HasSuffix(from, QqSuffix) {
		host = QqHost
		port = QqPort
		isSSL = QqIsSSL
	}

	receivers := strings.Split(to, ",")
	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		e.From = from
	}
	e.To = receivers
	e.Subject = subject
	if isHTML(body) {
		e.HTML = []byte(body)
	} else {
		e.Text = []byte(body)
	}

	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

// isHTML 检查字符串是否包含 HTML 标签
func isHTML(str string) bool {
	// 定义一个简单的正则表达式，用于检测 HTML 标签
	re := regexp.MustCompile(`(?i)<[a-z][\s\S]*>`)
	return re.MatchString(str)
}
