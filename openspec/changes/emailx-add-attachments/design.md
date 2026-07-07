## Context

The `emailx` package in `github.com/gomsr/atom-kits` provides SMTP email sending via a procedural API (`DoSend`, `DoSendType`, `DoSendTypeFunc`). It wraps `github.com/jordan-wright/email`, which supports attachments, CC/BCC, Reply-To, and inline content — but `emailx` exposes none of these.

The existing API already has function sprawl (4 send variants + 1 deprecated), all with long positional parameter lists. Adding attachments via more positional parameters would worsen this. The functional options pattern is idiomatic Go for this situation.

## Goals / Non-Goals

**Goals:**
- Provide a single, clean `Send` function using functional options
- Support file attachments (by path), in-memory data attachments, and inline HTML images
- Support CC, BCC, and Reply-To headers
- Support explicit SMTP server config and function-based server config
- Fail-fast validation (missing files, invalid params)
- Zero changes to existing code (`email.go`)
- Full backward compatibility

**Non-Goals:**
- Deprecating or removing existing API functions
- Attachment pools or connection pooling
- DKIM signing, S/MIME encryption, or other advanced email features
- Templating or email body building
- STARTTLS variant (the underlying lib supports it, but the existing API doesn't expose it either)

## Decisions

### Decision 1: New file (`send.go`) vs editing `email.go`

**Choice: New file `send.go`**

**Rationale:** Zero risk of breaking existing callers. Clean separation between old and new API. The existing `email.go` has grown organically and mixing patterns would confuse readers. New file signals "this is the new way".

**Alternative considered:** Editing `email.go` to add `Send` — rejected because it muddies the file and risks accidental breakage.

### Decision 2: Functional options vs struct config

**Choice: Functional options (`SendOption` interface)**

```go
func Send(from, nickname, secret string, to []string, subject, body string,
          opts ...SendOption) error
```

**Rationale:** Idiomatic Go (e.g., `http.Client`, `grpc.Dial`). Extensible without breaking callers. Reads naturally:

```go
Send(from, nick, secret, to, subj, body,
    WithAttachments("a.pdf"),
    WithCC("cc@test.com"),
)
```

**Alternative considered:** A `Message` or `SendRequest` struct — rejected because struct initialization is verbose for callers who DON'T want attachments (the common case). Options keep the simple case simple: `Send(from, nick, secret, to, subj, body)`.

### Decision 3: Attachment struct export

**Choice: Export `Attachment` struct**

```go
type Attachment struct {
    Filename    string
    ContentType string
    Data        []byte
    FilePath    string
    Inline      bool
}
```

**Rationale:** `WithAttachment(a Attachment)` gives advanced callers full control (e.g., programmatically building attachment lists). Simpler options (`WithAttachments`, `WithAttachmentData`, `WithInlineImage`) cover the common cases.

**Alternative considered:** Keeping `Attachment` unexported and only exposing option functions — rejected because it prevents callers from building attachment lists dynamically.

### Decision 4: Server resolution precedence

**Choice:** `WithServerFunc` > `WithServer` > auto-detect from `from` suffix

**Rationale:** Most specific to most general. Same order as the existing `defVal` pattern in `DoSend` but made explicit. When both `WithServerFunc` and `WithServer` are passed, `WithServerFunc` wins (it was set last by the options applier — each option overrides the previous).

### Decision 5: Shared SMTP logic

**Choice: Extract SMTP send logic into an internal helper, but otherwise keep `email.go` untouched**

**Rationale:** `RealDoSend` already handles the SMTP send pattern. Either `Send` calls `RealDoSend` (which would need attachment support added), or `Send` has its own send logic. The cleanest approach: refactor the SMTP send call (`email.NewEmail → set fields → Send/SendWithTLS`) into a small internal `sendEmail` helper that BOTH `RealDoSend` and the new code can use. But to truly keep `email.go` zero-diff, the new `send.go` will inline the ~15 lines of SMTP wiring. The duplication is minimal and intentional.

## Risks / Trade-offs

- **API confusion**: Having two APIs (old `DoSend*` + new `Send`) in the same package may confuse newcomers. → Mitigation: Add a doc comment on `DoSend` pointing to `Send` as the recommended entry point.
- **No STARTTLS**: The new `Send` doesn't expose `SendWithStartTLS`, only `Send` (STARTTLS) and `SendWithTLS` (implicit SSL) — matching the existing API surface. → Mitigation: `WithServer(host, port, false)` auto-detects whether to use TLS (based on existing `RealDoSend` logic via `isSSL` flag).
- **Large file size**: `send.go` may grow large as more options are added. → Mitigation: start with a single file; split into `send.go` + `attachment.go` later if needed.
