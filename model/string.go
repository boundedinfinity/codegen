package model

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenString struct {
	codeGenCommon
	Min   optioner.Option[int]    `json:"min,omitempty"`
	Max   optioner.Option[int]    `json:"max,omitempty"`
	Regex optioner.Option[string] `json:"regex,omitempty"`
}

var _ CodeGenType = &CodeGenString{}

func (t CodeGenString) BaseType() string {
	return "string"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenString) HasValidation() bool {
	return t.Common().HasValidation() || t.Min.Defined() || t.Max.Defined() || t.Regex.Defined()
}

var (
	ErrStringMinBelow1    = errorer.New("min below 1")
	ErrStringMaxBelowMin  = errorer.New("max below min")
	ErrStringInvalidRegex = errorer.New("invalid regex")
)

func (t CodeGenString) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	if t.Min.Defined() && t.Min.Get() < 1 {
		return ErrStringMinBelow1.WithValue(t.Min.Get())
	}

	if t.Min.Defined() && t.Max.Defined() {
		if t.Max.Get() < t.Min.Get() {
			return ErrStringMaxBelowMin.FormatFn("max %v, min %v")(t.Max.Get(), t.Min.Get())
		}
	}

	if t.Regex.Defined() {
		_, err := regexp.Compile(t.Regex.Get())
		if err != nil {
			return ErrStringInvalidRegex.WithValue(err.Error())
		}
	}

	if t.Regex.Defined() {
		if _, err := regexp.Compile(t.Regex.Get()); err != nil {
			return err
		}
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (t *CodeGenString) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId        string `json:"base-type"`
		CodeGenString `json:",inline"`
	}{
		TypeId:        t.BaseType(),
		CodeGenString: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewString() *CodeGenString {
	return &CodeGenString{}
}

func (t *CodeGenString) WithQName(v string) *CodeGenString {
	t.codeGenCommon.withQName(v)
	return t
}

func (t *CodeGenString) WithName(v string) *CodeGenString {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenString) WithPackage(v string) *CodeGenString {
	t.codeGenCommon.withPackage(v)
	return t
}

func (t *CodeGenString) WithDescription(v string) *CodeGenString {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenString) WithRequired(v bool) *CodeGenString {
	t.codeGenCommon.withRequired(v)
	return t
}

// func (t *CodeGenString) WithDefault(v CodeGenString) *CodeGenString {
// 	t.codeGenCommon.withDefault(&v)
// 	return t
// }

func (t *CodeGenString) WithMin(v int) *CodeGenString {
	t.Min = optioner.Some(v)
	return t
}

func (t *CodeGenString) WithMax(v int) *CodeGenString {
	t.Max = optioner.Some(v)
	return t
}

func (t *CodeGenString) WithRegex(v string) *CodeGenString {
	t.Regex = optioner.Some(v)
	return t
}
