package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// EnumValue
//////////////////////////////////////////////////////////////////

type CodeGenEnumValue struct {
	Name        optioner.Option[string] `json:"name,omitempty"`
	Value       optioner.Option[string] `json:"value,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenEnumValue) Validate() error {
	if t.Name.Empty() && t.Value.Empty() {
		return errors.New("name or value must be defined")
	}

	return nil
}

///////////////////////////////////////////////////////////////////
// Enum
//////////////////////////////////////////////////////////////////

type CodeGenEnum struct {
	CodeGenCommon
	Values []CodeGenEnumValue `json:"values,omitempty"`
}

func (t CodeGenEnum) TypeId() string {
	return "enum"
}

var _ CodeGenType = &CodeGenEnum{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenEnum) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	for i, value := range t.Values {
		if err := value.Validate(); err != nil {
			return errors.Join(fmt.Errorf("value[%v]", i))
		}
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenEnum) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId      string `json:"type-id"`
		CodeGenEnum `json:",inline"`
	}{
		TypeId:      t.TypeId(),
		CodeGenEnum: *t,
	}

	return json.Marshal(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *CodeGenEnum) WithName(v string) *CodeGenEnum {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *CodeGenEnum) WithDescription(v string) *CodeGenEnum {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *CodeGenEnum) WithRequired(v bool) *CodeGenEnum {
	t.CodeGenCommon.WithRequired(v)
	return t
}
