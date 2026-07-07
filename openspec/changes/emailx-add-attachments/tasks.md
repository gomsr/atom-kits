## 1. New types and core structure

- [x] 1.1 Create `emailx/send.go` with package declaration and imports
- [x] 1.2 Define exported `Attachment` struct (Filename, ContentType, Data, FilePath, Inline)
- [x] 1.3 Define unexported `sendConfig` struct (host, port, isSSL, serverFunc, cc, bcc, replyTo, attachments)
- [x] 1.4 Define exported `SendOption` type (`func(*sendConfig)`)

## 2. Attachment options

- [x] 2.1 Implement `WithAttachments(paths ...string)` — validates file existence, reads content, detects MIME type
- [x] 2.2 Implement `WithAttachmentData(filename, contentType string, data []byte)` — validates data non-empty
- [x] 2.3 Implement `WithInlineImage(path string)` — reads file, sets Inline=true, detects MIME type
- [x] 2.4 Implement `WithAttachment(a Attachment)` — validates Data XOR FilePath is set

## 3. Server configuration options

- [x] 3.1 Implement `WithServer(host string, port int, isSSL bool)` — sets host/port/isSSL on config
- [x] 3.2 Implement `WithServerFunc(fn func() (string, int, bool))` — sets serverFunc on config
- [x] 3.3 Implement server resolution logic: serverFunc > explicit host > auto-detect from `from` suffix

## 4. Recipient options

- [x] 4.1 Implement `WithCC(cc ...string)` — appends to config.cc
- [x] 4.2 Implement `WithBCC(bcc ...string)` — appends to config.bcc
- [x] 4.3 Implement `WithReplyTo(replyTo ...string)` — appends to config.replyTo

## 5. Core Send function

- [x] 5.1 Implement `Send(from, nickname, secret string, to []string, subject, body string, opts ...SendOption) error`
- [x] 5.2 Validate `from` non-empty, `to` non-empty, each Attachment valid — fail immediately on first error
- [x] 5.3 Build `*email.Email` with From, To, Subject, CC, BCC, ReplyTo, HTML/Text body (reuse HTML detection from `isHTML`)
- [x] 5.4 Attach each `Attachment` using `e.Attach` or `e.AttachFile`, setting `HTMLRelated` for inline attachments
- [x] 5.5 Send via `e.SendWithTLS` (if SSL) or `e.Send` (if not)

## 6. Tests

- [x] 6.1 Add `TestSendSimple` — sends a plain-text email via the new `Send` function
- [x] 6.2 Add `TestSendWithFileAttachments` — sends with one or more file attachments
- [x] 6.3 Add `TestSendWithAttachmentData` — sends with in-memory byte attachment
- [x] 6.4 Add `TestSendWithInlineImage` — sends HTML email with inline image
- [x] 6.5 Add `TestSendValidation` — verifies error on empty from, empty to, and missing file
- [x] 6.6 Add `TestSendWithCCBCCReplyTo` — verifies CC/BCC/Reply-To headers are set
- [x] 6.7 Add `TestSendWithCustomServer` — verifies WithServer and WithServerFunc

## 7. Verification

- [x] 7.1 Run `go build ./emailx/...` to verify compilation
- [x] 7.2 Run existing tests `go test ./emailx/...` to confirm backward compatibility (TestIsHTML passes; pre-existing SMTP auth failures unrelated)
- [x] 7.3 Run `go vet ./emailx/...` for static analysis
