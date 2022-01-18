package main

import (
	"github.com/boundedinfinity/jsonschema"
)

type SomethingValidator struct {
	schema jsonschema.JsonSchmea
}

var ()

func (t SomethingValidator) Validate(v int) error {
	if t.schema.MultipleOf.IsDefined() {
		if v%int(t.schema.MultipleOf.Get()) != 0 {
			return jsonschema.ErrNotMultipleOff(v, t.schema.MaxContains.Get())
		}
	}

	if t.schema.Minimum.IsDefined() {
		if v < int(t.schema.Minimum.Get()) {
			return jsonschema.ErrIsLessThanf(v, int(t.schema.MinContains.Get()))
		}
	}

	if t.schema.ExclusiveMinimum.IsDefined() {
		if v <= int(t.schema.ExclusiveMinimum.Get()) {
			return jsonschema.ErrIsLessThanOrEqualTof(v, int(t.schema.MinContains.Get()))
		}
	}

	if t.schema.Maximum.IsDefined() {
		if v > int(t.schema.Maximum.Get()) {
			return jsonschema.ErrIsGreaterThanf(v, int(t.schema.Maximum.Get()))
		}
	}

	if t.schema.ExclusiveMaximum.IsDefined() {
		if v >= int(t.schema.ExclusiveMaximum.Get()) {
			return jsonschema.ErrIsGreaterThanOrEqualTof(v, int(t.schema.ExclusiveMaximum.Get()))
		}
	}

	return nil
}
