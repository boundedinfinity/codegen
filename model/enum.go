package model

import (
	"boundedinfinity/codegen/model/type_id"
)

type Enum struct {
	Common
}

func (t Enum) TypeId() type_id.TypeId {
	return type_id.String
}

var _ Type = &Enum{}
