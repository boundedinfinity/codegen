package model

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	CodeGenCommon
}

var _ CodeGenType = &CodeGenBoolean{}

func (t CodeGenBoolean) GetType() string {
	return "boolean"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenBoolean) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (t *CodeGenBoolean) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"type"`
		CodeGenBoolean `json:",inline"`
	}{
		TypeId:         t.GetType(),
		CodeGenBoolean: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func BuildBoolean() BooleanBuilder {
	return &codeGenBooleanBuilder{}
}

type codeGenBooleanBuilder struct {
	obj CodeGenBoolean
}

var _ BooleanBuilder = &codeGenBooleanBuilder{}

// Build implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Build() *CodeGenBoolean {
	return &t.obj
}

// Description implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Description(v string) BooleanBuilder {
	return SetO(t, &t.obj.Description, v)
}

// Name implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Name(v string) BooleanBuilder {
	return SetO(t, &t.obj.Name, v)
}

// Package implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Package(v string) BooleanBuilder {
	return SetO(t, &t.obj.Package, v)
}

// Id implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Id(v string) BooleanBuilder {
	return SetO(t, &t.obj.Id, v)
}

// Required implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Required(v bool) BooleanBuilder {
	return SetO(t, &t.obj.Required, v)
}
