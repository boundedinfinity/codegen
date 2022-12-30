package template_delimiter

//go:generate enumer -path=./template_delimiter.go

type TemplateDelimiter string

const (
	Curly  TemplateDelimiter = "curly"
	Square TemplateDelimiter = "square"
	Parens TemplateDelimiter = "parens"
	Angle  TemplateDelimiter = "angle"
)
