package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeInteger struct {
	CodeGenTypeBase
	Min        o.Option[int64]                     `json:"min,omitempty"`
	Max        o.Option[int64]                     `json:"max,omitempty"`
	Ranges     o.Option[[]CodeGenTypeIntegerRange] `json:"ranges,omitempty"`
	MultipleOf o.Option[int64]                     `json:"multipleOf,omitempty"`
}

type CodeGenTypeIntegerRange struct {
	Min o.Option[float64] `json:"min,omitempty"`
	Max o.Option[float64] `json:"max,omitempty"`
}

func (t CodeGenTypeInteger) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

func (t CodeGenTypeInteger) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Integer
}

func (t CodeGenTypeInteger) ValidateSchema() error {
	if err := validateMinMax("integer", t.Min, t.Max); err != nil {
		return err
	}

	return nil
}

var _ CodeGenType = &CodeGenTypeInteger{}
