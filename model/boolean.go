package model

import "encoding/json"

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	CodeGenCommon
}

func (t CodeGenBoolean) CodeGenId() string {
	return "boolean"
}

var _ CodeGenType = &CodeGenBoolean{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenBoolean) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (t *CodeGenBoolean) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"codegen-id"`
		CodeGenBoolean `json:",inline"`
	}{
		TypeId:         t.CodeGenId(),
		CodeGenBoolean: *t,
	}

	return json.Marshal(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewBoolean() *CodeGenBoolean {
	return &CodeGenBoolean{}
}

func (t *CodeGenBoolean) WithSchemaId(v string) *CodeGenBoolean {
	t.CodeGenCommon.WithTypeId(v)
	return t
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
