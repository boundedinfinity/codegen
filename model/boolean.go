package model

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	CodeGenCommon
}

var _ CodeGenSchema = &CodeGenBoolean{}

func (t CodeGenBoolean) Schema() string {
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
		TypeId:         t.Schema(),
		CodeGenBoolean: *t,
	}

	return marshalCodeGenType(dto)
}
