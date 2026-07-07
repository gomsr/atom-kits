## Why

The `emailx` package currently only supports plain-text and HTML email bodies. There is no way to send file attachments, inline images, or use CC/BCC/Reply-To headers. The underlying library (`github.com/jordan-wright/email`) fully supports these features, and adding them would significantly expand `emailx`'s usefulness for common use cases like sending invoices, reports, and verification emails with embedded images.

## What Changes

- Add a new `Send` function using the functional options pattern as the recommended entry point (in a new `send.go` file — zero changes to existing `email.go`)
- Export an `Attachment` struct supporting both file-path-based and in-memory byte attachments
- Add attachment options: `WithAttachments` (file paths), `WithAttachmentData` (raw bytes), `WithInlineImage` (HTML `cid:` references), `WithAttachment` (full control)
- Add server options: `WithServer` (explicit host/port/SSL), `WithServerFunc` (function-based, replaces old `defVal` pattern)
- Add recipient options: `WithCC`, `WithBCC`, `WithReplyTo`
- All validation fails fast: missing `from`/`to`, missing attachment data, non-existent file paths
- Existing API (`DoSend`, `DoSendType`, `DoSendTypeFunc`, `RealDoSend`, `SendEmail`) remains unchanged and backward-compatible

## Capabilities

### New Capabilities

- `email-attachments`: Send emails with file attachments (disk or in-memory), inline images for HTML content, and CC/BCC/Reply-To recipient fields, via a new functional-options-based `Send` function

### Modified Capabilities

<!-- No existing capabilities are modified — this is a purely additive change -->

## Impact

- **Code**: New file `emailx/send.go` (~200 lines); no changes to `emailx/email.go`
- **API**: New exported symbols: `Send`, `SendOption`, `Attachment`, `WithAttachments`, `WithAttachmentData`, `WithInlineImage`, `WithAttachment`, `WithServer`, `WithServerFunc`, `WithCC`, `WithBCC`, `WithReplyTo`
- **Dependencies**: No new external dependencies (reuses `github.com/jordan-wright/email` and standard library)
- **Backward compatibility**: Full — existing callers are unaffected
