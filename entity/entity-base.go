package entity

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

type Entity interface {
	Type() EntityType
	GetQName() string
	Marshalable
	Validatable
	AsJsonSchema() (json_schema.JsonSchema, error)
}

type entityBase struct {
	entityType           EntityType
	required             bool
	additionalValidation bool
	defaultValue         map[string]any
	common
}

var (
	ErrEntityInvalidType = errorer.Errorf(
		"invalid entity type, must be one of %v",
		strings.Join(asString(validTypes...), ", "),
	)

	ErrEntityMissingQName = errorer.New("missing q-name")
)

func (this entityBase) Validate() error {
	if err := this.common.Validate(); err != nil {
		return err
	}

	if _, ok := type2StringMap[this.entityType]; !ok {
		return ErrEntityInvalidType.WithValue(this.entityType)
	}

	if this.qname == "" {
		return ErrEntityMissingQName
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

	if this.entityType != noEntityType {
		data["entity-type"] = type2StringMap[this.entityType]
	}

	bparam(data, "additional-validation", this.additionalValidation)
	bparam(data, "required", this.required)

	if this.defaultValue != nil && len(this.defaultValue) > 0 {
		data["default"] = this.defaultValue
	}

	return data, nil
}

func (this entityBase) Type() EntityType { return this.entityType }
