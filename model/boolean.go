package model

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	codeGenCommon
}

var _ CodeGenType = &CodeGenBoolean{}

func (t CodeGenBoolean) GetType() string {
	return "boolean"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t CodeGenBoolean) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (t *CodeGenBoolean) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"base-type"`
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
	return setO(t, &t.obj.Description, v)
}

// Name implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Name(v string) BooleanBuilder {
	return setO(t, &t.obj.Name, v)
}

// Package implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Package(v string) BooleanBuilder {
	return setO(t, &t.obj.Package, v)
}

// QName implements BooleanBuilder.
func (t *codeGenBooleanBuilder) QName(v string) BooleanBuilder {
	return setO(t, &t.obj.Name, v)
}

// Required implements BooleanBuilder.
func (t *codeGenBooleanBuilder) Required(v bool) BooleanBuilder {
	return setO(t, &t.obj.Required, v)
}
