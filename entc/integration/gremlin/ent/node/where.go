// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package node

import (
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Node {
	return predicate.Node(
		func(t *dsl.Traversal) {
			t.HasID(id)
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	},
	)
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.EQ(v))
	},
	)
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.EQ(v))
	},
	)
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.NEQ(v))
	},
	)
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.Within(v...))
	},
	)
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.Without(v...))
	},
	)
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.GT(v))
	},
	)
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.GTE(v))
	},
	)
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.LT(v))
	},
	)
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.Has(Label, FieldValue, p.LTE(v))
	},
	)
}

// ValueIsNil applies the IsNil predicate on the "value" field.
func ValueIsNil() predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldValue)
	},
	)
}

// ValueNotNil applies the NotNil predicate on the "value" field.
func ValueNotNil() predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldValue)
	},
	)
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.InE(PrevInverseLabel).InV()
	},
	)
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(PrevInverseLabel).Where(tr).InV()
	},
	)
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		t.OutE(NextLabel).OutV()
	},
	)
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(NextLabel).Where(tr).OutV()
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.And(trs...))
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.Or(trs...))
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(
		func(tr *dsl.Traversal) {
			t := __.New()
			p(t)
			tr.Where(__.Not(t))
		},
	)
}
