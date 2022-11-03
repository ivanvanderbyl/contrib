// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entgql/internal/todo/ent/friendship"
	"entgo.io/contrib/entgql/internal/todo/ent/predicate"
	"entgo.io/contrib/entgql/internal/todo/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FriendshipUpdate is the builder for updating Friendship entities.
type FriendshipUpdate struct {
	config
	hooks     []Hook
	mutation  *FriendshipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (fu *FriendshipUpdate) Where(ps ...predicate.Friendship) *FriendshipUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetCreatedAt sets the "created_at" field.
func (fu *FriendshipUpdate) SetCreatedAt(t time.Time) *FriendshipUpdate {
	fu.mutation.SetCreatedAt(t)
	return fu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fu *FriendshipUpdate) SetNillableCreatedAt(t *time.Time) *FriendshipUpdate {
	if t != nil {
		fu.SetCreatedAt(*t)
	}
	return fu
}

// SetUserID sets the "user_id" field.
func (fu *FriendshipUpdate) SetUserID(i int) *FriendshipUpdate {
	fu.mutation.SetUserID(i)
	return fu
}

// SetFriendID sets the "friend_id" field.
func (fu *FriendshipUpdate) SetFriendID(i int) *FriendshipUpdate {
	fu.mutation.SetFriendID(i)
	return fu
}

// SetUser sets the "user" edge to the User entity.
func (fu *FriendshipUpdate) SetUser(u *User) *FriendshipUpdate {
	return fu.SetUserID(u.ID)
}

// SetFriend sets the "friend" edge to the User entity.
func (fu *FriendshipUpdate) SetFriend(u *User) *FriendshipUpdate {
	return fu.SetFriendID(u.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fu *FriendshipUpdate) Mutation() *FriendshipMutation {
	return fu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fu *FriendshipUpdate) ClearUser() *FriendshipUpdate {
	fu.mutation.ClearUser()
	return fu
}

// ClearFriend clears the "friend" edge to the User entity.
func (fu *FriendshipUpdate) ClearFriend() *FriendshipUpdate {
	fu.mutation.ClearFriend()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FriendshipUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FriendshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FriendshipUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FriendshipUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FriendshipUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FriendshipUpdate) check() error {
	if _, ok := fu.mutation.UserID(); fu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.user"`)
	}
	if _, ok := fu.mutation.FriendID(); fu.mutation.FriendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fu *FriendshipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FriendshipUpdate {
	fu.modifiers = append(fu.modifiers, modifiers...)
	return fu
}

func (fu *FriendshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   friendship.Table,
			Columns: friendship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: friendship.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = fu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FriendshipUpdateOne is the builder for updating a single Friendship entity.
type FriendshipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FriendshipMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (fuo *FriendshipUpdateOne) SetCreatedAt(t time.Time) *FriendshipUpdateOne {
	fuo.mutation.SetCreatedAt(t)
	return fuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fuo *FriendshipUpdateOne) SetNillableCreatedAt(t *time.Time) *FriendshipUpdateOne {
	if t != nil {
		fuo.SetCreatedAt(*t)
	}
	return fuo
}

// SetUserID sets the "user_id" field.
func (fuo *FriendshipUpdateOne) SetUserID(i int) *FriendshipUpdateOne {
	fuo.mutation.SetUserID(i)
	return fuo
}

// SetFriendID sets the "friend_id" field.
func (fuo *FriendshipUpdateOne) SetFriendID(i int) *FriendshipUpdateOne {
	fuo.mutation.SetFriendID(i)
	return fuo
}

// SetUser sets the "user" edge to the User entity.
func (fuo *FriendshipUpdateOne) SetUser(u *User) *FriendshipUpdateOne {
	return fuo.SetUserID(u.ID)
}

// SetFriend sets the "friend" edge to the User entity.
func (fuo *FriendshipUpdateOne) SetFriend(u *User) *FriendshipUpdateOne {
	return fuo.SetFriendID(u.ID)
}

// Mutation returns the FriendshipMutation object of the builder.
func (fuo *FriendshipUpdateOne) Mutation() *FriendshipMutation {
	return fuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fuo *FriendshipUpdateOne) ClearUser() *FriendshipUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// ClearFriend clears the "friend" edge to the User entity.
func (fuo *FriendshipUpdateOne) ClearFriend() *FriendshipUpdateOne {
	fuo.mutation.ClearFriend()
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FriendshipUpdateOne) Select(field string, fields ...string) *FriendshipUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Friendship entity.
func (fuo *FriendshipUpdateOne) Save(ctx context.Context) (*Friendship, error) {
	var (
		err  error
		node *Friendship
	)
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FriendshipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Friendship)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FriendshipMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) SaveX(ctx context.Context) *Friendship {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FriendshipUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FriendshipUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FriendshipUpdateOne) check() error {
	if _, ok := fuo.mutation.UserID(); fuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.user"`)
	}
	if _, ok := fuo.mutation.FriendID(); fuo.mutation.FriendCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fuo *FriendshipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FriendshipUpdateOne {
	fuo.modifiers = append(fuo.modifiers, modifiers...)
	return fuo
}

func (fuo *FriendshipUpdateOne) sqlSave(ctx context.Context) (_node *Friendship, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   friendship.Table,
			Columns: friendship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: friendship.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friendship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for _, f := range fields {
			if !friendship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.UserTable,
			Columns: []string{friendship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FriendCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FriendIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   friendship.FriendTable,
			Columns: []string{friendship.FriendColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = fuo.modifiers
	_node = &Friendship{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
