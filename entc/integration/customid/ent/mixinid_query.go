// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/customid/ent/mixinid"
	"github.com/facebook/ent/entc/integration/customid/ent/predicate"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// MixinIDQuery is the builder for querying MixinID entities.
type MixinIDQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	predicates []predicate.MixinID
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (miq *MixinIDQuery) Where(ps ...predicate.MixinID) *MixinIDQuery {
	miq.predicates = append(miq.predicates, ps...)
	return miq
}

// Limit adds a limit step to the query.
func (miq *MixinIDQuery) Limit(limit int) *MixinIDQuery {
	miq.limit = &limit
	return miq
}

// Offset adds an offset step to the query.
func (miq *MixinIDQuery) Offset(offset int) *MixinIDQuery {
	miq.offset = &offset
	return miq
}

// Order adds an order step to the query.
func (miq *MixinIDQuery) Order(o ...OrderFunc) *MixinIDQuery {
	miq.order = append(miq.order, o...)
	return miq
}

// First returns the first MixinID entity in the query. Returns *NotFoundError when no mixinid was found.
func (miq *MixinIDQuery) First(ctx context.Context) (*MixinID, error) {
	nodes, err := miq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mixinid.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (miq *MixinIDQuery) FirstX(ctx context.Context) *MixinID {
	node, err := miq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MixinID id in the query. Returns *NotFoundError when no id was found.
func (miq *MixinIDQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = miq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mixinid.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (miq *MixinIDQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := miq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only MixinID entity in the query, returns an error if not exactly one entity was returned.
func (miq *MixinIDQuery) Only(ctx context.Context) (*MixinID, error) {
	nodes, err := miq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mixinid.Label}
	default:
		return nil, &NotSingularError{mixinid.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (miq *MixinIDQuery) OnlyX(ctx context.Context) *MixinID {
	node, err := miq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only MixinID id in the query, returns an error if not exactly one id was returned.
func (miq *MixinIDQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = miq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = &NotSingularError{mixinid.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (miq *MixinIDQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := miq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MixinIDs.
func (miq *MixinIDQuery) All(ctx context.Context) ([]*MixinID, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return miq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (miq *MixinIDQuery) AllX(ctx context.Context) []*MixinID {
	nodes, err := miq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MixinID ids.
func (miq *MixinIDQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := miq.Select(mixinid.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (miq *MixinIDQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := miq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (miq *MixinIDQuery) Count(ctx context.Context) (int, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return miq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (miq *MixinIDQuery) CountX(ctx context.Context) int {
	count, err := miq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (miq *MixinIDQuery) Exist(ctx context.Context) (bool, error) {
	if err := miq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return miq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (miq *MixinIDQuery) ExistX(ctx context.Context) bool {
	exist, err := miq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (miq *MixinIDQuery) Clone() *MixinIDQuery {
	if miq == nil {
		return nil
	}
	return &MixinIDQuery{
		config:     miq.config,
		limit:      miq.limit,
		offset:     miq.offset,
		order:      append([]OrderFunc{}, miq.order...),
		predicates: append([]predicate.MixinID{}, miq.predicates...),
		// clone intermediate query.
		sql:  miq.sql.Clone(),
		path: miq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SomeField string `json:"some_field,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MixinID.Query().
//		GroupBy(mixinid.FieldSomeField).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (miq *MixinIDQuery) GroupBy(field string, fields ...string) *MixinIDGroupBy {
	group := &MixinIDGroupBy{config: miq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return miq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		SomeField string `json:"some_field,omitempty"`
//	}
//
//	client.MixinID.Query().
//		Select(mixinid.FieldSomeField).
//		Scan(ctx, &v)
//
func (miq *MixinIDQuery) Select(field string, fields ...string) *MixinIDSelect {
	selector := &MixinIDSelect{config: miq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := miq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return miq.sqlQuery(), nil
	}
	return selector
}

func (miq *MixinIDQuery) prepareQuery(ctx context.Context) error {
	if miq.path != nil {
		prev, err := miq.path(ctx)
		if err != nil {
			return err
		}
		miq.sql = prev
	}
	return nil
}

func (miq *MixinIDQuery) sqlAll(ctx context.Context) ([]*MixinID, error) {
	var (
		nodes = []*MixinID{}
		_spec = miq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &MixinID{config: miq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, miq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (miq *MixinIDQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := miq.querySpec()
	return sqlgraph.CountNodes(ctx, miq.driver, _spec)
}

func (miq *MixinIDQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := miq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (miq *MixinIDQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mixinid.Table,
			Columns: mixinid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		},
		From:   miq.sql,
		Unique: true,
	}
	if ps := miq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := miq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := miq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := miq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, mixinid.ValidColumn)
			}
		}
	}
	return _spec
}

func (miq *MixinIDQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(miq.driver.Dialect())
	t1 := builder.Table(mixinid.Table)
	selector := builder.Select(t1.Columns(mixinid.Columns...)...).From(t1)
	if miq.sql != nil {
		selector = miq.sql
		selector.Select(selector.Columns(mixinid.Columns...)...)
	}
	for _, p := range miq.predicates {
		p(selector)
	}
	for _, p := range miq.order {
		p(selector, mixinid.ValidColumn)
	}
	if offset := miq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := miq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MixinIDGroupBy is the builder for group-by MixinID entities.
type MixinIDGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (migb *MixinIDGroupBy) Aggregate(fns ...AggregateFunc) *MixinIDGroupBy {
	migb.fns = append(migb.fns, fns...)
	return migb
}

// Scan applies the group-by query and scan the result into the given value.
func (migb *MixinIDGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := migb.path(ctx)
	if err != nil {
		return err
	}
	migb.sql = query
	return migb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (migb *MixinIDGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := migb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MixinIDGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (migb *MixinIDGroupBy) StringsX(ctx context.Context) []string {
	v, err := migb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = migb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (migb *MixinIDGroupBy) StringX(ctx context.Context) string {
	v, err := migb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MixinIDGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (migb *MixinIDGroupBy) IntsX(ctx context.Context) []int {
	v, err := migb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = migb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (migb *MixinIDGroupBy) IntX(ctx context.Context) int {
	v, err := migb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MixinIDGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (migb *MixinIDGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := migb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = migb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (migb *MixinIDGroupBy) Float64X(ctx context.Context) float64 {
	v, err := migb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(migb.fields) > 1 {
		return nil, errors.New("ent: MixinIDGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := migb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (migb *MixinIDGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := migb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (migb *MixinIDGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = migb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (migb *MixinIDGroupBy) BoolX(ctx context.Context) bool {
	v, err := migb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (migb *MixinIDGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range migb.fields {
		if !mixinid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := migb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := migb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (migb *MixinIDGroupBy) sqlQuery() *sql.Selector {
	selector := migb.sql
	columns := make([]string, 0, len(migb.fields)+len(migb.fns))
	columns = append(columns, migb.fields...)
	for _, fn := range migb.fns {
		columns = append(columns, fn(selector, mixinid.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(migb.fields...)
}

// MixinIDSelect is the builder for select fields of MixinID entities.
type MixinIDSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (mis *MixinIDSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := mis.path(ctx)
	if err != nil {
		return err
	}
	mis.sql = query
	return mis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mis *MixinIDSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MixinIDSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mis *MixinIDSelect) StringsX(ctx context.Context) []string {
	v, err := mis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mis *MixinIDSelect) StringX(ctx context.Context) string {
	v, err := mis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MixinIDSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mis *MixinIDSelect) IntsX(ctx context.Context) []int {
	v, err := mis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mis *MixinIDSelect) IntX(ctx context.Context) int {
	v, err := mis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MixinIDSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mis *MixinIDSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mis *MixinIDSelect) Float64X(ctx context.Context) float64 {
	v, err := mis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mis.fields) > 1 {
		return nil, errors.New("ent: MixinIDSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mis *MixinIDSelect) BoolsX(ctx context.Context) []bool {
	v, err := mis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (mis *MixinIDSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{mixinid.Label}
	default:
		err = fmt.Errorf("ent: MixinIDSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mis *MixinIDSelect) BoolX(ctx context.Context) bool {
	v, err := mis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mis *MixinIDSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mis.fields {
		if !mixinid.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := mis.sqlQuery().Query()
	if err := mis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mis *MixinIDSelect) sqlQuery() sql.Querier {
	selector := mis.sql
	selector.Select(selector.Columns(mis.fields...)...)
	return selector
}
