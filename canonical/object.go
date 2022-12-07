package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"encoding/json"
)

type CanonicalObject struct {
	CanonicalBase
	Properties []Canonical `json:"properties,omitempty" yaml:"properties,omitempty"`
}

func (t CanonicalObject) CType() canonical_type.CanonicalType {
	return canonical_type.Object
}

func (t CanonicalObject) HasValidation() bool {
	for _, property := range t.Properties {
		if property.HasValidation() {
			return true
		}
	}

	return false
}

var _ Canonical = &CanonicalObject{}

type canonicalObject struct {
	CanonicalBase
	Properties []json.RawMessage `json:"properties,omitempty" yaml:"properties,omitempty"`
}

func (t *CanonicalObject) UnmarshalJSON(data []byte) error {
	var d canonicalObject

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	t.CanonicalBase.Merge(d.CanonicalBase)

	for _, property := range d.Properties {
		if p, err := UnmarshalCanonicalSchemaJson(property); err != nil {
			return err
		} else {
			t.Properties = append(t.Properties, p)
		}
	}

	return nil
}
