package model

import (
	"encoding/json"
	"regexp"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenString struct {
	CodeGenCommon
	Min   optioner.Option[int]    `json:"min,omitempty"`
	Max   optioner.Option[int]    `json:"max,omitempty"`
	Regex optioner.Option[string] `json:"regex,omitempty"`
}

func (t CodeGenString) CodeGenId() string {
	return "string"
}

var _ CodeGenType = &CodeGenString{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenString) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if t.Regex.Defined() {
		if _, err := regexp.Compile(t.Regex.Get()); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (t *CodeGenString) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId        string `json:"codegen-id"`
		CodeGenString `json:",inline"`
	}{
		TypeId:        t.CodeGenId(),
		CodeGenString: *t,
	}

	return json.Marshal(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewString() *CodeGenString {
	return &CodeGenString{}
}

func (t *CodeGenString) WithSchemaId(v string) *CodeGenString {
	t.CodeGenCommon.WithTypeId(v)
	return t
}

func (t *CodeGenString) WithName(v string) *CodeGenString {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *CodeGenString) WithDescription(v string) *CodeGenString {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *CodeGenString) WithRequired(v bool) *CodeGenString {
	t.CodeGenCommon.WithRequired(v)
	return t
}

func (t *CodeGenString) WithDefault(v CodeGenString) *CodeGenString {
	t.CodeGenCommon.WithDefault(&v)
	return t
}

func (t *CodeGenString) WithMin(v int) *CodeGenString {
	t.Min = optioner.OfZero(v)
	return t
}

func (t *CodeGenString) WithMax(v int) *CodeGenString {
	t.Max = optioner.OfZero(v)
	return t
}

func (t *CodeGenString) WithRegex(v string) *CodeGenString {
	t.Regex = optioner.OfZero(v)
	return t
}
