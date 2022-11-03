// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/dependsonskipped"
	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DependsOnSkippedUpdate is the builder for updating DependsOnSkipped entities.
type DependsOnSkippedUpdate struct {
	config
	hooks    []Hook
	mutation *DependsOnSkippedMutation
}

// Where appends a list predicates to the DependsOnSkippedUpdate builder.
func (dosu *DependsOnSkippedUpdate) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedUpdate {
	dosu.mutation.Where(ps...)
	return dosu
}

// SetName sets the "name" field.
func (dosu *DependsOnSkippedUpdate) SetName(s string) *DependsOnSkippedUpdate {
	dosu.mutation.SetName(s)
	return dosu
}

// AddSkippedIDs adds the "skipped" edge to the ImplicitSkippedMessage entity by IDs.
func (dosu *DependsOnSkippedUpdate) AddSkippedIDs(ids ...int) *DependsOnSkippedUpdate {
	dosu.mutation.AddSkippedIDs(ids...)
	return dosu
}

// AddSkipped adds the "skipped" edges to the ImplicitSkippedMessage entity.
func (dosu *DependsOnSkippedUpdate) AddSkipped(i ...*ImplicitSkippedMessage) *DependsOnSkippedUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return dosu.AddSkippedIDs(ids...)
}

// Mutation returns the DependsOnSkippedMutation object of the builder.
func (dosu *DependsOnSkippedUpdate) Mutation() *DependsOnSkippedMutation {
	return dosu.mutation
}

// ClearSkipped clears all "skipped" edges to the ImplicitSkippedMessage entity.
func (dosu *DependsOnSkippedUpdate) ClearSkipped() *DependsOnSkippedUpdate {
	dosu.mutation.ClearSkipped()
	return dosu
}

// RemoveSkippedIDs removes the "skipped" edge to ImplicitSkippedMessage entities by IDs.
func (dosu *DependsOnSkippedUpdate) RemoveSkippedIDs(ids ...int) *DependsOnSkippedUpdate {
	dosu.mutation.RemoveSkippedIDs(ids...)
	return dosu
}

// RemoveSkipped removes "skipped" edges to ImplicitSkippedMessage entities.
func (dosu *DependsOnSkippedUpdate) RemoveSkipped(i ...*ImplicitSkippedMessage) *DependsOnSkippedUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return dosu.RemoveSkippedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dosu *DependsOnSkippedUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dosu.hooks) == 0 {
		affected, err = dosu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DependsOnSkippedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dosu.mutation = mutation
			affected, err = dosu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dosu.hooks) - 1; i >= 0; i-- {
			if dosu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dosu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dosu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dosu *DependsOnSkippedUpdate) SaveX(ctx context.Context) int {
	affected, err := dosu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dosu *DependsOnSkippedUpdate) Exec(ctx context.Context) error {
	_, err := dosu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dosu *DependsOnSkippedUpdate) ExecX(ctx context.Context) {
	if err := dosu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dosu *DependsOnSkippedUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dependsonskipped.Table,
			Columns: dependsonskipped.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dependsonskipped.FieldID,
			},
		},
	}
	if ps := dosu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dosu.mutation.Name(); ok {
		_spec.SetField(dependsonskipped.FieldName, field.TypeString, value)
	}
	if dosu.mutation.SkippedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implicitskippedmessage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dosu.mutation.RemovedSkippedIDs(); len(nodes) > 0 && !dosu.mutation.SkippedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implicitskippedmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dosu.mutation.SkippedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implicitskippedmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dosu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dependsonskipped.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DependsOnSkippedUpdateOne is the builder for updating a single DependsOnSkipped entity.
type DependsOnSkippedUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DependsOnSkippedMutation
}

// SetName sets the "name" field.
func (dosuo *DependsOnSkippedUpdateOne) SetName(s string) *DependsOnSkippedUpdateOne {
	dosuo.mutation.SetName(s)
	return dosuo
}

