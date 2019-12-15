// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/node"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	value *int
	prev  map[string]struct{}
	next  map[string]struct{}
}

// SetValue sets the value field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.value = &i
	return nc
}

// SetNillableValue sets the value field if the given value is not nil.
func (nc *NodeCreate) SetNillableValue(i *int) *NodeCreate {
	if i != nil {
		nc.SetValue(*i)
	}
	return nc
}

// SetPrevID sets the prev edge to Node by id.
func (nc *NodeCreate) SetPrevID(id string) *NodeCreate {
	if nc.prev == nil {
		nc.prev = make(map[string]struct{})
	}
	nc.prev[id] = struct{}{}
	return nc
}

// SetNillablePrevID sets the prev edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillablePrevID(id *string) *NodeCreate {
	if id != nil {
		nc = nc.SetPrevID(*id)
	}
	return nc
}

// SetPrev sets the prev edge to Node.
func (nc *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return nc.SetPrevID(n.ID)
}

// SetNextID sets the next edge to Node by id.
func (nc *NodeCreate) SetNextID(id string) *NodeCreate {
	if nc.next == nil {
		nc.next = make(map[string]struct{})
	}
	nc.next[id] = struct{}{}
	return nc
}

// SetNillableNextID sets the next edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillableNextID(id *string) *NodeCreate {
	if id != nil {
		nc = nc.SetNextID(*id)
	}
	return nc
}

// SetNext sets the next edge to Node.
func (nc *NodeCreate) SetNext(n *Node) *NodeCreate {
	return nc.SetNextID(n.ID)
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	if len(nc.prev) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"prev\"")
	}
	if len(nc.next) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"next\"")
	}
	return nc.gremlinSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NodeCreate) gremlinSave(ctx context.Context) (*Node, error) {
	res := &gremlin.Response{}
	query, bindings := nc.gremlin().Query()
	if err := nc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	n := &Node{config: nc.config}
	if err := n.FromResponse(res); err != nil {
		return nil, err
	}
	return n, nil
}

func (nc *NodeCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(node.Label)
	if nc.value != nil {
		v.Property(dsl.Single, node.FieldValue, *nc.value)
	}
	for id := range nc.prev {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	for id := range nc.next {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
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