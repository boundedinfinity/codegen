package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"encoding/json"
)

type CodeGenTypeObject struct {
	SourceMeta
	RenderNamespace
	CodeGenTypeBase
	Properties []CodeGenType `json:"properties,omitempty"`
}

func (t CodeGenTypeObject) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Object
}

func (t CodeGenTypeObject) HasValidation() bool {
	for _, property := range t.Properties {
		if property.HasValidation() {
			return true
		}
	}

	return false
}

func (t CodeGenTypeObject) ValidateSchema() error {
	for _, property := range t.Properties {
		if err := property.ValidateSchema(); err != nil {
			return err
		}
	}

	return nil
}

var _ CodeGenType = &CodeGenTypeObject{}

type codeGenObject struct {
	CodeGenTypeBase
	Properties []json.RawMessage `json:"properties,omitempty"`
}

func (t *CodeGenTypeObject) UnmarshalJSON(data []byte) error {
	var d codeGenObject

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	t.CodeGenTypeBase = d.CodeGenTypeBase

	for _, property := range d.Properties {
		var p CodeGenType
		if err := UnmarshalJson(property, &p); err != nil {
			return err
		} else {
			t.Properties = append(t.Properties, p)
		}
	}

	return nil
}
