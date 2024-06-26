package model

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	codeGenCommon
}

var _ CodeGenType = &CodeGenBoolean{}

func (t CodeGenBoolean) BaseType() string {
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
		TypeId:         t.BaseType(),
		CodeGenBoolean: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewBoolean() *CodeGenBoolean {
	return &CodeGenBoolean{}
}

func (t *CodeGenBoolean) WithQName(v string) *CodeGenBoolean {
	t.codeGenCommon.withQName(v)
	return t
}

func (t *CodeGenBoolean) WithName(v string) *CodeGenBoolean {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenBoolean) WithDescription(v string) *CodeGenBoolean {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenBoolean) WithRequired(v bool) *CodeGenBoolean {
	t.codeGenCommon.withRequired(v)
	return t
}

// func (t *CodeGenBoolean) WithDefault(v CodeGenBoolean) *CodeGenBoolean {
// 	t.codeGenCommon.withDefault(&v)
// 	return t
// }
