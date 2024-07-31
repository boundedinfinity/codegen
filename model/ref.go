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
