package canonical

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

type Canonical interface {
	// CType() conical_type.ConicalType
	HasValidation() bool
}

type CanonicalBase struct {
	Id          o.Option[string] `json:"id,omitempty" yaml:"id,omitempty"`
	Name        o.Option[string] `json:"name,omitempty" yaml:"name,omitempty"`
	Source      string           `json:"source,omitempty" yaml:"source,omitempty"`
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Imported    o.Option[bool]   `json:"imported,omitempty" yaml:"imported,omitempty"`
	Public      o.Option[bool]   `json:"public,omitempty" yaml:"public,omitempty"`
	Required    o.Option[bool]   `json:"required,omitempty" yaml:"required,omitempty"`
}

func (t CanonicalBase) HasValidation() bool {
	return false
}

func (t *CanonicalBase) Merge(o CanonicalBase) bool {
	t.Description = o.Description
	t.Id = o.Id
	t.Imported = o.Imported
	t.Name = o.Name
	t.Required = o.Required
	t.Source = o.Source

	return false
}
