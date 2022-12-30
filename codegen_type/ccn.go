package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

// https://ihateregex.io/expr/credit-card/

type CodeGenTypeCreditCardNumber struct {
	CodeGenTypeBase
}

func (t CodeGenTypeCreditCardNumber) HasValidation() bool {
	return true
}

func (t CodeGenTypeCreditCardNumber) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.CreditCardNumber
}

func (t CodeGenTypeCreditCardNumber) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypeCreditCardNumber{}
