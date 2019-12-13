// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package card

import (
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	},
	)
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	},
	)
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	},
	)
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	},
	)
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	},
	)
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	},
	)
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	},
	)
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	},
	)
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	},
	)
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	},
	)
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	},
	)
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	},
	)
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	},
	)
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	},
	)
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	},
	)
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	},
	)
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	},
	)
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	},
	)
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	},
	)
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	},
	)
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumber), v))
	},
	)
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumber), v...))
	},
	)
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumber), v...))
	},
	)
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumber), v))
	},
	)
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumber), v))
	},
	)
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumber), v))
	},
	)
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumber), v))
	},
	)
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNumber), v))
	},
	)
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNumber), v))
	},
	)
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNumber), v))
	},
	)
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNumber), v))
	},
	)
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNumber), v))
	},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Card(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	},
	)
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	},
	)
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	},
	)
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	},
	)
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	},
	)
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sql.NewStep(
			sql.From(Table, FieldID),
			sql.To(OwnerTable, FieldID),
			sql.Edge(sql.O2O, true, OwnerTable, OwnerColumn),
		)
		sql.HasNeighbors(s, step)
	},
	)
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.Card(func(s *sql.Selector) {
		step := sql.NewStep(
			sql.From(Table, FieldID),
			sql.To(OwnerInverseTable, FieldID),
			sql.Edge(sql.O2O, true, OwnerTable, OwnerColumn),
		)
		sql.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
