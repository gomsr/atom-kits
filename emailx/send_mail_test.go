package emailx

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ── Validation tests (no SMTP required) ───────────────────────
// Test that DoSendOption fails fast on invalid input

func TestDoSendOption_EmptyFrom(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithTo("to@test.com").WithSubject("s").WithBody("b")
	err := DoSendOption(cfg)
	if err == nil {
		t.Fatal("expected error for empty from address")
	}
	if !strings.Contains(err.Error(), "from") {
		t.Fatalf("expected 'from' in error, got: %v", err)
	}
}

func TestDoSendOption_EmptyTo(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithFrom("from@test.com").WithSubject("s").WithBody("b")
	err := DoSendOption(cfg)
	if err == nil {
		t.Fatal("expected error for empty to list")
	}
	if !strings.Contains(err.Error(), "to") {
		t.Fatalf("expected 'to' in error, got: %v", err)
	}
}

func TestDoSendOption_UnknownDomain(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithFrom("user@unknown-example.com").
		WithTo("to@test.com").
		WithSubject("s").WithBody("b")
	err := DoSendOption(cfg)
	if err == nil {
		t.Fatal("expected error for unknown email domain")
	}
}

func TestDoSendOption_MissingFile(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithFrom("from@test.com").
		WithTo("to@test.com").
		WithSubject("s").WithBody("b").
		WithAttachments("/nonexistent/path/file.pdf")
	err := DoSendOption(cfg)
	if err == nil {
		t.Fatal("expected error for non-existent attachment file")
	}
	if !strings.Contains(err.Error(), "file not found") {
		t.Fatalf("expected 'file not found' in error, got: %v", err)
	}
}

func TestDoSendOption_AttachmentNoDataNoFile(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithFrom("from@test.com").
		WithTo("to@test.com").
		WithSubject("s").WithBody("b").
		WithAttachment(Attachment{Filename: "empty.bin"})
	err := DoSendOption(cfg)
	if err == nil {
		t.Fatal("expected error for attachment with neither Data nor FilePath")
	}
}

func TestDoSendOption_AttachmentBothDataAndFile(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithFrom("from@test.com").
		WithTo("to@test.com").
		WithSubject("s").WithBody("b").
		WithAttachment(Attachment{
			Filename: "conflict.bin",
			Data:     []byte("x"),
			FilePath: "/tmp/x",
		})
	err := DoSendOption(cfg)
	if err == nil {
		t.Fatal("expected error for attachment with both Data and FilePath")
	}
	if !strings.Contains(err.Error(), "mutually exclusive") {
		t.Fatalf("expected 'mutually exclusive' in error, got: %v", err)
	}
}

// ── Chained API tests ─────────────────────────────────────────

func TestSendConfig_ChainedBuild(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithSender(SenderConfig{From: "a@b.com", Nickname: "nn", Secret: "s"}).
		WithRecipient(RecipientConfig{To: []string{"to@b.com"}, CC: []string{"cc@b.com"}}).
		WithContent(ContentConfig{Subject: "sub", Body: "body"}).
		WithReplyTo("reply@b.com")

	if cfg.From != "a@b.com" {
		t.Fatalf("From = %q, want %q", cfg.From, "a@b.com")
	}
	if cfg.Nickname != "nn" {
		t.Fatalf("Nickname = %q, want %q", cfg.Nickname, "nn")
	}
	if cfg.Secret != "s" {
		t.Fatalf("Secret = %q, want %q", cfg.Secret, "s")
	}
	if len(cfg.To) != 1 || cfg.To[0] != "to@b.com" {
		t.Fatalf("To = %v, want [to@b.com]", cfg.To)
	}
	if len(cfg.CC) != 1 || cfg.CC[0] != "cc@b.com" {
		t.Fatalf("CC = %v, want [cc@b.com]", cfg.CC)
	}
	if cfg.Subject != "sub" {
		t.Fatalf("Subject = %q, want %q", cfg.Subject, "sub")
	}
	if cfg.Body != "body" {
		t.Fatalf("Body = %q, want %q", cfg.Body, "body")
	}
	if len(cfg.ReplyTo) != 1 || cfg.ReplyTo[0] != "reply@b.com" {
		t.Fatalf("ReplyTo = %v, want [reply@b.com]", cfg.ReplyTo)
	}
}

