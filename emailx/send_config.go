package emailx

import (
	"path/filepath"
	"regexp"
)

// ── Sub-configs ────────────────────────────────────────────────

// SenderConfig holds sender identity and SMTP server configuration.
type SenderConfig struct {
	From       string                     `json:"from" yaml:"from" mapstructure:"from"`             // 发件人邮箱
	Nickname   string                     `json:"nickname" yaml:"nickname" mapstructure:"nickname"` // 发件人昵称
	Secret     string                     `json:"secret" yaml:"secret" mapstructure:"secret"`       // SMTP 密钥
	ServerFunc func() (string, int, bool) `json:"-" yaml:"-" mapstructure:"-"`                      // 函数式服务器配置（优先级最高，不可序列化）
}

// RecipientConfig holds all recipient addresses.
type RecipientConfig struct {
	To      []string `json:"to" yaml:"to" mapstructure:"to"`                   // 收件人列表
	CC      []string `json:"cc" yaml:"cc" mapstructure:"cc"`                   // 抄送
	BCC     []string `json:"bcc" yaml:"bcc" mapstructure:"bcc"`                // 密送
	ReplyTo []string `json:"reply_to" yaml:"reply_to" mapstructure:"reply_to"` // 回复地址
}

// ContentConfig holds email content.
type ContentConfig struct {
	Subject     string       `json:"subject" yaml:"subject" mapstructure:"subject"`             // 邮件主题
	Body        string       `json:"body" yaml:"body" mapstructure:"body"`                      // 邮件正文
	Attachments []Attachment `json:"attachments" yaml:"attachments" mapstructure:"attachments"` // 附件列表
}

// ── SendConfig ─────────────────────────────────────────────────

// SendOption is a functional option for Send.
type SendOption func(*SendConfig)

// SendConfig holds all parameters for sending an email.
// Embedded structs use mapstructure squash for flat serialization.
// Callers access fields directly: cfg.From, cfg.Subject, etc.
type SendConfig struct {
	SenderConfig    `mapstructure:",squash"`
	RecipientConfig `mapstructure:",squash"`
	ContentConfig   `mapstructure:",squash"`
}

// ── Chained methods (*SendConfig receivers) ────────────────────

// WithSender sets sender identity and SMTP server config.
func (c *SendConfig) WithSender(sc SenderConfig) *SendConfig {
	c.SenderConfig = sc
	return c
}

// WithRecipient sets all recipient addresses.
func (c *SendConfig) WithRecipient(rc RecipientConfig) *SendConfig {
	c.RecipientConfig = rc
	return c
}

// WithContent sets email content including attachments.
func (c *SendConfig) WithContent(cc ContentConfig) *SendConfig {
	c.ContentConfig = cc
	return c
}

func (c *SendConfig) WithFrom(from string) *SendConfig {
	c.From = from
	return c
}

func (c *SendConfig) WithNickname(nickname string) *SendConfig {
	c.Nickname = nickname
	return c
}

func (c *SendConfig) WithSecret(secret string) *SendConfig {
	c.Secret = secret
	return c
}

func (c *SendConfig) WithTo(to ...string) *SendConfig {
	c.To = append(c.To, to...)
	return c
}

func (c *SendConfig) WithSubject(subject string) *SendConfig {
	c.Subject = subject
	return c
}

func (c *SendConfig) WithBody(body string) *SendConfig {
	c.Body = body
	return c
}

func (c *SendConfig) WithAttachments(paths ...string) *SendConfig {
	for _, p := range paths {
		c.Attachments = append(c.Attachments, Attachment{
			FilePath: p,
			Filename: filepath.Base(p),
		})
	}
	return c
}

func (c *SendConfig) WithAttachmentData(filename, contentType string, data []byte) *SendConfig {
	c.Attachments = append(c.Attachments, Attachment{
		Filename:    filename,
		ContentType: contentType,
		Data:        data,
	})
	return c
}

func (c *SendConfig) WithInlineImage(path string) *SendConfig {
	c.Attachments = append(c.Attachments, Attachment{
		FilePath: path,
		Filename: filepath.Base(path),
		Inline:   true,
	})
	return c
}

func (c *SendConfig) WithAttachment(a Attachment) *SendConfig {
	c.Attachments = append(c.Attachments, a)
	return c
}

func (c *SendConfig) WithServerFunc(fn func() (string, int, bool)) *SendConfig {
	c.ServerFunc = fn
	return c
}

func (c *SendConfig) WithCC(cc ...string) *SendConfig {
	c.CC = append(c.CC, cc...)
	return c
}

func (c *SendConfig) WithBCC(bcc ...string) *SendConfig {
	c.BCC = append(c.BCC, bcc...)
	return c
}

func (c *SendConfig) WithReplyTo(replyTo ...string) *SendConfig {
	c.ReplyTo = append(c.ReplyTo, replyTo...)
	return c
}

// ── SendOption functions (standalone) ──────────────────────────

func WithSender(sc SenderConfig) SendOption {
	return func(c *SendConfig) { c.WithSender(sc) }
}

func WithRecipient(rc RecipientConfig) SendOption {
	return func(c *SendConfig) { c.WithRecipient(rc) }
}

func WithContent(cc ContentConfig) SendOption {
	return func(c *SendConfig) { c.WithContent(cc) }
}

func WithFrom(from string) SendOption {
	return func(c *SendConfig) { c.WithFrom(from) }
}

func WithNickname(nickname string) SendOption {
	return func(c *SendConfig) { c.WithNickname(nickname) }
}

func WithSecret(secret string) SendOption {
	return func(c *SendConfig) { c.WithSecret(secret) }
}

func WithTo(to ...string) SendOption {
	return func(c *SendConfig) { c.WithTo(to...) }
}

func WithSubject(subject string) SendOption {
	return func(c *SendConfig) { c.WithSubject(subject) }
}

func WithBody(body string) SendOption {
	return func(c *SendConfig) { c.WithBody(body) }
}

func WithAttachments(paths ...string) SendOption {
	return func(c *SendConfig) { c.WithAttachments(paths...) }
}

func WithAttachmentData(filename, contentType string, data []byte) SendOption {
	return func(c *SendConfig) { c.WithAttachmentData(filename, contentType, data) }
}

func WithInlineImage(path string) SendOption {
	return func(c *SendConfig) { c.WithInlineImage(path) }
}

func WithAttachment(a Attachment) SendOption {
	return func(c *SendConfig) { c.WithAttachment(a) }
}

func WithServerFunc(fn func() (string, int, bool)) SendOption {
	return func(c *SendConfig) { c.WithServerFunc(fn) }
}

func WithCC(cc ...string) SendOption {
	return func(c *SendConfig) { c.WithCC(cc...) }
}

func WithBCC(bcc ...string) SendOption {
	return func(c *SendConfig) { c.WithBCC(bcc...) }
}

func WithReplyTo(replyTo ...string) SendOption {
	return func(c *SendConfig) { c.WithReplyTo(replyTo...) }
}

func isHTML(str string) bool {
	re := regexp.MustCompile(`(?i)<[a-z][\s\S]*>`)
	return re.MatchString(str)
}

