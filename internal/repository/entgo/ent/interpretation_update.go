// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretationtranslation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/google/uuid"
)

// InterpretationUpdate is the builder for updating Interpretation entities.
type InterpretationUpdate struct {
	config
	hooks    []Hook
	mutation *InterpretationMutation
}

// Where appends a list predicates to the InterpretationUpdate builder.
func (iu *InterpretationUpdate) Where(ps ...predicate.Interpretation) *InterpretationUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *InterpretationUpdate) SetUpdateTime(t time.Time) *InterpretationUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetRange sets the "range" field.
func (iu *InterpretationUpdate) SetRange(f [2]float64) *InterpretationUpdate {
	iu.mutation.SetRange(f)
	return iu
}

// AddTranslationIDs adds the "translations" edge to the InterpretationTranslation entity by IDs.
func (iu *InterpretationUpdate) AddTranslationIDs(ids ...uuid.UUID) *InterpretationUpdate {
	iu.mutation.AddTranslationIDs(ids...)
	return iu
}

// AddTranslations adds the "translations" edges to the InterpretationTranslation entity.
func (iu *InterpretationUpdate) AddTranslations(i ...*InterpretationTranslation) *InterpretationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iu.AddTranslationIDs(ids...)
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (iu *InterpretationUpdate) SetScaleID(id uuid.UUID) *InterpretationUpdate {
	iu.mutation.SetScaleID(id)
	return iu
}

// SetNillableScaleID sets the "scale" edge to the Scale entity by ID if the given value is not nil.
func (iu *InterpretationUpdate) SetNillableScaleID(id *uuid.UUID) *InterpretationUpdate {
	if id != nil {
		iu = iu.SetScaleID(*id)
	}
	return iu
}

// SetScale sets the "scale" edge to the Scale entity.
func (iu *InterpretationUpdate) SetScale(s *Scale) *InterpretationUpdate {
	return iu.SetScaleID(s.ID)
}

// Mutation returns the InterpretationMutation object of the builder.
func (iu *InterpretationUpdate) Mutation() *InterpretationMutation {
	return iu.mutation
}

// ClearTranslations clears all "translations" edges to the InterpretationTranslation entity.
func (iu *InterpretationUpdate) ClearTranslations() *InterpretationUpdate {
	iu.mutation.ClearTranslations()
	return iu
}

// RemoveTranslationIDs removes the "translations" edge to InterpretationTranslation entities by IDs.
func (iu *InterpretationUpdate) RemoveTranslationIDs(ids ...uuid.UUID) *InterpretationUpdate {
	iu.mutation.RemoveTranslationIDs(ids...)
	return iu
}

// RemoveTranslations removes "translations" edges to InterpretationTranslation entities.
func (iu *InterpretationUpdate) RemoveTranslations(i ...*InterpretationTranslation) *InterpretationUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iu.RemoveTranslationIDs(ids...)
}

