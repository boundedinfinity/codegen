package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Interface
//////////////////////////////////////////////////////////////////

type CodeGenType interface {
	TypeId() string
}

///////////////////////////////////////////////////////////////////
// Common
//////////////////////////////////////////////////////////////////

type CodeGenCommon struct {
	Name        optioner.Option[string]        `json:"name,omitempty"`
	Description optioner.Option[string]        `json:"description,omitempty"`
	Required    optioner.Option[bool]          `json:"required,omitempty"`
	Default     optioner.Option[CodeGenType]   `json:"default,omitempty"`
	Inherit     optioner.Option[string]        `json:"inherit,omitempty"`
	Links       optioner.Option[[]CodeGenLink] `json:"links,omitempty"`
}

///////////////////////////////////////////////////////////////////
// Marshal
//////////////////////////////////////////////////////////////////

type codeGenCommonMarshal struct {
	Name        *string      `json:"name,omitempty"`
	Description *string      `json:"description,omitempty"`
	Required    *bool        `json:"required,omitempty"`
	Default     *CodeGenType `json:"default,omitempty"`
}

func (t *CodeGenCommon) MarshalJSON() ([]byte, error) {
	v := codeGenCommonMarshal{
		Name:        t.Name.OrNil(),
		Description: t.Description.OrNil(),
		Required:    t.Required.OrNil(),
		Default:     t.Default.OrNil(),
	}

	return json.Marshal(&v)
}

///////////////////////////////////////////////////////////////////
// Builder
///////////////////////////////////////////////////////////////////

func BuildCommon(v *CodeGenCommon) *codeGenCommonBuilder {
	return &codeGenCommonBuilder{v: v}
}

type codeGenCommonBuilder struct {
	v *CodeGenCommon
}

func (t *codeGenCommonBuilder) Name(v string) *codeGenCommonBuilder {
	t.v.Name = optioner.OfZero(v)
	return t
}

func (t *codeGenCommonBuilder) Description(v string) *codeGenCommonBuilder {
	t.v.Description = optioner.OfZero(v)
	return t
}

func (t *codeGenCommonBuilder) Required(v bool) *codeGenCommonBuilder {
	t.v.Required = optioner.OfZero(v)
	return t
}

func (t *codeGenCommonBuilder) Default(v CodeGenType) *codeGenCommonBuilder {
	t.v.Default = optioner.OfZero(v)
	return t
}
