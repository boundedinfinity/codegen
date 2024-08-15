package model

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Enum
//////////////////////////////////////////////////////////////////

type CodeGenEnum struct {
	CodeGenCommon
	Values optioner.Option[[]CodeGenEnumItem] `json:"values,omitempty"`
}

func (this CodeGenEnum) Schema() string {
	return "enum"
}

var _ CodeGenSchema = &CodeGenEnum{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (this CodeGenEnum) HasValidation() bool {
	return this.CodeGenCommon.HasValidation()
}

func (this CodeGenEnum) Validate() error {
	if err := this.CodeGenCommon.Validate(); err != nil {
		return err
	}

	for i, value := range this.Values.Get() {
		if err := value.Validate(); err != nil {
			return errors.Join(fmt.Errorf("value[%v]", i))
		}
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (this *CodeGenEnum) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId      string `json:"type"`
		CodeGenEnum `json:",inline"`
	}{
		TypeId:      this.Schema(),
		CodeGenEnum: *this,
	}

	return marshalCodeGenType(dto)
}

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

func (this CodeGenEnumItem) Validate() error {
	if this.Name.Empty() || this.Items.Empty() || len(this.Items.Get()) < 1 {
		return errCodeGenEnumValueMustBeDefined
	}

	return nil
}