// AddSkippedIDs adds the "skipped" edge to the ImplicitSkippedMessage entity by IDs.
func (dosuo *DependsOnSkippedUpdateOne) AddSkippedIDs(ids ...int) *DependsOnSkippedUpdateOne {
	dosuo.mutation.AddSkippedIDs(ids...)
	return dosuo
}

// AddSkipped adds the "skipped" edges to the ImplicitSkippedMessage entity.
func (dosuo *DependsOnSkippedUpdateOne) AddSkipped(i ...*ImplicitSkippedMessage) *DependsOnSkippedUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return dosuo.AddSkippedIDs(ids...)
}

// Mutation returns the DependsOnSkippedMutation object of the builder.
func (dosuo *DependsOnSkippedUpdateOne) Mutation() *DependsOnSkippedMutation {
	return dosuo.mutation
}

// ClearSkipped clears all "skipped" edges to the ImplicitSkippedMessage entity.
func (dosuo *DependsOnSkippedUpdateOne) ClearSkipped() *DependsOnSkippedUpdateOne {
	dosuo.mutation.ClearSkipped()
	return dosuo
}

// RemoveSkippedIDs removes the "skipped" edge to ImplicitSkippedMessage entities by IDs.
func (dosuo *DependsOnSkippedUpdateOne) RemoveSkippedIDs(ids ...int) *DependsOnSkippedUpdateOne {
	dosuo.mutation.RemoveSkippedIDs(ids...)
	return dosuo
}

// RemoveSkipped removes "skipped" edges to ImplicitSkippedMessage entities.
func (dosuo *DependsOnSkippedUpdateOne) RemoveSkipped(i ...*ImplicitSkippedMessage) *DependsOnSkippedUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return dosuo.RemoveSkippedIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dosuo *DependsOnSkippedUpdateOne) Select(field string, fields ...string) *DependsOnSkippedUpdateOne {
	dosuo.fields = append([]string{field}, fields...)
	return dosuo
}

// Save executes the query and returns the updated DependsOnSkipped entity.
func (dosuo *DependsOnSkippedUpdateOne) Save(ctx context.Context) (*DependsOnSkipped, error) {
	var (
		err  error
		node *DependsOnSkipped
	)
	if len(dosuo.hooks) == 0 {
		node, err = dosuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DependsOnSkippedMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dosuo.mutation = mutation
			node, err = dosuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dosuo.hooks) - 1; i >= 0; i-- {
			if dosuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dosuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dosuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DependsOnSkipped)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DependsOnSkippedMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dosuo *DependsOnSkippedUpdateOne) SaveX(ctx context.Context) *DependsOnSkipped {
	node, err := dosuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dosuo *DependsOnSkippedUpdateOne) Exec(ctx context.Context) error {
	_, err := dosuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dosuo *DependsOnSkippedUpdateOne) ExecX(ctx context.Context) {
	if err := dosuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dosuo *DependsOnSkippedUpdateOne) sqlSave(ctx context.Context) (_node *DependsOnSkipped, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dependsonskipped.Table,
			Columns: dependsonskipped.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dependsonskipped.FieldID,
			},
		},
	}
	id, ok := dosuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DependsOnSkipped.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dosuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dependsonskipped.FieldID)
		for _, f := range fields {
			if !dependsonskipped.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dependsonskipped.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dosuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dosuo.mutation.Name(); ok {
		_spec.SetField(dependsonskipped.FieldName, field.TypeString, value)
	}
	if dosuo.mutation.SkippedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implicitskippedmessage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dosuo.mutation.RemovedSkippedIDs(); len(nodes) > 0 && !dosuo.mutation.SkippedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implicitskippedmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dosuo.mutation.SkippedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dependsonskipped.SkippedTable,
			Columns: []string{dependsonskipped.SkippedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: implicitskippedmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &DependsOnSkipped{config: dosuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dosuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dependsonskipped.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
