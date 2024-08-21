package sql

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Table() *TableSchema {
	return &TableSchema{}
}

type TableSchema struct {
	Schema       optioner.Option[string]
	Name         string
	IfNotExists  bool
	WithoutRowId bool
	Columns      []*ColumnSchema
	ForeignKeys  []*ForeignKeySchema
	Formatted    bool
}

func (this *TableSchema) Generate() string {
	this.ensureId()

	var sb stringBuiler

	sb.Writef("CREATE TABLE")

	if this.IfNotExists {
		sb.Writef(" IF NOT EXISTS")
	}

	if this.Schema.Defined() {
		sb.Writef(" %s.%s", quote(this.Schema.Get()), quote(this.Name))
	} else {
		sb.Writef(" %s (", quote(this.Name))
	}

	if this.Formatted {
		sb.Writef("\n")
	}

	primaryKeys := this.primaryKeyColumns()

	switch len(primaryKeys) {
	case 0:
		// no primary key
	case 1:
		// only one primary key, set on column description
	default:
		for _, column := range primaryKeys {
			column.primayKeyDisabled = true
		}
	}

	var columnTexts []string

	for _, column := range this.Columns {
		columnTexts = append(columnTexts, column.Generate())
	}

	if len(primaryKeys) > 1 {
		names := stringer.Join(", ", getColumnNames(this.primaryKeyColumns())...)
		columnTexts = append(columnTexts, fmt.Sprintf("PRIMARY KEY (%s)", names))
	}

	for _, foreignKey := range this.ForeignKeys {
		columnTexts = append(columnTexts, foreignKey.Generate())
	}

	if this.Formatted {
		columnTexts = slicer.Map(func(_ int, s string) string { return "    " + s }, columnTexts...)
		sb.Writef(stringer.Join(",\n", columnTexts...))
		sb.Writef("\n")
	} else {
		sb.Writef(stringer.Join(", ", columnTexts...))
		sb.Writef(" ")
	}

	sb.Writef(")")

	if this.WithoutRowId {
		sb.Writef(" WITHOUT ROWID")
	}

	sb.Writef(";")

	return sb.String()
}

func (this TableSchema) ensureId() {
	for _, column := range this.Columns {
		if column.PrimaryKey {
			return
		}
	}

	this.Column(&ColumnSchema{Name: "id", Type: ColumnTypes.TEXT, PrimaryKey: true})
}

func (this *TableSchema) AddForeignKey(fk *ForeignKeySchema) *TableSchema {
	this.ForeignKeys = append(this.ForeignKeys, fk)
	return this
}

func (this TableSchema) primaryKeyColumns() []*ColumnSchema {
	return slicer.Filter(
		func(_ int, column *ColumnSchema) bool { return column.PrimaryKey },
		this.Columns...)
}

func (this TableSchema) indexedColumns() []*ColumnSchema {
	return slicer.Filter(
		func(_ int, column *ColumnSchema) bool { return column.Indexed },
		this.Columns...)
}

func (this *TableSchema) SetName(name string) *TableSchema {
	return setAndReturn(this, &this.Name, name)
}

func (this *TableSchema) SetWithoutRowId(enabed bool) *TableSchema {
	return setAndReturn(this, &this.WithoutRowId, enabed)
}

func (this *TableSchema) SetIsNotExists(enabed bool) *TableSchema {
	return setAndReturn(this, &this.IfNotExists, enabed)
}

func (this *TableSchema) GetColumn(name string) *ColumnSchema {
	var found *ColumnSchema

	for _, column := range this.Columns {
		if column.Name == name {
			found = column
			break
		}
	}

	if found == nil {
		panic(&ErrTableColumnNotFoundDetails{TableName: this.Name, ColumnName: name})
	}

	return found
}

func (this *TableSchema) Column(column *ColumnSchema) *TableSchema {
	column.Table = this
	this.Columns = append(this.Columns, column)
	return this
}

func (this *TableSchema) qualifiedNames() []string {
	return slicer.Map(func(_ int, column *ColumnSchema) string {
		return column.qualifiedName()
	}, this.Columns...)
}
