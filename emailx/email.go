package emailx

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

// Deprecated: Use Send instead.
type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // 收件人:多个以英文逗号分隔 例：a@qq.com b@qq.com 正式开发中请把此项目作为参数使用
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 发件人  你自己要发邮件的邮箱
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称    发件人昵称 通常为自己的邮箱
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`       // 是否SSL   是否开启SSL
}

// Deprecated: Use Send instead.
// SendEmail is the original email sending method. Prefer Send() from send_attach.go.
//
// Migration:
//
//	emailx.SendEmail(config, to, subject, body)
//	→ emailx.Send(emailx.SendConfig{
//	      From:     config.From,
//	      Nickname: config.Nickname,
//	      Secret:   config.Secret,
//	      To:       to,
//	      Subject:  subject,
//	      Body:     body,
//	      Host:     config.Host,
//	      Port:     config.Port,
//	      IsSSL:    config.IsSSL,
//	  })
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

// Deprecated: Use Send with WithServerFunc instead.
//
// Migration:
//
//	emailx.DoSendTypeFunc(to, subject, body, from, nickname, secret, fn)
//	→ emailx.Send(emailx.SendConfig{
//	      From:       from,
//	      Nickname:   nickname,
//	      Secret:     secret,
//	      To:         strings.Split(to, ","),
//	      Subject:    subject,
//	      Body:       body,
//	      ServerFunc: fn,
//	  })
func DoSendTypeFunc(to, subject, body, from, nickname, secret string, defVal func() (string, int, bool)) (err error) {
	host, port, isSSL := defVal()
	return RealDoSend(to, subject, body, from, nickname, secret, host, port, isSSL)
}

// Deprecated: Use Send instead (auto-detection handles provider types).
//
// Migration:
//
//	emailx.DoSendType("google", to, subject, body, from, nickname, secret)
//	→ emailx.Send(from, nickname, secret, strings.Split(to, ","), subject, body)
func DoSendType(types, to, subject, body, from, nickname, secret string) (err error) {
	host, port, isSSL := deterServer(types)
	return RealDoSend(to, subject, body, from, nickname, secret, host, port, isSSL)
}

// Deprecated: Use Send instead.
//
// Migration:
//
//	emailx.DoSend(to, subject, body, from, nickname, secret)
//	→ emailx.Send(from, nickname, secret, strings.Split(to, ","), subject, body)
//
// For custom servers:
//
//	emailx.DoSend(to, subject, body, from, nickname, secret, fn)
//	→ emailx.Send(from, nickname, secret, strings.Split(to, ","), subject, body, emailx.WithServerFunc(fn))
func DoSend(to, subject, body, from, nickname, secret string, defVal ...func() (string, int, bool)) (err error) {
	// parse host, ssl, port info
	host, port, isSSL := "", 0, false
	if strings.HasSuffix(from, GmailSuffix) {
		host, port, isSSL = GmailFunc()

	} else if strings.HasSuffix(from, NetSuffix) {
		host, port, isSSL = NetFunc()

	} else if strings.HasSuffix(from, QqSuffix) {
		host, port, isSSL = QqFunc()

	} else if strings.HasSuffix(from, ICloudSuffix) {
		host, port, isSSL = ICloudFunc()

	} else {
		if len(defVal) > 0 {
			host, port, isSSL = defVal[0]()
		}
	}

	return RealDoSend(to, subject, body, from, nickname, secret, host, port, isSSL)
}

// Deprecated: Use Send with WithServer instead.
//
// Migration:
//
//	emailx.RealDoSend(to, subject, body, from, nickname, secret, host, port, isSSL)
//	→ emailx.Send(from, nickname, secret, strings.Split(to, ","), subject, body, emailx.WithServer(host, port, isSSL))
func RealDoSend(to, subject, body, from, nickname, secret, host string, port int, isSSL bool) (err error) {
	if len(from) == 0 {
		return errors.New("函数配置的发件人不能为空")
	}
	if len(to) == 0 {
		return errors.New("函数配置的收件人不能为空")
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

func deterServer(types string) (string, int, bool) {
	if strings.EqualFold(types, GmailType) {
		return GmailFunc()
	}
	if strings.EqualFold(types, NetType) {
		return NetFunc()
	}
	if strings.EqualFold(types, QqType) {
		return QqFunc()
	}
	if strings.EqualFold(types, ICloudType) {
		return ICloudFunc()
	}

	return "", 0, false
}
