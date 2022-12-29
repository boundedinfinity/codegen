package codegen_type

import (
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type LoaderContext interface {
	GetFileInfo() *LoaderFileInfo
}

type LoaderFileInfo struct {
	Source   string
	Root     string
	IsFile   bool
	MimeType mime_type.MimeType
}
