package sql

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

func setAndReturn[B any, V any](builder B, field *V, value V) B {
	*field = value
	return builder
}

func appendAndReturn[B any, V any](builder B, field *[]V, values ...V) B {
	*field = append(*field, values...)
	return builder
}

func columnNames(columns []*ColumnSchema) []string {
	return slicer.Map(
		func(_ int, column *ColumnSchema) string { return column.Name },
		columns...)
}
