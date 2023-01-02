package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type LoaderContext interface {
	GetFileInfo() *SourceMeta
	GetNamespace() *RenderNamespace
}

type SourceMeta struct {
	SourcePath     o.Option[string]   `json:"source-path,omitempty"`
	RootPath       o.Option[string]   `json:"root-path,omitempty"`
	SourceMimeType mime_type.MimeType `json:"source-mime-type,omitempty"`
}

func (t *SourceMeta) Source() *SourceMeta {
	return t
}

type RenderNamespace struct {
	RootNs   string
	CurrNs   string
	SchemaNs string
	RelNs    string
}

func (t *RenderNamespace) Namespace() *RenderNamespace {
	return t
}
