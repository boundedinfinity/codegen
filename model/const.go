package model

import "strings"

const (
	TYPE_UNKNOWN      = "<UNKNOWN_TYPE>"
	TYPE_BUILTIN      = "<builtin>"
	NAMESPACE_BUILTIN = "builtin"
	NAMESPACE_UNKNOWN = "UNKNOWN"
	COLLECTION_SUFFIX = "[]"
	SUMMERY_SIZE      = 25
)

type TemplateExt string

const (
	TemplateExt_Unkown     TemplateExt = ".unknown"
	TemplateExt_GoTmpl     TemplateExt = ".gotmpl"
	TemplateExt_Handlebars TemplateExt = ".handlebars"
)

type TemplateType string

const (
	TemplateType_MODEL     TemplateType = "model"
	TemplateType_OPERATION TemplateType = "operation"
	TemplateType_NAMESPACE TemplateType = "namespace"
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

	LanguageExts = []LanguageExt{
		LanguageExt_Go,
		LanguageExt_GoMod,
		LanguageExt_Typescript,
		LanguageExt_Javascript,
		LanguageExt_Html,
		LanguageExt_Css,
	}
)
