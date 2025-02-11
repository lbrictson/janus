// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/janus/ent/jobconfig"
	"github.com/lbrictson/janus/ent/predicate"
)

// JobConfigQuery is the builder for querying JobConfig entities.
type JobConfigQuery struct {
	config
	ctx        *QueryContext
	order      []jobconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.JobConfig
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobConfigQuery builder.
func (jcq *JobConfigQuery) Where(ps ...predicate.JobConfig) *JobConfigQuery {
	jcq.predicates = append(jcq.predicates, ps...)
	return jcq
}

// Limit the number of records to be returned by this query.
func (jcq *JobConfigQuery) Limit(limit int) *JobConfigQuery {
	jcq.ctx.Limit = &limit
	return jcq
}

// Offset to start from.
func (jcq *JobConfigQuery) Offset(offset int) *JobConfigQuery {
	jcq.ctx.Offset = &offset
	return jcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jcq *JobConfigQuery) Unique(unique bool) *JobConfigQuery {
	jcq.ctx.Unique = &unique
	return jcq
}

// Order specifies how the records should be ordered.
func (jcq *JobConfigQuery) Order(o ...jobconfig.OrderOption) *JobConfigQuery {
	jcq.order = append(jcq.order, o...)
	return jcq
}

// First returns the first JobConfig entity from the query.
// Returns a *NotFoundError when no JobConfig was found.
func (jcq *JobConfigQuery) First(ctx context.Context) (*JobConfig, error) {
	nodes, err := jcq.Limit(1).All(setContextOp(ctx, jcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{jobconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jcq *JobConfigQuery) FirstX(ctx context.Context) *JobConfig {
	node, err := jcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobConfig ID from the query.
// Returns a *NotFoundError when no JobConfig ID was found.
func (jcq *JobConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jcq.Limit(1).IDs(setContextOp(ctx, jcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{jobconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jcq *JobConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := jcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobConfig entity is found.
// Returns a *NotFoundError when no JobConfig entities are found.
func (jcq *JobConfigQuery) Only(ctx context.Context) (*JobConfig, error) {
	nodes, err := jcq.Limit(2).All(setContextOp(ctx, jcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{jobconfig.Label}
	default:
		return nil, &NotSingularError{jobconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jcq *JobConfigQuery) OnlyX(ctx context.Context) *JobConfig {
	node, err := jcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobConfig ID in the query.
// Returns a *NotSingularError when more than one JobConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (jcq *JobConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jcq.Limit(2).IDs(setContextOp(ctx, jcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{jobconfig.Label}
	default:
		err = &NotSingularError{jobconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jcq *JobConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := jcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobConfigs.
func (jcq *JobConfigQuery) All(ctx context.Context) ([]*JobConfig, error) {
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryAll)
	if err := jcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*JobConfig, *JobConfigQuery]()
	return withInterceptors[[]*JobConfig](ctx, jcq, qr, jcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jcq *JobConfigQuery) AllX(ctx context.Context) []*JobConfig {
	nodes, err := jcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobConfig IDs.
func (jcq *JobConfigQuery) IDs(ctx context.Context) (ids []int, err error) {
	if jcq.ctx.Unique == nil && jcq.path != nil {
		jcq.Unique(true)
	}
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryIDs)
	if err = jcq.Select(jobconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jcq *JobConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := jcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jcq *JobConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryCount)
	if err := jcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jcq, querierCount[*JobConfigQuery](), jcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jcq *JobConfigQuery) CountX(ctx context.Context) int {
	count, err := jcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jcq *JobConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jcq.ctx, ent.OpQueryExist)
	switch _, err := jcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jcq *JobConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := jcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jcq *JobConfigQuery) Clone() *JobConfigQuery {
	if jcq == nil {
		return nil
	}
	return &JobConfigQuery{
		config:     jcq.config,
		ctx:        jcq.ctx.Clone(),
		order:      append([]jobconfig.OrderOption{}, jcq.order...),
		inters:     append([]Interceptor{}, jcq.inters...),
		predicates: append([]predicate.JobConfig{}, jcq.predicates...),
		// clone intermediate query.
		sql:  jcq.sql.Clone(),
		path: jcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MaxConcurrentJobs int `json:"max_concurrent_jobs,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobConfig.Query().
//		GroupBy(jobconfig.FieldMaxConcurrentJobs).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jcq *JobConfigQuery) GroupBy(field string, fields ...string) *JobConfigGroupBy {
	jcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobConfigGroupBy{build: jcq}
	grbuild.flds = &jcq.ctx.Fields
	grbuild.label = jobconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MaxConcurrentJobs int `json:"max_concurrent_jobs,omitempty"`
//	}
//
//	client.JobConfig.Query().
//		Select(jobconfig.FieldMaxConcurrentJobs).
//		Scan(ctx, &v)
func (jcq *JobConfigQuery) Select(fields ...string) *JobConfigSelect {
	jcq.ctx.Fields = append(jcq.ctx.Fields, fields...)
	sbuild := &JobConfigSelect{JobConfigQuery: jcq}
	sbuild.label = jobconfig.Label
	sbuild.flds, sbuild.scan = &jcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobConfigSelect configured with the given aggregations.
func (jcq *JobConfigQuery) Aggregate(fns ...AggregateFunc) *JobConfigSelect {
	return jcq.Select().Aggregate(fns...)
}

func (jcq *JobConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jcq); err != nil {
				return err
			}
		}
	}
	for _, f := range jcq.ctx.Fields {
		if !jobconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jcq.path != nil {
		prev, err := jcq.path(ctx)
		if err != nil {
			return err
		}
		jcq.sql = prev
	}
	return nil
}

func (jcq *JobConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*JobConfig, error) {
	var (
		nodes = []*JobConfig{}
		_spec = jcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*JobConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &JobConfig{config: jcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (jcq *JobConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jcq.querySpec()
	_spec.Node.Columns = jcq.ctx.Fields
	if len(jcq.ctx.Fields) > 0 {
		_spec.Unique = jcq.ctx.Unique != nil && *jcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jcq.driver, _spec)
}

func (jcq *JobConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(jobconfig.Table, jobconfig.Columns, sqlgraph.NewFieldSpec(jobconfig.FieldID, field.TypeInt))
	_spec.From = jcq.sql
	if unique := jcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jcq.path != nil {
		_spec.Unique = true
	}
	if fields := jcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, jobconfig.FieldID)
		for i := range fields {
			if fields[i] != jobconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jcq *JobConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jcq.driver.Dialect())
	t1 := builder.Table(jobconfig.Table)
	columns := jcq.ctx.Fields
	if len(columns) == 0 {
		columns = jobconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jcq.sql != nil {
		selector = jcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jcq.ctx.Unique != nil && *jcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jcq.predicates {
		p(selector)
	}
	for _, p := range jcq.order {
		p(selector)
	}
	if offset := jcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobConfigGroupBy is the group-by builder for JobConfig entities.
type JobConfigGroupBy struct {
	selector
	build *JobConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jcgb *JobConfigGroupBy) Aggregate(fns ...AggregateFunc) *JobConfigGroupBy {
	jcgb.fns = append(jcgb.fns, fns...)
	return jcgb
}

// Scan applies the selector query and scans the result into the given value.
func (jcgb *JobConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jcgb.build.ctx, ent.OpQueryGroupBy)
	if err := jcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobConfigQuery, *JobConfigGroupBy](ctx, jcgb.build, jcgb, jcgb.build.inters, v)
}

func (jcgb *JobConfigGroupBy) sqlScan(ctx context.Context, root *JobConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jcgb.fns))
	for _, fn := range jcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jcgb.flds)+len(jcgb.fns))
		for _, f := range *jcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobConfigSelect is the builder for selecting fields of JobConfig entities.
type JobConfigSelect struct {
	*JobConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (jcs *JobConfigSelect) Aggregate(fns ...AggregateFunc) *JobConfigSelect {
	jcs.fns = append(jcs.fns, fns...)
	return jcs
}

// Scan applies the selector query and scans the result into the given value.
func (jcs *JobConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jcs.ctx, ent.OpQuerySelect)
	if err := jcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobConfigQuery, *JobConfigSelect](ctx, jcs.JobConfigQuery, jcs, jcs.inters, v)
}

func (jcs *JobConfigSelect) sqlScan(ctx context.Context, root *JobConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(jcs.fns))
	for _, fn := range jcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*jcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
