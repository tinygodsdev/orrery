// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/interpretation"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/interpretationtranslation"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/scale"
	"github.com/google/uuid"
)

// InterpretationQuery is the builder for querying Interpretation entities.
type InterpretationQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Interpretation
	// eager-loading edges.
	withTranslations *InterpretationTranslationQuery
	withScale        *ScaleQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InterpretationQuery builder.
func (iq *InterpretationQuery) Where(ps ...predicate.Interpretation) *InterpretationQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InterpretationQuery) Limit(limit int) *InterpretationQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InterpretationQuery) Offset(offset int) *InterpretationQuery {
	iq.offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InterpretationQuery) Unique(unique bool) *InterpretationQuery {
	iq.unique = &unique
	return iq
}

// Order adds an order step to the query.
func (iq *InterpretationQuery) Order(o ...OrderFunc) *InterpretationQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryTranslations chains the current query on the "translations" edge.
func (iq *InterpretationQuery) QueryTranslations() *InterpretationTranslationQuery {
	query := &InterpretationTranslationQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(interpretation.Table, interpretation.FieldID, selector),
			sqlgraph.To(interpretationtranslation.Table, interpretationtranslation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, interpretation.TranslationsTable, interpretation.TranslationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScale chains the current query on the "scale" edge.
func (iq *InterpretationQuery) QueryScale() *ScaleQuery {
	query := &ScaleQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(interpretation.Table, interpretation.FieldID, selector),
			sqlgraph.To(scale.Table, scale.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, interpretation.ScaleTable, interpretation.ScaleColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Interpretation entity from the query.
// Returns a *NotFoundError when no Interpretation was found.
func (iq *InterpretationQuery) First(ctx context.Context) (*Interpretation, error) {
	nodes, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{interpretation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InterpretationQuery) FirstX(ctx context.Context) *Interpretation {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Interpretation ID from the query.
// Returns a *NotFoundError when no Interpretation ID was found.
func (iq *InterpretationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{interpretation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InterpretationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Interpretation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Interpretation entity is found.
// Returns a *NotFoundError when no Interpretation entities are found.
func (iq *InterpretationQuery) Only(ctx context.Context) (*Interpretation, error) {
	nodes, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{interpretation.Label}
	default:
		return nil, &NotSingularError{interpretation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InterpretationQuery) OnlyX(ctx context.Context) *Interpretation {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Interpretation ID in the query.
// Returns a *NotSingularError when more than one Interpretation ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InterpretationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{interpretation.Label}
	default:
		err = &NotSingularError{interpretation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InterpretationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Interpretations.
func (iq *InterpretationQuery) All(ctx context.Context) ([]*Interpretation, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InterpretationQuery) AllX(ctx context.Context) []*Interpretation {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Interpretation IDs.
func (iq *InterpretationQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := iq.Select(interpretation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InterpretationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InterpretationQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InterpretationQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InterpretationQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InterpretationQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InterpretationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InterpretationQuery) Clone() *InterpretationQuery {
	if iq == nil {
		return nil
	}
	return &InterpretationQuery{
		config:           iq.config,
		limit:            iq.limit,
		offset:           iq.offset,
		order:            append([]OrderFunc{}, iq.order...),
		predicates:       append([]predicate.Interpretation{}, iq.predicates...),
		withTranslations: iq.withTranslations.Clone(),
		withScale:        iq.withScale.Clone(),
		// clone intermediate query.
		sql:    iq.sql.Clone(),
		path:   iq.path,
		unique: iq.unique,
	}
}

// WithTranslations tells the query-builder to eager-load the nodes that are connected to
// the "translations" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InterpretationQuery) WithTranslations(opts ...func(*InterpretationTranslationQuery)) *InterpretationQuery {
	query := &InterpretationTranslationQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withTranslations = query
	return iq
}

// WithScale tells the query-builder to eager-load the nodes that are connected to
// the "scale" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InterpretationQuery) WithScale(opts ...func(*ScaleQuery)) *InterpretationQuery {
	query := &ScaleQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withScale = query
	return iq
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
//	client.Interpretation.Query().
//		GroupBy(interpretation.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *InterpretationQuery) GroupBy(field string, fields ...string) *InterpretationGroupBy {
	grbuild := &InterpretationGroupBy{config: iq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(ctx), nil
	}
	grbuild.label = interpretation.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
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
//	client.Interpretation.Query().
//		Select(interpretation.FieldCreateTime).
//		Scan(ctx, &v)
//
func (iq *InterpretationQuery) Select(fields ...string) *InterpretationSelect {
	iq.fields = append(iq.fields, fields...)
	selbuild := &InterpretationSelect{InterpretationQuery: iq}
	selbuild.label = interpretation.Label
	selbuild.flds, selbuild.scan = &iq.fields, selbuild.Scan
	return selbuild
}

func (iq *InterpretationQuery) prepareQuery(ctx context.Context) error {
	for _, f := range iq.fields {
		if !interpretation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InterpretationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Interpretation, error) {
	var (
		nodes       = []*Interpretation{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [2]bool{
			iq.withTranslations != nil,
			iq.withScale != nil,
		}
	)
	if iq.withScale != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, interpretation.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Interpretation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Interpretation{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withTranslations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Interpretation)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Translations = []*InterpretationTranslation{}
		}
		query.withFKs = true
		query.Where(predicate.InterpretationTranslation(func(s *sql.Selector) {
			s.Where(sql.InValues(interpretation.TranslationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.interpretation_translations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "interpretation_translations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "interpretation_translations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Translations = append(node.Edges.Translations, n)
		}
	}

	if query := iq.withScale; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*Interpretation)
		for i := range nodes {
			if nodes[i].scale_interpretations == nil {
				continue
			}
			fk := *nodes[i].scale_interpretations
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(scale.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_interpretations" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Scale = n
			}
		}
	}

	return nodes, nil
}

func (iq *InterpretationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.fields
	if len(iq.fields) > 0 {
		_spec.Unique = iq.unique != nil && *iq.unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InterpretationQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (iq *InterpretationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   interpretation.Table,
			Columns: interpretation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: interpretation.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if unique := iq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := iq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, interpretation.FieldID)
		for i := range fields {
			if fields[i] != interpretation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InterpretationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(interpretation.Table)
	columns := iq.fields
	if len(columns) == 0 {
		columns = interpretation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.unique != nil && *iq.unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InterpretationGroupBy is the group-by builder for Interpretation entities.
type InterpretationGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InterpretationGroupBy) Aggregate(fns ...AggregateFunc) *InterpretationGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scans the result into the given value.
func (igb *InterpretationGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

func (igb *InterpretationGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range igb.fields {
		if !interpretation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := igb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InterpretationGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql.Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(igb.fields)+len(igb.fns))
		for _, f := range igb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(igb.fields...)...)
}

// InterpretationSelect is the builder for selecting fields of Interpretation entities.
type InterpretationSelect struct {
	*InterpretationQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (is *InterpretationSelect) Scan(ctx context.Context, v interface{}) error {
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	is.sql = is.InterpretationQuery.sqlQuery(ctx)
	return is.sqlScan(ctx, v)
}

func (is *InterpretationSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sql.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
