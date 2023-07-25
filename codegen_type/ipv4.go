package codegen_type

// // https://ihateregex.io/expr/ip/

// type CodeGenTypeIpv4 struct {
// 	SourceMeta
// 	RenderNamespace
// 	CodeGenTypeBase
// 	Minimum o.Option[string] `json:"minimum,omitempty"`
// 	Maximum o.Option[string] `json:"maximum,omitempty"`
// 	Mask    o.Option[string] `json:"mask,omitempty"`
// }

// func (t CodeGenTypeIpv4) HasValidation() bool {
// 	return t.Minimum.Defined() || t.Maximum.Defined() || t.Mask.Defined()
// }

// func (t CodeGenTypeIpv4) SchemaType() codegen_type_id.CodgenTypeId {
// 	return codegen_type_id.Ipv4
// }

// func (t CodeGenTypeIpv4) ValidateSchema() error {
// 	return nil
// }

// var _ CodeGenType = &CodeGenTypeIpv4{}
