package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type EnumValue struct {
	Name        optioner.Option[string] `json:"name,omitempty"`
	Value       optioner.Option[string] `json:"value,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
}

type Enum struct {
	CodeGenCommon
}

func (t Enum) TypeId() string {
	return "enum"
}

var _ CodeGenType = &Enum{}

//////////////////////////////////////////////////////////////////
// Marshal
//////////////////////////////////////////////////////////////////

func (t *Enum) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId string `json:"type-id"`
		Enum   `json:",inline"`
	}{
		TypeId: t.TypeId(),
		Enum:   *t,
	}

	return json.Marshal(dto)
}

///////////////////////////////////////////////////////////////////
// Builders
//////////////////////////////////////////////////////////////////

func (t *Enum) WithName(v string) *Enum {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *Enum) WithDescription(v string) *Enum {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *Enum) WithRequired(v bool) *Enum {
	t.CodeGenCommon.WithRequired(v)
	return t
}
