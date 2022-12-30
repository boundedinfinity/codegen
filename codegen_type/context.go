package codegen_type

import (
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type LoaderContext interface {
	GetFileInfo() *FileInfo
	GetNamespace() *Namespace
}

type FileInfo struct {
	Source   string
	Root     string
	MimeType mime_type.MimeType
}

type Namespace struct {
	RootNs   string
	CurrNs   string
	SchemaNs string
	RelNs    string
}
