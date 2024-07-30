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

	Package optioner.Option[string] `json:"-"`

	ImportPath optioner.Option[string] `json:"-"`

	// Description description of the type
	Description optioner.Option[string] `json:"description,omitempty"`

	// Required true if this types is required, false otherwise
	Required optioner.Option[bool] `json:"required,omitempty"`

	// JsonName name used for serialization
	JsonName optioner.Option[string] `json:"json-name,omitempty"`

	// YamlName name used for serialization
	YamlName optioner.Option[string] `json:"yaml-name,omitempty"`

	// SqlName name used for serialization
	SqlName optioner.Option[string] `json:"sql-name,omitempty"`

	Imports optioner.Option[[]CodeGenImport]
}

type CodeGenImport struct {
	Name    string `json:"name,omitempty"`
	Package string `json:"package,omitempty"`
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
