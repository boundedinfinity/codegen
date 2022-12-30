package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// (?:(?:31(\/|-|\.)(?:0?[13578]|1[02]))\1|(?:(?:29|30)(\/|-|\.)(?:0?[13-9]|1[0-2])\2))(?:(?:1[6-9]|[2-9]\d)?\d{2})$|^(?:29(\/|-|\.)0?2\3(?:(?:(?:1[6-9]|[2-9]\d)?(?:0[48]|[2468][048]|[13579][26])|(?:(?:16|[2468][048]|[3579][26])00))))$|^(?:0?[1-9]|1\d|2[0-8])(\/|-|\.)(?:(?:0?[1-9])|(?:1[0-2]))\4(?:(?:1[6-9]|[2-9]\d)?\d{2})
// https://ihateregex.io/expr/date

type CodeGenTypeTime struct {
	CodeGenTypeBase
	Before o.Option[CodeGenTypeDate]     `json:"before,omitempty"`
	After  o.Option[CodeGenTypeDate]     `json:"after,omitempty"`
	Within o.Option[CodeGenTypeDuration] `json:"within,omitempty"`
	Ahead  o.Option[CodeGenTypeDuration] `json:"minimum,omitempty"`
	Behind o.Option[CodeGenTypeDuration] `json:"maximum,omitempty"`
}

func (t CodeGenTypeTime) HasValidation() bool {
	return t.Before.Defined() || t.After.Defined() || t.Ahead.Defined() ||
		t.Behind.Defined() || t.Within.Defined()
}

func (t CodeGenTypeTime) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Time
}

func (t CodeGenTypeTime) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypeTime{}
