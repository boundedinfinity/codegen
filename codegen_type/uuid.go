package codegen_type

// // https://ihateregex.io/expr/uuid

// type CodeGenTypeUuid struct {
// 	SourceMeta
// 	RenderNamespace
// 	CodeGenTypeBase
// 	CaseSensitive o.Option[bool]                     `json:"caseSensitive,omitempty"`
// 	Version       o.Option[uuid_version.UuidVersion] `json:"version,omitempty"`
// }

// func (t CodeGenTypeUuid) HasValidation() bool {
// 	return t.CaseSensitive.Defined()
// }

// func (t CodeGenTypeUuid) SchemaType() codegen_type_id.CodgenTypeId {
// 	return codegen_type_id.Uuid
// }

// func (t CodeGenTypeUuid) ValidateSchema() error {
// 	return nil
// }

// var _ CodeGenType = &CodeGenTypeUuid{}
