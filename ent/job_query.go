// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobhistory"
	"github.com/lbrictson/janus/ent/jobversion"
	"github.com/lbrictson/janus/ent/predicate"
	"github.com/lbrictson/janus/ent/project"
)

// JobQuery is the builder for querying Job entities.
type JobQuery struct {
	config
	ctx          *QueryContext
	order        []job.OrderOption
	inters       []Interceptor
	predicates   []predicate.Job
	withProject  *ProjectQuery
	withHistory  *JobHistoryQuery
	withVersions *JobVersionQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobQuery builder.
func (jq *JobQuery) Where(ps ...predicate.Job) *JobQuery {
	jq.predicates = append(jq.predicates, ps...)
	return jq
}

// Limit the number of records to be returned by this query.
func (jq *JobQuery) Limit(limit int) *JobQuery {
	jq.ctx.Limit = &limit
	return jq
}

// Offset to start from.
func (jq *JobQuery) Offset(offset int) *JobQuery {
	jq.ctx.Offset = &offset
	return jq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (jq *JobQuery) Unique(unique bool) *JobQuery {
	jq.ctx.Unique = &unique
	return jq
}

// Order specifies how the records should be ordered.
func (jq *JobQuery) Order(o ...job.OrderOption) *JobQuery {
	jq.order = append(jq.order, o...)
	return jq
}

// QueryProject chains the current query on the "project" edge.
func (jq *JobQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: jq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, job.ProjectTable, job.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(jq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHistory chains the current query on the "history" edge.
func (jq *JobQuery) QueryHistory() *JobHistoryQuery {
	query := (&JobHistoryClient{config: jq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, selector),
			sqlgraph.To(jobhistory.Table, jobhistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, job.HistoryTable, job.HistoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(jq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVersions chains the current query on the "versions" edge.
func (jq *JobQuery) QueryVersions() *JobVersionQuery {
	query := (&JobVersionClient{config: jq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := jq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := jq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(job.Table, job.FieldID, selector),
			sqlgraph.To(jobversion.Table, jobversion.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, job.VersionsTable, job.VersionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(jq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Job entity from the query.
// Returns a *NotFoundError when no Job was found.
func (jq *JobQuery) First(ctx context.Context) (*Job, error) {
	nodes, err := jq.Limit(1).All(setContextOp(ctx, jq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{job.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (jq *JobQuery) FirstX(ctx context.Context) *Job {
	node, err := jq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Job ID from the query.
// Returns a *NotFoundError when no Job ID was found.
func (jq *JobQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jq.Limit(1).IDs(setContextOp(ctx, jq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{job.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (jq *JobQuery) FirstIDX(ctx context.Context) int {
	id, err := jq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Job entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Job entity is found.
// Returns a *NotFoundError when no Job entities are found.
func (jq *JobQuery) Only(ctx context.Context) (*Job, error) {
	nodes, err := jq.Limit(2).All(setContextOp(ctx, jq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{job.Label}
	default:
		return nil, &NotSingularError{job.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (jq *JobQuery) OnlyX(ctx context.Context) *Job {
	node, err := jq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Job ID in the query.
// Returns a *NotSingularError when more than one Job ID is found.
// Returns a *NotFoundError when no entities are found.
func (jq *JobQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = jq.Limit(2).IDs(setContextOp(ctx, jq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{job.Label}
	default:
		err = &NotSingularError{job.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (jq *JobQuery) OnlyIDX(ctx context.Context) int {
	id, err := jq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Jobs.
func (jq *JobQuery) All(ctx context.Context) ([]*Job, error) {
	ctx = setContextOp(ctx, jq.ctx, ent.OpQueryAll)
	if err := jq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Job, *JobQuery]()
	return withInterceptors[[]*Job](ctx, jq, qr, jq.inters)
}

// AllX is like All, but panics if an error occurs.
func (jq *JobQuery) AllX(ctx context.Context) []*Job {
	nodes, err := jq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Job IDs.
func (jq *JobQuery) IDs(ctx context.Context) (ids []int, err error) {
	if jq.ctx.Unique == nil && jq.path != nil {
		jq.Unique(true)
	}
	ctx = setContextOp(ctx, jq.ctx, ent.OpQueryIDs)
	if err = jq.Select(job.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (jq *JobQuery) IDsX(ctx context.Context) []int {
	ids, err := jq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (jq *JobQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, jq.ctx, ent.OpQueryCount)
	if err := jq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, jq, querierCount[*JobQuery](), jq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (jq *JobQuery) CountX(ctx context.Context) int {
	count, err := jq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (jq *JobQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, jq.ctx, ent.OpQueryExist)
	switch _, err := jq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (jq *JobQuery) ExistX(ctx context.Context) bool {
	exist, err := jq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (jq *JobQuery) Clone() *JobQuery {
	if jq == nil {
		return nil
	}
	return &JobQuery{
		config:       jq.config,
		ctx:          jq.ctx.Clone(),
		order:        append([]job.OrderOption{}, jq.order...),
		inters:       append([]Interceptor{}, jq.inters...),
		predicates:   append([]predicate.Job{}, jq.predicates...),
		withProject:  jq.withProject.Clone(),
		withHistory:  jq.withHistory.Clone(),
		withVersions: jq.withVersions.Clone(),
		// clone intermediate query.
		sql:  jq.sql.Clone(),
		path: jq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (jq *JobQuery) WithProject(opts ...func(*ProjectQuery)) *JobQuery {
	query := (&ProjectClient{config: jq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jq.withProject = query
	return jq
}

// WithHistory tells the query-builder to eager-load the nodes that are connected to
// the "history" edge. The optional arguments are used to configure the query builder of the edge.
func (jq *JobQuery) WithHistory(opts ...func(*JobHistoryQuery)) *JobQuery {
	query := (&JobHistoryClient{config: jq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jq.withHistory = query
	return jq
}

// WithVersions tells the query-builder to eager-load the nodes that are connected to
// the "versions" edge. The optional arguments are used to configure the query builder of the edge.
func (jq *JobQuery) WithVersions(opts ...func(*JobVersionQuery)) *JobQuery {
	query := (&JobVersionClient{config: jq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	jq.withVersions = query
	return jq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Job.Query().
//		GroupBy(job.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (jq *JobQuery) GroupBy(field string, fields ...string) *JobGroupBy {
	jq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &JobGroupBy{build: jq}
	grbuild.flds = &jq.ctx.Fields
	grbuild.label = job.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Job.Query().
//		Select(job.FieldName).
//		Scan(ctx, &v)
func (jq *JobQuery) Select(fields ...string) *JobSelect {
	jq.ctx.Fields = append(jq.ctx.Fields, fields...)
	sbuild := &JobSelect{JobQuery: jq}
	sbuild.label = job.Label
	sbuild.flds, sbuild.scan = &jq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a JobSelect configured with the given aggregations.
func (jq *JobQuery) Aggregate(fns ...AggregateFunc) *JobSelect {
	return jq.Select().Aggregate(fns...)
}

func (jq *JobQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range jq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, jq); err != nil {
				return err
			}
		}
	}
	for _, f := range jq.ctx.Fields {
		if !job.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if jq.path != nil {
		prev, err := jq.path(ctx)
		if err != nil {
			return err
		}
		jq.sql = prev
	}
	return nil
}

func (jq *JobQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Job, error) {
	var (
		nodes       = []*Job{}
		withFKs     = jq.withFKs
		_spec       = jq.querySpec()
		loadedTypes = [3]bool{
			jq.withProject != nil,
			jq.withHistory != nil,
			jq.withVersions != nil,
		}
	)
	if jq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, job.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Job).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Job{config: jq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, jq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := jq.withProject; query != nil {
		if err := jq.loadProject(ctx, query, nodes, nil,
			func(n *Job, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := jq.withHistory; query != nil {
		if err := jq.loadHistory(ctx, query, nodes,
			func(n *Job) { n.Edges.History = []*JobHistory{} },
			func(n *Job, e *JobHistory) { n.Edges.History = append(n.Edges.History, e) }); err != nil {
			return nil, err
		}
	}
	if query := jq.withVersions; query != nil {
		if err := jq.loadVersions(ctx, query, nodes,
			func(n *Job) { n.Edges.Versions = []*JobVersion{} },
			func(n *Job, e *JobVersion) { n.Edges.Versions = append(n.Edges.Versions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (jq *JobQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Job, init func(*Job), assign func(*Job, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Job)
	for i := range nodes {
		if nodes[i].project_jobs == nil {
			continue
		}
		fk := *nodes[i].project_jobs
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_jobs" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (jq *JobQuery) loadHistory(ctx context.Context, query *JobHistoryQuery, nodes []*Job, init func(*Job), assign func(*Job, *JobHistory)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Job)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.JobHistory(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(job.HistoryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.job_history
		if fk == nil {
			return fmt.Errorf(`foreign-key "job_history" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "job_history" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (jq *JobQuery) loadVersions(ctx context.Context, query *JobVersionQuery, nodes []*Job, init func(*Job), assign func(*Job, *JobVersion)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Job)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.JobVersion(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(job.VersionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.job_versions
		if fk == nil {
			return fmt.Errorf(`foreign-key "job_versions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "job_versions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (jq *JobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := jq.querySpec()
	_spec.Node.Columns = jq.ctx.Fields
	if len(jq.ctx.Fields) > 0 {
		_spec.Unique = jq.ctx.Unique != nil && *jq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, jq.driver, _spec)
}

func (jq *JobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(job.Table, job.Columns, sqlgraph.NewFieldSpec(job.FieldID, field.TypeInt))
	_spec.From = jq.sql
	if unique := jq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if jq.path != nil {
		_spec.Unique = true
	}
	if fields := jq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, job.FieldID)
		for i := range fields {
			if fields[i] != job.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := jq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := jq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := jq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := jq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (jq *JobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(jq.driver.Dialect())
	t1 := builder.Table(job.Table)
	columns := jq.ctx.Fields
	if len(columns) == 0 {
		columns = job.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if jq.sql != nil {
		selector = jq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if jq.ctx.Unique != nil && *jq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range jq.predicates {
		p(selector)
	}
	for _, p := range jq.order {
		p(selector)
	}
	if offset := jq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := jq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobGroupBy is the group-by builder for Job entities.
type JobGroupBy struct {
	selector
	build *JobQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jgb *JobGroupBy) Aggregate(fns ...AggregateFunc) *JobGroupBy {
	jgb.fns = append(jgb.fns, fns...)
	return jgb
}

// Scan applies the selector query and scans the result into the given value.
func (jgb *JobGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, jgb.build.ctx, ent.OpQueryGroupBy)
	if err := jgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobQuery, *JobGroupBy](ctx, jgb.build, jgb, jgb.build.inters, v)
}

func (jgb *JobGroupBy) sqlScan(ctx context.Context, root *JobQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(jgb.fns))
	for _, fn := range jgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*jgb.flds)+len(jgb.fns))
		for _, f := range *jgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*jgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// JobSelect is the builder for selecting fields of Job entities.
type JobSelect struct {
	*JobQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (js *JobSelect) Aggregate(fns ...AggregateFunc) *JobSelect {
	js.fns = append(js.fns, fns...)
	return js
}

// Scan applies the selector query and scans the result into the given value.
func (js *JobSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, js.ctx, ent.OpQuerySelect)
	if err := js.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*JobQuery, *JobSelect](ctx, js.JobQuery, js, js.inters, v)
}

func (js *JobSelect) sqlScan(ctx context.Context, root *JobQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(js.fns))
	for _, fn := range js.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*js.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := js.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
