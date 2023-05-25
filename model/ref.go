package model

import "boundedinfinity/codegen/model/type_id"

type Ref struct {
	Common
	Ref string
}

func (t Ref) TypeId() type_id.TypeId {
	return type_id.Float
}

var _ Type = &Ref{}
