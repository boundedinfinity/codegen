package sql

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func ForeignKey() *ForeignKeySchema {
	return &ForeignKeySchema{}
}

type ForeignKeySchema struct {
	Foreign  *ColumnSchema
	Domestic *ColumnSchema
	OnUpdate optioner.Option[ForeignKeyAction]
	OnDelete optioner.Option[ForeignKeyAction]
}

func (this ForeignKeySchema) Generate() string {
	var sb stringBuiler

	sb.Writef(
		"FOREIGN KEY (%s) REFERENCES %s(%s)",
		quote(this.Domestic.Name),
		quote(this.Foreign.Table.Name),
		quote(this.Foreign.Name),
	)

	if this.OnDelete.Defined() {
		sb.Writef("ON DELETE %s", this.OnDelete.Get())
	}

	if this.OnUpdate.Defined() {
		sb.Writef("ON UPDATE %s", this.OnUpdate.Get())
	}

	return sb.String()
}

type ForeignKeyAction string

type foreignKeyActions struct {
	SET_NULL    ForeignKeyAction
	SET_DEFAULT ForeignKeyAction
	RESTRICT    ForeignKeyAction
	NO_ACTION   ForeignKeyAction
	CASCADE     ForeignKeyAction
}

var ForeignKeyActions = foreignKeyActions{
	SET_NULL:    "SET NULL",
	SET_DEFAULT: "SET DEFAULT",
	RESTRICT:    "RESTRICT",
	NO_ACTION:   "NO ACTION",
	CASCADE:     "CASCADE",
}
