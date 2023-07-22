package model

import (
	"boundedinfinity/codegen/model/type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type Source struct {
	Path string `json:"path,omitempty"`
}

type Type interface {
	TypeId() type_id.TypeId
}

type Common struct {
	Name         o.Option[string] `json:"name,omitempty"`
	Desc         o.Option[string] `json:"desc,omitempty"`
	Version      o.Option[string] `json:"version,omitempty"`
	Required     o.Option[bool]   `json:"required,omitempty"`
	Header       o.Option[string] `json:"header,omitempty"`
	FormatSource o.Option[bool]   `json:"format-source,omitempty"`
	Source       Source           `json:"source,omitempty"`
}

type commonBuilder struct {
	v *Common
}

func BuildCommon(v *Common) *commonBuilder {
	return &commonBuilder{v: v}
}

func (b *commonBuilder) Done() Common {
	return *b.v
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
