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
	"github.com/lbrictson/janus/ent/predicate"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/projectuser"
	"github.com/lbrictson/janus/ent/user"
)

// ProjectUserQuery is the builder for querying ProjectUser entities.
type ProjectUserQuery struct {
	config
	ctx         *QueryContext
	order       []projectuser.OrderOption
	inters      []Interceptor
	predicates  []predicate.ProjectUser
	withProject *ProjectQuery
	withUser    *UserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectUserQuery builder.
func (puq *ProjectUserQuery) Where(ps ...predicate.ProjectUser) *ProjectUserQuery {
	puq.predicates = append(puq.predicates, ps...)
	return puq
}

// Limit the number of records to be returned by this query.
func (puq *ProjectUserQuery) Limit(limit int) *ProjectUserQuery {
	puq.ctx.Limit = &limit
	return puq
}

// Offset to start from.
func (puq *ProjectUserQuery) Offset(offset int) *ProjectUserQuery {
	puq.ctx.Offset = &offset
	return puq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (puq *ProjectUserQuery) Unique(unique bool) *ProjectUserQuery {
	puq.ctx.Unique = &unique
	return puq
}

// Order specifies how the records should be ordered.
func (puq *ProjectUserQuery) Order(o ...projectuser.OrderOption) *ProjectUserQuery {
	puq.order = append(puq.order, o...)
	return puq
}

// QueryProject chains the current query on the "project" edge.
func (puq *ProjectUserQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: puq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := puq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := puq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectuser.Table, projectuser.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectuser.ProjectTable, projectuser.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(puq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (puq *ProjectUserQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: puq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := puq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := puq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectuser.Table, projectuser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, projectuser.UserTable, projectuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(puq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectUser entity from the query.
// Returns a *NotFoundError when no ProjectUser was found.
func (puq *ProjectUserQuery) First(ctx context.Context) (*ProjectUser, error) {
	nodes, err := puq.Limit(1).All(setContextOp(ctx, puq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projectuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (puq *ProjectUserQuery) FirstX(ctx context.Context) *ProjectUser {
	node, err := puq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProjectUser ID from the query.
// Returns a *NotFoundError when no ProjectUser ID was found.
func (puq *ProjectUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = puq.Limit(1).IDs(setContextOp(ctx, puq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{projectuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (puq *ProjectUserQuery) FirstIDX(ctx context.Context) int {
	id, err := puq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProjectUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProjectUser entity is found.
// Returns a *NotFoundError when no ProjectUser entities are found.
func (puq *ProjectUserQuery) Only(ctx context.Context) (*ProjectUser, error) {
	nodes, err := puq.Limit(2).All(setContextOp(ctx, puq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projectuser.Label}
	default:
		return nil, &NotSingularError{projectuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (puq *ProjectUserQuery) OnlyX(ctx context.Context) *ProjectUser {
	node, err := puq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProjectUser ID in the query.
// Returns a *NotSingularError when more than one ProjectUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (puq *ProjectUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = puq.Limit(2).IDs(setContextOp(ctx, puq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{projectuser.Label}
	default:
		err = &NotSingularError{projectuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (puq *ProjectUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := puq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProjectUsers.
func (puq *ProjectUserQuery) All(ctx context.Context) ([]*ProjectUser, error) {
	ctx = setContextOp(ctx, puq.ctx, ent.OpQueryAll)
	if err := puq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProjectUser, *ProjectUserQuery]()
	return withInterceptors[[]*ProjectUser](ctx, puq, qr, puq.inters)
}

// AllX is like All, but panics if an error occurs.
func (puq *ProjectUserQuery) AllX(ctx context.Context) []*ProjectUser {
	nodes, err := puq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProjectUser IDs.
func (puq *ProjectUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if puq.ctx.Unique == nil && puq.path != nil {
		puq.Unique(true)
	}
	ctx = setContextOp(ctx, puq.ctx, ent.OpQueryIDs)
	if err = puq.Select(projectuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (puq *ProjectUserQuery) IDsX(ctx context.Context) []int {
	ids, err := puq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (puq *ProjectUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, puq.ctx, ent.OpQueryCount)
	if err := puq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, puq, querierCount[*ProjectUserQuery](), puq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (puq *ProjectUserQuery) CountX(ctx context.Context) int {
	count, err := puq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (puq *ProjectUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, puq.ctx, ent.OpQueryExist)
	switch _, err := puq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (puq *ProjectUserQuery) ExistX(ctx context.Context) bool {
	exist, err := puq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (puq *ProjectUserQuery) Clone() *ProjectUserQuery {
	if puq == nil {
		return nil
	}
	return &ProjectUserQuery{
		config:      puq.config,
		ctx:         puq.ctx.Clone(),
		order:       append([]projectuser.OrderOption{}, puq.order...),
		inters:      append([]Interceptor{}, puq.inters...),
		predicates:  append([]predicate.ProjectUser{}, puq.predicates...),
		withProject: puq.withProject.Clone(),
		withUser:    puq.withUser.Clone(),
		// clone intermediate query.
		sql:  puq.sql.Clone(),
		path: puq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (puq *ProjectUserQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectUserQuery {
	query := (&ProjectClient{config: puq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	puq.withProject = query
	return puq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (puq *ProjectUserQuery) WithUser(opts ...func(*UserQuery)) *ProjectUserQuery {
	query := (&UserClient{config: puq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	puq.withUser = query
	return puq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CanEdit bool `json:"can_edit,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProjectUser.Query().
//		GroupBy(projectuser.FieldCanEdit).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (puq *ProjectUserQuery) GroupBy(field string, fields ...string) *ProjectUserGroupBy {
	puq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProjectUserGroupBy{build: puq}
	grbuild.flds = &puq.ctx.Fields
	grbuild.label = projectuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CanEdit bool `json:"can_edit,omitempty"`
//	}
//
//	client.ProjectUser.Query().
//		Select(projectuser.FieldCanEdit).
//		Scan(ctx, &v)
func (puq *ProjectUserQuery) Select(fields ...string) *ProjectUserSelect {
	puq.ctx.Fields = append(puq.ctx.Fields, fields...)
	sbuild := &ProjectUserSelect{ProjectUserQuery: puq}
	sbuild.label = projectuser.Label
	sbuild.flds, sbuild.scan = &puq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProjectUserSelect configured with the given aggregations.
func (puq *ProjectUserQuery) Aggregate(fns ...AggregateFunc) *ProjectUserSelect {
	return puq.Select().Aggregate(fns...)
}

func (puq *ProjectUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range puq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, puq); err != nil {
				return err
			}
		}
	}
	for _, f := range puq.ctx.Fields {
		if !projectuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if puq.path != nil {
		prev, err := puq.path(ctx)
		if err != nil {
			return err
		}
		puq.sql = prev
	}
	return nil
}

func (puq *ProjectUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProjectUser, error) {
	var (
		nodes       = []*ProjectUser{}
		withFKs     = puq.withFKs
		_spec       = puq.querySpec()
		loadedTypes = [2]bool{
			puq.withProject != nil,
			puq.withUser != nil,
		}
	)
	if puq.withProject != nil || puq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, projectuser.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProjectUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProjectUser{config: puq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, puq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := puq.withProject; query != nil {
		if err := puq.loadProject(ctx, query, nodes, nil,
			func(n *ProjectUser, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := puq.withUser; query != nil {
		if err := puq.loadUser(ctx, query, nodes, nil,
			func(n *ProjectUser, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (puq *ProjectUserQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*ProjectUser, init func(*ProjectUser), assign func(*ProjectUser, *Project)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectUser)
	for i := range nodes {
		if nodes[i].project_project_users == nil {
			continue
		}
		fk := *nodes[i].project_project_users
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
			return fmt.Errorf(`unexpected foreign-key "project_project_users" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (puq *ProjectUserQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ProjectUser, init func(*ProjectUser), assign func(*ProjectUser, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProjectUser)
	for i := range nodes {
		if nodes[i].user_project_users == nil {
			continue
		}
		fk := *nodes[i].user_project_users
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_project_users" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (puq *ProjectUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := puq.querySpec()
	_spec.Node.Columns = puq.ctx.Fields
	if len(puq.ctx.Fields) > 0 {
		_spec.Unique = puq.ctx.Unique != nil && *puq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, puq.driver, _spec)
}

func (puq *ProjectUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(projectuser.Table, projectuser.Columns, sqlgraph.NewFieldSpec(projectuser.FieldID, field.TypeInt))
	_spec.From = puq.sql
	if unique := puq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if puq.path != nil {
		_spec.Unique = true
	}
	if fields := puq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, projectuser.FieldID)
		for i := range fields {
			if fields[i] != projectuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := puq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := puq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := puq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := puq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (puq *ProjectUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(puq.driver.Dialect())
	t1 := builder.Table(projectuser.Table)
	columns := puq.ctx.Fields
	if len(columns) == 0 {
		columns = projectuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if puq.sql != nil {
		selector = puq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if puq.ctx.Unique != nil && *puq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range puq.predicates {
		p(selector)
	}
	for _, p := range puq.order {
		p(selector)
	}
	if offset := puq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := puq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectUserGroupBy is the group-by builder for ProjectUser entities.
type ProjectUserGroupBy struct {
	selector
	build *ProjectUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pugb *ProjectUserGroupBy) Aggregate(fns ...AggregateFunc) *ProjectUserGroupBy {
	pugb.fns = append(pugb.fns, fns...)
	return pugb
}

// Scan applies the selector query and scans the result into the given value.
func (pugb *ProjectUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pugb.build.ctx, ent.OpQueryGroupBy)
	if err := pugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectUserQuery, *ProjectUserGroupBy](ctx, pugb.build, pugb, pugb.build.inters, v)
}

func (pugb *ProjectUserGroupBy) sqlScan(ctx context.Context, root *ProjectUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pugb.fns))
	for _, fn := range pugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pugb.flds)+len(pugb.fns))
		for _, f := range *pugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProjectUserSelect is the builder for selecting fields of ProjectUser entities.
type ProjectUserSelect struct {
	*ProjectUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pus *ProjectUserSelect) Aggregate(fns ...AggregateFunc) *ProjectUserSelect {
	pus.fns = append(pus.fns, fns...)
	return pus
}

// Scan applies the selector query and scans the result into the given value.
func (pus *ProjectUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pus.ctx, ent.OpQuerySelect)
	if err := pus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectUserQuery, *ProjectUserSelect](ctx, pus.ProjectUserQuery, pus, pus.inters, v)
}

func (pus *ProjectUserSelect) sqlScan(ctx context.Context, root *ProjectUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pus.fns))
	for _, fn := range pus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
