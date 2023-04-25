// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github-bot/ent/predicate"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PullsQuery is the builder for querying Pulls entities.
type PullsQuery struct {
	config
	ctx         *QueryContext
	order       []pulls.OrderOption
	inters      []Interceptor
	predicates  []predicate.Pulls
	withCreator *UserQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PullsQuery builder.
func (pq *PullsQuery) Where(ps ...predicate.Pulls) *PullsQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PullsQuery) Limit(limit int) *PullsQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PullsQuery) Offset(offset int) *PullsQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PullsQuery) Unique(unique bool) *PullsQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PullsQuery) Order(o ...pulls.OrderOption) *PullsQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryCreator chains the current query on the "Creator" edge.
func (pq *PullsQuery) QueryCreator() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pulls.Table, pulls.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, pulls.CreatorTable, pulls.CreatorColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pulls entity from the query.
// Returns a *NotFoundError when no Pulls was found.
func (pq *PullsQuery) First(ctx context.Context) (*Pulls, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pulls.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PullsQuery) FirstX(ctx context.Context) *Pulls {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pulls ID from the query.
// Returns a *NotFoundError when no Pulls ID was found.
func (pq *PullsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pulls.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PullsQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pulls entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pulls entity is found.
// Returns a *NotFoundError when no Pulls entities are found.
func (pq *PullsQuery) Only(ctx context.Context) (*Pulls, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pulls.Label}
	default:
		return nil, &NotSingularError{pulls.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PullsQuery) OnlyX(ctx context.Context) *Pulls {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pulls ID in the query.
// Returns a *NotSingularError when more than one Pulls ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PullsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pulls.Label}
	default:
		err = &NotSingularError{pulls.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PullsQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PullsSlice.
func (pq *PullsQuery) All(ctx context.Context) ([]*Pulls, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Pulls, *PullsQuery]()
	return withInterceptors[[]*Pulls](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PullsQuery) AllX(ctx context.Context) []*Pulls {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pulls IDs.
func (pq *PullsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(pulls.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PullsQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PullsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PullsQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PullsQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PullsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PullsQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PullsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PullsQuery) Clone() *PullsQuery {
	if pq == nil {
		return nil
	}
	return &PullsQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]pulls.OrderOption{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Pulls{}, pq.predicates...),
		withCreator: pq.withCreator.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithCreator tells the query-builder to eager-load the nodes that are connected to
// the "Creator" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PullsQuery) WithCreator(opts ...func(*UserQuery)) *PullsQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCreator = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Owner string `json:"owner,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pulls.Query().
//		GroupBy(pulls.FieldOwner).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PullsQuery) GroupBy(field string, fields ...string) *PullsGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PullsGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = pulls.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Owner string `json:"owner,omitempty"`
//	}
//
//	client.Pulls.Query().
//		Select(pulls.FieldOwner).
//		Scan(ctx, &v)
func (pq *PullsQuery) Select(fields ...string) *PullsSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PullsSelect{PullsQuery: pq}
	sbuild.label = pulls.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PullsSelect configured with the given aggregations.
func (pq *PullsQuery) Aggregate(fns ...AggregateFunc) *PullsSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PullsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !pulls.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PullsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pulls, error) {
	var (
		nodes       = []*Pulls{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withCreator != nil,
		}
	)
	if pq.withCreator != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pulls.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pulls).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pulls{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withCreator; query != nil {
		if err := pq.loadCreator(ctx, query, nodes, nil,
			func(n *Pulls, e *User) { n.Edges.Creator = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PullsQuery) loadCreator(ctx context.Context, query *UserQuery, nodes []*Pulls, init func(*Pulls), assign func(*Pulls, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Pulls)
	for i := range nodes {
		if nodes[i].user_pull_requests == nil {
			continue
		}
		fk := *nodes[i].user_pull_requests
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
			return fmt.Errorf(`unexpected foreign-key "user_pull_requests" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PullsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PullsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(pulls.Table, pulls.Columns, sqlgraph.NewFieldSpec(pulls.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pulls.FieldID)
		for i := range fields {
			if fields[i] != pulls.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PullsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pulls.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = pulls.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PullsGroupBy is the group-by builder for Pulls entities.
type PullsGroupBy struct {
	selector
	build *PullsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PullsGroupBy) Aggregate(fns ...AggregateFunc) *PullsGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PullsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PullsQuery, *PullsGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PullsGroupBy) sqlScan(ctx context.Context, root *PullsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PullsSelect is the builder for selecting fields of Pulls entities.
type PullsSelect struct {
	*PullsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PullsSelect) Aggregate(fns ...AggregateFunc) *PullsSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PullsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PullsQuery, *PullsSelect](ctx, ps.PullsQuery, ps, ps.inters, v)
}

func (ps *PullsSelect) sqlScan(ctx context.Context, root *PullsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
