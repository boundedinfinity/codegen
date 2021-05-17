package util

import (
	"boundedinfinity/codegen/model"
)

var (
	SchemaTypePrimitives = []model.SchemaTypeEnum{
		model.SchemaType_String,
		model.SchemaType_Long,
		model.SchemaType_Int,
		model.SchemaType_Float,
		model.SchemaType_Double,
		model.SchemaType_Byte,
		model.SchemaType_Boolean,
	}
)

func IsSchemaPrimitive(v model.SchemaTypeEnum) bool {
	for _, x := range SchemaTypePrimitives {
		if v == x {
			return true
		}
	}

	return false
}

func IsSchemaPrimitiveS(v string) bool {
	for _, x := range SchemaTypePrimitives {
		if v == string(x) {
			return true
		}
	}

	return false
}
