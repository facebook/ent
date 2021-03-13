// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package file

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "fsize"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// FieldOp holds the string denoting the op field in the database.
	FieldOp = "op"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// EdgeField holds the string denoting the field edge name in mutations.
	EdgeField = "field"
	// Table holds the table name of the file in the database.
	Table = "files"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "files"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_files"
	// TypeTable is the table the holds the type relation/edge.
	TypeTable = "files"
	// TypeInverseTable is the table name for the FileType entity.
	// It exists in this package in order to avoid circular dependency with the "filetype" package.
	TypeInverseTable = "file_types"
	// TypeColumn is the table column denoting the type relation/edge.
	TypeColumn = "file_type_files"
	// FieldTable is the table the holds the field relation/edge.
	FieldTable = "field_types"
	// FieldInverseTable is the table name for the FieldType entity.
	// It exists in this package in order to avoid circular dependency with the "fieldtype" package.
	FieldInverseTable = "field_types"
	// FieldColumn is the table column denoting the field relation/edge.
	FieldColumn = "file_field"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldSize,
	FieldName,
	FieldUser,
	FieldGroup,
	FieldOp,
}

// Columns holds all SQL unique columns for file fields that can be used in conflict constraints.
var UniqueColumns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "files"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"file_type_files",
	"group_files",
	"user_files",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// ValidConstraintColumn reports if the column name is valid for use as a conflict column.
func ValidConstraintColumn(column string) bool {
	for i := range UniqueColumns {
		if column == UniqueColumns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSize holds the default value on creation for the "size" field.
	DefaultSize int
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int) error
)

// comment from another template.
