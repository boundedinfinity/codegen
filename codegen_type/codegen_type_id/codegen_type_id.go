package codegen_type_id

//go:generate enumer -path=./codegen_type_id.go

type CodgenTypeId string

const (
	Ref              CodgenTypeId = "ref"
	Array            CodgenTypeId = "array"
	Coordinate       CodgenTypeId = "coordinate"
	CreditCardNumber CodgenTypeId = "credit-card-number"
	Date             CodgenTypeId = "date"
	Datetime         CodgenTypeId = "date-time"
	Duration         CodgenTypeId = "duration"
	Email            CodgenTypeId = "email"
	Enum             CodgenTypeId = "enum"
	Float            CodgenTypeId = "float"
	Integer          CodgenTypeId = "integer"
	Ipv4             CodgenTypeId = "ipv4"
	Ipv6             CodgenTypeId = "ipv6"
	Mac              CodgenTypeId = "mac"
	Object           CodgenTypeId = "object"
	Phone            CodgenTypeId = "phone"
	SemanticVersion  CodgenTypeId = "semantic-version"
	String           CodgenTypeId = "string"
	Time             CodgenTypeId = "time"
	Uuid             CodgenTypeId = "uuid"
	Url              CodgenTypeId = "url"
)
