package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Interfaces
//////////////////////////////////////////////////////////////////

type CodeGenType interface {
	GetType() string
	GetName() string
	GetPackage() string
	Validate() error
	HasValidation() bool
}

///////////////////////////////////////////////////////////////////
// Common Type
//////////////////////////////////////////////////////////////////

type codeGenCommon struct {
	// Type ID is the type of this type definition.
	QName optioner.Option[string] `json:"q-name,omitempty"`

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
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t codeGenCommon) GetName() string {
	return t.Name.Get()
}

func (t codeGenCommon) GetPackage() string {
	return t.Package.Get()
}

func (t codeGenCommon) Validate() error {
	return nil
}

func (t codeGenCommon) HasValidation() bool {
	return false
}

//----------------------------------------------------------------
// Helpers
//----------------------------------------------------------------

func setV[T any, V any](t T, c *V, n V) T {
	*c = n
	return t
}
