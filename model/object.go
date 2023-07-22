package model

import "boundedinfinity/codegen/model/type_id"

type Object struct {
	Common
	Properties []Type
}

func (t Object) TypeId() type_id.TypeId {
	return type_id.Object
}

var _ Type = &Object{}
