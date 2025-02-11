// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/janus/ent/dataconfig"
	"github.com/lbrictson/janus/ent/predicate"
)

// DataConfigUpdate is the builder for updating DataConfig entities.
type DataConfigUpdate struct {
	config
	hooks    []Hook
	mutation *DataConfigMutation
}

// Where appends a list predicates to the DataConfigUpdate builder.
func (dcu *DataConfigUpdate) Where(ps ...predicate.DataConfig) *DataConfigUpdate {
	dcu.mutation.Where(ps...)
	return dcu
}

// SetDaysToKeep sets the "days_to_keep" field.
func (dcu *DataConfigUpdate) SetDaysToKeep(i int) *DataConfigUpdate {
	dcu.mutation.ResetDaysToKeep()
	dcu.mutation.SetDaysToKeep(i)
	return dcu
}

// SetNillableDaysToKeep sets the "days_to_keep" field if the given value is not nil.
func (dcu *DataConfigUpdate) SetNillableDaysToKeep(i *int) *DataConfigUpdate {
	if i != nil {
		dcu.SetDaysToKeep(*i)
	}
	return dcu
}

// AddDaysToKeep adds i to the "days_to_keep" field.
func (dcu *DataConfigUpdate) AddDaysToKeep(i int) *DataConfigUpdate {
	dcu.mutation.AddDaysToKeep(i)
	return dcu
}

// Mutation returns the DataConfigMutation object of the builder.
func (dcu *DataConfigUpdate) Mutation() *DataConfigMutation {
	return dcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dcu *DataConfigUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dcu.sqlSave, dcu.mutation, dcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dcu *DataConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := dcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dcu *DataConfigUpdate) Exec(ctx context.Context) error {
	_, err := dcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcu *DataConfigUpdate) ExecX(ctx context.Context) {
	if err := dcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dcu *DataConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dataconfig.Table, dataconfig.Columns, sqlgraph.NewFieldSpec(dataconfig.FieldID, field.TypeInt))
	if ps := dcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcu.mutation.DaysToKeep(); ok {
		_spec.SetField(dataconfig.FieldDaysToKeep, field.TypeInt, value)
	}
	if value, ok := dcu.mutation.AddedDaysToKeep(); ok {
		_spec.AddField(dataconfig.FieldDaysToKeep, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dataconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dcu.mutation.done = true
	return n, nil
}

// DataConfigUpdateOne is the builder for updating a single DataConfig entity.
type DataConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DataConfigMutation
}

// SetDaysToKeep sets the "days_to_keep" field.
func (dcuo *DataConfigUpdateOne) SetDaysToKeep(i int) *DataConfigUpdateOne {
	dcuo.mutation.ResetDaysToKeep()
	dcuo.mutation.SetDaysToKeep(i)
	return dcuo
}

// SetNillableDaysToKeep sets the "days_to_keep" field if the given value is not nil.
func (dcuo *DataConfigUpdateOne) SetNillableDaysToKeep(i *int) *DataConfigUpdateOne {
	if i != nil {
		dcuo.SetDaysToKeep(*i)
	}
	return dcuo
}

// AddDaysToKeep adds i to the "days_to_keep" field.
func (dcuo *DataConfigUpdateOne) AddDaysToKeep(i int) *DataConfigUpdateOne {
	dcuo.mutation.AddDaysToKeep(i)
	return dcuo
}

// Mutation returns the DataConfigMutation object of the builder.
func (dcuo *DataConfigUpdateOne) Mutation() *DataConfigMutation {
	return dcuo.mutation
}

// Where appends a list predicates to the DataConfigUpdate builder.
func (dcuo *DataConfigUpdateOne) Where(ps ...predicate.DataConfig) *DataConfigUpdateOne {
	dcuo.mutation.Where(ps...)
	return dcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dcuo *DataConfigUpdateOne) Select(field string, fields ...string) *DataConfigUpdateOne {
	dcuo.fields = append([]string{field}, fields...)
	return dcuo
}

// Save executes the query and returns the updated DataConfig entity.
func (dcuo *DataConfigUpdateOne) Save(ctx context.Context) (*DataConfig, error) {
	return withHooks(ctx, dcuo.sqlSave, dcuo.mutation, dcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dcuo *DataConfigUpdateOne) SaveX(ctx context.Context) *DataConfig {
	node, err := dcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dcuo *DataConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := dcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcuo *DataConfigUpdateOne) ExecX(ctx context.Context) {
	if err := dcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dcuo *DataConfigUpdateOne) sqlSave(ctx context.Context) (_node *DataConfig, err error) {
	_spec := sqlgraph.NewUpdateSpec(dataconfig.Table, dataconfig.Columns, sqlgraph.NewFieldSpec(dataconfig.FieldID, field.TypeInt))
	id, ok := dcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DataConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dataconfig.FieldID)
		for _, f := range fields {
			if !dataconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dataconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dcuo.mutation.DaysToKeep(); ok {
		_spec.SetField(dataconfig.FieldDaysToKeep, field.TypeInt, value)
	}
	if value, ok := dcuo.mutation.AddedDaysToKeep(); ok {
		_spec.AddField(dataconfig.FieldDaysToKeep, field.TypeInt, value)
	}
	_node = &DataConfig{config: dcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dataconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dcuo.mutation.done = true
	return _node, nil
}
