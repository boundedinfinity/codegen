package header_strategy

//go:generate enumer -path=./header.go

type HeaderStrategy string

const (
	GlobalFirst   HeaderStrategy = "GlobalFirst"
	TemplateFirst HeaderStrategy = "TemplateFirst"
	TemplateOnly  HeaderStrategy = "TemplateOnly"
	GlobalOnly    HeaderStrategy = "GlobalOnly"
	Ignore        HeaderStrategy = "Ignore"
	UseDefault    HeaderStrategy = "UseDefault"
)
