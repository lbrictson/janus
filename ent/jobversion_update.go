// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobversion"
	"github.com/lbrictson/janus/ent/predicate"
	"github.com/lbrictson/janus/ent/schema"
)

// JobVersionUpdate is the builder for updating JobVersion entities.
type JobVersionUpdate struct {
	config
	hooks    []Hook
	mutation *JobVersionMutation
}

// Where appends a list predicates to the JobVersionUpdate builder.
func (jvu *JobVersionUpdate) Where(ps ...predicate.JobVersion) *JobVersionUpdate {
	jvu.mutation.Where(ps...)
	return jvu
}

// SetCreatedAt sets the "created_at" field.
func (jvu *JobVersionUpdate) SetCreatedAt(t time.Time) *JobVersionUpdate {
	jvu.mutation.SetCreatedAt(t)
	return jvu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableCreatedAt(t *time.Time) *JobVersionUpdate {
	if t != nil {
		jvu.SetCreatedAt(*t)
	}
	return jvu
}

// SetName sets the "name" field.
func (jvu *JobVersionUpdate) SetName(s string) *JobVersionUpdate {
	jvu.mutation.SetName(s)
	return jvu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableName(s *string) *JobVersionUpdate {
	if s != nil {
		jvu.SetName(*s)
	}
	return jvu
}

// SetDescription sets the "description" field.
func (jvu *JobVersionUpdate) SetDescription(s string) *JobVersionUpdate {
	jvu.mutation.SetDescription(s)
	return jvu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableDescription(s *string) *JobVersionUpdate {
	if s != nil {
		jvu.SetDescription(*s)
	}
	return jvu
}

// SetScript sets the "script" field.
func (jvu *JobVersionUpdate) SetScript(s string) *JobVersionUpdate {
	jvu.mutation.SetScript(s)
	return jvu
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableScript(s *string) *JobVersionUpdate {
	if s != nil {
		jvu.SetScript(*s)
	}
	return jvu
}

// SetCronSchedule sets the "cron_schedule" field.
func (jvu *JobVersionUpdate) SetCronSchedule(s string) *JobVersionUpdate {
	jvu.mutation.SetCronSchedule(s)
	return jvu
}

// SetNillableCronSchedule sets the "cron_schedule" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableCronSchedule(s *string) *JobVersionUpdate {
	if s != nil {
		jvu.SetCronSchedule(*s)
	}
	return jvu
}

// ClearCronSchedule clears the value of the "cron_schedule" field.
func (jvu *JobVersionUpdate) ClearCronSchedule() *JobVersionUpdate {
	jvu.mutation.ClearCronSchedule()
	return jvu
}

// SetScheduleEnabled sets the "schedule_enabled" field.
func (jvu *JobVersionUpdate) SetScheduleEnabled(b bool) *JobVersionUpdate {
	jvu.mutation.SetScheduleEnabled(b)
	return jvu
}

// SetNillableScheduleEnabled sets the "schedule_enabled" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableScheduleEnabled(b *bool) *JobVersionUpdate {
	if b != nil {
		jvu.SetScheduleEnabled(*b)
	}
	return jvu
}

// SetAllowConcurrentRuns sets the "allow_concurrent_runs" field.
func (jvu *JobVersionUpdate) SetAllowConcurrentRuns(b bool) *JobVersionUpdate {
	jvu.mutation.SetAllowConcurrentRuns(b)
	return jvu
}

// SetNillableAllowConcurrentRuns sets the "allow_concurrent_runs" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableAllowConcurrentRuns(b *bool) *JobVersionUpdate {
	if b != nil {
		jvu.SetAllowConcurrentRuns(*b)
	}
	return jvu
}

// SetArguments sets the "arguments" field.
func (jvu *JobVersionUpdate) SetArguments(sa []schema.JobArgument) *JobVersionUpdate {
	jvu.mutation.SetArguments(sa)
	return jvu
}

// AppendArguments appends sa to the "arguments" field.
func (jvu *JobVersionUpdate) AppendArguments(sa []schema.JobArgument) *JobVersionUpdate {
	jvu.mutation.AppendArguments(sa)
	return jvu
}

// SetRequiresFileUpload sets the "requires_file_upload" field.
func (jvu *JobVersionUpdate) SetRequiresFileUpload(b bool) *JobVersionUpdate {
	jvu.mutation.SetRequiresFileUpload(b)
	return jvu
}

// SetNillableRequiresFileUpload sets the "requires_file_upload" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableRequiresFileUpload(b *bool) *JobVersionUpdate {
	if b != nil {
		jvu.SetRequiresFileUpload(*b)
	}
	return jvu
}

