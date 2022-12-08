package canonical_type

//go:generate enumer -path=./canonical_type.go

type CanonicalType string

const (
	Array    CanonicalType = "array"
	Date     CanonicalType = "date"
	Datetime CanonicalType = "datetime"
	Duration CanonicalType = "duration"
	Enum     CanonicalType = "enum"
	Float    CanonicalType = "float"
	Integer  CanonicalType = "integer"
	Object   CanonicalType = "object"
	String   CanonicalType = "string"
	Uuid     CanonicalType = "uuid"
	Url      CanonicalType = "url"
	Email    CanonicalType = "email"
)
