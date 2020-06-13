// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetOptionalInt sets the optional_int field.
func (uc *UserCreate) SetOptionalInt(i int) *UserCreate {
	uc.mutation.SetOptionalInt(i)
	return uc
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (uc *UserCreate) SetNillableOptionalInt(i *int) *UserCreate {
	if i != nil {
		uc.SetOptionalInt(*i)
	}
	return uc
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.mutation.SetAge(i)
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetLast sets the last field.
func (uc *UserCreate) SetLast(s string) *UserCreate {
	uc.mutation.SetLast(s)
	return uc
}

// SetNillableLast sets the last field if the given value is not nil.
func (uc *UserCreate) SetNillableLast(s *string) *UserCreate {
	if s != nil {
		uc.SetLast(*s)
	}
	return uc
}

// SetNickname sets the nickname field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.mutation.SetNickname(s)
	return uc
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetPhone sets the phone field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.mutation.SetPhone(s)
	return uc
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNillablePassword sets the password field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetRole sets the role field.
func (uc *UserCreate) SetRole(u user.Role) *UserCreate {
	uc.mutation.SetRole(u)
	return uc
}

// SetNillableRole sets the role field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(u *user.Role) *UserCreate {
	if u != nil {
		uc.SetRole(*u)
	}
	return uc
}

// SetSSOCert sets the SSOCert field.
func (uc *UserCreate) SetSSOCert(s string) *UserCreate {
	uc.mutation.SetSSOCert(s)
	return uc
}

// SetNillableSSOCert sets the SSOCert field if the given value is not nil.
func (uc *UserCreate) SetNillableSSOCert(s *string) *UserCreate {
	if s != nil {
		uc.SetSSOCert(*s)
	}
	return uc
}

// SetCardID sets the card edge to Card by id.
func (uc *UserCreate) SetCardID(id string) *UserCreate {
	uc.mutation.SetCardID(id)
	return uc
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uc *UserCreate) SetNillableCardID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetCardID(*id)
	}
	return uc
}

// SetCard sets the card edge to Card.
func (uc *UserCreate) SetCard(c *Card) *UserCreate {
	return uc.SetCardID(c.ID)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uc *UserCreate) AddPetIDs(ids ...string) *UserCreate {
	uc.mutation.AddPetIDs(ids...)
	return uc
}

// AddPets adds the pets edges to Pet.
func (uc *UserCreate) AddPets(p ...*Pet) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPetIDs(ids...)
}

// AddFileIDs adds the files edge to File by ids.
func (uc *UserCreate) AddFileIDs(ids ...string) *UserCreate {
	uc.mutation.AddFileIDs(ids...)
	return uc
}

// AddFiles adds the files edges to File.
func (uc *UserCreate) AddFiles(f ...*File) *UserCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uc.AddFileIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uc *UserCreate) AddGroupIDs(ids ...string) *UserCreate {
	uc.mutation.AddGroupIDs(ids...)
	return uc
}

// AddGroups adds the groups edges to Group.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// AddFriendIDs adds the friends edge to User by ids.
func (uc *UserCreate) AddFriendIDs(ids ...string) *UserCreate {
	uc.mutation.AddFriendIDs(ids...)
	return uc
}

// AddFriends adds the friends edges to User.
func (uc *UserCreate) AddFriends(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFriendIDs(ids...)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uc *UserCreate) AddFollowerIDs(ids ...string) *UserCreate {
	uc.mutation.AddFollowerIDs(ids...)
	return uc
}

// AddFollowers adds the followers edges to User.
func (uc *UserCreate) AddFollowers(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uc *UserCreate) AddFollowingIDs(ids ...string) *UserCreate {
	uc.mutation.AddFollowingIDs(ids...)
	return uc
}

// AddFollowing adds the following edges to User.
func (uc *UserCreate) AddFollowing(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowingIDs(ids...)
}

// SetTeamID sets the team edge to Pet by id.
func (uc *UserCreate) SetTeamID(id string) *UserCreate {
	uc.mutation.SetTeamID(id)
	return uc
}

// SetNillableTeamID sets the team edge to Pet by id if the given value is not nil.
func (uc *UserCreate) SetNillableTeamID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetTeamID(*id)
	}
	return uc
}

// SetTeam sets the team edge to Pet.
func (uc *UserCreate) SetTeam(p *Pet) *UserCreate {
	return uc.SetTeamID(p.ID)
}

// SetSpouseID sets the spouse edge to User by id.
func (uc *UserCreate) SetSpouseID(id string) *UserCreate {
	uc.mutation.SetSpouseID(id)
	return uc
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableSpouseID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetSpouseID(*id)
	}
	return uc
}

