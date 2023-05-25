package model

import "boundedinfinity/codegen/model/type_id"

type Array struct {
	Common
	Items Type
}

func (t Array) TypeId() type_id.TypeId {
	return type_id.String
}

var _ Type = &Array{}
