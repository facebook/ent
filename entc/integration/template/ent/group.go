// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/template/ent/group"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // max_users
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(values ...interface{}) error {
	if m, n := len(values), len(group.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	gr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field max_users", values[0])
	} else if value.Valid {
		gr.MaxUsers = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", max_users=")
	builder.WriteString(fmt.Sprintf("%v", gr.MaxUsers))
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