// SetSpouse sets the spouse edge to User.
func (uc *UserCreate) SetSpouse(u *User) *UserCreate {
	return uc.SetSpouseID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uc *UserCreate) AddChildIDs(ids ...string) *UserCreate {
	uc.mutation.AddChildIDs(ids...)
	return uc
}

// AddChildren adds the children edges to User.
func (uc *UserCreate) AddChildren(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddChildIDs(ids...)
}

// SetParentID sets the parent edge to User by id.
func (uc *UserCreate) SetParentID(id string) *UserCreate {
	uc.mutation.SetParentID(id)
	return uc
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableParentID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetParentID(*id)
	}
	return uc
}

// SetParent sets the parent edge to User.
func (uc *UserCreate) SetParent(u *User) *UserCreate {
	return uc.SetParentID(u.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if v, ok := uc.mutation.OptionalInt(); ok {
		if err := user.OptionalIntValidator(v); err != nil {
			return nil, &ValidationError{Name: "optional_int", err: fmt.Errorf("ent: validator failed for field \"optional_int\": %w", err)}
		}
	}
	if _, ok := uc.mutation.Age(); !ok {
		return nil, &ValidationError{Name: "age", err: errors.New("ent: missing required field \"age\"")}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := uc.mutation.Last(); !ok {
		v := user.DefaultLast
		uc.mutation.SetLast(v)
	}
	if _, ok := uc.mutation.Role(); !ok {
		v := user.DefaultRole
		uc.mutation.SetRole(v)
	}
	if v, ok := uc.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return nil, &ValidationError{Name: "role", err: fmt.Errorf("ent: validator failed for field \"role\": %w", err)}
		}
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) gremlinSave(ctx context.Context) (*User, error) {
	res := &gremlin.Response{}
	query, bindings := uc.gremlin().Query()
	if err := uc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	u := &User{config: uc.config}
	if err := u.FromResponse(res); err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 8)
	v := g.AddV(user.Label)
	if value, ok := uc.mutation.OptionalInt(); ok {
		v.Property(dsl.Single, user.FieldOptionalInt, value)
	}
	if value, ok := uc.mutation.Age(); ok {
		v.Property(dsl.Single, user.FieldAge, value)
	}
	if value, ok := uc.mutation.Name(); ok {
		v.Property(dsl.Single, user.FieldName, value)
	}
	if value, ok := uc.mutation.Last(); ok {
		v.Property(dsl.Single, user.FieldLast, value)
	}
	if value, ok := uc.mutation.Nickname(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldNickname, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldNickname, value)),
		})
		v.Property(dsl.Single, user.FieldNickname, value)
	}
	if value, ok := uc.mutation.Phone(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldPhone, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldPhone, value)),
		})
		v.Property(dsl.Single, user.FieldPhone, value)
	}
	if value, ok := uc.mutation.Password(); ok {
		v.Property(dsl.Single, user.FieldPassword, value)
	}
	if value, ok := uc.mutation.Role(); ok {
		v.Property(dsl.Single, user.FieldRole, value)
	}
	if value, ok := uc.mutation.SSOCert(); ok {
		v.Property(dsl.Single, user.FieldSSOCert, value)
	}
	for _, id := range uc.mutation.CardIDs() {
		v.AddE(user.CardLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.CardLabel, id)),
		})
	}
	for _, id := range uc.mutation.PetsIDs() {
		v.AddE(user.PetsLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.PetsLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.PetsLabel, id)),
		})
	}
	for _, id := range uc.mutation.FilesIDs() {
		v.AddE(user.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.FilesLabel, id)),
		})
	}
	for _, id := range uc.mutation.GroupsIDs() {
		v.AddE(user.GroupsLabel).To(g.V(id)).OutV()
	}
	for _, id := range uc.mutation.FriendsIDs() {
		v.AddE(user.FriendsLabel).To(g.V(id)).OutV()
	}
	for _, id := range uc.mutation.FollowersIDs() {
		v.AddE(user.FollowingLabel).From(g.V(id)).InV()
	}
	for _, id := range uc.mutation.FollowingIDs() {
		v.AddE(user.FollowingLabel).To(g.V(id)).OutV()
	}
	for _, id := range uc.mutation.TeamIDs() {
		v.AddE(user.TeamLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.TeamLabel, id)),
		})
	}
	for _, id := range uc.mutation.SpouseIDs() {
		v.AddE(user.SpouseLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.SpouseLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.SpouseLabel, id)),
		})
	}
	for _, id := range uc.mutation.ChildrenIDs() {
		v.AddE(user.ParentLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.ParentLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.ParentLabel, id)),
		})
	}
	for _, id := range uc.mutation.ParentIDs() {
		v.AddE(user.ParentLabel).To(g.V(id)).OutV()
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
