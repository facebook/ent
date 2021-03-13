// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package fieldtype

import (
	"fmt"
	"net"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/role"
)

const (
	// Label holds the string label denoting the fieldtype type in the database.
	Label = "field_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInt holds the string denoting the int field in the database.
	FieldInt = "int"
	// FieldInt8 holds the string denoting the int8 field in the database.
	FieldInt8 = "int8"
	// FieldInt16 holds the string denoting the int16 field in the database.
	FieldInt16 = "int16"
	// FieldInt32 holds the string denoting the int32 field in the database.
	FieldInt32 = "int32"
	// FieldInt64 holds the string denoting the int64 field in the database.
	FieldInt64 = "int64"
	// FieldOptionalInt holds the string denoting the optional_int field in the database.
	FieldOptionalInt = "optional_int"
	// FieldOptionalInt8 holds the string denoting the optional_int8 field in the database.
	FieldOptionalInt8 = "optional_int8"
	// FieldOptionalInt16 holds the string denoting the optional_int16 field in the database.
	FieldOptionalInt16 = "optional_int16"
	// FieldOptionalInt32 holds the string denoting the optional_int32 field in the database.
	FieldOptionalInt32 = "optional_int32"
	// FieldOptionalInt64 holds the string denoting the optional_int64 field in the database.
	FieldOptionalInt64 = "optional_int64"
	// FieldNillableInt holds the string denoting the nillable_int field in the database.
	FieldNillableInt = "nillable_int"
	// FieldNillableInt8 holds the string denoting the nillable_int8 field in the database.
	FieldNillableInt8 = "nillable_int8"
	// FieldNillableInt16 holds the string denoting the nillable_int16 field in the database.
	FieldNillableInt16 = "nillable_int16"
	// FieldNillableInt32 holds the string denoting the nillable_int32 field in the database.
	FieldNillableInt32 = "nillable_int32"
	// FieldNillableInt64 holds the string denoting the nillable_int64 field in the database.
	FieldNillableInt64 = "nillable_int64"
	// FieldValidateOptionalInt32 holds the string denoting the validate_optional_int32 field in the database.
	FieldValidateOptionalInt32 = "validate_optional_int32"
	// FieldOptionalUint holds the string denoting the optional_uint field in the database.
	FieldOptionalUint = "optional_uint"
	// FieldOptionalUint8 holds the string denoting the optional_uint8 field in the database.
	FieldOptionalUint8 = "optional_uint8"
	// FieldOptionalUint16 holds the string denoting the optional_uint16 field in the database.
	FieldOptionalUint16 = "optional_uint16"
	// FieldOptionalUint32 holds the string denoting the optional_uint32 field in the database.
	FieldOptionalUint32 = "optional_uint32"
	// FieldOptionalUint64 holds the string denoting the optional_uint64 field in the database.
	FieldOptionalUint64 = "optional_uint64"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldOptionalFloat holds the string denoting the optional_float field in the database.
	FieldOptionalFloat = "optional_float"
	// FieldOptionalFloat32 holds the string denoting the optional_float32 field in the database.
	FieldOptionalFloat32 = "optional_float32"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"
	// FieldDecimal holds the string denoting the decimal field in the database.
	FieldDecimal = "decimal"
	// FieldDir holds the string denoting the dir field in the database.
	FieldDir = "dir"
	// FieldNdir holds the string denoting the ndir field in the database.
	FieldNdir = "ndir"
	// FieldStr holds the string denoting the str field in the database.
	FieldStr = "str"
	// FieldNullStr holds the string denoting the null_str field in the database.
	FieldNullStr = "null_str"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldLinkOther holds the string denoting the link_other field in the database.
	FieldLinkOther = "link_other"
	// FieldNullLink holds the string denoting the null_link field in the database.
	FieldNullLink = "null_link"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldNullActive holds the string denoting the null_active field in the database.
	FieldNullActive = "null_active"
	// FieldDeleted holds the string denoting the deleted field in the database.
	FieldDeleted = "deleted"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldNullInt64 holds the string denoting the null_int64 field in the database.
	FieldNullInt64 = "null_int64"
	// FieldSchemaInt holds the string denoting the schema_int field in the database.
	FieldSchemaInt = "schema_int"
	// FieldSchemaInt8 holds the string denoting the schema_int8 field in the database.
	FieldSchemaInt8 = "schema_int8"
	// FieldSchemaInt64 holds the string denoting the schema_int64 field in the database.
	FieldSchemaInt64 = "schema_int64"
	// FieldSchemaFloat holds the string denoting the schema_float field in the database.
	FieldSchemaFloat = "schema_float"
	// FieldSchemaFloat32 holds the string denoting the schema_float32 field in the database.
	FieldSchemaFloat32 = "schema_float32"
	// FieldNullFloat holds the string denoting the null_float field in the database.
	FieldNullFloat = "null_float"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldMAC holds the string denoting the mac field in the database.
	FieldMAC = "mac"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// Table holds the table name of the fieldtype in the database.
	Table = "field_types"
)

