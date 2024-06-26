package entity

func String() *stringEntity {
	return &stringEntity{
		entityBase: entityBase{entityType: StringType},
	}
}

type StringEntity interface {
	Entity
	Min(int) StringEntity
	Max(int) StringEntity
	Regex(string) StringEntity
	Includes(string) StringEntity
	StartsWith(string) StringEntity
	EndsWith(string) StringEntity
}

type stringEntity struct {
	entityBase
	min        int
	max        int
	regex      string
	includes   string
	startsWith string
	endsWith   string
}

var _ Entity = &stringEntity{}
var _ StringEntity = &stringEntity{}

func (t *stringEntity) Min(n int) StringEntity {
	t.min = n
	return t
}

func (t *stringEntity) Max(n int) StringEntity {
	t.max = n
	return t
}

func (t *stringEntity) Length(n int) StringEntity {
	t.max = n
	t.min = n
	return t
}

func (t *stringEntity) Regex(s string) StringEntity {
	t.regex = s
	return t
}

func (t *stringEntity) Includes(s string) StringEntity {
	t.includes = s
	return t
}

func (t *stringEntity) StartsWith(r string) StringEntity {
	t.startsWith = r
	return t
}

func (t *stringEntity) EndsWith(r string) StringEntity {
	t.endsWith = r
	return t
}
