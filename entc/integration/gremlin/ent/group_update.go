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

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebook/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebook/ent/entc/integration/gremlin/ent/group"
	"github.com/facebook/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebook/ent/entc/integration/gremlin/ent/user"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks      []Hook
	mutation   *GroupMutation
	predicates []predicate.Group
}

// Where adds a new predicate for the builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetActive sets the active field.
func (gu *GroupUpdate) SetActive(b bool) *GroupUpdate {
	gu.mutation.SetActive(b)
	return gu
}

// SetNillableActive sets the active field if the given value is not nil.
func (gu *GroupUpdate) SetNillableActive(b *bool) *GroupUpdate {
	if b != nil {
		gu.SetActive(*b)
	}
	return gu
}

// SetExpire sets the expire field.
func (gu *GroupUpdate) SetExpire(t time.Time) *GroupUpdate {
	gu.mutation.SetExpire(t)
	return gu
}

// SetType sets the type field.
func (gu *GroupUpdate) SetType(s string) *GroupUpdate {
	gu.mutation.SetType(s)
	return gu
}

// SetNillableType sets the type field if the given value is not nil.
func (gu *GroupUpdate) SetNillableType(s *string) *GroupUpdate {
	if s != nil {
		gu.SetType(*s)
	}
	return gu
}

// ClearType clears the value of type.
func (gu *GroupUpdate) ClearType() *GroupUpdate {
	gu.mutation.ClearType()
	return gu
}

// SetMaxUsers sets the max_users field.
func (gu *GroupUpdate) SetMaxUsers(i int) *GroupUpdate {
	gu.mutation.ResetMaxUsers()
	gu.mutation.SetMaxUsers(i)
	return gu
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (gu *GroupUpdate) SetNillableMaxUsers(i *int) *GroupUpdate {
	if i != nil {
		gu.SetMaxUsers(*i)
	}
	return gu
}

// AddMaxUsers adds i to max_users.
func (gu *GroupUpdate) AddMaxUsers(i int) *GroupUpdate {
	gu.mutation.AddMaxUsers(i)
	return gu
}

// ClearMaxUsers clears the value of max_users.
func (gu *GroupUpdate) ClearMaxUsers() *GroupUpdate {
	gu.mutation.ClearMaxUsers()
	return gu
}

// SetName sets the name field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// AddFileIDs adds the files edge to File by ids.
func (gu *GroupUpdate) AddFileIDs(ids ...string) *GroupUpdate {
	gu.mutation.AddFileIDs(ids...)
	return gu
}

// AddFiles adds the files edges to File.
func (gu *GroupUpdate) AddFiles(f ...*File) *GroupUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gu.AddFileIDs(ids...)
}

// AddBlockedIDs adds the blocked edge to User by ids.
func (gu *GroupUpdate) AddBlockedIDs(ids ...string) *GroupUpdate {
	gu.mutation.AddBlockedIDs(ids...)
	return gu
}

