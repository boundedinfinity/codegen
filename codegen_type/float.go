package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeFloat struct {
	SourceMeta
	RenderNamespace
	CodeGenTypeBase
	Min        o.Option[float64]                 `json:"min,omitempty"`
	Max        o.Option[float64]                 `json:"max,omitempty"`
	MultipleOf o.Option[float64]                 `json:"multipleOf,omitempty"`
	Ranges     o.Option[[]CodeGenTypeFloatRange] `json:"ranges,omitempty"`
}

type CodeGenTypeFloatRange struct {
	Min o.Option[float64] `json:"min,omitempty"`
	Max o.Option[float64] `json:"max,omitempty"`
}

func (t CodeGenTypeFloat) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

func (t CodeGenTypeFloat) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Float
}

func (t CodeGenTypeFloat) ValidateSchema() error {
	if err := validateMinMax("float", t.Min, t.Max); err != nil {
		return err
	}

	return nil
}

var _ CodeGenType = &CodeGenTypeFloat{}
