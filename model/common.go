package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
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
	Common() *codeGenCommon
}

///////////////////////////////////////////////////////////////////
// Common Type
//////////////////////////////////////////////////////////////////

type codeGenCommon struct {
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

	// // Default is the default value of this type if it's not set
	// // Note that this value is mutually exclusive with the Required option.
	// Default optioner.Option[CodeGenType] `json:"default,omitempty"`

	// // Eager will load this type if it's containted inside another type.
	// Eager optioner.Option[bool] `json:"eager,omitempty"`

	CodeGenMeta
}

func (t *codeGenCommon) QName() optioner.Option[string] {
	return t.QName_
}

func (t *codeGenCommon) Meta() *CodeGenMeta {
	return &t.CodeGenMeta
}

func (t *codeGenCommon) Common() *codeGenCommon {
	return t
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t codeGenCommon) Validate() error {
	if err := t.Meta().Validate(); err != nil {
		return err
	}

	return nil
}

func (t codeGenCommon) HasValidation() bool {
	// return t.Required.Defined() &&
	// 	t.Meta().HasValidation()
	return t.Meta().HasValidation()
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *codeGenCommon) withQName(v string) *codeGenCommon {
	t.QName_ = optioner.Some(v)
	t.withName(pather.Paths.Base(v))
	t.withPackage(pather.Paths.Base(v))
	return t
}

func (t *codeGenCommon) withPackage(v string) *codeGenCommon {
	t.Package = optioner.Some(v)
	return t
}

func (t *codeGenCommon) withName(v string) *codeGenCommon {
	t.Name = optioner.Some(v)
	return t
}

func (t *codeGenCommon) withDescription(v string) *codeGenCommon {
	t.Description = optioner.Some(v)
	return t
}

func (t *codeGenCommon) withRequired(v bool) *codeGenCommon {
	t.Required = optioner.Some(v)
	return t
}

// func (t *codeGenCommon) withDefault(v CodeGenType) *codeGenCommon {
// 	t.Default = optioner.Some(v)
// 	return t
// }

// func (t *codeGenCommon) withEager(v bool) *codeGenCommon {
// 	t.Eager = optioner.Some(v)
// 	return t
// }