// AddBlocked adds the blocked edges to User.
func (gu *GroupUpdate) AddBlocked(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddBlockedIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (gu *GroupUpdate) AddUserIDs(ids ...string) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the users edges to User.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetInfoID sets the info edge to GroupInfo by id.
func (gu *GroupUpdate) SetInfoID(id string) *GroupUpdate {
	gu.mutation.SetInfoID(id)
	return gu
}

// SetInfo sets the info edge to GroupInfo.
func (gu *GroupUpdate) SetInfo(g *GroupInfo) *GroupUpdate {
	return gu.SetInfoID(g.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// RemoveFileIDs removes the files edge to File by ids.
func (gu *GroupUpdate) RemoveFileIDs(ids ...string) *GroupUpdate {
	gu.mutation.RemoveFileIDs(ids...)
	return gu
}

// RemoveFiles removes files edges to File.
func (gu *GroupUpdate) RemoveFiles(f ...*File) *GroupUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gu.RemoveFileIDs(ids...)
}

// RemoveBlockedIDs removes the blocked edge to User by ids.
func (gu *GroupUpdate) RemoveBlockedIDs(ids ...string) *GroupUpdate {
	gu.mutation.RemoveBlockedIDs(ids...)
	return gu
}

// RemoveBlocked removes blocked edges to User.
func (gu *GroupUpdate) RemoveBlocked(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveBlockedIDs(ids...)
}

// RemoveUserIDs removes the users edge to User by ids.
func (gu *GroupUpdate) RemoveUserIDs(ids ...string) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes users edges to User.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearInfo clears the info edge to GroupInfo.
func (gu *GroupUpdate) ClearInfo() *GroupUpdate {
	gu.mutation.ClearInfo()
	return gu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := gu.mutation.GetType(); ok {
		if err := group.TypeValidator(v); err != nil {
			return 0, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := gu.mutation.MaxUsers(); ok {
		if err := group.MaxUsersValidator(v); err != nil {
			return 0, &ValidationError{Name: "max_users", err: fmt.Errorf("ent: validator failed for field \"max_users\": %w", err)}
		}
	}
	if v, ok := gu.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	if _, ok := gu.mutation.InfoID(); gu.mutation.InfoCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"info\"")
	}
	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gu.gremlin().Query()
	if err := gu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (gu *GroupUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(group.Label)
	for _, p := range gu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := gu.mutation.Active(); ok {
		v.Property(dsl.Single, group.FieldActive, value)
	}
	if value, ok := gu.mutation.Expire(); ok {
		v.Property(dsl.Single, group.FieldExpire, value)
	}
	if value, ok := gu.mutation.GetType(); ok {
		v.Property(dsl.Single, group.FieldType, value)
	}
	if value, ok := gu.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, group.FieldMaxUsers, value)
	}
	if value, ok := gu.mutation.AddedMaxUsers(); ok {
		v.Property(dsl.Single, group.FieldMaxUsers, __.Union(__.Values(group.FieldMaxUsers), __.Constant(value)).Sum())
	}
	if value, ok := gu.mutation.Name(); ok {
		v.Property(dsl.Single, group.FieldName, value)
	}
	var properties []interface{}
	if gu.mutation.TypeCleared() {
		properties = append(properties, group.FieldType)
	}
	if gu.mutation.MaxUsersCleared() {
		properties = append(properties, group.FieldMaxUsers)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	for _, id := range gu.mutation.RemovedFilesIDs() {
		tr := rv.Clone().OutE(group.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range gu.mutation.FilesIDs() {
		v.AddE(group.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.FilesLabel, id)),
		})
	}
	for _, id := range gu.mutation.RemovedBlockedIDs() {
		tr := rv.Clone().OutE(group.BlockedLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range gu.mutation.BlockedIDs() {
		v.AddE(group.BlockedLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.BlockedLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.BlockedLabel, id)),
		})
	}
	for _, id := range gu.mutation.RemovedUsersIDs() {
		tr := rv.Clone().InE(user.GroupsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range gu.mutation.UsersIDs() {
		v.AddE(user.GroupsLabel).From(g.V(id)).InV()
	}
	if gu.mutation.InfoCleared() {
		tr := rv.Clone().OutE(group.InfoLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range gu.mutation.InfoIDs() {
		v.AddE(group.InfoLabel).To(g.V(id)).OutV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// SetActive sets the active field.
func (guo *GroupUpdateOne) SetActive(b bool) *GroupUpdateOne {
	guo.mutation.SetActive(b)
	return guo
}

// SetNillableActive sets the active field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableActive(b *bool) *GroupUpdateOne {
	if b != nil {
		guo.SetActive(*b)
	}
	return guo
}

// SetExpire sets the expire field.
func (guo *GroupUpdateOne) SetExpire(t time.Time) *GroupUpdateOne {
	guo.mutation.SetExpire(t)
	return guo
}

// SetType sets the type field.
func (guo *GroupUpdateOne) SetType(s string) *GroupUpdateOne {
	guo.mutation.SetType(s)
	return guo
}

// SetNillableType sets the type field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableType(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetType(*s)
	}
	return guo
}

// ClearType clears the value of type.
func (guo *GroupUpdateOne) ClearType() *GroupUpdateOne {
	guo.mutation.ClearType()
	return guo
}

// SetMaxUsers sets the max_users field.
func (guo *GroupUpdateOne) SetMaxUsers(i int) *GroupUpdateOne {
	guo.mutation.ResetMaxUsers()
	guo.mutation.SetMaxUsers(i)
	return guo
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableMaxUsers(i *int) *GroupUpdateOne {
	if i != nil {
		guo.SetMaxUsers(*i)
	}
	return guo
}

// AddMaxUsers adds i to max_users.
func (guo *GroupUpdateOne) AddMaxUsers(i int) *GroupUpdateOne {
	guo.mutation.AddMaxUsers(i)
	return guo
}

// ClearMaxUsers clears the value of max_users.
func (guo *GroupUpdateOne) ClearMaxUsers() *GroupUpdateOne {
	guo.mutation.ClearMaxUsers()
	return guo
}

// SetName sets the name field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// AddFileIDs adds the files edge to File by ids.
func (guo *GroupUpdateOne) AddFileIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.AddFileIDs(ids...)
	return guo
}

// AddFiles adds the files edges to File.
func (guo *GroupUpdateOne) AddFiles(f ...*File) *GroupUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return guo.AddFileIDs(ids...)
}

// AddBlockedIDs adds the blocked edge to User by ids.
func (guo *GroupUpdateOne) AddBlockedIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.AddBlockedIDs(ids...)
	return guo
}

// AddBlocked adds the blocked edges to User.
func (guo *GroupUpdateOne) AddBlocked(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddBlockedIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (guo *GroupUpdateOne) AddUserIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the users edges to User.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetInfoID sets the info edge to GroupInfo by id.
func (guo *GroupUpdateOne) SetInfoID(id string) *GroupUpdateOne {
	guo.mutation.SetInfoID(id)
	return guo
}

// SetInfo sets the info edge to GroupInfo.
func (guo *GroupUpdateOne) SetInfo(g *GroupInfo) *GroupUpdateOne {
	return guo.SetInfoID(g.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// RemoveFileIDs removes the files edge to File by ids.
func (guo *GroupUpdateOne) RemoveFileIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.RemoveFileIDs(ids...)
	return guo
}

// RemoveFiles removes files edges to File.
func (guo *GroupUpdateOne) RemoveFiles(f ...*File) *GroupUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return guo.RemoveFileIDs(ids...)
}

// RemoveBlockedIDs removes the blocked edge to User by ids.
func (guo *GroupUpdateOne) RemoveBlockedIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.RemoveBlockedIDs(ids...)
	return guo
}

// RemoveBlocked removes blocked edges to User.
func (guo *GroupUpdateOne) RemoveBlocked(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveBlockedIDs(ids...)
}

// RemoveUserIDs removes the users edge to User by ids.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes users edges to User.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearInfo clears the info edge to GroupInfo.
func (guo *GroupUpdateOne) ClearInfo() *GroupUpdateOne {
	guo.mutation.ClearInfo()
	return guo
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	if v, ok := guo.mutation.GetType(); ok {
		if err := group.TypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := guo.mutation.MaxUsers(); ok {
		if err := group.MaxUsersValidator(v); err != nil {
			return nil, &ValidationError{Name: "max_users", err: fmt.Errorf("ent: validator failed for field \"max_users\": %w", err)}
		}
	}
	if v, ok := guo.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}

	if _, ok := guo.mutation.InfoID(); guo.mutation.InfoCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"info\"")
	}
	var (
		err  error
		node *Group
	)
	if len(guo.hooks) == 0 {
		node, err = guo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	gr, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) gremlinSave(ctx context.Context) (*Group, error) {
	res := &gremlin.Response{}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Group.ID for update")}
	}
	query, bindings := guo.gremlin(id).Query()
	if err := guo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gr := &Group{config: guo.config}
	if err := gr.FromResponse(res); err != nil {
		return nil, err
	}
	return gr, nil
}

