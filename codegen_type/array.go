package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"encoding/json"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeArray struct {
	CodeGenTypeBase
	Items CodeGenType   `json:"items,omitempty"`
	Min   o.Option[int] `json:"min,omitempty"`
	Max   o.Option[int] `json:"max,omitempty"`
}

func (t CodeGenTypeArray) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Array
}

func (t CodeGenTypeArray) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.Items.HasValidation()
}

var _ CodeGenType = &CodeGenTypeArray{}

type marshalArray struct {
	CodeGenTypeBase
	Items json.RawMessage `json:"items,omitempty"`
	Min   o.Option[int]   `json:"min,omitempty"`
	Max   o.Option[int]   `json:"max,omitempty"`
}

func (t *CodeGenTypeArray) UnmarshalJSON(data []byte) error {
	var d marshalArray

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	t.CodeGenTypeBase = d.CodeGenTypeBase
	t.Max = d.Max
	t.Min = d.Min

	var items CodeGenType

	if err := UnmarshalJson(d.Items, &items); err != nil {
		return err
	} else {
		t.Items = items
	}

	return nil
}