// ClearScale clears the "scale" edge to the Scale entity.
func (iu *InterpretationUpdate) ClearScale() *InterpretationUpdate {
	iu.mutation.ClearScale()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InterpretationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InterpretationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InterpretationUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InterpretationUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InterpretationUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *InterpretationUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := interpretation.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

func (iu *InterpretationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   interpretation.Table,
			Columns: interpretation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: interpretation.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: interpretation.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Range(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: interpretation.FieldRange,
		})
	}
	if iu.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   interpretation.TranslationsTable,
			Columns: []string{interpretation.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretationtranslation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedTranslationsIDs(); len(nodes) > 0 && !iu.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   interpretation.TranslationsTable,
			Columns: []string{interpretation.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretationtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   interpretation.TranslationsTable,
			Columns: []string{interpretation.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretationtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.ScaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   interpretation.ScaleTable,
			Columns: []string{interpretation.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   interpretation.ScaleTable,
			Columns: []string{interpretation.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{interpretation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// InterpretationUpdateOne is the builder for updating a single Interpretation entity.
type InterpretationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InterpretationMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *InterpretationUpdateOne) SetUpdateTime(t time.Time) *InterpretationUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetRange sets the "range" field.
func (iuo *InterpretationUpdateOne) SetRange(f [2]float64) *InterpretationUpdateOne {
	iuo.mutation.SetRange(f)
	return iuo
}

// AddTranslationIDs adds the "translations" edge to the InterpretationTranslation entity by IDs.
func (iuo *InterpretationUpdateOne) AddTranslationIDs(ids ...uuid.UUID) *InterpretationUpdateOne {
	iuo.mutation.AddTranslationIDs(ids...)
	return iuo
}

// AddTranslations adds the "translations" edges to the InterpretationTranslation entity.
func (iuo *InterpretationUpdateOne) AddTranslations(i ...*InterpretationTranslation) *InterpretationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iuo.AddTranslationIDs(ids...)
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (iuo *InterpretationUpdateOne) SetScaleID(id uuid.UUID) *InterpretationUpdateOne {
	iuo.mutation.SetScaleID(id)
	return iuo
}

// SetNillableScaleID sets the "scale" edge to the Scale entity by ID if the given value is not nil.
func (iuo *InterpretationUpdateOne) SetNillableScaleID(id *uuid.UUID) *InterpretationUpdateOne {
	if id != nil {
		iuo = iuo.SetScaleID(*id)
	}
	return iuo
}

// SetScale sets the "scale" edge to the Scale entity.
func (iuo *InterpretationUpdateOne) SetScale(s *Scale) *InterpretationUpdateOne {
	return iuo.SetScaleID(s.ID)
}

// Mutation returns the InterpretationMutation object of the builder.
func (iuo *InterpretationUpdateOne) Mutation() *InterpretationMutation {
	return iuo.mutation
}

// ClearTranslations clears all "translations" edges to the InterpretationTranslation entity.
func (iuo *InterpretationUpdateOne) ClearTranslations() *InterpretationUpdateOne {
	iuo.mutation.ClearTranslations()
	return iuo
}

// RemoveTranslationIDs removes the "translations" edge to InterpretationTranslation entities by IDs.
func (iuo *InterpretationUpdateOne) RemoveTranslationIDs(ids ...uuid.UUID) *InterpretationUpdateOne {
	iuo.mutation.RemoveTranslationIDs(ids...)
	return iuo
}

// RemoveTranslations removes "translations" edges to InterpretationTranslation entities.
func (iuo *InterpretationUpdateOne) RemoveTranslations(i ...*InterpretationTranslation) *InterpretationUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return iuo.RemoveTranslationIDs(ids...)
}

// ClearScale clears the "scale" edge to the Scale entity.
func (iuo *InterpretationUpdateOne) ClearScale() *InterpretationUpdateOne {
	iuo.mutation.ClearScale()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InterpretationUpdateOne) Select(field string, fields ...string) *InterpretationUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Interpretation entity.
func (iuo *InterpretationUpdateOne) Save(ctx context.Context) (*Interpretation, error) {
	var (
		err  error
		node *Interpretation
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InterpretationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Interpretation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InterpretationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InterpretationUpdateOne) SaveX(ctx context.Context) *Interpretation {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InterpretationUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InterpretationUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *InterpretationUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := interpretation.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

func (iuo *InterpretationUpdateOne) sqlSave(ctx context.Context) (_node *Interpretation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   interpretation.Table,
			Columns: interpretation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: interpretation.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Interpretation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, interpretation.FieldID)
		for _, f := range fields {
			if !interpretation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != interpretation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: interpretation.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Range(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: interpretation.FieldRange,
		})
	}
	if iuo.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   interpretation.TranslationsTable,
			Columns: []string{interpretation.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretationtranslation.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedTranslationsIDs(); len(nodes) > 0 && !iuo.mutation.TranslationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   interpretation.TranslationsTable,
			Columns: []string{interpretation.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretationtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   interpretation.TranslationsTable,
			Columns: []string{interpretation.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretationtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.ScaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   interpretation.ScaleTable,
			Columns: []string{interpretation.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   interpretation.ScaleTable,
			Columns: []string{interpretation.ScaleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: scale.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Interpretation{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{interpretation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