// Columns holds all SQL columns for fieldtype fields.
var Columns = []string{
	FieldID,
	FieldInt,
	FieldInt8,
	FieldInt16,
	FieldInt32,
	FieldInt64,
	FieldOptionalInt,
	FieldOptionalInt8,
	FieldOptionalInt16,
	FieldOptionalInt32,
	FieldOptionalInt64,
	FieldNillableInt,
	FieldNillableInt8,
	FieldNillableInt16,
	FieldNillableInt32,
	FieldNillableInt64,
	FieldValidateOptionalInt32,
	FieldOptionalUint,
	FieldOptionalUint8,
	FieldOptionalUint16,
	FieldOptionalUint32,
	FieldOptionalUint64,
	FieldDuration,
	FieldState,
	FieldOptionalFloat,
	FieldOptionalFloat32,
	FieldDatetime,
	FieldDecimal,
	FieldDir,
	FieldNdir,
	FieldStr,
	FieldNullStr,
	FieldLink,
	FieldLinkOther,
	FieldNullLink,
	FieldActive,
	FieldNullActive,
	FieldDeleted,
	FieldDeletedAt,
	FieldIP,
	FieldNullInt64,
	FieldSchemaInt,
	FieldSchemaInt8,
	FieldSchemaInt64,
	FieldSchemaFloat,
	FieldSchemaFloat32,
	FieldNullFloat,
	FieldRole,
	FieldMAC,
	FieldUUID,
}

// Columns holds all SQL unique columns for fieldtype fields that can be used in conflict constraints.
var UniqueColumns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "field_types"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"file_field",
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
	// ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	ValidateOptionalInt32Validator func(int32) error
	// NdirValidator is a validator for the "ndir" field. It is called by the builders before save.
	NdirValidator func(string) error
	// DefaultStr holds the default value on creation for the "str" field.
	DefaultStr func() sql.NullString
	// DefaultNullStr holds the default value on creation for the "null_str" field.
	DefaultNullStr func() sql.NullString
	// LinkValidator is a validator for the "link" field. It is called by the builders before save.
	LinkValidator func(string) error
	// DefaultIP holds the default value on creation for the "ip" field.
	DefaultIP func() net.IP
	// MACValidator is a validator for the "mac" field. It is called by the builders before save.
	MACValidator func(string) error
)

// State defines the type for the "state" enum field.
type State string

// State values.
const (
	StateOn  State = "on"
	StateOff State = "off"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateOn, StateOff:
		return nil
	default:
		return fmt.Errorf("fieldtype: invalid enum value for state field: %q", s)
	}
}

const DefaultRole role.Role = "READ"

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r role.Role) error {
	switch r {
	case "ADMIN", "OWNER", "USER", "READ", "WRITE":
		return nil
	default:
		return fmt.Errorf("fieldtype: invalid enum value for role field: %q", r)
	}
}

// Ptr returns a new pointer to the enum value.
func (s State) Ptr() *State {
	return &s
}

// comment from another template.
