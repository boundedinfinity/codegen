package thing

type Thing interface {
	GetType() string
	Validate() error
	HasValidation() bool
}

type IntBuilder interface {
	Name(v string) IntBuilder
	Type(v string) IntBuilder
	Min(v int) IntBuilder
	Max(v int) IntBuilder
	Build() *IntThing
}

func setV[T any, V any](t T, c *V, n V) T {
	*c = n
	return t
}

type commonThing struct {
	Name string
	Type string
}

type IntThing struct {
	commonThing
	Min int
	Max int
}

func (t IntThing) GetType() string {
	return t.Type
}

func (t IntThing) Validate() error {
	return nil
}

func (t IntThing) HasValidation() bool {
	return false
}

var _ Thing = &IntThing{}

type intThingBuilder struct {
	thing IntThing
}

func (t *intThingBuilder) Build() *IntThing {
	return &t.thing
}

func (t *intThingBuilder) Name(v string) IntBuilder {
	return setV(t, &t.thing.Name, v)
}

func (t *intThingBuilder) Type(v string) IntBuilder {
	return setV(t, &t.thing.Type, v)
}

func (t *intThingBuilder) Min(v int) IntBuilder {
	return setV(t, &t.thing.Min, v)
}

func (t *intThingBuilder) Max(v int) IntBuilder {
	return setV(t, &t.thing.Max, v)
}

func BuildInt() IntBuilder {
	return &intThingBuilder{}
}

// type strThing struct {
// 	commonThing
// 	Regex string
// 	Abnfr string
// }

// func (t *strThing) WithRegex(v string) *strThing {
// 	t.Regex = v
// 	return t
// }

// func (t *strThing) WithAbnfr(v string) *strThing {
// 	t.Abnfr = v
// 	return t
// }

// type StrBuilder interface {
// 	CommonBuilder
// 	WithRegex(string) StrBuilder
// 	WithAbnfr(string) StrBuilder
// }

// func BuildStr() StrBuilder {
// 	return &strThing{}
// }
