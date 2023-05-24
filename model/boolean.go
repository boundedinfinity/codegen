package model

import "boundedinfinity/codegen/model/type_id"

type Boolean struct {
	common
}

func (t Boolean) TypeId() type_id.TypeId {
	return type_id.Boolean
}

var _ Type = &Boolean{}