func TestSendConfig_ChainedAttachments(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithFrom("a@b.com").WithTo("to@b.com").
		WithAttachments("send_config.go", "send_mail.go").
		WithAttachmentData("gen.txt", "text/plain", []byte("hello")).
		WithAttachment(Attachment{Filename: "custom.bin", Data: []byte("world")})

	if len(cfg.Attachments) != 4 {
		t.Fatalf("Attachments count = %d, want 4", len(cfg.Attachments))
	}
	if cfg.Attachments[0].Filename != "send_config.go" {
		t.Fatalf("Att[0].Filename = %q, want send_config.go", cfg.Attachments[0].Filename)
	}
	if cfg.Attachments[2].Filename != "gen.txt" {
		t.Fatalf("Att[2].Filename = %q, want gen.txt", cfg.Attachments[2].Filename)
	}
	if string(cfg.Attachments[2].Data) != "hello" {
		t.Fatalf("Att[2].Data = %q, want hello", cfg.Attachments[2].Data)
	}
	if cfg.Attachments[3].Filename != "custom.bin" {
		t.Fatalf("Att[3].Filename = %q, want custom.bin", cfg.Attachments[3].Filename)
	}
}

func TestSendConfig_WithInlineImage(t *testing.T) {
	cfg := &SendConfig{}
	cfg.WithInlineImage("send_config.go")

	if len(cfg.Attachments) != 1 {
		t.Fatal("expected 1 attachment")
	}
	if !cfg.Attachments[0].Inline {
		t.Fatal("expected Inline=true for WithInlineImage")
	}
	if cfg.Attachments[0].Filename != "send_config.go" {
		t.Fatalf("Filename = %q, want send_config.go", cfg.Attachments[0].Filename)
	}
}

func TestSendConfig_WithServerFunc(t *testing.T) {
	cfg := &SendConfig{}
	fn := func() (string, int, bool) { return "smtp.test.com", 25, false }
	cfg.WithServerFunc(fn)

	if cfg.ServerFunc == nil {
		t.Fatal("ServerFunc is nil")
	}
	h, p, s := cfg.ServerFunc()
	if h != "smtp.test.com" || p != 25 || s != false {
		t.Fatalf("ServerFunc() = (%q, %d, %v), want (smtp.test.com, 25, false)", h, p, s)
	}
}

// ── Attachment validation tests ────────────────────────────────

