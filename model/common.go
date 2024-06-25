package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Interface
//////////////////////////////////////////////////////////////////

type CodeGenType interface {
	BaseType() string
	QName() optioner.Option[string]
	Validate() error
	HasValidation() bool
	Meta() *CodeGenMeta
	Common() *CodeGenCommon
}

///////////////////////////////////////////////////////////////////
// Common Type
//////////////////////////////////////////////////////////////////

type CodeGenCommon struct {
	// Type ID is the type of this type definition.
	QName_ optioner.Option[string] `json:"q-name,omitempty"`

	// Name is the name of the type.
	Name optioner.Option[string] `json:"name,omitempty"`

	// Package
	Package optioner.Option[string] `json:"package,omitempty"`

	// Description description of the type
	Description optioner.Option[string] `json:"description,omitempty"`

	// Required true if this types is required, false otherwise
	Required optioner.Option[bool] `json:"required,omitempty"`

	// Default is the default value of this type if it's not set
	// Note that this value is mutually exclusive with the Required option.
	Default optioner.Option[CodeGenType] `json:"default,omitempty"`

	// Eager will load this type if it's containted inside another type.
	Eager optioner.Option[bool] `json:"eager,omitempty"`

	CodeGenMeta
}

func (t *CodeGenCommon) QName() optioner.Option[string] {
	return t.QName_
}

func (t *CodeGenCommon) Meta() *CodeGenMeta {
	return &t.CodeGenMeta
}

func (t *CodeGenCommon) Common() *CodeGenCommon {
	return t
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenCommon) Validate() error {
	if err := t.Meta().Validate(); err != nil {
		return err
	}

	return nil
}

func (t CodeGenCommon) HasValidation() bool {
	return t.Required.Defined() &&
		t.Meta().HasValidation()
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *CodeGenCommon) WithQName(v string) *CodeGenCommon {
	t.QName_ = optioner.Some(v)
	return t
}

func (t *CodeGenCommon) WithPackage(v string) *CodeGenCommon {
	t.Package = optioner.Some(v)
	return t
}

func (t *CodeGenCommon) WithName(v string) *CodeGenCommon {
	t.Name = optioner.Some(v)
	return t
}

func (t *CodeGenCommon) WithDescription(v string) *CodeGenCommon {
	t.Description = optioner.Some(v)
	return t
}

func (t *CodeGenCommon) WithRequired(v bool) *CodeGenCommon {
	t.Required = optioner.Some(v)
	return t
}

func (t *CodeGenCommon) WithDefault(v CodeGenType) *CodeGenCommon {
	t.Default = optioner.Some(v)
	return t
}

func (t *CodeGenCommon) WithEager(v bool) *CodeGenCommon {
	t.Eager = optioner.Some(v)
	return t
}
