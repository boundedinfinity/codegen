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
	CodeGenCommon
	Properties []CodeGenType `json:"properties"`
}

func (t CodeGenObject) BaseType() string {
	return "object"
}

var _ CodeGenType = &CodeGenObject{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenObject) HasValidation() bool {
	return t.Common().HasValidation()
}

func (t CodeGenObject) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
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

	return json.Marshal(dto)
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

func (t *CodeGenObject) WithSchemaId(v string) *CodeGenObject {
	t.CodeGenCommon.WithQName(v)
	return t
}

func (t *CodeGenObject) WithName(v string) *CodeGenObject {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *CodeGenObject) WithDescription(v string) *CodeGenObject {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *CodeGenObject) WithRequired(v bool) *CodeGenObject {
	t.CodeGenCommon.WithRequired(v)
	return t
}

func (t *CodeGenObject) WithDefault(v CodeGenObject) *CodeGenObject {
	t.CodeGenCommon.WithDefault(&v)
	return t
}

func (t *CodeGenObject) WithProperties(v ...CodeGenType) *CodeGenObject {
	t.Properties = append(t.Properties, v...)
	return t
}
