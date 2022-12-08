package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
)

// https://ihateregex.io/expr/credit-card/

type CanonicalCreditCardNumber struct {
	CanonicalBase
}

func (t CanonicalCreditCardNumber) CType() canonical_type.CanonicalType {
	return canonical_type.CreditCardNumber
}

func (t CanonicalCreditCardNumber) HasValidation() bool {
	return true
}

func (t CanonicalCreditCardNumber) SchemaType() canonical_type.CanonicalType {
	return canonical_type.CreditCardNumber
}

var _ Canonical = &CanonicalCreditCardNumber{}
