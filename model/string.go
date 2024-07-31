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
	CodeGenCommon
	Min      optioner.Option[int]      `json:"min,omitempty"`
	Max      optioner.Option[int]      `json:"max,omitempty"`
	Regex    optioner.Option[string]   `json:"regex,omitempty"`
	Abnf     optioner.Option[string]   `json:"abnf,omitempty"`
	Includes optioner.Option[[]string] `json:"includes,omitempty"`
	Excludes optioner.Option[[]string] `json:"excludes,omitempty"`
	OneOf    optioner.Option[[]string] `json:"one-of,omitempty"`
	NoneOf   optioner.Option[[]string] `json:"none-of,omitempty"`
	// upper case, lower case, snake case, kebab case, etc...
}

var _ CodeGenType = &CodeGenString{}

func (t CodeGenString) GetType() string {
	return "string"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenString) HasValidation() bool {
	return t.CodeGenCommon.HasValidation() || t.Min.Defined() || t.Max.Defined() || t.Regex.Defined() ||
		t.Includes.Defined() || t.Excludes.Defined() || t.OneOf.Defined() || t.NoneOf.Defined()
}

var (
	ErrStringMinBelow1    = errorer.New("min below 1")
	ErrStringMaxBelowMin  = errorer.New("max below min")
	ErrStringInvalidRegex = errorer.New("invalid regex")
)

func (t CodeGenString) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
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
		TypeId        string `json:"type"`
		CodeGenString `json:",inline"`
	}{
		TypeId:        t.GetType(),
		CodeGenString: *t,
	}

	return marshalCodeGenType(dto)
}
