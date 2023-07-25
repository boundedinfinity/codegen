package model

// type Operation struct {
// 	Meta
// 	Name        o.Option[string] `json:"name,omitempty"`
// 	Description o.Option[string] `json:"description,omitempty"`
// 	Input       o.Option[Type]   `json:"input,omitempty"`
// 	Output      o.Option[Type]   `json:"output,omitempty"`
// }

// func (t Operation) TypeId() type_id.TypeId {
// 	return type_id.Operation
// }

// var _ Type = &Operation{}

// type marshalOperation struct {
// 	Meta
// 	Name        o.Option[string]          `json:"name,omitempty"`
// 	Description o.Option[string]          `json:"description,omitempty"`
// 	Input       o.Option[json.RawMessage] `json:"input,omitempty"`
// 	Output      o.Option[json.RawMessage] `json:"output,omitempty"`
// }

// func (t *Operation) UnmarshalJSON(data []byte) error {
// 	var d marshalOperation

// 	if err := json.Unmarshal(data, &d); err != nil {
// 		return err
// 	}

// 	t.Meta.Source = d.Meta.Source
// 	t.Meta.Namespace = d.Meta.Namespace
// 	t.Name = d.Name
// 	t.Description = d.Description

// 	if d.Input.Defined() {
// 		var typ CodeGenType

// 		if err := UnmarshalJson(d.Input.Get(), &typ); err != nil {
// 			return err
// 		}

// 		t.Input = o.Some(typ)
// 	}

// 	if d.Output.Defined() {
// 		var typ CodeGenType

// 		if err := UnmarshalJson(d.Output.Get(), &typ); err != nil {
// 			return err
// 		}

// 		t.Output = o.Some(typ)
// 	}

// 	return nil
// }
