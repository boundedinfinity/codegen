package codegen_type

// type CodeGenTypeDateTime struct {
// 	SourceMeta
// 	RenderNamespace
// 	CodeGenTypeBase
// 	Before o.Option[CodeGenTypeDate]     `json:"before,omitempty"`
// 	After  o.Option[CodeGenTypeDate]     `json:"after,omitempty"`
// 	Within o.Option[CodeGenTypeDuration] `json:"within,omitempty"`
// 	Ahead  o.Option[CodeGenTypeDuration] `json:"minimum,omitempty"`
// 	Behind o.Option[CodeGenTypeDuration] `json:"maximum,omitempty"`
// }

// func (t CodeGenTypeDateTime) HasValidation() bool {
// 	return t.Before.Defined() || t.After.Defined() || t.Ahead.Defined() ||
// 		t.Behind.Defined() || t.Within.Defined()
// }

// func (t CodeGenTypeDateTime) SchemaType() codegen_type_id.CodgenTypeId {
// 	return codegen_type_id.Datetime
// }

// func (t CodeGenTypeDateTime) ValidateSchema() error {
// 	return nil
// }

// var _ CodeGenType = &CodeGenTypeDateTime{}
