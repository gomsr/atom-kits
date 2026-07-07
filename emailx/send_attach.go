package emailx

import (
	"fmt"
)

// Attachment represents an email attachment.
// Either FilePath or Data must be set, but not both.
type Attachment struct {
	Filename    string `json:"filename" yaml:"filename" mapstructure:"filename"`             // displayed filename in the email client
	ContentType string `json:"content_type" yaml:"content_type" mapstructure:"content_type"` // MIME type; auto-detected from extension if empty
	Data        []byte `json:"data" yaml:"data" mapstructure:"data"`                         // raw content (mutually exclusive with FilePath)
	FilePath    string `json:"file_path" yaml:"file_path" mapstructure:"file_path"`          // read from disk (mutually exclusive with Data)
	Inline      bool   `json:"inline" yaml:"inline" mapstructure:"inline"`                   // if true, sets Content-Disposition: inline for HTML cid: references
}

func (a Attachment) validate() error {
	if a.FilePath == "" && len(a.Data) == 0 {
		return fmt.Errorf("attachment %q: must have either FilePath or Data set", a.Filename)
	}
	if a.FilePath != "" && len(a.Data) > 0 {
		return fmt.Errorf("attachment %q: FilePath and Data are mutually exclusive", a.Filename)
	}
	return nil
}
