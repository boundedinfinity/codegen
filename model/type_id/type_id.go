package type_id

//go:generate enumer -path=./type_id.go

type TypeId string

const (
	Ref     TypeId = "ref"
	Array   TypeId = "array"
	Boolean TypeId = "boolean"
	Float   TypeId = "float"
	Integer TypeId = "integer"
	Object  TypeId = "object"
	String  TypeId = "string"
)
