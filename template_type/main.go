package template_type

//go:generate enumer -path=./main.go

type TemplateType string

const (
	Model     TemplateType = "model"
	Namespace TemplateType = "namespace"
	Operation TemplateType = "operation"
)