func (guo *GroupUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := guo.mutation.Active(); ok {
		v.Property(dsl.Single, group.FieldActive, value)
	}
	if value, ok := guo.mutation.Expire(); ok {
		v.Property(dsl.Single, group.FieldExpire, value)
	}
	if value, ok := guo.mutation.GetType(); ok {
		v.Property(dsl.Single, group.FieldType, value)
	}
	if value, ok := guo.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, group.FieldMaxUsers, value)
	}
	if value, ok := guo.mutation.AddedMaxUsers(); ok {
		v.Property(dsl.Single, group.FieldMaxUsers, __.Union(__.Values(group.FieldMaxUsers), __.Constant(value)).Sum())
	}
	if value, ok := guo.mutation.Name(); ok {
		v.Property(dsl.Single, group.FieldName, value)
	}
	var properties []interface{}
	if guo.mutation.TypeCleared() {
		properties = append(properties, group.FieldType)
	}
	if guo.mutation.MaxUsersCleared() {
		properties = append(properties, group.FieldMaxUsers)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	for _, id := range guo.mutation.RemovedFilesIDs() {
		tr := rv.Clone().OutE(group.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range guo.mutation.FilesIDs() {
		v.AddE(group.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.FilesLabel, id)),
		})
	}
	for _, id := range guo.mutation.RemovedBlockedIDs() {
		tr := rv.Clone().OutE(group.BlockedLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range guo.mutation.BlockedIDs() {
		v.AddE(group.BlockedLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.BlockedLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.BlockedLabel, id)),
		})
	}
	for _, id := range guo.mutation.RemovedUsersIDs() {
		tr := rv.Clone().InE(user.GroupsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range guo.mutation.UsersIDs() {
		v.AddE(user.GroupsLabel).From(g.V(id)).InV()
	}
	if guo.mutation.InfoCleared() {
		tr := rv.Clone().OutE(group.InfoLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range guo.mutation.InfoIDs() {
		v.AddE(group.InfoLabel).To(g.V(id)).OutV()
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
