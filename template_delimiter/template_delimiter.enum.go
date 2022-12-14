//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package template_delimiter

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slicer" // v1.0.15
)

var (
	All = []TemplateDelimiter{
		Curly,
		Square,
		Parens,
	}
)

func (t TemplateDelimiter) String() string {
	return string(t)
}

func Parse(v string) (TemplateDelimiter, error) {
	f, ok := slicer.FindFn(All, func(x TemplateDelimiter) bool {
		return TemplateDelimiter(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slicer.ContainsFn(All, func(v TemplateDelimiter) bool {
		return string(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, slicer.Join(All, ","),
	)
}

func (t TemplateDelimiter) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *TemplateDelimiter) UnmarshalJSON(data []byte) error {
	var v string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t TemplateDelimiter) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *TemplateDelimiter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
