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
	codeGenCommon `json:",inline,omitempty"`
	Min           optioner.Option[int] `json:"min,omitempty"`
	ExclusiveMin  optioner.Option[int] `json:"exclusive-min,omitempty"`
	Max           optioner.Option[int] `json:"max,omitempty"`
	ExclusiveMax  optioner.Option[int] `json:"exclusive-max,omitempty"`
	Items         CodeGenType          `json:"items,omitempty"`
}

var _ CodeGenType = &CodeGenArray{}

func (t CodeGenArray) BaseType() string {
	return "array"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	errCodeGenArrayMinAndExclusiveMinMutuallyExclusive = errorer.New("min and exclusive-min are multually exclusive")
	errCodeGenArrayMaxAndExclusiveMaxMutuallyExclusive = errorer.New("max and exclusive-max are multually exclusive")
)

func (t CodeGenArray) HasValidation() bool {
	return t.Common().HasValidation() ||
		t.Min.Defined() || t.ExclusiveMin.Defined() ||
		t.Max.Defined() || t.ExclusiveMax.Defined()
}

func (t CodeGenArray) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	if err := t.Items.Validate(); err != nil {
		return err
	}

	if t.Min.Defined() && t.ExclusiveMin.Defined() {
		return errCodeGenArrayMinAndExclusiveMinMutuallyExclusive
	}

	if t.Max.Defined() && t.ExclusiveMax.Defined() {
		return errCodeGenArrayMaxAndExclusiveMaxMutuallyExclusive
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenArray) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"base-type"`
		CodeGenArray `json:",inline"`
	}{
		TypeId:       t.BaseType(),
		CodeGenArray: *t,
	}

	return marshalCodeGenType(dto)
}

func (t *CodeGenArray) UnmarshalJSON(data []byte) error {
	dto := struct {
		codeGenCommon
		Min          optioner.Option[int] `json:"min,omitempty"`
		ExclusiveMin optioner.Option[int] `json:"exclusive-min,omitempty"`
		Max          optioner.Option[int] `json:"max,omitempty"`
		ExclusiveMax optioner.Option[int] `json:"exclusive-max,omitempty"`
		Items        json.RawMessage      `json:"items,omitempty"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.codeGenCommon = dto.codeGenCommon
		t.Min = dto.Min
		t.Max = dto.Max
		t.ExclusiveMax = dto.ExclusiveMax
		t.ExclusiveMin = dto.ExclusiveMin
	}

	if items, err := UnmarshalCodeGenType(dto.Items); err != nil {
		return err
	} else {
		t.Items = items
	}

	return nil
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewArray() *CodeGenArray {
	return &CodeGenArray{}
}

func (t *CodeGenArray) WithSchemaId(v string) *CodeGenArray {
	t.codeGenCommon.withQName(v)
	return t
}

func (t *CodeGenArray) WithName(v string) *CodeGenArray {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenArray) WithDescription(v string) *CodeGenArray {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenArray) WithRequired(v bool) *CodeGenArray {
	t.codeGenCommon.withRequired(v)
	return t
}

// func (t *CodeGenArray) WithDefault(v CodeGenArray) *CodeGenArray {
// 	t.codeGenCommon.withDefault(&v)
// 	return t
// }

// func (t *CodeGenArray) WithEager(v bool) *CodeGenArray {
// 	t.codeGenCommon.withEager(v)
// 	return t
// }

func (t *CodeGenArray) WithMin(v int) *CodeGenArray {
	t.Min = optioner.OfZero(v)
	return t
}

func (t *CodeGenArray) WithMax(v int) *CodeGenArray {
	t.Max = optioner.OfZero(v)
	return t
}

func (t *CodeGenArray) WithExclusiveMin(v int) *CodeGenArray {
	t.ExclusiveMin = optioner.OfZero(v)
	return t
}

func (t *CodeGenArray) WithExclusiveMax(v int) *CodeGenArray {
	t.ExclusiveMax = optioner.OfZero(v)
	return t
}

func (t *CodeGenArray) WithItems(v CodeGenType) *CodeGenArray {
	t.Items = v
	return t
}
