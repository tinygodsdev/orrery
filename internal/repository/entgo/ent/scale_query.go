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
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/norm"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaleitem"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scaletranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/test"
	"github.com/google/uuid"
)

// ScaleQuery is the builder for querying Scale entities.
type ScaleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Scale
	// eager-loading edges.
	withItems           *ItemQuery
	withInterpretations *InterpretationQuery
	withTranslations    *ScaleTranslationQuery
	withNorms           *NormQuery
	withTest            *TestQuery
	withScaleItem       *ScaleItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScaleQuery builder.
func (sq *ScaleQuery) Where(ps ...predicate.Scale) *ScaleQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ScaleQuery) Limit(limit int) *ScaleQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ScaleQuery) Offset(offset int) *ScaleQuery {
	sq.offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *ScaleQuery) Unique(unique bool) *ScaleQuery {
	sq.unique = &unique
	return sq
}

// Order adds an order step to the query.
func (sq *ScaleQuery) Order(o ...OrderFunc) *ScaleQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryItems chains the current query on the "items" edge.
func (sq *ScaleQuery) QueryItems() *ItemQuery {
	query := &ItemQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, selector),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, scale.ItemsTable, scale.ItemsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInterpretations chains the current query on the "interpretations" edge.
func (sq *ScaleQuery) QueryInterpretations() *InterpretationQuery {
	query := &InterpretationQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, selector),
			sqlgraph.To(interpretation.Table, interpretation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scale.InterpretationsTable, scale.InterpretationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTranslations chains the current query on the "translations" edge.
func (sq *ScaleQuery) QueryTranslations() *ScaleTranslationQuery {
	query := &ScaleTranslationQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, selector),
			sqlgraph.To(scaletranslation.Table, scaletranslation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scale.TranslationsTable, scale.TranslationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNorms chains the current query on the "norms" edge.
func (sq *ScaleQuery) QueryNorms() *NormQuery {
	query := &NormQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, selector),
			sqlgraph.To(norm.Table, norm.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, scale.NormsTable, scale.NormsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTest chains the current query on the "test" edge.
func (sq *ScaleQuery) QueryTest() *TestQuery {
	query := &TestQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, selector),
			sqlgraph.To(test.Table, test.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, scale.TestTable, scale.TestPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScaleItem chains the current query on the "scale_item" edge.
func (sq *ScaleQuery) QueryScaleItem() *ScaleItemQuery {
	query := &ScaleItemQuery{config: sq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scale.Table, scale.FieldID, selector),
			sqlgraph.To(scaleitem.Table, scaleitem.ScaleColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, scale.ScaleItemTable, scale.ScaleItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Scale entity from the query.
// Returns a *NotFoundError when no Scale was found.
func (sq *ScaleQuery) First(ctx context.Context) (*Scale, error) {
	nodes, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{scale.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ScaleQuery) FirstX(ctx context.Context) *Scale {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Scale ID from the query.
// Returns a *NotFoundError when no Scale ID was found.
func (sq *ScaleQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scale.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *ScaleQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Scale entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Scale entity is found.
// Returns a *NotFoundError when no Scale entities are found.
func (sq *ScaleQuery) Only(ctx context.Context) (*Scale, error) {
	nodes, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{scale.Label}
	default:
		return nil, &NotSingularError{scale.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ScaleQuery) OnlyX(ctx context.Context) *Scale {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Scale ID in the query.
// Returns a *NotSingularError when more than one Scale ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *ScaleQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scale.Label}
	default:
		err = &NotSingularError{scale.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *ScaleQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Scales.
func (sq *ScaleQuery) All(ctx context.Context) ([]*Scale, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ScaleQuery) AllX(ctx context.Context) []*Scale {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Scale IDs.
func (sq *ScaleQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := sq.Select(scale.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ScaleQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ScaleQuery) Count(ctx context.Context) (int, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ScaleQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ScaleQuery) Exist(ctx context.Context) (bool, error) {
	if err := sq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ScaleQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScaleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ScaleQuery) Clone() *ScaleQuery {
	if sq == nil {
		return nil
	}
	return &ScaleQuery{
		config:              sq.config,
		limit:               sq.limit,
		offset:              sq.offset,
		order:               append([]OrderFunc{}, sq.order...),
		predicates:          append([]predicate.Scale{}, sq.predicates...),
		withItems:           sq.withItems.Clone(),
		withInterpretations: sq.withInterpretations.Clone(),
		withTranslations:    sq.withTranslations.Clone(),
		withNorms:           sq.withNorms.Clone(),
		withTest:            sq.withTest.Clone(),
		withScaleItem:       sq.withScaleItem.Clone(),
		// clone intermediate query.
		sql:    sq.sql.Clone(),
		path:   sq.path,
		unique: sq.unique,
	}
}

// WithItems tells the query-builder to eager-load the nodes that are connected to
// the "items" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScaleQuery) WithItems(opts ...func(*ItemQuery)) *ScaleQuery {
	query := &ItemQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withItems = query
	return sq
}

// WithInterpretations tells the query-builder to eager-load the nodes that are connected to
// the "interpretations" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScaleQuery) WithInterpretations(opts ...func(*InterpretationQuery)) *ScaleQuery {
	query := &InterpretationQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withInterpretations = query
	return sq
}

// WithTranslations tells the query-builder to eager-load the nodes that are connected to
// the "translations" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScaleQuery) WithTranslations(opts ...func(*ScaleTranslationQuery)) *ScaleQuery {
	query := &ScaleTranslationQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withTranslations = query
	return sq
}

// WithNorms tells the query-builder to eager-load the nodes that are connected to
// the "norms" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScaleQuery) WithNorms(opts ...func(*NormQuery)) *ScaleQuery {
	query := &NormQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withNorms = query
	return sq
}

// WithTest tells the query-builder to eager-load the nodes that are connected to
// the "test" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScaleQuery) WithTest(opts ...func(*TestQuery)) *ScaleQuery {
	query := &TestQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withTest = query
	return sq
}

// WithScaleItem tells the query-builder to eager-load the nodes that are connected to
// the "scale_item" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *ScaleQuery) WithScaleItem(opts ...func(*ScaleItemQuery)) *ScaleQuery {
	query := &ScaleItemQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withScaleItem = query
	return sq
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
//	client.Scale.Query().
//		GroupBy(scale.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *ScaleQuery) GroupBy(field string, fields ...string) *ScaleGroupBy {
	grbuild := &ScaleGroupBy{config: sq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sq.sqlQuery(ctx), nil
	}
	grbuild.label = scale.Label
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
//	client.Scale.Query().
//		Select(scale.FieldCreateTime).
//		Scan(ctx, &v)
//
func (sq *ScaleQuery) Select(fields ...string) *ScaleSelect {
	sq.fields = append(sq.fields, fields...)
	selbuild := &ScaleSelect{ScaleQuery: sq}
	selbuild.label = scale.Label
	selbuild.flds, selbuild.scan = &sq.fields, selbuild.Scan
	return selbuild
}

func (sq *ScaleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sq.fields {
		if !scale.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *ScaleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Scale, error) {
	var (
		nodes       = []*Scale{}
		_spec       = sq.querySpec()
		loadedTypes = [6]bool{
			sq.withItems != nil,
			sq.withInterpretations != nil,
			sq.withTranslations != nil,
			sq.withNorms != nil,
			sq.withTest != nil,
			sq.withScaleItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Scale).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Scale{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withItems; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[uuid.UUID]*Scale)
		nids := make(map[uuid.UUID]map[*Scale]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Items = []*Item{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(scale.ItemsTable)
			s.Join(joinT).On(s.C(item.FieldID), joinT.C(scale.ItemsPrimaryKey[1]))
			s.Where(sql.InValues(joinT.C(scale.ItemsPrimaryKey[0]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(scale.ItemsPrimaryKey[0]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Scale]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "items" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Items = append(kn.Edges.Items, n)
			}
		}
	}

	if query := sq.withInterpretations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Scale)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Interpretations = []*Interpretation{}
		}
		query.withFKs = true
		query.Where(predicate.Interpretation(func(s *sql.Selector) {
			s.Where(sql.InValues(scale.InterpretationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.scale_interpretations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "scale_interpretations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_interpretations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Interpretations = append(node.Edges.Interpretations, n)
		}
	}

	if query := sq.withTranslations; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Scale)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Translations = []*ScaleTranslation{}
		}
		query.withFKs = true
		query.Where(predicate.ScaleTranslation(func(s *sql.Selector) {
			s.Where(sql.InValues(scale.TranslationsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.scale_translations
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "scale_translations" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_translations" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Translations = append(node.Edges.Translations, n)
		}
	}

	if query := sq.withNorms; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Scale)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Norms = []*Norm{}
		}
		query.withFKs = true
		query.Where(predicate.Norm(func(s *sql.Selector) {
			s.Where(sql.InValues(scale.NormsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.scale_norms
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "scale_norms" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_norms" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Norms = append(node.Edges.Norms, n)
		}
	}

	if query := sq.withTest; query != nil {
		edgeids := make([]driver.Value, len(nodes))
		byid := make(map[uuid.UUID]*Scale)
		nids := make(map[uuid.UUID]map[*Scale]struct{})
		for i, node := range nodes {
			edgeids[i] = node.ID
			byid[node.ID] = node
			node.Edges.Test = []*Test{}
		}
		query.Where(func(s *sql.Selector) {
			joinT := sql.Table(scale.TestTable)
			s.Join(joinT).On(s.C(test.FieldID), joinT.C(scale.TestPrimaryKey[0]))
			s.Where(sql.InValues(joinT.C(scale.TestPrimaryKey[1]), edgeids...))
			columns := s.SelectedColumns()
			s.Select(joinT.C(scale.TestPrimaryKey[1]))
			s.AppendSelect(columns...)
			s.SetDistinct(false)
		})
		neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]interface{}, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]interface{}{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []interface{}) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Scale]struct{}{byid[outValue]: struct{}{}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byid[outValue]] = struct{}{}
				return nil
			}
		})
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "test" node returned %v`, n.ID)
			}
			for kn := range nodes {
				kn.Edges.Test = append(kn.Edges.Test, n)
			}
		}
	}

	if query := sq.withScaleItem; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Scale)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.ScaleItem = []*ScaleItem{}
		}
		query.Where(predicate.ScaleItem(func(s *sql.Selector) {
			s.Where(sql.InValues(scale.ScaleItemColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ScaleID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scale_id" returned %v for node %v`, fk, n)
			}
			node.Edges.ScaleItem = append(node.Edges.ScaleItem, n)
		}
	}

	return nodes, nil
}

func (sq *ScaleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.fields
	if len(sq.fields) > 0 {
		_spec.Unique = sq.unique != nil && *sq.unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ScaleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sq *ScaleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scale.Table,
			Columns: scale.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: scale.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scale.FieldID)
		for i := range fields {
			if fields[i] != scale.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ScaleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(scale.Table)
	columns := sq.fields
	if len(columns) == 0 {
		columns = scale.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.unique != nil && *sq.unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScaleGroupBy is the group-by builder for Scale entities.
type ScaleGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ScaleGroupBy) Aggregate(fns ...AggregateFunc) *ScaleGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scans the result into the given value.
func (sgb *ScaleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sgb.path(ctx)
	if err != nil {
		return err
	}
	sgb.sql = query
	return sgb.sqlScan(ctx, v)
}

func (sgb *ScaleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sgb.fields {
		if !scale.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ScaleGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql.Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
		for _, f := range sgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(sgb.fields...)...)
}

// ScaleSelect is the builder for selecting fields of Scale entities.
type ScaleSelect struct {
	*ScaleQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ss *ScaleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	ss.sql = ss.ScaleQuery.sqlQuery(ctx)
	return ss.sqlScan(ctx, v)
}

func (ss *ScaleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sql.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
