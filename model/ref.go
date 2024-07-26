package model

import (
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

func (t CodeGenRef) GetType() string {
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
		TypeId string `json:"type"`
		CodeGenRef
	}{
		TypeId:     t.GetType(),
		CodeGenRef: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builder
//----------------------------------------------------------------

func RefFromType(typ CodeGenType) (RefBuilder, error) {
	return &codeGenRefBuilder{}, nil
	// if typ.QName().Empty() {
	// 	return nil, errors.New("invalid ref target")
	// }

	// return NewRef().WithRef(typ.QName().Get()), nil
}

type codeGenRefBuilder struct {
	obj CodeGenRef
}

var _ RefBuilder = &codeGenRefBuilder{}

func BuildRef() *codeGenRefBuilder {
	return &codeGenRefBuilder{}
}

// Build implements RefBuilder.
func (t *codeGenRefBuilder) Build() CodeGenRef {
	return t.obj
}

// Ref implements RefBuilder.
func (t *codeGenRefBuilder) Ref(v string) RefBuilder {
	return setO(t, &t.obj.Ref, v)
}

// Description implements RefBuilder.
func (t *codeGenRefBuilder) Description(v string) RefBuilder {
	return setO(t, &t.obj.Description, v)
}

// Name implements RefBuilder.
func (t *codeGenRefBuilder) Name(v string) RefBuilder {
	return setO(t, &t.obj.Name, v)
}

// Package implements RefBuilder.
func (t *codeGenRefBuilder) Package(v string) RefBuilder {
	return setO(t, &t.obj.Package, v)
}

// QName implements RefBuilder.
func (t *codeGenRefBuilder) QName(v string) RefBuilder {
	panic("unimplemented")
}

// Required implements RefBuilder.
func (t *codeGenRefBuilder) Required(v bool) RefBuilder {
	return setO(t, &t.obj.Required, v)
}
