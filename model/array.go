package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenArray struct {
	CodeGenCommon
	Items CodeGenType          `json:"items,omitempty"`
	Min   optioner.Option[int] `json:"min,omitempty"`
	Max   optioner.Option[int] `json:"max,omitempty"`
}

func (t CodeGenArray) TypeId() string {
	return "array"
}

var _ CodeGenType = &CodeGenArray{}

//////////////////////////////////////////////////////////////////
// Marshal
//////////////////////////////////////////////////////////////////

type codeGenArray struct {
	CodeGenCommon
	Items CodeGenType `json:"items,omitempty"`
}

func (t *CodeGenArray) MarshalJSON() ([]byte, error) {
	dto := typedDto{
		TypeId: t.TypeId(),
		Value: codeGenArray{
			CodeGenCommon: t.CodeGenCommon,
			Items:         t.Items,
		},
	}

	return json.Marshal(dto)
}

//////////////////////////////////////////////////////////////////
// Builder
//////////////////////////////////////////////////////////////////

func BuildArray() *codeGenArrayBuilder {
	return &codeGenArrayBuilder{}
}

type codeGenArrayBuilder struct {
	v CodeGenArray
}

func (t *codeGenArrayBuilder) Value() CodeGenType {
	return &t.v
}

func (t *codeGenArrayBuilder) Name(v string) *codeGenArrayBuilder {
	t.v.Name = optioner.OfZero(v)
	return t
}

func (t *codeGenArrayBuilder) Description(v string) *codeGenArrayBuilder {
	t.v.Description = optioner.OfZero(v)
	return t
}

func (t *codeGenArrayBuilder) Required(v bool) *codeGenArrayBuilder {
	t.v.Required = optioner.OfZero(v)
	return t
}

func (t *codeGenArrayBuilder) Items(v CodeGenType) *codeGenArrayBuilder {
	t.v.Items = v
	return t
}
