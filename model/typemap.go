package model

type TypeInfo struct {
	SpecType           string
	InNamespaceType    string
	OutOfNamespaceType string
	Namespace          string
	JsonStructure      map[string]interface{}
	Example            string
}
