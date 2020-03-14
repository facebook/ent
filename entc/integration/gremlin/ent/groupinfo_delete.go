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
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	hooks      []Hook
	mutation   *GroupInfoMutation
	predicates []predicate.GroupInfo
}

// Where adds a new predicate to the delete builder.
func (gid *GroupInfoDelete) Where(ps ...predicate.GroupInfo) *GroupInfoDelete {
	gid.predicates = append(gid.predicates, ps...)
	return gid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gid *GroupInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gid.hooks) == 0 {
		affected, err = gid.gremlinExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gid.mutation = mutation
			affected, err = gid.gremlinExec(ctx)
			return affected, err
		})
		for i := len(gid.hooks); i > 0; i-- {
			mut = gid.hooks[i-1](mut)
		}
		if _, err := mut.Mutate(ctx, gid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gid *GroupInfoDelete) ExecX(ctx context.Context) int {
	n, err := gid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gid *GroupInfoDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gid.gremlin().Query()
	if err := gid.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (gid *GroupInfoDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(groupinfo.Label)
	for _, p := range gid.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// GroupInfoDeleteOne is the builder for deleting a single GroupInfo entity.
type GroupInfoDeleteOne struct {
	gid *GroupInfoDelete
}

// Exec executes the deletion query.
func (gido *GroupInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := gido.gid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gido *GroupInfoDeleteOne) ExecX(ctx context.Context) {
	gido.gid.ExecX(ctx)
}
