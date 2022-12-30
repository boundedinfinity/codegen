package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/codegen_type/uuid_version"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// https://ihateregex.io/expr/uuid

type CodeGenTypeUuid struct {
	CodeGenTypeBase
	CaseSensitive o.Option[bool]                     `json:"caseSensitive,omitempty"`
	Version       o.Option[uuid_version.UuidVersion] `json:"version,omitempty"`
}

func (t CodeGenTypeUuid) HasValidation() bool {
	return t.CaseSensitive.Defined()
}

func (t CodeGenTypeUuid) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Uuid
}

func (t CodeGenTypeUuid) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypeUuid{}
