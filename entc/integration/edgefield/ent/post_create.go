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
	"entgo.io/ent/entc/integration/edgefield/ent/post"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation         *PostMutation
	hooks            []Hook
	constraintFields []string
}

// SetText sets the "text" field.
func (pc *PostCreate) SetText(s string) *PostCreate {
	pc.mutation.SetText(s)
	return pc
}

// SetAuthorID sets the "author_id" field.
func (pc *PostCreate) SetAuthorID(i int) *PostCreate {
	pc.mutation.SetAuthorID(i)
	return pc
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableAuthorID(i *int) *PostCreate {
	if i != nil {
		pc.SetAuthorID(*i)
	}
	return pc
}

// SetAuthor sets the "author" edge to the User entity.
func (pc *PostCreate) SetAuthor(u *User) *PostCreate {
	return pc.SetAuthorID(u.ID)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	var (
		err  error
		node *Post
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New("ent: missing required field \"text\"")}
	}
	return nil
}

// OnConflict specifies how to handle inserts that conflict with a unique constraint on Post entities.
func (pc *PostCreate) OnConflict(constraintField string, otherFields ...string) *PostCreate {
	pc.constraintFields = []string{constraintField}
	pc.constraintFields = append(pc.constraintFields, otherFields...)
	return pc
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	err := pc.validateUpsertConstraints()
	if err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: post.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		}
	)

	if pc.constraintFields != nil {
		_spec.ConstraintFields = pc.constraintFields
	}
	if value, ok := pc.mutation.Text(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: post.FieldText,
		})
		_node.Text = value
	}
	if nodes := pc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   post.AuthorTable,
			Columns: []string{post.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AuthorID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// validateUpsertConstraints validates the columns specified in OnConflict are valid for
// handling conflicts on Post entities.
func (pc *PostCreate) validateUpsertConstraints() error {
	for _, f := range pc.constraintFields {
		if !post.ValidConstraintColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for upsert conflict resolution, valid fields are: %+v", f, post.UniqueColumns)}
		}
	}
	return nil
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders              []*PostCreate
	batchConstraintFields []string
}

// OnConflict specifies how to handle bulk inserts that conflict with a unique constraint on Post entities.
func (pcb *PostCreateBulk) OnConflict(constraintField string, otherFields ...string) *PostCreateBulk {
	pcb.batchConstraintFields = []string{constraintField}
	pcb.batchConstraintFields = append(pcb.batchConstraintFields, otherFields...)

	return pcb
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	err := pcb.validateUpsertConstraints()
	if err != nil {
		return nil, err
	}

	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs, BatchConstraintFields: pcb.batchConstraintFields}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// validateUpsertConstraints validates the columns specified in OnConflict are valid for
// handling conflicts on batch inserted Post entities.
func (pcb *PostCreateBulk) validateUpsertConstraints() error {
	for _, f := range pcb.batchConstraintFields {
		if !post.ValidConstraintColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for upsert conflict resolution, valid fields are: %+v", f, post.UniqueColumns)}
		}
	}
	return nil
}
