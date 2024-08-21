package sql

import (
	"errors"
	"fmt"
)

// ====================================================================================
// Database Errors
// ====================================================================================

var ErrDatabaseTableDuplicate = errors.New("duplicate table")

type ErrDatabaseTableDuplicateDetails struct {
	tables []*TableSchema
}

func (this ErrDatabaseTableDuplicateDetails) Error() string {
	return fmt.Sprintf("%s : %d duplicates",
		ErrDatabaseTableDuplicate.Error(),
		len(this.tables),
	)
}

func (this ErrDatabaseTableDuplicateDetails) Unwrap() error {
	return ErrDatabaseTableDuplicate
}

var ErrDatabaseTableNotFound = errors.New("table not found")

type ErrDatabaseTableNotFoundDetails struct {
	TableName string
}

func (e ErrDatabaseTableNotFoundDetails) Error() string {
	return fmt.Sprintf("%s : %s", ErrDatabaseTableNotFound.Error(), e.TableName)
}

func (e ErrDatabaseTableNotFoundDetails) Unwrap() error {
	return ErrDatabaseTableNotFound
}

// ====================================================================================
// Table Errors
// ====================================================================================

var ErrTableColumnNotFound = errors.New("column not found")

type ErrTableColumnNotFoundDetails struct {
	TableName  string
	ColumnName string
}

func (e ErrTableColumnNotFoundDetails) Error() string {
	return fmt.Sprintf("%s : %s.%s", ErrTableColumnNotFound.Error(), e.TableName, e.ColumnName)
}

func (e ErrTableColumnNotFoundDetails) Unwrap() error {
	return ErrTableColumnNotFound
}

// ====================================================================================
// Select Errors
// ====================================================================================

var ErrSelectNoColumns = errors.New("no table or columns")
