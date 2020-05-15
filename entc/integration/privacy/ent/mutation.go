// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGalaxy = "Galaxy"
	TypePlanet = "Planet"
)

// GalaxyMutation represents an operation that mutate the Galaxies
// nodes in the graph.
type GalaxyMutation struct {
	config
	op             Op
	typ            string
	id             *int
	name           *string
	_type          *galaxy.Type
	clearedFields  map[string]struct{}
	planets        map[int]struct{}
	removedplanets map[int]struct{}
	oldValue       func(context.Context) (*Galaxy, error)
}

var _ ent.Mutation = (*GalaxyMutation)(nil)

// galaxyOption allows to manage the mutation configuration using functional options.
type galaxyOption func(*GalaxyMutation)

// newGalaxyMutation creates new mutation for $n.Name.
func newGalaxyMutation(c config, op Op, opts ...galaxyOption) *GalaxyMutation {
	m := &GalaxyMutation{
		config:        c,
		op:            op,
		typ:           TypeGalaxy,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGalaxyID sets the id field of the mutation.
func withGalaxyID(id int) galaxyOption {
	return func(m *GalaxyMutation) {
		var (
			err   error
			once  sync.Once
			value *Galaxy
		)
		m.oldValue = func(ctx context.Context) (*Galaxy, error) {
			once.Do(func() {
				value, err = m.Client().Galaxy.Get(ctx, id)
			})
			return value, err
		}
		m.id = &id
	}
}

// withGalaxy sets the old Galaxy of the mutation.
func withGalaxy(node *Galaxy) galaxyOption {
	return func(m *GalaxyMutation) {
		m.oldValue = func(context.Context) (*Galaxy, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GalaxyMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GalaxyMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *GalaxyMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *GalaxyMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *GalaxyMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GalaxyMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *GalaxyMutation) ResetName() {
	m.name = nil
}

// SetType sets the type field.
func (m *GalaxyMutation) SetType(ga galaxy.Type) {
	m._type = &ga
}

// GetType returns the type value in the mutation.
func (m *GalaxyMutation) GetType() (r galaxy.Type, exists bool) {
	v := m._type
	if v == nil {
		return
	}
	return *v, true
}

// OldType returns the old type value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *GalaxyMutation) OldType(ctx context.Context) (v galaxy.Type, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldType is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldType: %w", err)
	}
	return oldValue.Type, nil
}

// ResetType reset all changes of the "type" field.
func (m *GalaxyMutation) ResetType() {
	m._type = nil
}

// AddPlanetIDs adds the planets edge to Planet by ids.
func (m *GalaxyMutation) AddPlanetIDs(ids ...int) {
	if m.planets == nil {
		m.planets = make(map[int]struct{})
	}
	for i := range ids {
		m.planets[ids[i]] = struct{}{}
	}
}

// RemovePlanetIDs removes the planets edge to Planet by ids.
func (m *GalaxyMutation) RemovePlanetIDs(ids ...int) {
	if m.removedplanets == nil {
		m.removedplanets = make(map[int]struct{})
	}
	for i := range ids {
		m.removedplanets[ids[i]] = struct{}{}
	}
}

// RemovedPlanets returns the removed ids of planets.
func (m *GalaxyMutation) RemovedPlanetsIDs() (ids []int) {
	for id := range m.removedplanets {
		ids = append(ids, id)
	}
	return
}

// PlanetsIDs returns the planets ids in the mutation.
func (m *GalaxyMutation) PlanetsIDs() (ids []int) {
	for id := range m.planets {
		ids = append(ids, id)
	}
	return
}

// ResetPlanets reset all changes of the "planets" edge.
func (m *GalaxyMutation) ResetPlanets() {
	m.planets = nil
	m.removedplanets = nil
}

// Op returns the operation name.
func (m *GalaxyMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Galaxy).
func (m *GalaxyMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *GalaxyMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, galaxy.FieldName)
	}
	if m._type != nil {
		fields = append(fields, galaxy.FieldType)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *GalaxyMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case galaxy.FieldName:
		return m.Name()
	case galaxy.FieldType:
		return m.GetType()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *GalaxyMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case galaxy.FieldName:
		return m.OldName(ctx)
	case galaxy.FieldType:
		return m.OldType(ctx)
	}
	return nil, fmt.Errorf("unknown Galaxy field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GalaxyMutation) SetField(name string, value ent.Value) error {
	switch name {
	case galaxy.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case galaxy.FieldType:
		v, ok := value.(galaxy.Type)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetType(v)
		return nil
	}
	return fmt.Errorf("unknown Galaxy field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *GalaxyMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *GalaxyMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *GalaxyMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Galaxy numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *GalaxyMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *GalaxyMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *GalaxyMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Galaxy nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *GalaxyMutation) ResetField(name string) error {
	switch name {
	case galaxy.FieldName:
		m.ResetName()
		return nil
	case galaxy.FieldType:
		m.ResetType()
		return nil
	}
	return fmt.Errorf("unknown Galaxy field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *GalaxyMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.planets != nil {
		edges = append(edges, galaxy.EdgePlanets)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *GalaxyMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case galaxy.EdgePlanets:
		ids := make([]ent.Value, 0, len(m.planets))
		for id := range m.planets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *GalaxyMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedplanets != nil {
		edges = append(edges, galaxy.EdgePlanets)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *GalaxyMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case galaxy.EdgePlanets:
		ids := make([]ent.Value, 0, len(m.removedplanets))
		for id := range m.removedplanets {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *GalaxyMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *GalaxyMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *GalaxyMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Galaxy unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *GalaxyMutation) ResetEdge(name string) error {
	switch name {
	case galaxy.EdgePlanets:
		m.ResetPlanets()
		return nil
	}
	return fmt.Errorf("unknown Galaxy edge %s", name)
}

// PlanetMutation represents an operation that mutate the Planets
// nodes in the graph.
type PlanetMutation struct {
	config
	op               Op
	typ              string
	id               *int
	name             *string
	age              *uint
	addage           *uint
	clearedFields    map[string]struct{}
	neighbors        map[int]struct{}
	removedneighbors map[int]struct{}
	oldValue         func(context.Context) (*Planet, error)
}

var _ ent.Mutation = (*PlanetMutation)(nil)

// planetOption allows to manage the mutation configuration using functional options.
type planetOption func(*PlanetMutation)

// newPlanetMutation creates new mutation for $n.Name.
func newPlanetMutation(c config, op Op, opts ...planetOption) *PlanetMutation {
	m := &PlanetMutation{
		config:        c,
		op:            op,
		typ:           TypePlanet,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPlanetID sets the id field of the mutation.
func withPlanetID(id int) planetOption {
	return func(m *PlanetMutation) {
		var (
			err   error
			once  sync.Once
			value *Planet
		)
		m.oldValue = func(ctx context.Context) (*Planet, error) {
			once.Do(func() {
				value, err = m.Client().Planet.Get(ctx, id)
			})
			return value, err
		}
		m.id = &id
	}
}

// withPlanet sets the old Planet of the mutation.
func withPlanet(node *Planet) planetOption {
	return func(m *PlanetMutation) {
		m.oldValue = func(context.Context) (*Planet, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PlanetMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PlanetMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *PlanetMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *PlanetMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *PlanetMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PlanetMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *PlanetMutation) ResetName() {
	m.name = nil
}

// SetAge sets the age field.
func (m *PlanetMutation) SetAge(u uint) {
	m.age = &u
	m.addage = nil
}

// Age returns the age value in the mutation.
func (m *PlanetMutation) Age() (r uint, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old age value, if exists.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *PlanetMutation) OldAge(ctx context.Context) (v uint, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAge is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds u to age.
func (m *PlanetMutation) AddAge(u uint) {
	if m.addage != nil {
		*m.addage += u
	} else {
		m.addage = &u
	}
}

// AddedAge returns the value that was added to the age field in this mutation.
func (m *PlanetMutation) AddedAge() (r uint, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ClearAge clears the value of age.
func (m *PlanetMutation) ClearAge() {
	m.age = nil
	m.addage = nil
	m.clearedFields[planet.FieldAge] = struct{}{}
}

// AgeCleared returns if the field age was cleared in this mutation.
func (m *PlanetMutation) AgeCleared() bool {
	_, ok := m.clearedFields[planet.FieldAge]
	return ok
}

// ResetAge reset all changes of the "age" field.
func (m *PlanetMutation) ResetAge() {
	m.age = nil
	m.addage = nil
	delete(m.clearedFields, planet.FieldAge)
}

// AddNeighborIDs adds the neighbors edge to Planet by ids.
func (m *PlanetMutation) AddNeighborIDs(ids ...int) {
	if m.neighbors == nil {
		m.neighbors = make(map[int]struct{})
	}
	for i := range ids {
		m.neighbors[ids[i]] = struct{}{}
	}
}

// RemoveNeighborIDs removes the neighbors edge to Planet by ids.
func (m *PlanetMutation) RemoveNeighborIDs(ids ...int) {
	if m.removedneighbors == nil {
		m.removedneighbors = make(map[int]struct{})
	}
	for i := range ids {
		m.removedneighbors[ids[i]] = struct{}{}
	}
}

// RemovedNeighbors returns the removed ids of neighbors.
func (m *PlanetMutation) RemovedNeighborsIDs() (ids []int) {
	for id := range m.removedneighbors {
		ids = append(ids, id)
	}
	return
}

// NeighborsIDs returns the neighbors ids in the mutation.
func (m *PlanetMutation) NeighborsIDs() (ids []int) {
	for id := range m.neighbors {
		ids = append(ids, id)
	}
	return
}

// ResetNeighbors reset all changes of the "neighbors" edge.
func (m *PlanetMutation) ResetNeighbors() {
	m.neighbors = nil
	m.removedneighbors = nil
}

// Op returns the operation name.
func (m *PlanetMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Planet).
func (m *PlanetMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *PlanetMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, planet.FieldName)
	}
	if m.age != nil {
		fields = append(fields, planet.FieldAge)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *PlanetMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case planet.FieldName:
		return m.Name()
	case planet.FieldAge:
		return m.Age()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *PlanetMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case planet.FieldName:
		return m.OldName(ctx)
	case planet.FieldAge:
		return m.OldAge(ctx)
	}
	return nil, fmt.Errorf("unknown Planet field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PlanetMutation) SetField(name string, value ent.Value) error {
	switch name {
	case planet.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case planet.FieldAge:
		v, ok := value.(uint)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	}
	return fmt.Errorf("unknown Planet field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *PlanetMutation) AddedFields() []string {
	var fields []string
	if m.addage != nil {
		fields = append(fields, planet.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *PlanetMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case planet.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *PlanetMutation) AddField(name string, value ent.Value) error {
	switch name {
	case planet.FieldAge:
		v, ok := value.(uint)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown Planet numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *PlanetMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(planet.FieldAge) {
		fields = append(fields, planet.FieldAge)
	}
	return fields
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *PlanetMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *PlanetMutation) ClearField(name string) error {
	switch name {
	case planet.FieldAge:
		m.ClearAge()
		return nil
	}
	return fmt.Errorf("unknown Planet nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *PlanetMutation) ResetField(name string) error {
	switch name {
	case planet.FieldName:
		m.ResetName()
		return nil
	case planet.FieldAge:
		m.ResetAge()
		return nil
	}
	return fmt.Errorf("unknown Planet field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *PlanetMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.neighbors != nil {
		edges = append(edges, planet.EdgeNeighbors)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *PlanetMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case planet.EdgeNeighbors:
		ids := make([]ent.Value, 0, len(m.neighbors))
		for id := range m.neighbors {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *PlanetMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedneighbors != nil {
		edges = append(edges, planet.EdgeNeighbors)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *PlanetMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case planet.EdgeNeighbors:
		ids := make([]ent.Value, 0, len(m.removedneighbors))
		for id := range m.removedneighbors {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *PlanetMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *PlanetMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *PlanetMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Planet unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *PlanetMutation) ResetEdge(name string) error {
	switch name {
	case planet.EdgeNeighbors:
		m.ResetNeighbors()
		return nil
	}
	return fmt.Errorf("unknown Planet edge %s", name)
}
