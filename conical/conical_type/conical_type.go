package conical_type

//go:generate enumer -path=./conical_type.go

type ConicalType string

const (
	Object   ConicalType = "Object"
	Integer  ConicalType = "Integer"
	Float    ConicalType = "Float"
	String   ConicalType = "String"
	Date     ConicalType = "Date"
	Duration ConicalType = "Duration"
)
