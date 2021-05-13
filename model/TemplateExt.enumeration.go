////////////////////////////////////////////////////////////////////////
// Generated by bounded/enumeration
////////////////////////////////////////////////////////////////////////

package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type TemplateExtEnum string

const (
	TemplateExt_Unknown    TemplateExtEnum = "unknown"
	TemplateExt_Gotmpl     TemplateExtEnum = "gotmpl"
	TemplateExt_Handlebars TemplateExtEnum = "handlebars"
)

var (
	TemplateExtEnums = []TemplateExtEnum{
		TemplateExt_Unknown,
		TemplateExt_Gotmpl,
		TemplateExt_Handlebars,
	}
)

func IsTemplateExtEnum(v string) bool {
	var f bool

	for _, e := range TemplateExtEnums {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func TemplateExtEnumParse(v string) (TemplateExtEnum, error) {
	var o TemplateExtEnum
	var f bool
	n := strings.ToLower(v)

	for _, e := range TemplateExtEnums {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, ErrTemplateExtEnumNotFound(v)
	}

	return o, nil
}

func ErrTemplateExtEnumNotFound(v string) error {
	var ss []string

	for _, e := range TemplateExtEnums {
		ss = append(ss, string(e))
	}

	return fmt.Errorf(
		"invalid enumeration type '%v', must be one of %v",
		v, strings.Join(ss, ","),
	)
}

func (t TemplateExtEnum) String() string {
	return string(t)
}

func (t TemplateExtEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal([]byte(t))
}

func (t *TemplateExtEnum) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := TemplateExtEnumParse(s)

	if err != nil {
		return err
	}

	t = &e

	return nil
}
