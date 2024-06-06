package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenOperation struct {
	Name        optioner.Option[string]        `json:"name,omitempty"`
	Description optioner.Option[string]        `json:"description,omitempty"`
	Inputs      optioner.Option[[]CodeGenType] `json:"inputs,omitempty"`
	Outputs     optioner.Option[[]CodeGenType] `json:"outputs,omitempty"`
}

func (t CodeGenOperation) TypeId() string {
	return "operation"
}

var _ CodeGenType = &CodeGenOperation{}

//////////////////////////////////////////////////////////////////
// Marshal
//////////////////////////////////////////////////////////////////

func (t *CodeGenOperation) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId           string `json:"type-id"`
		CodeGenOperation `json:",inline"`
	}{
		TypeId:           t.TypeId(),
		CodeGenOperation: *t,
	}

	return json.Marshal(dto)
}

func (t *CodeGenOperation) UnmarshalJSON(data []byte) error {
	dto := struct {
		Name        optioner.Option[string] `json:"name,omitempty"`
		Description optioner.Option[string] `json:"description,omitempty"`
		Inputs      []json.RawMessage       `json:"inputs,omitempty"`
		Outputs     []json.RawMessage       `json:"outputs,omitempty"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.Name = dto.Name
		t.Description = dto.Description
	}

	for i, rawInput := range dto.Inputs {
		if input, err := UnmarshalCodeGenType(rawInput); err != nil {
			return errors.Join(fmt.Errorf("input[%v]", i), err)
		} else {
			t.Inputs = optioner.Some(append(t.Inputs.Get(), input))
		}
	}

	for i, rawOutput := range dto.Inputs {
		if output, err := UnmarshalCodeGenType(rawOutput); err != nil {
			return errors.Join(fmt.Errorf("input[%v]", i), err)
		} else {
			t.Outputs = optioner.Some(append(t.Outputs.Get(), output))
		}
	}

	return nil
}
