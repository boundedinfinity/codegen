package model

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

func (this CodeGenArray) Schema() string {
	return "array"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	errCodeGenArrayMinGreaterThanMax = errorer.New("min greater than max")
)

func (this CodeGenArray) HasValidation() bool {
	return this.CodeGenCommon.HasValidation() ||
		this.Items.Defined() && this.Items.Get().HasValidation() ||
		this.Min.Defined() || this.Max.Defined()
}

func (this CodeGenArray) Validate() error {
	if err := this.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if err := this.Items.Get().Validate(); err != nil {
		return err
	}

	if this.Min.Defined() && this.Max.Defined() && this.Min.Get() > this.Max.Get() {
		return errCodeGenArrayMinGreaterThanMax.FormatFn("min: %v, max: %v")(this.Min.Get(), this.Max.Get())
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (this *CodeGenArray) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"type"`
		CodeGenArray `json:",inline"`
	}{
		TypeId:       this.Schema(),
		CodeGenArray: *this,
	}

	return marshalCodeGenType(dto)
}

func (this *CodeGenArray) UnmarshalJSON(data []byte) error {
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
		this.CodeGenCommon = dto.CodeGenCommon
		this.Min = dto.Min
		this.Max = dto.Max
	}

	if items, err := UnmarshalCodeGenType(dto.Items); err != nil {
		return err
	} else {
		this.Items = optioner.Some(items)
	}

	return nil
}
