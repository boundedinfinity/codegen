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

func getTableNames(tables []*TableSchema) []string {
	return slicer.Map(
		func(_ int, table *TableSchema) string { return table.Name },
		tables...)
}

func getColumnName(qualified bool, column *ColumnSchema) string {
	if qualified {
		return column.Table.Name + "." + column.Name
	}
	return column.Name
}

func getColumnNames(qualified bool, columns []*ColumnSchema) []string {
	return slicer.Map(
		func(_ int, column *ColumnSchema) string { return getColumnName(qualified, column) },
		columns...)
}
