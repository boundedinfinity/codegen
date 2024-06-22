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
	Min           optioner.Option[int]  `json:"min,omitempty"`
	ExclusiveMin  optioner.Option[int]  `json:"exclusive-min,omitempty"`
	Max           optioner.Option[int]  `json:"max,omitempty"`
	ExclusiveMax  optioner.Option[int]  `json:"exclusive-max,omitempty"`
	Items         CodeGenType           `json:"items,omitempty"`
	ManyToMany    optioner.Option[bool] `json:"many-to-many,omitempty"`
}

func (t CodeGenArray) CodeGenId() string {
	return "array"
}

var _ CodeGenType = &CodeGenArray{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	errCodeGenArrayMinAndExclusiveMinMutuallyExclusive = errorer.New("min and exclusive-min are multually exclusive")
	errCodeGenArrayMaxAndExclusiveMaxMutuallyExclusive = errorer.New("max and exclusive-max are multually exclusive")
)

func (t CodeGenArray) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
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
		TypeId       string `json:"codegen-id"`
		CodeGenArray `json:",inline"`
	}{
		TypeId:       t.CodeGenId(),
		CodeGenArray: *t,
	}

	return json.Marshal(dto)
}

func (t *CodeGenArray) UnmarshalJSON(data []byte) error {
	dto := struct {
		CodeGenCommon
		Min          optioner.Option[int]  `json:"min,omitempty"`
		ExclusiveMin optioner.Option[int]  `json:"exclusive-min,omitempty"`
		Max          optioner.Option[int]  `json:"max,omitempty"`
		ExclusiveMax optioner.Option[int]  `json:"exclusive-max,omitempty"`
		Items        json.RawMessage       `json:"items,omitempty"`
		ManyToMany   optioner.Option[bool] `json:"many-to-many,omitempty"`
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.CodeGenCommon = dto.CodeGenCommon
		t.Min = dto.Min
		t.Max = dto.Max
		t.ExclusiveMax = dto.ExclusiveMax
		t.ExclusiveMin = dto.ExclusiveMin
		t.ManyToMany = dto.ManyToMany
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
	t.CodeGenCommon.WithTypeId(v)
	return t
}

func (t *CodeGenArray) WithName(v string) *CodeGenArray {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *CodeGenArray) WithDescription(v string) *CodeGenArray {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *CodeGenArray) WithRequired(v bool) *CodeGenArray {
	t.CodeGenCommon.WithRequired(v)
	return t
}

func (t *CodeGenArray) WithDefault(v CodeGenArray) *CodeGenArray {
	t.CodeGenCommon.WithDefault(&v)
	return t
}

func (t *CodeGenArray) WithEager(v bool) *CodeGenArray {
	t.CodeGenCommon.WithEager(v)
	return t
}

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
