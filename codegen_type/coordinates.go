package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

// https://ihateregex.io/expr/lat-long

type CodeGenTypeCoordinate struct {
	CodeGenTypeBase
}

func (t CodeGenTypeCoordinate) HasValidation() bool {
	return true
}

func (t CodeGenTypeCoordinate) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Coordinate
}

func (t CodeGenTypeCoordinate) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypeCoordinate{}
