package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeRef struct {
	SourceMeta
	RenderNamespace
	CodeGenTypeBase
	Ref      o.Option[string] `json:"ref,omitempty"`
	Resolved CodeGenType
}

func (t CodeGenTypeRef) HasValidation() bool {
	return true
}

func (t CodeGenTypeRef) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Ref
}

func (t CodeGenTypeRef) ValidateSchema() error {
	return nil
}

func (t CodeGenTypeRef) String() string {
	return t.Ref.Get()
}

var _ CodeGenType = &CodeGenTypeRef{}
