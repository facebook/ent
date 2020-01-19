// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/card"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Expired holds the value of the "expired" field.
	Expired time.Time `json:"expired,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardQuery when eager-loading is set.
	Edges struct {
		// Owner holds the value of the owner edge.
		Owner *User
	} `json:"edges"`
	owner_id *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Card) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // expired
		&sql.NullString{}, // number
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Card) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Card fields.
func (c *Card) assignValues(values ...interface{}) error {
	if m, n := len(values), len(card.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field expired", values[0])
	} else if value.Valid {
		c.Expired = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field number", values[1])
	} else if value.Valid {
		c.Number = value.String
	}
	values = values[2:]
	if len(values) == len(card.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			c.owner_id = new(int)
			*c.owner_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Card.
func (c *Card) QueryOwner() *UserQuery {
	return (&CardClient{c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Card.
// Note that, you need to call Card.Unwrap() before calling this method, if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return (&CardClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", expired=")
	builder.WriteString(c.Expired.Format(time.ANSIC))
	builder.WriteString(", number=")
	builder.WriteString(c.Number)
	builder.WriteByte(')')
	return builder.String()
}

// Cards is a parsable slice of Card.
type Cards []*Card

func (c Cards) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
