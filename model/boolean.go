package model

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenBoolean struct {
	CodeGenCommon
}

var _ CodeGenSchema = &CodeGenBoolean{}

func (this CodeGenBoolean) Schema() string {
	return "boolean"
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (this CodeGenBoolean) Validate() error {
	if err := this.CodeGenCommon.Validate(); err != nil {
		return err
	}

	return nil
}

// ----------------------------------------------------------------
// Marshal
// ----------------------------------------------------------------

func (this *CodeGenBoolean) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"type"`
		CodeGenBoolean `json:",inline"`
	}{
		TypeId:         this.Schema(),
		CodeGenBoolean: *this,
	}

	return marshalCodeGenType(dto)
}
