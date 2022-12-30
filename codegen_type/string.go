package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeString struct {
	CodeGenTypeBase
	Min                   o.Option[int]    `json:"min,omitempty"`
	Max                   o.Option[int]    `json:"max,omitempty"`
	IntegerMin            o.Option[int]    `json:"integer-min,omitempty"`
	IntegerMax            o.Option[int]    `json:"integer-max,omitempty"`
	LetterMin             o.Option[int]    `json:"letter-min,omitempty"`
	LetterMax             o.Option[int]    `json:"letter-max,omitempty"`
	LowerCaseMin          o.Option[int]    `json:"lower-case-min,omitempty"`
	LowerCaseMax          o.Option[int]    `json:"lower-case-max,omitempty"`
	UpperCaseMin          o.Option[int]    `json:"upper-case-min,omitempty"`
	UpperCaseMax          o.Option[int]    `json:"upper-case-max,omitempty"`
	SymbolMin             o.Option[int]    `json:"symbol-min,omitempty"`
	SymbolMax             o.Option[int]    `json:"symbol-max,omitempty"`
	Regex                 o.Option[string] `json:"regex,omitempty"`
	RegexErrorDescription o.Option[string] `json:"regex-error-description,omitempty"`
}

func (t CodeGenTypeString) HasValidation() bool {
	return o.FirstOf(
		t.Min, t.Max,
		t.IntegerMin, t.IntegerMax,
		t.LetterMin, t.LetterMax,
		t.LowerCaseMin, t.LowerCaseMax,
		t.UpperCaseMax, t.UpperCaseMin,
		t.SymbolMin, t.SymbolMax,
	).Defined() || o.FirstOf(
		t.Regex,
	).Defined()
}

func (t CodeGenTypeString) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.String
}

func (t CodeGenTypeString) ValidateSchema() error {
	if err := validateMinMax("length", t.Min, t.Max); err != nil {
		return err
	}

	if err := validateMinMax("letter count", t.Min, t.Max); err != nil {
		return err
	}

	if err := validateMinMax("lower case count", t.Min, t.Max); err != nil {
		return err
	}

	if err := validateMinMax("upper case count", t.Min, t.Max); err != nil {
		return err
	}

	if err := validateMinMax("symbol count", t.Min, t.Max); err != nil {
		return err
	}

	return nil
}

var _ CodeGenType = &CodeGenTypeString{}
