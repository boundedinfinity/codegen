package model

import (
	"encoding/json"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenRef struct {
	Ref           optioner.Option[string] `json:"ref,omitempty"`
	Found         CodeGenType             `json:"-"`
	CodeGenCommon `json:",inline,omitempty"`
}

func (t CodeGenRef) BaseType() string {
	return "ref"
}

var _ CodeGenType = &CodeGenRef{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenRef) HashValidation() bool {
	return false
}

func (t CodeGenRef) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if t.Found != nil {
		if err := t.Found.Validate(); err != nil {
			return err
		}
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenRef) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId string `json:"base-type"`
		CodeGenRef
	}{
		TypeId:     t.BaseType(),
		CodeGenRef: *t,
	}

	return json.Marshal(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewRef() *CodeGenRef {
	return &CodeGenRef{}
}

func (t CodeGenCommon) NewRefFromType(typ CodeGenType) (CodeGenType, error) {
	if typ.QName().Empty() {
		return nil, errors.New("invalid ref target")
	}

	return NewRef().WithRef(typ.QName().Get()), nil
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
