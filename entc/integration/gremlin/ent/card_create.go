// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/card"
	"entgo.io/ent/entc/integration/gremlin/ent/spec"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	mutation         *CardMutation
	hooks            []Hook
	constraintFields []string
}

// SetCreateTime sets the "create_time" field.
func (cc *CardCreate) SetCreateTime(t time.Time) *CardCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableCreateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *CardCreate) SetUpdateTime(t time.Time) *CardCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *CardCreate) SetNillableUpdateTime(t *time.Time) *CardCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetBalance sets the "balance" field.
func (cc *CardCreate) SetBalance(f float64) *CardCreate {
	cc.mutation.SetBalance(f)
	return cc
}

// SetNillableBalance sets the "balance" field if the given value is not nil.
func (cc *CardCreate) SetNillableBalance(f *float64) *CardCreate {
	if f != nil {
		cc.SetBalance(*f)
	}
	return cc
}

// SetNumber sets the "number" field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.mutation.SetNumber(s)
	return cc
}

// SetName sets the "name" field.
func (cc *CardCreate) SetName(s string) *CardCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CardCreate) SetNillableName(s *string) *CardCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CardCreate) SetOwnerID(id string) *CardCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cc *CardCreate) SetNillableOwnerID(id *string) *CardCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// AddSpecIDs adds the "spec" edge to the Spec entity by IDs.
func (cc *CardCreate) AddSpecIDs(ids ...string) *CardCreate {
	cc.mutation.AddSpecIDs(ids...)
	return cc
}

// AddSpec adds the "spec" edges to the Spec entity.
func (cc *CardCreate) AddSpec(s ...*Spec) *CardCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddSpecIDs(ids...)
}

// Mutation returns the CardMutation object of the builder.
func (cc *CardCreate) Mutation() *CardMutation {
	return cc.mutation
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	var (
		err  error
		node *Card
	)
	err = cc.validateUpsertConstraints()
	if err != nil {
		return nil, err
	}
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.gremlinSave(ctx)
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
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (cc *CardCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := card.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := card.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Balance(); !ok {
		v := card.DefaultBalance
		cc.mutation.SetBalance(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CardCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New("ent: missing required field \"create_time\"")}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New("ent: missing required field \"update_time\"")}
	}
	if _, ok := cc.mutation.Balance(); !ok {
		return &ValidationError{Name: "balance", err: errors.New("ent: missing required field \"balance\"")}
	}
	if _, ok := cc.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New("ent: missing required field \"number\"")}
	}
	if v, ok := cc.mutation.Number(); ok {
		if err := card.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf("ent: validator failed for field \"number\": %w", err)}
		}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := card.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (cc *CardCreate) gremlinSave(ctx context.Context) (*Card, error) {
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Card{config: cc.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CardCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(card.Label)
	if value, ok := cc.mutation.CreateTime(); ok {
		v.Property(dsl.Single, card.FieldCreateTime, value)
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, card.FieldUpdateTime, value)
	}
	if value, ok := cc.mutation.Balance(); ok {
		v.Property(dsl.Single, card.FieldBalance, value)
	}
	if value, ok := cc.mutation.Number(); ok {
		v.Property(dsl.Single, card.FieldNumber, value)
	}
	if value, ok := cc.mutation.Name(); ok {
		v.Property(dsl.Single, card.FieldName, value)
	}
	for _, id := range cc.mutation.OwnerIDs() {
		v.AddE(user.CardLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(card.Label, user.CardLabel, id)),
		})
	}
	for _, id := range cc.mutation.SpecIDs() {
		v.AddE(spec.CardLabel).From(g.V(id)).InV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// CardCreateBulk is the builder for creating many Card entities in bulk.
type CardCreateBulk struct {
	config
	builders              []*CardCreate
	batchConstraintFields []string
}
