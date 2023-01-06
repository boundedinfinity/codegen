package codegen_type

import (
	"encoding/json"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenProjectOperation struct {
	SourceMeta
	RenderNamespace
	Name        o.Option[string]      `json:"name,omitempty"`
	Description o.Option[string]      `json:"description,omitempty"`
	Input       o.Option[CodeGenType] `json:"input,omitempty"`
	Output      o.Option[CodeGenType] `json:"output,omitempty"`
}

var _ LoaderContext = &CodeGenProjectOperation{}

type marshalOperation struct {
	SourceMeta
	RenderNamespace
	Name        o.Option[string]          `json:"name,omitempty"`
	Description o.Option[string]          `json:"description,omitempty"`
	Input       o.Option[json.RawMessage] `json:"input,omitempty"`
	Output      o.Option[json.RawMessage] `json:"output,omitempty"`
}

func (t *CodeGenProjectOperation) UnmarshalJSON(data []byte) error {
	var d marshalOperation

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	t.SourceMeta = d.SourceMeta
	t.RenderNamespace = d.RenderNamespace
	t.Name = d.Name
	t.Description = d.Description

	if d.Input.Defined() {
		var typ CodeGenType

		if err := UnmarshalJson(d.Input.Get(), &typ); err != nil {
			return err
		}

		t.Input = o.Some(typ)
	}

	if d.Output.Defined() {
		var typ CodeGenType

		if err := UnmarshalJson(d.Output.Get(), &typ); err != nil {
			return err
		}

		t.Output = o.Some(typ)
	}

	return nil
}