// SetChangedByEmail sets the "changed_by_email" field.
func (jvu *JobVersionUpdate) SetChangedByEmail(s string) *JobVersionUpdate {
	jvu.mutation.SetChangedByEmail(s)
	return jvu
}

// SetNillableChangedByEmail sets the "changed_by_email" field if the given value is not nil.
func (jvu *JobVersionUpdate) SetNillableChangedByEmail(s *string) *JobVersionUpdate {
	if s != nil {
		jvu.SetChangedByEmail(*s)
	}
	return jvu
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jvu *JobVersionUpdate) SetJobID(id int) *JobVersionUpdate {
	jvu.mutation.SetJobID(id)
	return jvu
}

// SetJob sets the "job" edge to the Job entity.
func (jvu *JobVersionUpdate) SetJob(j *Job) *JobVersionUpdate {
	return jvu.SetJobID(j.ID)
}

// Mutation returns the JobVersionMutation object of the builder.
func (jvu *JobVersionUpdate) Mutation() *JobVersionMutation {
	return jvu.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jvu *JobVersionUpdate) ClearJob() *JobVersionUpdate {
	jvu.mutation.ClearJob()
	return jvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (jvu *JobVersionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, jvu.sqlSave, jvu.mutation, jvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jvu *JobVersionUpdate) SaveX(ctx context.Context) int {
	affected, err := jvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (jvu *JobVersionUpdate) Exec(ctx context.Context) error {
	_, err := jvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jvu *JobVersionUpdate) ExecX(ctx context.Context) {
	if err := jvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jvu *JobVersionUpdate) check() error {
	if jvu.mutation.JobCleared() && len(jvu.mutation.JobIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "JobVersion.job"`)
	}
	return nil
}

func (jvu *JobVersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := jvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(jobversion.Table, jobversion.Columns, sqlgraph.NewFieldSpec(jobversion.FieldID, field.TypeInt))
	if ps := jvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jvu.mutation.CreatedAt(); ok {
		_spec.SetField(jobversion.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jvu.mutation.Name(); ok {
		_spec.SetField(jobversion.FieldName, field.TypeString, value)
	}
	if value, ok := jvu.mutation.Description(); ok {
		_spec.SetField(jobversion.FieldDescription, field.TypeString, value)
	}
	if value, ok := jvu.mutation.Script(); ok {
		_spec.SetField(jobversion.FieldScript, field.TypeString, value)
	}
	if value, ok := jvu.mutation.CronSchedule(); ok {
		_spec.SetField(jobversion.FieldCronSchedule, field.TypeString, value)
	}
	if jvu.mutation.CronScheduleCleared() {
		_spec.ClearField(jobversion.FieldCronSchedule, field.TypeString)
	}
	if value, ok := jvu.mutation.ScheduleEnabled(); ok {
		_spec.SetField(jobversion.FieldScheduleEnabled, field.TypeBool, value)
	}
	if value, ok := jvu.mutation.AllowConcurrentRuns(); ok {
		_spec.SetField(jobversion.FieldAllowConcurrentRuns, field.TypeBool, value)
	}
	if value, ok := jvu.mutation.Arguments(); ok {
		_spec.SetField(jobversion.FieldArguments, field.TypeJSON, value)
	}
	if value, ok := jvu.mutation.AppendedArguments(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, jobversion.FieldArguments, value)
		})
	}
	if value, ok := jvu.mutation.RequiresFileUpload(); ok {
		_spec.SetField(jobversion.FieldRequiresFileUpload, field.TypeBool, value)
	}
	if value, ok := jvu.mutation.ChangedByEmail(); ok {
		_spec.SetField(jobversion.FieldChangedByEmail, field.TypeString, value)
	}
	if jvu.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobversion.JobTable,
			Columns: []string{jobversion.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jvu.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobversion.JobTable,
			Columns: []string{jobversion.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, jvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	jvu.mutation.done = true
	return n, nil
}

// JobVersionUpdateOne is the builder for updating a single JobVersion entity.
type JobVersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobVersionMutation
}

// SetCreatedAt sets the "created_at" field.
func (jvuo *JobVersionUpdateOne) SetCreatedAt(t time.Time) *JobVersionUpdateOne {
	jvuo.mutation.SetCreatedAt(t)
	return jvuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableCreatedAt(t *time.Time) *JobVersionUpdateOne {
	if t != nil {
		jvuo.SetCreatedAt(*t)
	}
	return jvuo
}

// SetName sets the "name" field.
func (jvuo *JobVersionUpdateOne) SetName(s string) *JobVersionUpdateOne {
	jvuo.mutation.SetName(s)
	return jvuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableName(s *string) *JobVersionUpdateOne {
	if s != nil {
		jvuo.SetName(*s)
	}
	return jvuo
}

// SetDescription sets the "description" field.
func (jvuo *JobVersionUpdateOne) SetDescription(s string) *JobVersionUpdateOne {
	jvuo.mutation.SetDescription(s)
	return jvuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableDescription(s *string) *JobVersionUpdateOne {
	if s != nil {
		jvuo.SetDescription(*s)
	}
	return jvuo
}

// SetScript sets the "script" field.
func (jvuo *JobVersionUpdateOne) SetScript(s string) *JobVersionUpdateOne {
	jvuo.mutation.SetScript(s)
	return jvuo
}

// SetNillableScript sets the "script" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableScript(s *string) *JobVersionUpdateOne {
	if s != nil {
		jvuo.SetScript(*s)
	}
	return jvuo
}

// SetCronSchedule sets the "cron_schedule" field.
func (jvuo *JobVersionUpdateOne) SetCronSchedule(s string) *JobVersionUpdateOne {
	jvuo.mutation.SetCronSchedule(s)
	return jvuo
}

// SetNillableCronSchedule sets the "cron_schedule" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableCronSchedule(s *string) *JobVersionUpdateOne {
	if s != nil {
		jvuo.SetCronSchedule(*s)
	}
	return jvuo
}

// ClearCronSchedule clears the value of the "cron_schedule" field.
func (jvuo *JobVersionUpdateOne) ClearCronSchedule() *JobVersionUpdateOne {
	jvuo.mutation.ClearCronSchedule()
	return jvuo
}

// SetScheduleEnabled sets the "schedule_enabled" field.
func (jvuo *JobVersionUpdateOne) SetScheduleEnabled(b bool) *JobVersionUpdateOne {
	jvuo.mutation.SetScheduleEnabled(b)
	return jvuo
}

// SetNillableScheduleEnabled sets the "schedule_enabled" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableScheduleEnabled(b *bool) *JobVersionUpdateOne {
	if b != nil {
		jvuo.SetScheduleEnabled(*b)
	}
	return jvuo
}

// SetAllowConcurrentRuns sets the "allow_concurrent_runs" field.
func (jvuo *JobVersionUpdateOne) SetAllowConcurrentRuns(b bool) *JobVersionUpdateOne {
	jvuo.mutation.SetAllowConcurrentRuns(b)
	return jvuo
}

// SetNillableAllowConcurrentRuns sets the "allow_concurrent_runs" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableAllowConcurrentRuns(b *bool) *JobVersionUpdateOne {
	if b != nil {
		jvuo.SetAllowConcurrentRuns(*b)
	}
	return jvuo
}

// SetArguments sets the "arguments" field.
func (jvuo *JobVersionUpdateOne) SetArguments(sa []schema.JobArgument) *JobVersionUpdateOne {
	jvuo.mutation.SetArguments(sa)
	return jvuo
}

// AppendArguments appends sa to the "arguments" field.
func (jvuo *JobVersionUpdateOne) AppendArguments(sa []schema.JobArgument) *JobVersionUpdateOne {
	jvuo.mutation.AppendArguments(sa)
	return jvuo
}

// SetRequiresFileUpload sets the "requires_file_upload" field.
func (jvuo *JobVersionUpdateOne) SetRequiresFileUpload(b bool) *JobVersionUpdateOne {
	jvuo.mutation.SetRequiresFileUpload(b)
	return jvuo
}

// SetNillableRequiresFileUpload sets the "requires_file_upload" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableRequiresFileUpload(b *bool) *JobVersionUpdateOne {
	if b != nil {
		jvuo.SetRequiresFileUpload(*b)
	}
	return jvuo
}

// SetChangedByEmail sets the "changed_by_email" field.
func (jvuo *JobVersionUpdateOne) SetChangedByEmail(s string) *JobVersionUpdateOne {
	jvuo.mutation.SetChangedByEmail(s)
	return jvuo
}

// SetNillableChangedByEmail sets the "changed_by_email" field if the given value is not nil.
func (jvuo *JobVersionUpdateOne) SetNillableChangedByEmail(s *string) *JobVersionUpdateOne {
	if s != nil {
		jvuo.SetChangedByEmail(*s)
	}
	return jvuo
}

// SetJobID sets the "job" edge to the Job entity by ID.
func (jvuo *JobVersionUpdateOne) SetJobID(id int) *JobVersionUpdateOne {
	jvuo.mutation.SetJobID(id)
	return jvuo
}

// SetJob sets the "job" edge to the Job entity.
func (jvuo *JobVersionUpdateOne) SetJob(j *Job) *JobVersionUpdateOne {
	return jvuo.SetJobID(j.ID)
}

// Mutation returns the JobVersionMutation object of the builder.
func (jvuo *JobVersionUpdateOne) Mutation() *JobVersionMutation {
	return jvuo.mutation
}

// ClearJob clears the "job" edge to the Job entity.
func (jvuo *JobVersionUpdateOne) ClearJob() *JobVersionUpdateOne {
	jvuo.mutation.ClearJob()
	return jvuo
}

// Where appends a list predicates to the JobVersionUpdate builder.
func (jvuo *JobVersionUpdateOne) Where(ps ...predicate.JobVersion) *JobVersionUpdateOne {
	jvuo.mutation.Where(ps...)
	return jvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (jvuo *JobVersionUpdateOne) Select(field string, fields ...string) *JobVersionUpdateOne {
	jvuo.fields = append([]string{field}, fields...)
	return jvuo
}

// Save executes the query and returns the updated JobVersion entity.
func (jvuo *JobVersionUpdateOne) Save(ctx context.Context) (*JobVersion, error) {
	return withHooks(ctx, jvuo.sqlSave, jvuo.mutation, jvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (jvuo *JobVersionUpdateOne) SaveX(ctx context.Context) *JobVersion {
	node, err := jvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (jvuo *JobVersionUpdateOne) Exec(ctx context.Context) error {
	_, err := jvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jvuo *JobVersionUpdateOne) ExecX(ctx context.Context) {
	if err := jvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jvuo *JobVersionUpdateOne) check() error {
	if jvuo.mutation.JobCleared() && len(jvuo.mutation.JobIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "JobVersion.job"`)
	}
	return nil
}

func (jvuo *JobVersionUpdateOne) sqlSave(ctx context.Context) (_node *JobVersion, err error) {
	if err := jvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(jobversion.Table, jobversion.Columns, sqlgraph.NewFieldSpec(jobversion.FieldID, field.TypeInt))
	id, ok := jvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "JobVersion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := jvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobversion.FieldID)
		for _, f := range fields {
			if !jobversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != jobversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := jvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := jvuo.mutation.CreatedAt(); ok {
		_spec.SetField(jobversion.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := jvuo.mutation.Name(); ok {
		_spec.SetField(jobversion.FieldName, field.TypeString, value)
	}
	if value, ok := jvuo.mutation.Description(); ok {
		_spec.SetField(jobversion.FieldDescription, field.TypeString, value)
	}
	if value, ok := jvuo.mutation.Script(); ok {
		_spec.SetField(jobversion.FieldScript, field.TypeString, value)
	}
	if value, ok := jvuo.mutation.CronSchedule(); ok {
		_spec.SetField(jobversion.FieldCronSchedule, field.TypeString, value)
	}
	if jvuo.mutation.CronScheduleCleared() {
		_spec.ClearField(jobversion.FieldCronSchedule, field.TypeString)
	}
	if value, ok := jvuo.mutation.ScheduleEnabled(); ok {
		_spec.SetField(jobversion.FieldScheduleEnabled, field.TypeBool, value)
	}
	if value, ok := jvuo.mutation.AllowConcurrentRuns(); ok {
		_spec.SetField(jobversion.FieldAllowConcurrentRuns, field.TypeBool, value)
	}
	if value, ok := jvuo.mutation.Arguments(); ok {
		_spec.SetField(jobversion.FieldArguments, field.TypeJSON, value)
	}
	if value, ok := jvuo.mutation.AppendedArguments(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, jobversion.FieldArguments, value)
		})
	}
	if value, ok := jvuo.mutation.RequiresFileUpload(); ok {
		_spec.SetField(jobversion.FieldRequiresFileUpload, field.TypeBool, value)
	}
	if value, ok := jvuo.mutation.ChangedByEmail(); ok {
		_spec.SetField(jobversion.FieldChangedByEmail, field.TypeString, value)
	}
	if jvuo.mutation.JobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobversion.JobTable,
			Columns: []string{jobversion.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := jvuo.mutation.JobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   jobversion.JobTable,
			Columns: []string{jobversion.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(job.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &JobVersion{config: jvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, jvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{jobversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	jvuo.mutation.done = true
	return _node, nil
}
