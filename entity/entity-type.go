package entity

type EntityType int

const (
	noEntityType EntityType = iota
	StringType   EntityType = iota
	IntegerType  EntityType = iota
	FloatType    EntityType = iota
	ObjectType   EntityType = iota
	EnumType     EntityType = iota
	ArrayType    EntityType = iota
	RefType      EntityType = iota
	BooleanType  EntityType = iota
	DateType     EntityType = iota
	DateTimeType EntityType = iota
	TimeType     EntityType = iota
	DurationType EntityType = iota
	UnionType    EntityType = iota
)

func (e EntityType) String() string {
	return type2StringMap[e]
}

var (
	type2StringMap = map[EntityType]string{
		StringType:   "string",
		IntegerType:  "integer",
		FloatType:    "float",
		ObjectType:   "object",
		EnumType:     "enum",
		ArrayType:    "array",
		RefType:      "ref",
		BooleanType:  "boolean",
		DateType:     "date",
		DateTimeType: "date-time",
		TimeType:     "time",
		DurationType: "duration",
		UnionType:    "union",
	}

	validTypes = []EntityType{
		StringType, IntegerType, FloatType, ObjectType, DurationType, DateType, DateTimeType,
		TimeType, EnumType, ArrayType, RefType, BooleanType, UnionType,
	}
)