func TestAttachment_FilepathOnly(t *testing.T) {
	a := Attachment{Filename: "f.txt", FilePath: "/tmp/f.txt"}
	if err := a.validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAttachment_DataOnly(t *testing.T) {
	a := Attachment{Filename: "f.txt", Data: []byte("x")}
	if err := a.validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAttachment_Empty(t *testing.T) {
	a := Attachment{Filename: "empty.bin"}
	if err := a.validate(); err == nil {
		t.Fatal("expected error for empty attachment")
	}
}

func TestAttachment_BothSet(t *testing.T) {
	a := Attachment{Filename: "x.bin", Data: []byte("x"), FilePath: "/tmp/x"}
	if err := a.validate(); err == nil {
		t.Fatal("expected error when both Data and FilePath are set")
	}
}

// ── Send() with chained config via SendEmail() ─────────────────

func TestSendConfig_SendEmail_Validation(t *testing.T) {
	cfg := &SendConfig{}
	err := cfg.WithSubject("s").WithBody("b").
		WithTo("to@test.com").
		SendEmail()
	if err == nil {
		t.Fatal("expected error for empty from in SendEmail")
	}
}

func TestSendConfig_SendEmail_MissingFile(t *testing.T) {
	cfg := &SendConfig{}
	err := cfg.WithFrom("from@test.com").WithTo("to@test.com").
		WithSubject("s").WithBody("b").
		WithAttachments("/nonexistent/file.pdf").
		SendEmail()
	if err == nil {
		t.Fatal("expected error for missing attachment file in SendEmail")
	}
}

// ── Integration tests (require SMTP credentials) ──────────────

func TestSend_Simple(t *testing.T) {
	err := Send(SenderConfig{
		From:     "zzhang_xz@163.com",
		Nickname: "zack",
		Secret:   secretNet,
	},
		WithRecipient(RecipientConfig{
			To: []string{"1252068782@qq.com"},
		}),
		WithContent(ContentConfig{
			Subject: "TestSend_Simple",
			Body:    "Hello from Send() via emailx!",
		}),
	)
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}
}

func TestSend_ChainedCall(t *testing.T) {

	err := (&SendConfig{}).
		WithSender(SenderConfig{
			From:     "zzhang_xz@163.com",
			Nickname: "zack",
			Secret:   secretNet,
		}).
		WithRecipient(RecipientConfig{To: []string{"1252068782@qq.com"}}).
		WithContent(ContentConfig{Subject: "TestSend_ChainedCall", Body: "chained SendEmail()"}).
		SendEmail()
	if err != nil {
		t.Fatalf("SendEmail failed: %v", err)
	}
}

func TestSend_WithFileAttachment(t *testing.T) {

	err := Send(SenderConfig{
		From:     "zzhang_xz@163.com",
		Nickname: "zack",
		Secret:   secretNet,
	},
		WithRecipient(RecipientConfig{
			To: []string{"1252068782@qq.com"},
		}),
		WithContent(ContentConfig{
			Subject: "TestSend_WithFileAttachment",
			Body:    "See attached file.",
		}),
		WithAttachments("send_config.go", "send_mail.go"),
	)
	if err != nil {
		t.Fatalf("Send with file attachments failed: %v", err)
	}
}

func TestSend_WithAttachmentData(t *testing.T) {

	err := Send(SenderConfig{
		From:     "zzhang_xz@163.com",
		Nickname: "zack",
		Secret:   secretNet,
	},
		WithRecipient(RecipientConfig{
			To: []string{"1252068782@qq.com"},
		}),
		WithContent(ContentConfig{
			Subject: "TestSend_WithAttachmentData",
			Body:    "See attached generated file.",
		}),
		WithAttachmentData("hello.txt", "text/plain", []byte("Hello, World!")),
	)
	if err != nil {
		t.Fatalf("Send with attachment data failed: %v", err)
	}
}

func TestSend_WithInlineImage(t *testing.T) {

	tmpDir := t.TempDir()
	imgPath := filepath.Join(tmpDir, "test.png")
	pngData := []byte{
		0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
		0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
		0x08, 0x02, 0x00, 0x00, 0x00, 0x90, 0x77, 0x53,
		0xDE,
		0x00, 0x00, 0x00, 0x0C, 0x49, 0x44, 0x41, 0x54,
		0x08, 0xD7, 0x63, 0x68, 0x00, 0x00, 0x00, 0x82,
		0x00, 0x81, 0x93, 0x9E, 0xE0,
		0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, 0x44,
		0xAE, 0x42, 0x60, 0x82,
	}
	if err := os.WriteFile(imgPath, pngData, 0644); err != nil {
		t.Fatalf("failed to create test PNG: %v", err)
	}

	htmlBody := `<html><body><h1>Test</h1><img src="cid:test.png"/></body></html>`
	err := Send(SenderConfig{
		From:     "zzhang_xz@163.com",
		Nickname: "zack",
		Secret:   secretNet,
	},
		WithRecipient(RecipientConfig{To: []string{"1252068782@qq.com"}}),
		WithContent(ContentConfig{Subject: "TestSend_WithInlineImage", Body: htmlBody}),
		WithInlineImage(imgPath),
	)
	if err != nil {
		t.Fatalf("Send with inline image failed: %v", err)
	}
}

func TestSend_WithCCAndReplyTo(t *testing.T) {

	err := Send(SenderConfig{
		From:     "zzhang_xz@163.com",
		Nickname: "zack",
		Secret:   secretNet,
	},
		WithRecipient(RecipientConfig{To: []string{"1252068782@qq.com"}}),
		WithContent(ContentConfig{
			Subject: "TestSend_WithCCAndReplyTo",
			Body:    "Testing CC and Reply-To.",
		}),
		WithCC("cc@example.com"),
		WithBCC("bcc@example.com"),
		WithReplyTo("reply@example.com"),
	)
	if err != nil {
		t.Fatalf("Send with CC/BCC/ReplyTo failed: %v", err)
	}
}

func TestDoSendFrom(t *testing.T) {

	err := DoSendFrom(
		"1252068782@qq.com",
		"TestDoSendFrom",
		"Hello from DoSendFrom!",
		"zzhang_xz@163.com",
		"zack",
		secretNet,
	)
	if err != nil {
		t.Fatalf("DoSendFrom failed: %v", err)
	}
}

func TestDoSendFrom_WithAttachment(t *testing.T) {

	err := DoSendFrom(
		"1252068782@qq.com",
		"TestDoSendFrom_WithAttachment",
		"See attached.",
		"zzhang_xz@163.com",
		"zack",
		secretNet,
		WithAttachments("send_config.go", "send_mail.go"),
		WithAttachmentData("note.txt", "text/plain", []byte("generated content")),
	)
	if err != nil {
		t.Fatalf("DoSendFrom with attachments failed: %v", err)
	}
}

func TestSend_WithServerFunc(t *testing.T) {

	err := Send(SenderConfig{
		From:     "support@feturax.com",
		Nickname: "feturax",
		Secret:   "l&%iTi%Dx)&$!G2R)%Q3YI24",
	},
		WithRecipient(RecipientConfig{To: []string{"zzhang_xz@163.com"}}),
		WithContent(ContentConfig{
			Subject: "TestSend_WithServerFunc",
			Body:    "Testing WithServerFunc.",
		}),
		WithServerFunc(ZohoFunc),
	)
	if err != nil {
		t.Fatalf("Send with ServerFunc failed: %v", err)
	}
}
