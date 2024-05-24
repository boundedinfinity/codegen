package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	CodeGenCommon
}

func (t CodeGenBoolean) TypeId() string {
	return "boolean"
}

var _ CodeGenType = &CodeGenBoolean{}

//////////////////////////////////////////////////////////////////
// Marshal
//////////////////////////////////////////////////////////////////

type codeGenBoolean struct {
	CodeGenCommon
}

func (t *CodeGenBoolean) MarshalJSON() ([]byte, error) {
	dto := typedDto{
		TypeId: t.TypeId(),
		Value: codeGenBoolean{
			CodeGenCommon: t.CodeGenCommon,
		},
	}

	return json.Marshal(dto)
}

//////////////////////////////////////////////////////////////////
// Builder
//////////////////////////////////////////////////////////////////

func BuildBoolean() *codeGenBooleanBuilder {
	return &codeGenBooleanBuilder{}
}

type codeGenBooleanBuilder struct {
	v CodeGenBoolean
}

func (t *codeGenBooleanBuilder) Value() CodeGenType {
	return &t.v
}

func (t *codeGenBooleanBuilder) Name(v string) *codeGenBooleanBuilder {
	t.v.Name = optioner.OfZero(v)
	return t
}

func (t *codeGenBooleanBuilder) Description(v string) *codeGenBooleanBuilder {
	t.v.Description = optioner.OfZero(v)
	return t
}

func (t *codeGenBooleanBuilder) Required(v bool) *codeGenBooleanBuilder {
	t.v.Required = optioner.OfZero(v)
	return t
}
