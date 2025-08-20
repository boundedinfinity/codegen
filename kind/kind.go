// Package kind contains the enumeration of kind names
package kind

import (
	"boundedinfinity/codegen/kind/name"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

type ValidatorConfig struct {
	prefix string
}

type Kind interface {
	KindName() name.KindName
	Validate(ValidatorConfig) error
}

var (
	ErrKindMissingQName = errorer.New("missing q-name")
)

// //////////////////////////////////////////////////////////////////////////
// KindCommon
// //////////////////////////////////////////////////////////////////////////

type KindCommon struct {
	Name optioner.Option[string]
	Type optioner.Option[string]
}

func (this KindCommon) Validate() error {
	return nil
}

func (this KindCommon) HasValidation() bool {
	return false
}

// //////////////////////////////////////////////////////////////////////////
// Kind Builder
// //////////////////////////////////////////////////////////////////////////

type kindBuilder[K any] interface {
	Done() Kind
	Name(v string) *K
	Type(v string) *K
}
