package model

import "strings"

const (
	TYPE_UNKNOWN      = "<UNKNOWN_TYPE>"
	TYPE_BUILTIN      = "<builtin>"
	NAMESPACE_BUILTIN = "builtin"
	NAMESPACE_UNKNOWN = "UNKNOWN"
	NAMESPACE_CUSTOM  = "CUSTOM"
	COLLECTION_SUFFIX = "[]"
	SUMMERY_SIZE      = 35
)

type TemplateExt string

const (
	TemplateExt_Unkown     TemplateExt = ".unknown"
	TemplateExt_GoTmpl     TemplateExt = ".gotmpl"
	TemplateExt_Handlebars TemplateExt = ".handlebars"
)

func (t TemplateExt) String() string {
	return string(t)
}

func IsTemplateExt(v string) bool {
	for _, e := range TemplateExts {
		if v == e.String() || v == strings.TrimPrefix(e.String(), ".") {
			return true
		}
	}

	return false
}

type TemplateType string

const (
	TemplateType_MODEL     TemplateType = "model"
	TemplateType_OPERATION TemplateType = "operation"
	TemplateType_NAMESPACE TemplateType = "namespace"
)

func (t TemplateType) String() string {
	return string(t)
}

func IsTemplateType(v string) bool {
	for _, e := range TemplateTypes {
		if v == e.String() {
			return true
		}
	}

	return false
}

func TemplateTypeStrings() []string {
	ss := make([]string, 0)

	for _, s := range TemplateTypes {
		ss = append(ss, s.String())
	}

	return ss
}

type LanguageExt string

const (
	LanguageExt_Unkown     LanguageExt = ".unknown"
	LanguageExt_Go         LanguageExt = ".go"
	LanguageExt_GoMod      LanguageExt = ".mod"
	LanguageExt_Typescript LanguageExt = ".ts"
	LanguageExt_Javascript LanguageExt = ".js"
	LanguageExt_Html       LanguageExt = ".html"
	LanguageExt_Css        LanguageExt = ".css"
)

func (t LanguageExt) String() string {
	return string(t)
}

func IsLanguageExt(v string) bool {
	for _, e := range LanguageExts {
		if v == e.String() || v == strings.TrimPrefix(e.String(), ".") {
			return true
		}
	}

	return false
}

var (
	TemplateExts = []TemplateExt{
		TemplateExt_GoTmpl,
		TemplateExt_Handlebars,
	}

	TemplateTypes = []TemplateType{
		TemplateType_MODEL,
		TemplateType_NAMESPACE,
		TemplateType_OPERATION,
	}

	LanguageExts = []LanguageExt{
		LanguageExt_Go,
		LanguageExt_GoMod,
		LanguageExt_Typescript,
		LanguageExt_Javascript,
		LanguageExt_Html,
		LanguageExt_Css,
	}
)
