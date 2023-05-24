package type_id

//go:generate enumer -path=./type_id.go

type TypeId string

const (
	Array   TypeId = "array"
	Boolean TypeId = "boolean"
	Enum    TypeId = "enum"
	Float   TypeId = "float"
	Integer TypeId = "integer"
	Object  TypeId = "object"
	Ref     TypeId = "ref"
	String  TypeId = "string"
)
