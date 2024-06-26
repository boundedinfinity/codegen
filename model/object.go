package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenObject struct {
	codeGenCommon
	Properties []CodeGenType `json:"properties"`
}

var _ CodeGenType = &CodeGenObject{}

func (t CodeGenObject) BaseType() string {
	return "object"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenObject) HasValidation() bool {
	return t.Common().HasValidation()
}

func (t CodeGenObject) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	for i, prop := range t.Properties {
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
		TypeId        string `json:"base-type"`
		CodeGenObject `json:",inline"`
	}{
		TypeId:        t.BaseType(),
		CodeGenObject: *t,
	}

	return marshalCodeGenType(dto)
}

func (t *CodeGenObject) UnmarshalJSON(data []byte) error {
	dto := struct {
		codeGenCommon
		Properties []json.RawMessage `json:"properties"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.codeGenCommon = dto.codeGenCommon
	}

	for i, rawProp := range dto.Properties {
		if prop, err := UnmarshalCodeGenType(rawProp); err != nil {
			return errors.Join(fmt.Errorf("property[%v]", i), err)
		} else {
			t.Properties = append(t.Properties, prop)
		}
	}

	return nil
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewObject() *CodeGenObject {
	return &CodeGenObject{
		Properties: []CodeGenType{},
	}
}

func (t *CodeGenObject) WithQName(v string) *CodeGenObject {
	t.codeGenCommon.withQName(v)
	return t
}

func (t *CodeGenObject) WithName(v string) *CodeGenObject {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenObject) WithDescription(v string) *CodeGenObject {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenObject) WithRequired(v bool) *CodeGenObject {
	t.codeGenCommon.withRequired(v)
	return t
}

// func (t *CodeGenObject) WithDefault(v CodeGenObject) *CodeGenObject {
// 	t.codeGenCommon.withDefault(&v)
// 	return t
// }

func (t *CodeGenObject) WithProperties(v ...CodeGenType) *CodeGenObject {
	t.Properties = append(t.Properties, v...)
	return t
}
