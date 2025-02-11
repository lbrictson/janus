// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/janus/ent/jobconfig"
)

// JobConfigCreate is the builder for creating a JobConfig entity.
type JobConfigCreate struct {
	config
	mutation *JobConfigMutation
	hooks    []Hook
}

// SetMaxConcurrentJobs sets the "max_concurrent_jobs" field.
func (jcc *JobConfigCreate) SetMaxConcurrentJobs(i int) *JobConfigCreate {
	jcc.mutation.SetMaxConcurrentJobs(i)
	return jcc
}

// SetNillableMaxConcurrentJobs sets the "max_concurrent_jobs" field if the given value is not nil.
func (jcc *JobConfigCreate) SetNillableMaxConcurrentJobs(i *int) *JobConfigCreate {
	if i != nil {
		jcc.SetMaxConcurrentJobs(*i)
	}
	return jcc
}

// SetDefaultTimeoutSeconds sets the "default_timeout_seconds" field.
func (jcc *JobConfigCreate) SetDefaultTimeoutSeconds(i int) *JobConfigCreate {
	jcc.mutation.SetDefaultTimeoutSeconds(i)
	return jcc
}

// SetNillableDefaultTimeoutSeconds sets the "default_timeout_seconds" field if the given value is not nil.
func (jcc *JobConfigCreate) SetNillableDefaultTimeoutSeconds(i *int) *JobConfigCreate {
	if i != nil {
		jcc.SetDefaultTimeoutSeconds(*i)
	}
	return jcc
}

// Mutation returns the JobConfigMutation object of the builder.
func (jcc *JobConfigCreate) Mutation() *JobConfigMutation {
	return jcc.mutation
}

// Save creates the JobConfig in the database.
func (jcc *JobConfigCreate) Save(ctx context.Context) (*JobConfig, error) {
	jcc.defaults()
	return withHooks(ctx, jcc.sqlSave, jcc.mutation, jcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jcc *JobConfigCreate) SaveX(ctx context.Context) *JobConfig {
	v, err := jcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcc *JobConfigCreate) Exec(ctx context.Context) error {
	_, err := jcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcc *JobConfigCreate) ExecX(ctx context.Context) {
	if err := jcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jcc *JobConfigCreate) defaults() {
	if _, ok := jcc.mutation.MaxConcurrentJobs(); !ok {
		v := jobconfig.DefaultMaxConcurrentJobs
		jcc.mutation.SetMaxConcurrentJobs(v)
	}
	if _, ok := jcc.mutation.DefaultTimeoutSeconds(); !ok {
		v := jobconfig.DefaultDefaultTimeoutSeconds
		jcc.mutation.SetDefaultTimeoutSeconds(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jcc *JobConfigCreate) check() error {
	if _, ok := jcc.mutation.MaxConcurrentJobs(); !ok {
		return &ValidationError{Name: "max_concurrent_jobs", err: errors.New(`ent: missing required field "JobConfig.max_concurrent_jobs"`)}
	}
	if _, ok := jcc.mutation.DefaultTimeoutSeconds(); !ok {
		return &ValidationError{Name: "default_timeout_seconds", err: errors.New(`ent: missing required field "JobConfig.default_timeout_seconds"`)}
	}
	return nil
}

func (jcc *JobConfigCreate) sqlSave(ctx context.Context) (*JobConfig, error) {
	if err := jcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jcc.mutation.id = &_node.ID
	jcc.mutation.done = true
	return _node, nil
}

func (jcc *JobConfigCreate) createSpec() (*JobConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &JobConfig{config: jcc.config}
		_spec = sqlgraph.NewCreateSpec(jobconfig.Table, sqlgraph.NewFieldSpec(jobconfig.FieldID, field.TypeInt))
	)
	if value, ok := jcc.mutation.MaxConcurrentJobs(); ok {
		_spec.SetField(jobconfig.FieldMaxConcurrentJobs, field.TypeInt, value)
		_node.MaxConcurrentJobs = value
	}
	if value, ok := jcc.mutation.DefaultTimeoutSeconds(); ok {
		_spec.SetField(jobconfig.FieldDefaultTimeoutSeconds, field.TypeInt, value)
		_node.DefaultTimeoutSeconds = value
	}
	return _node, _spec
}

// JobConfigCreateBulk is the builder for creating many JobConfig entities in bulk.
type JobConfigCreateBulk struct {
	config
	err      error
	builders []*JobConfigCreate
}

// Save creates the JobConfig entities in the database.
func (jccb *JobConfigCreateBulk) Save(ctx context.Context) ([]*JobConfig, error) {
	if jccb.err != nil {
		return nil, jccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jccb.builders))
	nodes := make([]*JobConfig, len(jccb.builders))
	mutators := make([]Mutator, len(jccb.builders))
	for i := range jccb.builders {
		func(i int, root context.Context) {
			builder := jccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jccb *JobConfigCreateBulk) SaveX(ctx context.Context) []*JobConfig {
	v, err := jccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jccb *JobConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := jccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jccb *JobConfigCreateBulk) ExecX(ctx context.Context) {
	if err := jccb.Exec(ctx); err != nil {
		panic(err)
	}
}
