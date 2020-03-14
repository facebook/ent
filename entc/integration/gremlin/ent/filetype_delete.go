// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// FileTypeDelete is the builder for deleting a FileType entity.
type FileTypeDelete struct {
	config
	hooks      []Hook
	mutation   *FileTypeMutation
	predicates []predicate.FileType
}

// Where adds a new predicate to the delete builder.
func (ftd *FileTypeDelete) Where(ps ...predicate.FileType) *FileTypeDelete {
	ftd.predicates = append(ftd.predicates, ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FileTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftd.hooks) == 0 {
		affected, err = ftd.gremlinExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftd.mutation = mutation
			affected, err = ftd.gremlinExec(ctx)
			return affected, err
		})
		for i := len(ftd.hooks); i > 0; i-- {
			mut = ftd.hooks[i-1](mut)
		}
		if _, err := mut.Mutate(ctx, ftd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FileTypeDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftd.gremlin().Query()
	if err := ftd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftd *FileTypeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(filetype.Label)
	for _, p := range ftd.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// FileTypeDeleteOne is the builder for deleting a single FileType entity.
type FileTypeDeleteOne struct {
	ftd *FileTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FileTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{filetype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FileTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
