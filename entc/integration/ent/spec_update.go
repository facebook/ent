// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/spec"
	"github.com/facebookincubator/ent/schema/field"
)

// SpecUpdate is the builder for updating Spec entities.
type SpecUpdate struct {
	config
	card        map[string]struct{}
	removedCard map[string]struct{}
	predicates  []predicate.Spec
}

// Where adds a new predicate for the builder.
func (su *SpecUpdate) Where(ps ...predicate.Spec) *SpecUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// AddCardIDs adds the card edge to Card by ids.
func (su *SpecUpdate) AddCardIDs(ids ...string) *SpecUpdate {
	if su.card == nil {
		su.card = make(map[string]struct{})
	}
	for i := range ids {
		su.card[ids[i]] = struct{}{}
	}
	return su
}

// AddCard adds the card edges to Card.
func (su *SpecUpdate) AddCard(c ...*Card) *SpecUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.AddCardIDs(ids...)
}

// RemoveCardIDs removes the card edge to Card by ids.
func (su *SpecUpdate) RemoveCardIDs(ids ...string) *SpecUpdate {
	if su.removedCard == nil {
		su.removedCard = make(map[string]struct{})
	}
	for i := range ids {
		su.removedCard[ids[i]] = struct{}{}
	}
	return su
}

// RemoveCard removes card edges to Card.
func (su *SpecUpdate) RemoveCard(c ...*Card) *SpecUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return su.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *SpecUpdate) Save(ctx context.Context) (int, error) {
	return su.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpecUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpecUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpecUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SpecUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   spec.Table,
			Columns: spec.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: spec.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nodes := su.removedCard; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: card.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.card; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: card.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spec.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SpecUpdateOne is the builder for updating a single Spec entity.
type SpecUpdateOne struct {
	config
	id          string
	card        map[string]struct{}
	removedCard map[string]struct{}
}

// AddCardIDs adds the card edge to Card by ids.
func (suo *SpecUpdateOne) AddCardIDs(ids ...string) *SpecUpdateOne {
	if suo.card == nil {
		suo.card = make(map[string]struct{})
	}
	for i := range ids {
		suo.card[ids[i]] = struct{}{}
	}
	return suo
}

// AddCard adds the card edges to Card.
func (suo *SpecUpdateOne) AddCard(c ...*Card) *SpecUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.AddCardIDs(ids...)
}

// RemoveCardIDs removes the card edge to Card by ids.
func (suo *SpecUpdateOne) RemoveCardIDs(ids ...string) *SpecUpdateOne {
	if suo.removedCard == nil {
		suo.removedCard = make(map[string]struct{})
	}
	for i := range ids {
		suo.removedCard[ids[i]] = struct{}{}
	}
	return suo
}

// RemoveCard removes card edges to Card.
func (suo *SpecUpdateOne) RemoveCard(c ...*Card) *SpecUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return suo.RemoveCardIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *SpecUpdateOne) Save(ctx context.Context) (*Spec, error) {
	return suo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpecUpdateOne) SaveX(ctx context.Context) *Spec {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *SpecUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpecUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SpecUpdateOne) sqlSave(ctx context.Context) (s *Spec, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   spec.Table,
			Columns: spec.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  suo.id,
				Type:   field.TypeString,
				Column: spec.FieldID,
			},
		},
	}
	if nodes := suo.removedCard; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: card.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.card; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   spec.CardTable,
			Columns: spec.CardPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: card.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Spec{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spec.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
