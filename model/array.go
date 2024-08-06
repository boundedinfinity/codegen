package model

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenArray struct {
	CodeGenCommon `json:",inline,omitempty"`
	Min           optioner.Option[int]           `json:"min,omitempty"`
	Max           optioner.Option[int]           `json:"max,omitempty"`
	Items         optioner.Option[CodeGenSchema] `json:"items,omitempty"`
}

var _ CodeGenSchema = &CodeGenArray{}

func (t CodeGenArray) Schema() string {
	return "array"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	errCodeGenArrayMinGreaterThanMax = errorer.New("min greater than max")
)

func (t CodeGenArray) HasValidation() bool {
	return t.CodeGenCommon.HasValidation() ||
		t.Items.Defined() && t.Items.Get().HasValidation() ||
		t.Min.Defined() || t.Max.Defined()
}

func (t CodeGenArray) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if err := t.Items.Get().Validate(); err != nil {
		return err
	}

	if t.Min.Defined() && t.Max.Defined() && t.Min.Get() > t.Max.Get() {
		return errCodeGenArrayMinGreaterThanMax.FormatFn("min: %v, max: %v")(t.Min.Get(), t.Max.Get())
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenArray) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"type"`
		CodeGenArray `json:",inline"`
	}{
		TypeId:       t.Schema(),
		CodeGenArray: *t,
	}

	return marshalCodeGenType(dto)
}

func (t *CodeGenArray) UnmarshalJSON(data []byte) error {
	dto := struct {
		CodeGenCommon
		Min          optioner.Option[int] `json:"min,omitempty"`
		ExclusiveMin optioner.Option[int] `json:"exclusive-min,omitempty"`
		Max          optioner.Option[int] `json:"max,omitempty"`
		ExclusiveMax optioner.Option[int] `json:"exclusive-max,omitempty"`
		Items        json.RawMessage      `json:"items,omitempty"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.CodeGenCommon = dto.CodeGenCommon
		t.Min = dto.Min
		t.Max = dto.Max
	}

	if items, err := UnmarshalCodeGenType(dto.Items); err != nil {
		return err
	} else {
		t.Items = optioner.Some(items)
	}

	return nil
}
