package model

import "github.com/boundedinfinity/go-commoner/functional/optioner"

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenString struct {
	CodeGenCommon
	Min   optioner.Option[int]    `json:"min,omitempty"`
	Max   optioner.Option[int]    `json:"max,omitempty"`
	Regex optioner.Option[string] `json:"regex,omitempty"`
}

func (t CodeGenString) TypeId() string {
	return "string"
}

var _ CodeGenType = &CodeGenString{}

// /////////////////////////////////////////////////////////////////
// Builder
// ////////////////////////////////////////////////////////////////

func BuildString() *codeGenStringBuilder {
	return &codeGenStringBuilder{}
}

type codeGenStringBuilder struct {
	v CodeGenString
}

func (t *codeGenStringBuilder) Value() CodeGenType {
	return &t.v
}

func (t *codeGenStringBuilder) Name(v string) *codeGenStringBuilder {
	t.v.CodeGenCommon.Name = optioner.OfZero(v)
	return t
}

func (t *codeGenStringBuilder) Description(v string) *codeGenStringBuilder {
	t.v.CodeGenCommon.Description = optioner.OfZero(v)
	return t
}

func (t *codeGenStringBuilder) Required(v bool) *codeGenStringBuilder {
	t.v.CodeGenCommon.Required = optioner.OfZero(v)
	return t
}

func (t *codeGenStringBuilder) Min(v int) *codeGenStringBuilder {
	t.v.Min = optioner.OfZero(v)
	return t
}

func (t *codeGenStringBuilder) Max(v int) *codeGenStringBuilder {
	t.v.Max = optioner.OfZero(v)
	return t
}

func (t *codeGenStringBuilder) Regex(v string) *codeGenStringBuilder {
	t.v.Regex = optioner.OfZero(v)
	return t
}
