package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenRef struct {
	CodeGenCommon `json:",inline,omitempty"`
	Ref           optioner.Option[string] `json:"ref,omitempty"`
}

func (t CodeGenRef) TypeId() string {
	return "ref"
}

var _ CodeGenType = &CodeGenRef{}

//////////////////////////////////////////////////////////////////
// Marshal
//////////////////////////////////////////////////////////////////

func (t *CodeGenRef) MarshalJSON() ([]byte, error) {
	dto := struct {
		CodeGenRef
		TypeId string `json:"type-id"`
	}{
		TypeId:     t.TypeId(),
		CodeGenRef: *t,
	}

	return json.Marshal(dto)
}

//////////////////////////////////////////////////////////////////
// Builders
//////////////////////////////////////////////////////////////////

func NewRef() *CodeGenRef {
	return &CodeGenRef{}
}

func (t *CodeGenRef) WithName(v string) *CodeGenRef {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *CodeGenRef) WithDescription(v string) *CodeGenRef {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *CodeGenRef) WithRequired(v bool) *CodeGenRef {
	t.CodeGenCommon.WithRequired(v)
	return t
}

func (t *CodeGenRef) WithDefault(v CodeGenRef) *CodeGenRef {
	t.CodeGenCommon.WithDefault(&v)
	return t
}

func (t *CodeGenRef) WithEager(v bool) *CodeGenRef {
	t.CodeGenCommon.WithEager(v)
	return t
}

func (t *CodeGenRef) WithRef(v string) *CodeGenRef {
	t.Ref = optioner.OfZero(v)
	return t
}
