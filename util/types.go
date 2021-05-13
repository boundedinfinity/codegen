package util

import (
	"boundedinfinity/codegen/model"
	"strings"
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

func IsSchemaPrimitive(v model.InputModel) bool {
	if v.Type == "" {
		return false
	}

	var z string

	if IsSchemaArray(v) {
		z = TrimSchemaArray(v)
	} else {
		z = v.Type
	}

	for _, x := range SchemaTypePrimitives {
		if z == string(x) {
			return true
		}
	}

	return false
}

func IsSchemaRef(v model.InputModel) bool {
	return strings.Contains(v.Type, model.NAMESPACE_SEP)
}

func TrimSchemaArray(v model.InputModel) string {
	return strings.TrimSuffix(v.Type, model.COLLECTION_SUFFIX)
}

func IsSchemaArray(v model.InputModel) bool {
	return strings.HasSuffix(v.Type, model.COLLECTION_SUFFIX)
}

func IsSchemaRecord(v model.InputModel) bool {
	return v.Properties != nil
}

func IsSchemaEnum(v model.InputModel) bool {
	return v.Symbols != nil
}
