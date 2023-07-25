package model

import (
	"boundedinfinity/codegen/model/type_id"

	type_visibility "boundedinfinity/codegen/codegen_type/type_visability"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type Type interface {
	TypeId() type_id.TypeId
}

type Meta struct {
	Source    Source    `json:"source,omitempty"`
	Namespace Namespace `json:"namespace,omitempty"`
}

type Source struct {
	SourcePath     o.Option[string]   `json:"source-path,omitempty"`
	RootPath       o.Option[string]   `json:"root-path,omitempty"`
	SourceMimeType mime_type.MimeType `json:"source-mime-type,omitempty"`
}

type Namespace struct {
	RootNs   string
	SchemaNs string
	RelNs    string
	CurrNs   string
}

type Common struct {
	Name         o.Option[string]                         `json:"name,omitempty"`
	Desc         o.Option[string]                         `json:"desc,omitempty"`
	Version      o.Option[string]                         `json:"version,omitempty"`
	Required     o.Option[bool]                           `json:"required,omitempty"`
	Header       o.Option[string]                         `json:"header,omitempty"`
	FormatSource o.Option[bool]                           `json:"format-source,omitempty"`
	Deprecated   o.Option[bool]                           `json:"deprecated,omitempty"`
	Visibility   o.Option[type_visibility.TypeVisibility] `json:"visibility,omitempty"`
	Meta         Meta                                     `json:"meta,omitempty"`
}

type commonBuilder struct {
	v Common
}

func BuildCommon() *commonBuilder {
	return &commonBuilder{}
}

func (b *commonBuilder) Done() Common {
	return b.v
}

func (b *commonBuilder) Name(v string) *commonBuilder {
	b.v.Name = o.Some(v)
	return b
}

func (b *commonBuilder) Desc(v string) *commonBuilder {
	b.v.Desc = o.Some(v)
	return b
}

func (b *commonBuilder) Required(v bool) *commonBuilder {
	b.v.Required = o.Some(v)
	return b
}

func (b *commonBuilder) Version(v string) *commonBuilder {
	b.v.Version = o.Some(v)
	return b
}

func (b *commonBuilder) Header(v string) *commonBuilder {
	b.v.Header = o.Some(v)
	return b
}

func (b *commonBuilder) FormatSource(v bool) *commonBuilder {
	b.v.FormatSource = o.Some(v)
	return b
}
