package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenRef struct {
	CodeGenCommon `json:",inline,omitempty"`
	Ref           optioner.Option[string] `json:"ref,omitempty"`
	Resolved      CodeGenType             `json:"-"`
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
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if t.Resolved != nil {
		if err := t.Resolved.Validate(); err != nil {
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

type codeGenRefBuilder struct {
	obj CodeGenRef
}

var _ RefBuilder = &codeGenRefBuilder{}

func BuildRef() RefBuilder {
	return &codeGenRefBuilder{}
}

func BuildRefWithResolved(typ CodeGenType) RefBuilder {
	return BuildRef().Resolved(typ).Ref(typ.Common().Id.Get())
}

// Build implements RefBuilder.
func (t *codeGenRefBuilder) Build() *CodeGenRef {
	return &t.obj
}

func (t *codeGenRefBuilder) Resolved(typ CodeGenType) RefBuilder {
	return setV(t, &t.obj.Resolved, typ)
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

// Id implements RefBuilder.
func (t *codeGenRefBuilder) Id(v string) RefBuilder {
	return setO(t, &t.obj.Id, v)
}

// Required implements RefBuilder.
func (t *codeGenRefBuilder) Required(v bool) RefBuilder {
	return setO(t, &t.obj.Required, v)
}
