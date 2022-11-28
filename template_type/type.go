package template_type

//go:generate enumer -path=./type.go

type TemplateType string

const (
	Model     TemplateType = "model"
	Namespace TemplateType = "namespace"
	Operation TemplateType = "operation"
)
