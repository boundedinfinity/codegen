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
	Name        optioner.Option[string] `json:"name,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
	Inputs      []CodeGenType           `json:"inputs,omitempty"`
	Outputs     []CodeGenType           `json:"outputs,omitempty"`
	CodeGenMeta
}

func (t CodeGenOperation) TypeId() string {
	return "operation"
}

//----------------------------------------------------------------
// Validate
//----------------------------------------------------------------

var (
	ErrCodeGenOperationEmptyName     = errors.New("empty name")
	ErrCodeGenOperationInvalidInput  = errors.New("invalid input")
	ErrCodeGenOperationInvalidOutput = errors.New("invalid output")
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

	if err := t.CodeGenMeta.Validate(); err != nil {
		return err
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
		TypeId           string `json:"codegen-id"`
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

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewOperation() *CodeGenOperation {
	return &CodeGenOperation{}
}

func (t *CodeGenOperation) WithName(v string) *CodeGenOperation {
	t.Name = optioner.OfZero(v)
	return t
}

func (t *CodeGenOperation) WithDescription(v string) *CodeGenOperation {
	t.Description = optioner.OfZero(v)
	return t
}

func (t *CodeGenOperation) WithInputs(v ...CodeGenType) *CodeGenOperation {
	t.Inputs = append(t.Inputs, v...)
	return t
}

func (t *CodeGenOperation) WithOutputs(v ...CodeGenType) *CodeGenOperation {
	t.Outputs = append(t.Outputs, v...)
	return t
}
