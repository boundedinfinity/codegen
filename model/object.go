package model

import "boundedinfinity/codegen/model/type_id"

type Object struct {
	common
	Properties []Type
}

func (t Object) TypeId() type_id.TypeId {
	return type_id.Float
}

var _ Type = &Object{}
