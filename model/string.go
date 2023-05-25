package model

import (
	"boundedinfinity/codegen/model/type_id"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type String struct {
	Common
	Min   optioner.Option[int]    `json:"min,omitempty"`
	Max   optioner.Option[int]    `json:"max,omitempty"`
	Regex optioner.Option[string] `json:"regex,omitempty"`
}

func NewString(params ...ParamFunc[String]) String {
	t := String{}

	// handleParams(&t, params)

	return t
}

func (t String) TypeId() type_id.TypeId {
	return type_id.String
}

var _ Type = &String{}
