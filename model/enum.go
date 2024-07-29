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

func (t CodeGenEnum) GetType() string {
	return "enum"
}

var _ CodeGenType = &CodeGenEnum{}

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
		TypeId      string `json:"base-type"`
		CodeGenEnum `json:",inline"`
	}{
		TypeId:      t.GetType(),
		CodeGenEnum: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func BuildEnum() EnumBuilder {
	return &codeGenEnumBuilder{}
}

type codeGenEnumBuilder struct {
	obj CodeGenEnum
}

var _ EnumBuilder = &codeGenEnumBuilder{}

// Ref implements EnumBuilder.
func (t *codeGenEnumBuilder) Ref() RefBuilder {
	panic("unimplemented")
}

// Build implements EnumBuilder.
func (t *codeGenEnumBuilder) Build() *CodeGenEnum {
	return &t.obj
}

// Description implements EnumBuilder.
func (t *codeGenEnumBuilder) Description(v string) EnumBuilder {
	return setO(t, &t.obj.Description, v)
}

// Items implements EnumBuilder.
func (t *codeGenEnumBuilder) Values(v ...CodeGenEnumItem) EnumBuilder {
	return setO(t, &t.obj.Values, v)
}

// Name implements EnumBuilder.
func (t *codeGenEnumBuilder) Name(v string) EnumBuilder {
	return setO(t, &t.obj.Name, v)
}

// Package implements EnumBuilder.
func (t *codeGenEnumBuilder) Package(v string) EnumBuilder {
	return setO(t, &t.obj.Package, v)
}

// Id implements EnumBuilder.
func (t *codeGenEnumBuilder) Id(v string) EnumBuilder {
	return setO(t, &t.obj.Id, v)
}

// Required implements EnumBuilder.
func (t *codeGenEnumBuilder) Required(v bool) EnumBuilder {
	return setO(t, &t.obj.Required, v)
}
