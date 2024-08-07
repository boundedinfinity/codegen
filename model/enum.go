package model

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// EnumValue
//////////////////////////////////////////////////////////////////

type CodeGenEnumItem struct {
	Name        optioner.Option[string]   `json:"name,omitempty"`
	Items       optioner.Option[[]string] `json:"items,omitempty"`
	Description optioner.Option[string]   `json:"description,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	errCodeGenEnumValueMustBeDefined = errorer.New("name or value must be defined")
)

func (t CodeGenEnumItem) Validate() error {
	if t.Name.Empty() || t.Items.Empty() || len(t.Items.Get()) < 1 {
		return errCodeGenEnumValueMustBeDefined
	}

	return nil
}

///////////////////////////////////////////////////////////////////
// Enum
//////////////////////////////////////////////////////////////////

type CodeGenEnum struct {
	CodeGenCommon
	Values optioner.Option[[]CodeGenEnumItem] `json:"values,omitempty"`
}

func (t CodeGenEnum) Schema() string {
	return "enum"
}

var _ CodeGenSchema = &CodeGenEnum{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenEnum) HasValidation() bool {
	return t.CodeGenCommon.HasValidation()
}

func (t CodeGenEnum) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	for i, value := range t.Values.Get() {
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
		TypeId      string `json:"type"`
		CodeGenEnum `json:",inline"`
	}{
		TypeId:      t.Schema(),
		CodeGenEnum: *t,
	}

	return marshalCodeGenType(dto)
}
