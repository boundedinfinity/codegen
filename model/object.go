package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenObject struct {
	CodeGenCommon
	Properties optioner.Option[[]CodeGenType] `json:"properties"`
}

var _ CodeGenType = &CodeGenObject{}

func (t CodeGenObject) GetType() string {
	return "object"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenObject) HasValidation() bool {
	return t.CodeGenCommon.HasValidation()
}

func (t CodeGenObject) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	for i, prop := range t.Properties.Get() {
		if err := prop.Validate(); err != nil {
			return errors.Join(fmt.Errorf("prop[%v]", i))
		}
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenObject) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId        string `json:"type"`
		CodeGenObject `json:",inline"`
	}{
		TypeId:        t.GetType(),
		CodeGenObject: *t,
	}

	return marshalCodeGenType(dto)
}

func (t *CodeGenObject) UnmarshalJSON(data []byte) error {
	dto := struct {
		CodeGenCommon
		Properties []json.RawMessage `json:"properties"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.CodeGenCommon = dto.CodeGenCommon
	}

	for i, rawProp := range dto.Properties {
		if prop, err := UnmarshalCodeGenType(rawProp); err != nil {
			return errors.Join(fmt.Errorf("property[%v]", i), err)
		} else {
			if t.Properties.Defined() {
				t.Properties = optioner.Some(append(t.Properties.Get(), prop))
			}
		}
	}

	return nil
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func BuildObject() ObjectBuilder {
	return &codeGenObjectBuilder{}
}

type codeGenObjectBuilder struct {
	obj CodeGenObject
}

var _ ObjectBuilder = &codeGenObjectBuilder{}

// Build implements ObjectBuilder.
func (t *codeGenObjectBuilder) Build() *CodeGenObject {
	return &t.obj
}

// Description implements ObjectBuilder.
func (t *codeGenObjectBuilder) Description(v string) ObjectBuilder {
	return SetO(t, &t.obj.Description, v)
}

// Name implements ObjectBuilder.
func (t *codeGenObjectBuilder) Name(v string) ObjectBuilder {
	return SetO(t, &t.obj.Name, v)
}

// Package implements ObjectBuilder.
func (t *codeGenObjectBuilder) Package(v string) ObjectBuilder {
	return SetO(t, &t.obj.Package, v)
}

// Property implements ObjectBuilder.
func (t *codeGenObjectBuilder) Properties(v ...CodeGenType) ObjectBuilder {
	return SetO(t, &t.obj.Properties, v)
}

// Id implements ObjectBuilder.
func (t *codeGenObjectBuilder) Id(v string) ObjectBuilder {
	return SetO(t, &t.obj.Id, v)
}

// Required implements ObjectBuilder.
func (t *codeGenObjectBuilder) Required(v bool) ObjectBuilder {
	return SetO(t, &t.obj.Required, v)
}
