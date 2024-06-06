package model

import "encoding/json"

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

func (t *CodeGenBoolean) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"type-id"`
		CodeGenBoolean `json:",inline"`
	}{
		TypeId:         t.TypeId(),
		CodeGenBoolean: *t,
	}

	return json.Marshal(dto)
}

//////////////////////////////////////////////////////////////////
// Builders
//////////////////////////////////////////////////////////////////

func NewBoolean() *CodeGenBoolean {
	return &CodeGenBoolean{}
}

func (t *CodeGenBoolean) WithName(v string) *CodeGenBoolean {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *CodeGenBoolean) WithDescription(v string) *CodeGenBoolean {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *CodeGenBoolean) WithRequired(v bool) *CodeGenBoolean {
	t.CodeGenCommon.WithRequired(v)
	return t
}

func (t *CodeGenBoolean) WithDefault(v CodeGenBoolean) *CodeGenBoolean {
	t.CodeGenCommon.WithDefault(&v)
	return t
}
