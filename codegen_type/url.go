package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// // https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)
// https://ihateregex.io/expr/url

type CodeGenTypeUrl struct {
	CodeGenTypeBase
	SchemesAllowed o.Option[[]string] `json:"schemes-allowed,omitempty"`
	SchemesDenied  o.Option[[]string] `json:"schemes-denied,omitempty"`
	SchemesRegex   o.Option[[]string] `json:"schemes-regex,omitempty"`
}

func (t CodeGenTypeUrl) HasValidation() bool {
	// Must always be a valid URL
	return true
}

func (t CodeGenTypeUrl) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Url
}

var _ CodeGenType = &CodeGenTypeUrl{}
