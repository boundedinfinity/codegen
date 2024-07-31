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

	if len(dto.Properties) > 0 {
		t.Properties = optioner.Some([]CodeGenType{})

		for i, rawProp := range dto.Properties {
			if prop, err := UnmarshalCodeGenType(rawProp); err != nil {
				return errors.Join(fmt.Errorf("property[%v]", i), err)
			} else {
				if t.Properties.Defined() {
					t.Properties = optioner.Some(append(t.Properties.Get(), prop))
				}
			}
		}
	}

	return nil
}
