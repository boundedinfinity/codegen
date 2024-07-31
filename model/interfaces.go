package model

type CodeGenType interface {
	Common() *CodeGenCommon
	GetType() string
	Validate() error
	HasValidation() bool
}
