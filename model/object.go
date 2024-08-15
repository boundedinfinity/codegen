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
	Properties optioner.Option[[]CodeGenSchema] `json:"properties"`
}

var _ CodeGenSchema = &CodeGenObject{}

func (_ CodeGenObject) Schema() string {
	return "object"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (this CodeGenObject) HasValidation() bool {
	hasValidation := this.CodeGenCommon.HasValidation()

	if !hasValidation {
		for _, property := range this.Properties.Get() {
			if property.HasValidation() {
				hasValidation = true
				break
			}
		}
	}

	return hasValidation
}

func (this CodeGenObject) Validate() error {
	if err := this.CodeGenCommon.Validate(); err != nil {
		return err
	}

	for i, prop := range this.Properties.Get() {
		if err := prop.Validate(); err != nil {
			return errors.Join(fmt.Errorf("prop[%v]", i))
		}
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (this *CodeGenObject) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId        string `json:"type"`
		CodeGenObject `json:",inline"`
	}{
		TypeId:        this.Schema(),
		CodeGenObject: *this,
	}

	return marshalCodeGenType(dto)
}

func (this *CodeGenObject) UnmarshalJSON(data []byte) error {
	dto := struct {
		CodeGenCommon
		Properties []json.RawMessage `json:"properties"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		this.CodeGenCommon = dto.CodeGenCommon
	}

	if len(dto.Properties) > 0 {
		this.Properties = optioner.Some([]CodeGenSchema{})

		for i, rawProp := range dto.Properties {
			if prop, err := UnmarshalCodeGenType(rawProp); err != nil {
				return errors.Join(fmt.Errorf("property[%v]", i), err)
			} else {
				if this.Properties.Defined() {
					this.Properties = optioner.Some(append(this.Properties.Get(), prop))
				}
			}
		}
	}

	return nil
}
