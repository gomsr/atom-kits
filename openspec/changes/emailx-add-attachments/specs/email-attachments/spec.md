## ADDED Requirements

### Requirement: Send email with functional options API

The system SHALL provide a `Send` function that accepts functional options for configuring attachments, server settings, and recipient fields.

#### Scenario: Send simple email without options

- **WHEN** caller invokes `Send(from, nickname, secret, to, subject, body)`
- **THEN** the email is sent with the provided parameters and server settings auto-detected from the `from` address suffix

#### Scenario: Send with empty from address

- **WHEN** caller invokes `Send` with an empty `from` string
- **THEN** the function returns an error immediately

#### Scenario: Send with empty to list

- **WHEN** caller invokes `Send` with an empty `to` slice
- **THEN** the function returns an error immediately

### Requirement: File path attachments

The system SHALL support attaching files from disk via the `WithAttachments` option, with automatic MIME type detection from file extension.

#### Scenario: Attach a single file

- **WHEN** caller passes `WithAttachments("report.pdf")`
- **THEN** the file `report.pdf` is attached with MIME type `application/pdf`

#### Scenario: Attach multiple files

- **WHEN** caller passes `WithAttachments("a.pdf", "b.png")`
- **THEN** both files are attached with their respective MIME types

#### Scenario: Attachment file does not exist

- **WHEN** caller passes a file path that does not exist
- **THEN** the function returns an error immediately before attempting to send

### Requirement: In-memory data attachments

The system SHALL support attaching raw byte data via the `WithAttachmentData` option, with caller-specified filename and content type.

#### Scenario: Attach generated content

- **WHEN** caller passes `WithAttachmentData("report.pdf", "application/pdf", generatedPdfBytes)`
- **THEN** the bytes are attached as `report.pdf` with MIME type `application/pdf`

### Requirement: Inline image attachments

The system SHALL support attaching images for HTML `cid:` references via the `WithInlineImage` option.

#### Scenario: Attach inline image

- **WHEN** caller passes `WithInlineImage("logo.png")`
- **THEN** the file is attached with `Content-Disposition: inline` and a `Content-ID` header set to the filename, making it referenceable as `<img src="cid:logo.png">` in the HTML body

### Requirement: Custom attachment

The system SHALL support passing a pre-built `Attachment` struct via the `WithAttachment` option for full caller control.

#### Scenario: Attach with full custom settings

- **WHEN** caller passes `WithAttachment(Attachment{Filename: "data.bin", ContentType: "application/octet-stream", Data: bytes, Inline: false})`
- **THEN** the attachment is included with exactly the specified properties

#### Scenario: Attachment has neither Data nor FilePath

- **WHEN** caller passes an `Attachment` with both `Data` nil and `FilePath` empty
- **THEN** the function returns an error immediately

### Requirement: CC, BCC, and Reply-To recipients

The system SHALL support CC, BCC, and Reply-To headers via functional options.

#### Scenario: Send with CC

- **WHEN** caller passes `WithCC("cc@example.com")`
- **THEN** the email includes a CC header with the specified address

#### Scenario: Send with BCC

- **WHEN** caller passes `WithBCC("bcc@example.com")`
- **THEN** the email includes a BCC header with the specified address

#### Scenario: Send with Reply-To

- **WHEN** caller passes `WithReplyTo("reply@example.com")`
- **THEN** the email includes a Reply-To header with the specified address

### Requirement: Explicit SMTP server configuration

The system SHALL support overriding auto-detected server settings via `WithServer` and `WithServerFunc` options.

#### Scenario: Explicit server

- **WHEN** caller passes `WithServer("smtp.custom.com", 465, true)`
- **THEN** the email is sent via `smtp.custom.com:465` using implicit SSL

#### Scenario: Function-based server

- **WHEN** caller passes `WithServerFunc(func() (string, int, bool) { return "smtp.custom.com", 587, false })`
- **THEN** the function is called to determine host, port, and SSL settings

### Requirement: Backward compatibility

The system SHALL preserve all existing exported functions (`DoSend`, `DoSendType`, `DoSendTypeFunc`, `RealDoSend`, `SendEmail`) with unchanged signatures and behavior.

#### Scenario: Existing DoSend callers unaffected

- **WHEN** existing code calls `DoSend(to, subject, body, from, nickname, secret)`
- **THEN** the email is sent exactly as before, with no change in behavior

### Requirement: All new code in a separate file

The system SHALL place all new implementation in `send.go`, leaving `email.go` unchanged.

#### Scenario: Zero diff on email.go

- **WHEN** the change is applied
- **THEN** `email.go` has zero modifications relative to its pre-change state
