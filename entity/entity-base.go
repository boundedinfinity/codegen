package entity

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/errorer"
)

type Entity interface {
	Type() EntityType
	GetQName() string
	Marshalable
	Validatable
	ToJsonSchema() ([]byte, error)
	ToJsonSchemaIndent() ([]byte, error)
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

func (t entityBase) Validate() error {
	if err := t.common.Validate(); err != nil {
		return err
	}

	if _, ok := type2StringMap[t.entityType]; !ok {
		return ErrEntityInvalidType.WithValue(t.entityType)
	}

	if t.qname == "" {
		return ErrEntityMissingQName
	}

	return nil
}

func (t entityBase) HasValidation() bool {
	return t.common.HasValidation() && t.required && t.additionalValidation
}

func (t entityBase) ToMap() (map[string]any, error) {
	data, err := t.common.ToMap()

	if err != nil {
		return data, err
	}

	if t.entityType != noEntityType {
		data["entity-type"] = type2StringMap[t.entityType]
	}

	bparam(data, "additional-validation", t.additionalValidation)
	bparam(data, "required", t.required)

	if t.defaultValue != nil && len(t.defaultValue) > 0 {
		data["default"] = t.defaultValue
	}

	return data, nil
}

func (t entityBase) Type() EntityType { return t.entityType }
