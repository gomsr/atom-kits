package emailx

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"mime"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

func (c *SendConfig) SendEmail(opts ...SendOption) error {
	for _, opt := range opts {
		opt(c)
	}
	return DoSendOption(c)
}

func (c *SendConfig) SendEmailFrom(to, subject, body, from, nickname, secret string, opts ...SendOption) error {
	return DoSendFrom(to, subject, body, from, nickname, secret, opts...)
}

func Send(sender SenderConfig, opts ...SendOption) error {
	cfg := &SendConfig{SenderConfig: sender}
	// Apply options (they override/extend the config)
	for _, opt := range opts {
		opt(cfg)
	}

	return DoSendOption(cfg)
}

func DoSendFrom(to, subject, body, from, nickname, secret string, opts ...SendOption) error {
	cfg := &SendConfig{
		SenderConfig: SenderConfig{
			From:     from,
			Nickname: nickname,
			Secret:   secret,
		},
		RecipientConfig: RecipientConfig{To: strings.Split(to, ",")},
		ContentConfig:   ContentConfig{Subject: subject, Body: body},
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return DoSendOption(cfg)
}

func DoSendOption(cfg *SendConfig) error {
	// Validate
	if len(cfg.From) == 0 {
		return errors.New("emailx.Send: from address cannot be empty")
	}
	if len(cfg.To) == 0 {
		return errors.New("emailx.Send: to address cannot be empty")
	}
	for _, a := range cfg.Attachments {
		if err := a.validate(); err != nil {
			return fmt.Errorf("emailx.Send: %w", err)
		}
		if a.FilePath != "" {
			_, err := os.Stat(a.FilePath)
			if os.IsNotExist(err) {
				return fmt.Errorf("emailx.Send: attachment file not found: %s", a.FilePath)
			}
			if err != nil {
				return fmt.Errorf("emailx.Send: cannot access attachment file %s: %w", a.FilePath, err)
			}
		}
	}

	// Resolve server
	host, port, isSSL := "", 0, false
	if nil == cfg.ServerFunc {
		host, port, isSSL = resolveServer(cfg.From)
	} else {
		host, port, isSSL = cfg.ServerFunc()
	}

	if host == "" {
		return fmt.Errorf("emailx.Send: unable to determine SMTP server for %q; set Host or use WithServer/WithServerFunc", cfg.From)
	}

	// Build email
	auth := smtp.PlainAuth("", cfg.From, cfg.Secret, host)
	e := email.NewEmail()
	if cfg.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", cfg.Nickname, cfg.From)
	} else {
		e.From = cfg.From
	}
	e.To = cfg.To
	e.Subject = cfg.Subject
	if isHTML(cfg.Body) {
		e.HTML = []byte(cfg.Body)
	} else {
		e.Text = []byte(cfg.Body)
	}

	if len(cfg.CC) > 0 {
		e.Cc = cfg.CC
	}
	if len(cfg.BCC) > 0 {
		e.Bcc = cfg.BCC
	}
	if len(cfg.ReplyTo) > 0 {
		e.ReplyTo = cfg.ReplyTo
	}

	// Attach files
	for _, a := range cfg.Attachments {
		var att *email.Attachment
		var err error

		if a.FilePath != "" {
			att, err = e.AttachFile(a.FilePath)
		} else {
			ct := a.ContentType
			if ct == "" {
				ct = mime.TypeByExtension(filepath.Ext(a.Filename))
			}
			if ct == "" {
				ct = "application/octet-stream"
			}
			att, err = e.Attach(bytes.NewReader(a.Data), a.Filename, ct)
		}
		if err != nil {
			return fmt.Errorf("emailx.Send: failed to attach %q: %w", a.Filename, err)
		}
		if a.Inline {
			att.HTMLRelated = true
		}
	}

	// Send
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		return e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	}
	return e.Send(hostAddr, auth)
}
