// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.\n// This source code is licensed under the Apache 2.0 license found\n// in the LICENSE file in the root directory of this source tree.\n\n// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldCustom holds the string denoting the custom vertex property in the database.
	FieldCustom = "custom"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCustom,
}
