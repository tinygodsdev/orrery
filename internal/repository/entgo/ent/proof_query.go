// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/challenge"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/proof"
	"github.com/google/uuid"
)

// ProofQuery is the builder for querying Proof entities.
type ProofQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Proof
	// eager-loading edges.
	withChallenge *ChallengeQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProofQuery builder.
func (pq *ProofQuery) Where(ps ...predicate.Proof) *ProofQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *ProofQuery) Limit(limit int) *ProofQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *ProofQuery) Offset(offset int) *ProofQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProofQuery) Unique(unique bool) *ProofQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *ProofQuery) Order(o ...OrderFunc) *ProofQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryChallenge chains the current query on the "challenge" edge.
func (pq *ProofQuery) QueryChallenge() *ChallengeQuery {
	query := &ChallengeQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(proof.Table, proof.FieldID, selector),
			sqlgraph.To(challenge.Table, challenge.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, proof.ChallengeTable, proof.ChallengeColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Proof entity from the query.
// Returns a *NotFoundError when no Proof was found.
func (pq *ProofQuery) First(ctx context.Context) (*Proof, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{proof.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProofQuery) FirstX(ctx context.Context) *Proof {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Proof ID from the query.
// Returns a *NotFoundError when no Proof ID was found.
func (pq *ProofQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{proof.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProofQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Proof entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Proof entity is found.
// Returns a *NotFoundError when no Proof entities are found.
func (pq *ProofQuery) Only(ctx context.Context) (*Proof, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{proof.Label}
	default:
		return nil, &NotSingularError{proof.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProofQuery) OnlyX(ctx context.Context) *Proof {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Proof ID in the query.
// Returns a *NotSingularError when more than one Proof ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProofQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = &NotSingularError{proof.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProofQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Proofs.
func (pq *ProofQuery) All(ctx context.Context) ([]*Proof, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProofQuery) AllX(ctx context.Context) []*Proof {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Proof IDs.
func (pq *ProofQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := pq.Select(proof.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProofQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProofQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProofQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProofQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProofQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProofQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProofQuery) Clone() *ProofQuery {
	if pq == nil {
		return nil
	}
	return &ProofQuery{
		config:        pq.config,
		limit:         pq.limit,
		offset:        pq.offset,
		order:         append([]OrderFunc{}, pq.order...),
		predicates:    append([]predicate.Proof{}, pq.predicates...),
		withChallenge: pq.withChallenge.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithChallenge tells the query-builder to eager-load the nodes that are connected to
// the "challenge" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProofQuery) WithChallenge(opts ...func(*ChallengeQuery)) *ProofQuery {
	query := &ChallengeQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withChallenge = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Proof.Query().
//		GroupBy(proof.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *ProofQuery) GroupBy(field string, fields ...string) *ProofGroupBy {
	group := &ProofGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Proof.Query().
//		Select(proof.FieldCreateTime).
//		Scan(ctx, &v)
//
func (pq *ProofQuery) Select(fields ...string) *ProofSelect {
	pq.fields = append(pq.fields, fields...)
	return &ProofSelect{ProofQuery: pq}
}

func (pq *ProofQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !proof.ValidColumn(f) {
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

func (pq *ProofQuery) sqlAll(ctx context.Context) ([]*Proof, error) {
	var (
		nodes       = []*Proof{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withChallenge != nil,
		}
	)
	if pq.withChallenge != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, proof.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Proof{config: pq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := pq.withChallenge; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Proof)
		for i := range nodes {
			if nodes[i].challenge_proofs == nil {
				continue
			}
			fk := *nodes[i].challenge_proofs
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(challenge.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "challenge_proofs" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Challenge = n
			}
		}
	}

	return nodes, nil
}

func (pq *ProofQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProofQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *ProofQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   proof.Table,
			Columns: proof.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: proof.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, proof.FieldID)
		for i := range fields {
			if fields[i] != proof.FieldID {
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
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
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

func (pq *ProofQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(proof.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = proof.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProofGroupBy is the group-by builder for Proof entities.
type ProofGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProofGroupBy) Aggregate(fns ...AggregateFunc) *ProofGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *ProofGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *ProofGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProofGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *ProofGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *ProofGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProofGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *ProofGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *ProofGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProofGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *ProofGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *ProofGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProofGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *ProofGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProofGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *ProofGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *ProofGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !proof.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *ProofGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// ProofSelect is the builder for selecting fields of Proof entities.
type ProofSelect struct {
	*ProofQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProofSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.ProofQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *ProofSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProofSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *ProofSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *ProofSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProofSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *ProofSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *ProofSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProofSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *ProofSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *ProofSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProofSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *ProofSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ps *ProofSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{proof.Label}
	default:
		err = fmt.Errorf("ent: ProofSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *ProofSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *ProofSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
