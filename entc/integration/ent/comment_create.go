// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/comment"
	"entgo.io/ent/schema/field"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation         *CommentMutation
	hooks            []Hook
	constraintFields []string
}

// SetUniqueInt sets the "unique_int" field.
func (cc *CommentCreate) SetUniqueInt(i int) *CommentCreate {
	cc.mutation.SetUniqueInt(i)
	return cc
}

// SetUniqueFloat sets the "unique_float" field.
func (cc *CommentCreate) SetUniqueFloat(f float64) *CommentCreate {
	cc.mutation.SetUniqueFloat(f)
	return cc
}

// SetNillableInt sets the "nillable_int" field.
func (cc *CommentCreate) SetNillableInt(i int) *CommentCreate {
	cc.mutation.SetNillableInt(i)
	return cc
}

// SetNillableNillableInt sets the "nillable_int" field if the given value is not nil.
func (cc *CommentCreate) SetNillableNillableInt(i *int) *CommentCreate {
	if i != nil {
		cc.SetNillableInt(*i)
	}
	return cc
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	err = cc.validateUpsertConstraints()
	if err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.UniqueInt(); !ok {
		return &ValidationError{Name: "unique_int", err: errors.New("ent: missing required field \"unique_int\"")}
	}
	if _, ok := cc.mutation.UniqueFloat(); !ok {
		return &ValidationError{Name: "unique_float", err: errors.New("ent: missing required field \"unique_float\"")}
	}
	return nil
}

// OnConflict specifies how to handle inserts that conflict with a unique constraint on Comment entities.
func (cc *CommentCreate) OnConflict(constraintField string, otherFields ...string) *CommentCreate {
	cc.constraintFields = []string{constraintField}
	cc.constraintFields = append(cc.constraintFields, otherFields...)
	return cc
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: comment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		}
	)

	if cc.constraintFields != nil {
		_spec.ConstraintFields = cc.constraintFields
	}
	if value, ok := cc.mutation.UniqueInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUniqueInt,
		})
		_node.UniqueInt = value
	}
	if value, ok := cc.mutation.UniqueFloat(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: comment.FieldUniqueFloat,
		})
		_node.UniqueFloat = value
	}
	if value, ok := cc.mutation.NillableInt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldNillableInt,
		})
		_node.NillableInt = &value
	}
	return _node, _spec
}

// validateUpsertConstraints validates the columns specified in OnConflict are valid for
// handling conflicts on Comment entities.
func (cc *CommentCreate) validateUpsertConstraints() error {
	for _, f := range cc.constraintFields {
		if !comment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for upsert conflict resolution", f)}
		}
	}
	return nil
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	builders              []*CommentCreate
	batchConstraintFields []string
}

// OnConflict specifies how to handle bulk inserts that conflict with a unique constraint on Comment entities.
func (ccb *CommentCreateBulk) OnConflict(constraintField string, otherFields ...string) *CommentCreateBulk {
	ccb.batchConstraintFields = []string{constraintField}
	ccb.batchConstraintFields = append(ccb.batchConstraintFields, otherFields...)

	return ccb
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	err := ccb.validateUpsertConstraints()
	if err != nil {
		return nil, err
	}

	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs, BatchConstraintFields: ccb.batchConstraintFields}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// validateUpsertConstraints validates the columns specified in OnConflict are valid for
// handling conflicts on batch inserted Comment entities.
func (ccb *CommentCreateBulk) validateUpsertConstraints() error {
	for _, f := range ccb.batchConstraintFields {
		if !comment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for upsert conflict resolution", f)}
		}
	}
	return nil
}
