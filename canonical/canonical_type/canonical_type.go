package canonical_type

//go:generate enumer -path=./canonical_type.go

type CanonicalType string

const (
	Ref              CanonicalType = "ref"
	Array            CanonicalType = "array"
	Coordinate       CanonicalType = "coordinate"
	CreditCardNumber CanonicalType = "credit-card-number"
	Date             CanonicalType = "date"
	Datetime         CanonicalType = "date-time"
	Duration         CanonicalType = "duration"
	Email            CanonicalType = "email"
	Enum             CanonicalType = "enum"
	Float            CanonicalType = "float"
	Integer          CanonicalType = "integer"
	Ipv4             CanonicalType = "ipv4"
	Ipv6             CanonicalType = "ipv6"
	Mac              CanonicalType = "mac"
	Object           CanonicalType = "object"
	Phone            CanonicalType = "phone"
	SemanticVersion  CanonicalType = "semantic-version"
	String           CanonicalType = "string"
	Time             CanonicalType = "time"
	Uuid             CanonicalType = "uuid"
	Url              CanonicalType = "url"
)
