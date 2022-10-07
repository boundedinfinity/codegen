package template_ext

//go:generate enumer -path=./main.go

type TemplateExt string

const (
	Gotmpl     TemplateExt = "gotmpl"
	Handlebars TemplateExt = "handlebars"
)
