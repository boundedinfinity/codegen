package util

import (
	"boundedinfinity/codegen/model"
)

var (
	SchemaTypePrimitives = []model.SchemaTypeEnum{
		model.SchemaType_Byte,
		model.SchemaType_String,
		model.SchemaType_Boolean,
		model.SchemaType_Long,
		model.SchemaType_Int,
		model.SchemaType_Float,
		model.SchemaType_Double,
	}
)

func IsSchemaSimpleType(v model.SchemaTypeEnum) bool {
	for _, x := range SchemaTypePrimitives {
		if v == x {
			return true
		}
	}

	return false
}

func IsSchemaSimpleTypeS(v string) bool {
	t, _ := model.SchemaTypeEnumParse(v)
	return IsSchemaSimpleType(t)
}
