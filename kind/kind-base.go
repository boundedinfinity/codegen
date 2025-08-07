// Package kind contains the enumeration of kind names
package kind

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

type Kind interface {
	GetQName() string
	Marshalable
	Validatable
	AsJsonSchema() (json_schema.JsonSchema, error)
}

type entityBase struct {
	required             bool
	additionalValidation bool
	defaultValue         map[string]any
	common
}

var (
	ErrKindMissingQName = errorer.New("missing q-name")
)

func (this entityBase) Validate() error {
	if err := this.common.Validate(); err != nil {
		return err
	}

	if this.qname == "" {
		return ErrKindMissingQName
	}

	return nil
}

func (this entityBase) HasValidation() bool {
	return this.common.HasValidation() && this.required && this.additionalValidation
}

func (this entityBase) ToMap() (map[string]any, error) {
	data, err := this.common.ToMap()

	if err != nil {
		return data, err
	}

	bparam(data, "additional-validation", this.additionalValidation)
	bparam(data, "required", this.required)

	if this.defaultValue != nil && len(this.defaultValue) > 0 {
		data["default"] = this.defaultValue
	}

	return data, nil
}
