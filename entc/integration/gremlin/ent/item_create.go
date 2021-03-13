// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/item"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation         *ItemMutation
	hooks            []Hook
	constraintFields []string
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			node, err = ic.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	return nil
}

func (ic *ItemCreate) gremlinSave(ctx context.Context) (*Item, error) {
	res := &gremlin.Response{}
	query, bindings := ic.gremlin().Query()
	if err := ic.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	i := &Item{config: ic.config}
	if err := i.FromResponse(res); err != nil {
		return nil, err
	}
	return i, nil
}

func (ic *ItemCreate) gremlin() *dsl.Traversal {
	v := g.AddV(item.Label)
	return v.ValueMap(true)
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	builders              []*ItemCreate
	batchConstraintFields []string
}
