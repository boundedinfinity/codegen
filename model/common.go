package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Interface
//////////////////////////////////////////////////////////////////

type CodeGenType interface {
	CodeGenId() string
	TypeId() optioner.Option[string]
	Validate() error
	Meta() *CodeGenMeta
	Common() *CodeGenCommon
}

///////////////////////////////////////////////////////////////////
// Common Type
//////////////////////////////////////////////////////////////////

type CodeGenCommon struct {
	// Type ID is the type of this type definition.
	Type_Id optioner.Option[string] `json:"type-id,omitempty"`

	// Name is the name of the type.
	Name optioner.Option[string] `json:"name,omitempty"`

	// Description description of the type
	Description optioner.Option[string] `json:"description,omitempty"`

	// Required true if this types is required, false otherwise
	Required optioner.Option[bool] `json:"required,omitempty"`

	// Default is the default value of this type if it's not set
	// Note that this value is mutually exclusive with the Required option.
	Default optioner.Option[CodeGenType] `json:"default,omitempty"`

	// Eager will load this type if it's containted inside another type.
	Eager optioner.Option[bool] `json:"eager,omitempty"`

	// Package is the language pack designation used during code generation.
	//  This will be translated into a language appropriate formatted name.
	Package optioner.Option[string] `json:"package,omitempty"`

	// QualifiedName
	QualifiedName optioner.Option[string] `json:"-"`

	CodeGenMeta
}

func (t *CodeGenCommon) TypeId() optioner.Option[string] {
	return t.Type_Id
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
	return nil
}

//----------------------------------------------------------------
// Functions
//----------------------------------------------------------------

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *CodeGenCommon) WithTypeId(v string) *CodeGenCommon {
	t.Type_Id = optioner.OfZero(v)
	return t
}

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

///////////////////////////////////////////////////////////////////
// CodeGenMeta
//////////////////////////////////////////////////////////////////

type CodeGenMeta struct {
}

//----------------------------------------------------------------
// Merge
//----------------------------------------------------------------

func (t *CodeGenMeta) Merge(obj CodeGenMeta) error {
	return nil
}

//----------------------------------------------------------------
// Validate
//----------------------------------------------------------------

func (t *CodeGenMeta) Validate() error {
	return nil
}
