package sql

import "fmt"

func Column() *ColumnSchema {
	return &ColumnSchema{}
}

type ColumnSchema struct {
	Table             *TableSchema
	Name              string
	Type              ColumnType
	NotNull           bool
	PrimaryKey        bool
	primayKeyDisabled bool
	Unique            bool
	Check             bool
	AutoIncrement     bool
	Indexed           bool
	UniqueIndexed     bool
}

func (this *ColumnSchema) qualifiedName() string {
	return fmt.Sprintf("%s.%s", this.Table.Name, this.Name)
}

func (this ColumnSchema) Generate() string {
	var sb stringBuiler

	sb.Writef("%s %s", quote(this.Name), this.Type)

	if this.PrimaryKey && !this.primayKeyDisabled {
		sb.Writef(" PRIMARY KEY")
	}

	if this.NotNull {
		sb.Writef(" NOT NULL")
	}

	if this.Unique {
		sb.Writef(" UNIQUE")
	}

	if this.Check {
		sb.Writef(" CHECK")
	}

	if this.AutoIncrement {
		sb.Writef(" AUTOINCREMENT")
	}

	return sb.String()
}

func (this *ColumnSchema) SetName(name string) *ColumnSchema {
	return setAndReturn(this, &this.Name, name)
}

func (this *ColumnSchema) SetType(typ ColumnType) *ColumnSchema {
	return setAndReturn(this, &this.Type, typ)
}

func (this *ColumnSchema) SetNotNull(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.NotNull, enabed)
}

func (this *ColumnSchema) SetPrimaryKey(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.PrimaryKey, enabed)
}

func (this *ColumnSchema) SetUnique(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.Unique, enabed)
}

func (this *ColumnSchema) SetCheck(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.Check, enabed)
}

func (this *ColumnSchema) SetAutoIncrement(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.AutoIncrement, enabed)
}

func (this *ColumnSchema) SetIndexed(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.Indexed, enabed)
}

func (this *ColumnSchema) SetUniqueIndexed(enabed bool) *ColumnSchema {
	return setAndReturn(this, &this.UniqueIndexed, enabed)
}

type ColumnType string

type columnTypes struct {
	NULL    ColumnType
	INTEGER ColumnType
	REAL    ColumnType
	TEXT    ColumnType
	BLOB    ColumnType
}

var ColumnTypes = columnTypes{
	NULL:    "NULL",
	INTEGER: "INTEGER",
	REAL:    "REAL",
	TEXT:    "TEXT",
	BLOB:    "BLOB",
}
