// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/ent/filetype"
)

// FileType is the model entity for the FileType schema.
type FileType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type filetype.Type `json:"type,omitempty"`
	// State holds the value of the "state" field.
	State filetype.State `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileTypeQuery when eager-loading is set.
	Edges FileTypeEdges `json:"edges"`
}

// FileTypeEdges holds the relations/edges for other nodes in the graph.
type FileTypeEdges struct {
	// Files holds the value of the files edge.
	Files []*File
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e FileTypeEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[0] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FileType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // type
		&sql.NullString{}, // state
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FileType fields.
func (ft *FileType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(filetype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ft.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		ft.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[1])
	} else if value.Valid {
		ft.Type = filetype.Type(value.String)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[2])
	} else if value.Valid {
		ft.State = filetype.State(value.String)
	}
	return nil
}

// QueryFiles queries the files edge of the FileType.
func (ft *FileType) QueryFiles() *FileQuery {
	return (&FileTypeClient{config: ft.config}).QueryFiles(ft)
}

// Update returns a builder for updating this FileType.
// Note that, you need to call FileType.Unwrap() before calling this method, if this FileType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FileType) Update() *FileTypeUpdateOne {
	return (&FileTypeClient{config: ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ft *FileType) Unwrap() *FileType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FileType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FileType) String() string {
	var builder strings.Builder
	builder.WriteString("FileType(")
	builder.WriteString(fmt.Sprintf("id=%v", ft.ID))
	builder.WriteString(", name=")
	builder.WriteString(ft.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", ft.Type))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", ft.State))
	builder.WriteByte(')')
	return builder.String()
}

// FileTypes is a parsable slice of FileType.
type FileTypes []*FileType

func (ft FileTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
