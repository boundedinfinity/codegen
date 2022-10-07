package model

import (
	"boundedinfinity/codegen/uri_scheme"

	"github.com/boundedinfinity/mimetyper/mime_type"
)

type SourceInfo struct {
	InputUri  string
	Path      string
	Scheme    uri_scheme.Scheme
	LocalPath string
	MimeType  mime_type.MimeType
}
