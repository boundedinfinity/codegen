package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"encoding/json"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalArray struct {
	CanonicalBase
	Items Canonical     `json:"items,omitempty" yaml:"items,omitempty"`
	Min   o.Option[int] `json:"min,omitempty" yaml:"min,omitempty"`
	Max   o.Option[int] `json:"max,omitempty" yaml:"max,omitempty"`
}

func (t CanonicalArray) CType() canonical_type.CanonicalType {
	return canonical_type.Array
}

func (t CanonicalArray) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.Items.HasValidation()
}

func (t CanonicalArray) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Array
}

var _ Canonical = &CanonicalArray{}

type canonicalArray struct {
	CanonicalBase
	Items json.RawMessage `json:"items,omitempty" yaml:"items,omitempty"`
	Min   o.Option[int]   `json:"min,omitempty" yaml:"min,omitempty"`
	Max   o.Option[int]   `json:"max,omitempty" yaml:"max,omitempty"`
}

func (t *CanonicalArray) UnmarshalJSON(data []byte) error {
	var d canonicalArray

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	t.CanonicalBase.Merge(d.CanonicalBase)
	t.Max = d.Max
	t.Min = d.Min

	if i, err := UnmarshalCanonicalSchemaJson(d.Items); err != nil {
		return err
	} else {
		t.Items = i
	}

	return nil
}
