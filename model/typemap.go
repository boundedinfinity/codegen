package model

type TypeInfo struct {
	SpecType      string
	BaseName      string
	ImportName    string
	QName         string
	Namespace     string
	BuiltIn       bool
	JsonStructure map[string]interface{}
}
