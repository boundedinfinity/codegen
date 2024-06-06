package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Interface
//////////////////////////////////////////////////////////////////

type CodeGenType interface {
	TypeId() string
	Validate() error
}

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenCommon struct {
	// Name is the name of the type.
	Name optioner.Option[string] `json:"name,omitempty"`

	// Description description of the type
	Description optioner.Option[string] `json:"description,omitempty"`

	// Required true if this types is required, false otherwise
	Required optioner.Option[bool] `json:"required,omitempty"`

	// Default is the default value of this type if it's not set
	// Note that this value is mutually exclusive with the Required option.
	Default optioner.Option[CodeGenType] `json:"default,omitempty"`

	// Ref inherits all properties from the base type.
	// Any items can be overridden in this type.
	Ref optioner.Option[string] `json:"ref,omitempty"`

	// Eager will load this type if it's containted inside another type.
	Eager optioner.Option[bool] `json:"eager,omitempty"`

	// Package is the language pack designation used during code generation.
	//  This will be translated into a language appropriate formatted name.
	Package optioner.Option[string] `json:"package,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenCommon) Validate() error {
	return nil
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *CodeGenCommon) WithName(v string) *CodeGenCommon {
	t.Name = optioner.OfZero(v)
	return t
}

func (t *CodeGenCommon) WithDescription(v string) *CodeGenCommon {
	t.Description = optioner.OfZero(v)
	return t
}

func (t *CodeGenCommon) WithRequired(v bool) *CodeGenCommon {
	t.Required = optioner.OfZero(v)
	return t
}

func (t *CodeGenCommon) WithDefault(v CodeGenType) *CodeGenCommon {
	t.Default = optioner.OfZero(v)
	return t
}

func (t *CodeGenCommon) WithEager(v bool) *CodeGenCommon {
	t.Eager = optioner.OfZero(v)
	return t
}
