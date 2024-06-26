package entity

type EntityType int

const (
	noneType    EntityType = iota
	StringType  EntityType = iota
	IntegerType EntityType = iota
	FloatType   EntityType = iota
	ObjectType  EntityType = iota
	EnumType    EntityType = iota
	ArrayType   EntityType = iota
	RefType     EntityType = iota
)

var (
	type2String = map[EntityType]string{
		StringType:  "string",
		IntegerType: "integer",
		FloatType:   "float",
		ObjectType:  "object",
		EnumType:    "enum",
		ArrayType:   "array",
		RefType:     "ref",
	}
)

type Entity interface {
	Type() EntityType
	Name(string) Entity
	Required(bool) Entity
	Description(string) Entity
	Default(map[string]any) Entity
	ToJson() ([]byte, error)
	ToJsonIndent() ([]byte, error)
	ToYaml() ([]byte, error)
}

func newEntity(entityType EntityType) Entity {
	return &entityBase{
		entityType: entityType,
	}
}

type entityBase struct {
	entityType   EntityType
	name         string
	description  string
	required     bool
	defaultValue map[string]any
}

func (t *entityBase) Type() EntityType {
	return t.entityType
}

func (t *entityBase) Name(name string) Entity {
	t.name = name
	return t
}

func (t *entityBase) Required(required bool) Entity {
	t.required = required
	return t
}

func (t *entityBase) Description(description string) Entity {
	t.description = description
	return t
}

func (t *entityBase) Default(value map[string]any) Entity {
	t.defaultValue = value
	return t
}
