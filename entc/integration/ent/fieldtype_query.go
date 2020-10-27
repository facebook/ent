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
	"github.com/facebook/ent/entc/integration/ent/fieldtype"
	"github.com/facebook/ent/entc/integration/ent/predicate"
	"github.com/facebook/ent/schema/field"
)

// FieldTypeQuery is the builder for querying FieldType entities.
type FieldTypeQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.FieldType
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (ftq *FieldTypeQuery) Where(ps ...predicate.FieldType) *FieldTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit adds a limit step to the query.
func (ftq *FieldTypeQuery) Limit(limit int) *FieldTypeQuery {
	ftq.limit = &limit
	return ftq
}

// Offset adds an offset step to the query.
func (ftq *FieldTypeQuery) Offset(offset int) *FieldTypeQuery {
	ftq.offset = &offset
	return ftq
}

// Order adds an order step to the query.
func (ftq *FieldTypeQuery) Order(o ...OrderFunc) *FieldTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// First returns the first FieldType entity in the query. Returns *NotFoundError when no fieldtype was found.
func (ftq *FieldTypeQuery) First(ctx context.Context) (*FieldType, error) {
	nodes, err := ftq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fieldtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstX(ctx context.Context) *FieldType {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FieldType id in the query. Returns *NotFoundError when no id was found.
func (ftq *FieldTypeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ftq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fieldtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstIDX(ctx context.Context) int {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only FieldType entity in the query, returns an error if not exactly one entity was returned.
func (ftq *FieldTypeQuery) Only(ctx context.Context) (*FieldType, error) {
	nodes, err := ftq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fieldtype.Label}
	default:
		return nil, &NotSingularError{fieldtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyX(ctx context.Context) *FieldType {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only FieldType id in the query, returns an error if not exactly one id was returned.
func (ftq *FieldTypeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ftq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = &NotSingularError{fieldtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyIDX(ctx context.Context) int {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FieldTypes.
func (ftq *FieldTypeQuery) All(ctx context.Context) ([]*FieldType, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ftq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FieldTypeQuery) AllX(ctx context.Context) []*FieldType {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FieldType ids.
func (ftq *FieldTypeQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ftq.Select(fieldtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FieldTypeQuery) IDsX(ctx context.Context) []int {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FieldTypeQuery) Count(ctx context.Context) (int, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ftq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FieldTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FieldTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ftq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ftq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FieldTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FieldTypeQuery) Clone() *FieldTypeQuery {
	if ftq == nil {
		return nil
	}
	return &FieldTypeQuery{
		config:     ftq.config,
		limit:      ftq.limit,
		offset:     ftq.offset,
		order:      append([]OrderFunc{}, ftq.order...),
		unique:     append([]string{}, ftq.unique...),
		predicates: append([]predicate.FieldType{}, ftq.predicates...),
		// clone intermediate query.
		sql:  ftq.sql.Clone(),
		path: ftq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FieldType.Query().
//		GroupBy(fieldtype.FieldInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ftq *FieldTypeQuery) GroupBy(field string, fields ...string) *FieldTypeGroupBy {
	group := &FieldTypeGroupBy{config: ftq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ftq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//	}
//
//	client.FieldType.Query().
//		Select(fieldtype.FieldInt).
//		Scan(ctx, &v)
//
func (ftq *FieldTypeQuery) Select(field string, fields ...string) *FieldTypeSelect {
	selector := &FieldTypeSelect{config: ftq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ftq.sqlQuery(), nil
	}
	return selector
}

func (ftq *FieldTypeQuery) prepareQuery(ctx context.Context) error {
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.sql = prev
	}
	return nil
}

func (ftq *FieldTypeQuery) sqlAll(ctx context.Context) ([]*FieldType, error) {
	var (
		nodes   = []*FieldType{}
		withFKs = ftq.withFKs
		_spec   = ftq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, fieldtype.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &FieldType{config: ftq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, ftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ftq *FieldTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftq.querySpec()
	return sqlgraph.CountNodes(ctx, ftq.driver, _spec)
}

func (ftq *FieldTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ftq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ftq *FieldTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fieldtype.Table,
			Columns: fieldtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fieldtype.FieldID,
			},
		},
		From:   ftq.sql,
		Unique: true,
	}
	if ps := ftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, fieldtype.ValidColumn)
			}
		}
	}
	return _spec
}

func (ftq *FieldTypeQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(ftq.driver.Dialect())
	t1 := builder.Table(fieldtype.Table)
	selector := builder.Select(t1.Columns(fieldtype.Columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(fieldtype.Columns...)...)
	}
	for _, p := range ftq.predicates {
		p(selector)
	}
	for _, p := range ftq.order {
		p(selector, fieldtype.ValidColumn)
	}
	if offset := ftq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FieldTypeGroupBy is the builder for group-by FieldType entities.
type FieldTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FieldTypeGroupBy) Aggregate(fns ...AggregateFunc) *FieldTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the group-by query and scan the result into the given value.
func (ftgb *FieldTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ftgb.path(ctx)
	if err != nil {
		return err
	}
	ftgb.sql = query
	return ftgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ftgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := ftgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ftgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) StringX(ctx context.Context) string {
	v, err := ftgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := ftgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ftgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) IntX(ctx context.Context) int {
	v, err := ftgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ftgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ftgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ftgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ftgb.fields) > 1 {
		return nil, errors.New("ent: FieldTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ftgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ftgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (ftgb *FieldTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ftgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ftgb *FieldTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := ftgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftgb *FieldTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ftgb.fields {
		if !fieldtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ftgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ftgb *FieldTypeGroupBy) sqlQuery() *sql.Selector {
	selector := ftgb.sql
	columns := make([]string, 0, len(ftgb.fields)+len(ftgb.fns))
	columns = append(columns, ftgb.fields...)
	for _, fn := range ftgb.fns {
		columns = append(columns, fn(selector, fieldtype.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ftgb.fields...)
}

// FieldTypeSelect is the builder for select fields of FieldType entities.
type FieldTypeSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (fts *FieldTypeSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := fts.path(ctx)
	if err != nil {
		return err
	}
	fts.sql = query
	return fts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fts *FieldTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fts *FieldTypeSelect) StringsX(ctx context.Context) []string {
	v, err := fts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fts *FieldTypeSelect) StringX(ctx context.Context) string {
	v, err := fts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fts *FieldTypeSelect) IntsX(ctx context.Context) []int {
	v, err := fts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fts *FieldTypeSelect) IntX(ctx context.Context) int {
	v, err := fts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fts *FieldTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fts *FieldTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := fts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fts.fields) > 1 {
		return nil, errors.New("ent: FieldTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fts *FieldTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := fts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (fts *FieldTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = fmt.Errorf("ent: FieldTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fts *FieldTypeSelect) BoolX(ctx context.Context) bool {
	v, err := fts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fts *FieldTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fts.fields {
		if !fieldtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := fts.sqlQuery().Query()
	if err := fts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fts *FieldTypeSelect) sqlQuery() sql.Querier {
	selector := fts.sql
	selector.Select(selector.Columns(fts.fields...)...)
	return selector
}
