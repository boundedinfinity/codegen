package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Database() *DatabaseSchema {
	return &DatabaseSchema{
		Tables: map[string]*TableSchema{},
	}
}

type DatabaseSchema struct {
	Tables      map[string]*TableSchema
	Selects     []*SelectSchema
	ForeignKeys bool
	Formatted   bool
	errs        []error
}

func (this *DatabaseSchema) Error() error {
	return errors.Join(this.errs...)
}

func (this *DatabaseSchema) checkError() {
	if this.Error() != nil {
		panic(this.Error())
	}
}

func (this *DatabaseSchema) ManyToMany(domestic, foreign *TableSchema) *DatabaseSchema {
	link_table := &TableSchema{
		Name:         fmt.Sprintf("%s_%s", domestic.Name, foreign.Name),
		WithoutRowId: true,
	}
	this.Table(link_table)
	this.ReferencedTo(link_table, domestic)
	this.ReferencedTo(link_table, foreign)

	return this
}

func (this *DatabaseSchema) OneToMany(domestic, foreign *TableSchema) *DatabaseSchema {
	var domesticForeignKeys []*ForeignKeySchema
	var foreignDomesticKeys []*ForeignKeySchema

	for _, primaryKey := range domestic.primaryKeyColumns() {
		foreignColumn := &ColumnSchema{
			Name: fmt.Sprintf("%s_%s", primaryKey.Table.Name, primaryKey.Name),
			Type: primaryKey.Type,
		}

		domesticForeignKey := &ForeignKeySchema{
			Foreign:  foreignColumn,
			Domestic: primaryKey,
		}

		foreignDomesticKey := &ForeignKeySchema{
			Foreign:  primaryKey,
			Domestic: foreignColumn,
		}

		domesticForeignKeys = append(domesticForeignKeys, domesticForeignKey)
		foreignDomesticKeys = append(foreignDomesticKeys, foreignDomesticKey)
	}

	for _, foreignKey := range domesticForeignKeys {
		foreign.Column(foreignKey.Foreign)
		domestic.AddForeignKey(foreignKey)
	}

	for _, foreignKey := range foreignDomesticKeys {
		foreign.AddForeignKey(foreignKey)
	}

	return this
}

func (this *DatabaseSchema) OneToOne(domestic, foreign *TableSchema) *DatabaseSchema {
	var domesticForeignKeys []*ForeignKeySchema
	var foreignDomesticKeys []*ForeignKeySchema

	for _, primaryKey := range domestic.primaryKeyColumns() {
		foreignColumn := &ColumnSchema{
			Name:   fmt.Sprintf("%s_%s", primaryKey.Table.Name, primaryKey.Name),
			Type:   primaryKey.Type,
			Unique: true,
		}

		domesticForeignKey := &ForeignKeySchema{
			Foreign:  foreignColumn,
			Domestic: primaryKey,
		}

		foreignDomesticKey := &ForeignKeySchema{
			Foreign:  primaryKey,
			Domestic: foreignColumn,
		}

		domesticForeignKeys = append(domesticForeignKeys, domesticForeignKey)
		foreignDomesticKeys = append(foreignDomesticKeys, foreignDomesticKey)
	}

	for _, foreignKey := range domesticForeignKeys {
		foreign.Column(foreignKey.Foreign)
		domestic.AddForeignKey(foreignKey)
	}

	for _, foreignKey := range foreignDomesticKeys {
		foreign.AddForeignKey(foreignKey)
	}

	return this
}

func (this *DatabaseSchema) ReferencedTo(domestic, foreign *TableSchema) *DatabaseSchema {
	var foreignKeys []*ForeignKeySchema

	for _, primaryKey := range foreign.primaryKeyColumns() {
		domesticColumn := &ColumnSchema{
			Name:    fmt.Sprintf("%s_%s", primaryKey.Table.Name, primaryKey.Name),
			Type:    primaryKey.Type,
			Indexed: true,
		}

		foreignKey := &ForeignKeySchema{
			Foreign:  primaryKey,
			Domestic: domesticColumn,
		}

		foreignKeys = append(foreignKeys, foreignKey)
	}

	for _, foreignKey := range foreignKeys {
		domestic.Column(foreignKey.Domestic)
		domestic.AddForeignKey(foreignKey)
	}

	return this
}

func (this *DatabaseSchema) Generate() string {
	this.checkError()

	var sb stringBuiler

	var tableTexts []string
	sep := " "

	if this.Formatted {
		sep = "\n\n"
	}

	if this.ForeignKeys {
		tableTexts = append(tableTexts, "PRAGMA foreign_keys = ON;")
	} else {
		tableTexts = append(tableTexts, "PRAGMA foreign_keys = OFF;")
	}

	for _, table := range this.Tables {
		tableTexts = append(tableTexts, table.Generate())

		for _, column := range table.indexedColumns() {
			indexed := &IndexSchema{Column: column}
			tableTexts = append(tableTexts, indexed.Generate())
		}
	}

	sb.Writef(stringer.Join(sep, tableTexts...))

	return sb.String()
}

func (this DatabaseSchema) WriteSqlFile(path string) error {
	content := this.Generate()
	return os.WriteFile(path, []byte(content), os.FileMode(0755))
}

func (this DatabaseSchema) CreateTables(db sql.DB) error {
	if this.Error() != nil {
		return this.Error()
	}

	return nil
}

func (this *DatabaseSchema) GetTable(name string) *TableSchema {
	var found *TableSchema

	for _, table := range this.Tables {
		if table.Name == name {
			found = table
			break
		}
	}

	if found == nil {
		panic(&ErrDatabaseTableNotFoundDetails{TableName: name})
	}

	return found
}

func (this *DatabaseSchema) Table(table *TableSchema) *DatabaseSchema {
	if this.Formatted {
		table.Formatted = this.Formatted
	}

	this.Tables[table.Name] = table
	return this
}

func (this *DatabaseSchema) SetForeignKeys(enabled bool) *DatabaseSchema {
	return setAndReturn(this, &this.ForeignKeys, enabled)
}

func (this *DatabaseSchema) SetFormatted(enabled bool) *DatabaseSchema {
	return setAndReturn(this, &this.Formatted, enabled)
}
