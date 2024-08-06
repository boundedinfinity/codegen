package model

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

import (
	"boundedinfinity/codegen/utils"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenRef struct {
	CodeGenCommon `json:",inline,omitempty"`
	Ref           optioner.Option[string] `json:"ref,omitempty"`
	Resolved      CodeGenSchema           `json:"-"`
}

func (this CodeGenRef) Schema() string {
	return "ref"
}

var _ CodeGenSchema = &CodeGenRef{}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (this CodeGenRef) HashValidation() bool {
	return this.CodeGenCommon.HasValidation() ||
		this.Resolved != nil && this.Resolved.HasValidation()
}

func (this CodeGenRef) Validate() error {
	if err := this.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if this.Ref.Empty() {
		return errCodeGenRefEmpty.New(this)
	}

	if this.Resolved != nil {
		if err := this.Resolved.Validate(); err != nil {
			return err
		}
	}

	return nil
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (this *CodeGenRef) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId string `json:"type"`
		CodeGenRef
	}{
		TypeId:     this.Schema(),
		CodeGenRef: *this,
	}

	return marshalCodeGenType(dto)
}

var ErrCodeGenRef = errors.New("ref error")

var errCodeGenRefEmpty = utils.ErrorFactory(ErrCodeGenRef, func(ref CodeGenRef) string {
	return fmt.Sprintf("reference object %s doesn't have ID", ref.Common().Name.Get())
})
