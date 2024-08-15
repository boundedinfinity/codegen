// Package model contains schema typing information
package model

type CodeGenSchema interface {
	Common() *CodeGenCommon
	Schema() string
	Validate() error
	HasValidation() bool
}
