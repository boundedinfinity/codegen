package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenOperation struct {
	Id          optioner.Option[string] `json:"id,omitempty"`
	Name        optioner.Option[string] `json:"name,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
	Inputs      []CodeGenSchema         `json:"inputs,omitempty"`
	Outputs     []CodeGenSchema         `json:"outputs,omitempty"`
}

func (t CodeGenOperation) TypeId() string {
	return "operation"
}

//----------------------------------------------------------------
// Validate
//----------------------------------------------------------------

var (
	ErrCodeGenOperationEmptyName     = errorer.New("empty name")
	ErrCodeGenOperationInvalidInput  = errorer.New("invalid input")
	ErrCodeGenOperationInvalidOutput = errorer.New("invalid output")
)

func (t *CodeGenOperation) Validate() error {
	if t.Name.Empty() {
		return ErrCodeGenOperationEmptyName
	}

	for i, input := range t.Inputs {
		if err := input.Validate(); err != nil {
			return fmt.Errorf("inputs[%v]: %w", i, ErrCodeGenOperationInvalidInput)
		}
	}

	for i, output := range t.Outputs {
		if err := output.Validate(); err != nil {
			return fmt.Errorf("outputs[%v]: %w", i, ErrCodeGenOperationInvalidOutput)
		}
	}

	return nil
}

//----------------------------------------------------------------
// Merge
//----------------------------------------------------------------

func (t *CodeGenOperation) Merge(obj CodeGenOperation) error {
	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenOperation) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId           string `json:"type"`
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
			t.Inputs = append(t.Inputs, input)
		}
	}

	for i, rawOutput := range dto.Inputs {
		if output, err := UnmarshalCodeGenType(rawOutput); err != nil {
			return errors.Join(fmt.Errorf("input[%v]", i), err)
		} else {
			t.Outputs = append(t.Outputs, output)
		}
	}

	return nil
}
