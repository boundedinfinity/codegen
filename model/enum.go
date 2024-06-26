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

type CodeGenEnumValue struct {
	Name        optioner.Option[string] `json:"name,omitempty"`
	Value       optioner.Option[string] `json:"value,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	errCodeGenEnumValueMustBeDefined = errorer.New("name or value must be defined")
)

func (t CodeGenEnumValue) Validate() error {
	if t.Name.Empty() && t.Value.Empty() {
		return errCodeGenEnumValueMustBeDefined
	}

	return nil
}

///////////////////////////////////////////////////////////////////
// Enum
//////////////////////////////////////////////////////////////////

type CodeGenEnum struct {
	codeGenCommon
	Values []CodeGenEnumValue `json:"values,omitempty"`
}

func (t CodeGenEnum) BaseType() string {
	return "enum"
}

var _ CodeGenType = &CodeGenEnum{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenEnum) HasValidation() bool {
	return t.Common().HasValidation()
}

func (t CodeGenEnum) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
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
		TypeId      string `json:"base-type"`
		CodeGenEnum `json:",inline"`
	}{
		TypeId:      t.BaseType(),
		CodeGenEnum: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *CodeGenEnum) WithQName(v string) *CodeGenEnum {
	t.codeGenCommon.withQName(v)
	return t
}

func (t *CodeGenEnum) WithName(v string) *CodeGenEnum {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenEnum) WithDescription(v string) *CodeGenEnum {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenEnum) WithRequired(v bool) *CodeGenEnum {
	t.codeGenCommon.withRequired(v)
	return t
}
