package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Common Type
//////////////////////////////////////////////////////////////////

type codeGenCommon struct {
	// Type ID is the type of this type definition.
	ID optioner.Option[string] `json:"id,omitempty"`

	// QName is the qualified name of this type definition.
	QName optioner.Option[string] `json:"q-name,omitempty"`

	//Name The unqualified (or base) name for this type
	Name optioner.Option[string] `json:"name,omitempty"`

	// The package for this type
	Package optioner.Option[string] `json:"package,omitempty"`

	// Description description of the type
	Description optioner.Option[string] `json:"description,omitempty"`

	// Required true if this types is required, false otherwise
	Required optioner.Option[bool] `json:"required,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t codeGenCommon) GetQName() optioner.Option[string] {
	return t.QName
}

func (t codeGenCommon) GetName() optioner.Option[string] {
	return t.Name
}

func (t codeGenCommon) GetPackage() optioner.Option[string] {
	return t.Package
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

func setO[T any, V any](t T, c *optioner.Option[V], n V) T {
	*c = optioner.Some(n)
	return t
}
