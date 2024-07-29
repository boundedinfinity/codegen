package model

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Common Type
//////////////////////////////////////////////////////////////////

type CodeGenCommon struct {
	//Id is the type of this type definition.
	Id optioner.Option[string] `json:"id,omitempty"`

	//Name The unqualified (or base) name for this type
	Name optioner.Option[string] `json:"name,omitempty"`

	//Package The unqualified (or base) name for this type
	Package optioner.Option[string] `json:"package,omitempty"`

	// Description description of the type
	Description optioner.Option[string] `json:"description,omitempty"`

	// Required true if this types is required, false otherwise
	Required optioner.Option[bool] `json:"required,omitempty"`

	// AdditionalValidation enables
	AdditionalValidation optioner.Option[bool] `json:"additional-validation,omitempty"`

	// JsonName name used for serialization
	JsonName optioner.Option[bool] `json:"json-name,omitempty"`

	// YamlName name used for serialization
	YamlName optioner.Option[bool] `json:"yaml-name,omitempty"`

	// SqlName name used for serialization
	SqlName optioner.Option[bool] `json:"sql-name,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t *CodeGenCommon) Common() *CodeGenCommon {
	return t
}

func (t CodeGenCommon) Validate() error {
	return nil
}

func (t CodeGenCommon) HasValidation() bool {
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
