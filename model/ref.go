package model

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenRef struct {
	Ref           optioner.Option[string] `json:"ref,omitempty"`
	Found         CodeGenType             `json:"-"`
	codeGenCommon `json:",inline,omitempty"`
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
	if err := t.codeGenCommon.Validate(); err != nil {
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

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewRef() *CodeGenRef {
	return &CodeGenRef{}
}

func (t codeGenCommon) NewRefFromType(typ CodeGenType) (CodeGenType, error) {
	if typ.QName().Empty() {
		return nil, errors.New("invalid ref target")
	}

	return NewRef().WithRef(typ.QName().Get()), nil
}

func (t *CodeGenRef) WithName(v string) *CodeGenRef {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenRef) WithDescription(v string) *CodeGenRef {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenRef) WithRequired(v bool) *CodeGenRef {
	t.codeGenCommon.withRequired(v)
	return t
}

// func (t *CodeGenRef) WithDefault(v CodeGenRef) *CodeGenRef {
// 	t.codeGenCommon.withDefault(&v)
// 	return t
// }

// func (t *CodeGenRef) WithEager(v bool) *CodeGenRef {
// 	t.codeGenCommon.withEager(v)
// 	return t
// }

func (t *CodeGenRef) WithRef(v string) *CodeGenRef {
	t.Ref = optioner.OfZero(v)
	return t
}
