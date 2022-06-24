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
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/norm"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/sample"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/scale"
	"github.com/google/uuid"
)

// NormUpdate is the builder for updating Norm entities.
type NormUpdate struct {
	config
	hooks    []Hook
	mutation *NormMutation
}

// Where appends a list predicates to the NormUpdate builder.
func (nu *NormUpdate) Where(ps ...predicate.Norm) *NormUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdateTime sets the "update_time" field.
func (nu *NormUpdate) SetUpdateTime(t time.Time) *NormUpdate {
	nu.mutation.SetUpdateTime(t)
	return nu
}

// SetName sets the "name" field.
func (nu *NormUpdate) SetName(s string) *NormUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetBase sets the "base" field.
func (nu *NormUpdate) SetBase(i int) *NormUpdate {
	nu.mutation.ResetBase()
	nu.mutation.SetBase(i)
	return nu
}

// SetNillableBase sets the "base" field if the given value is not nil.
func (nu *NormUpdate) SetNillableBase(i *int) *NormUpdate {
	if i != nil {
		nu.SetBase(*i)
	}
	return nu
}

// AddBase adds i to the "base" field.
func (nu *NormUpdate) AddBase(i int) *NormUpdate {
	nu.mutation.AddBase(i)
	return nu
}

// SetMean sets the "mean" field.
func (nu *NormUpdate) SetMean(f float64) *NormUpdate {
	nu.mutation.ResetMean()
	nu.mutation.SetMean(f)
	return nu
}

// AddMean adds f to the "mean" field.
func (nu *NormUpdate) AddMean(f float64) *NormUpdate {
	nu.mutation.AddMean(f)
	return nu
}

// SetSigma sets the "sigma" field.
func (nu *NormUpdate) SetSigma(f float64) *NormUpdate {
	nu.mutation.ResetSigma()
	nu.mutation.SetSigma(f)
	return nu
}

// AddSigma adds f to the "sigma" field.
func (nu *NormUpdate) AddSigma(f float64) *NormUpdate {
	nu.mutation.AddSigma(f)
	return nu
}

// SetRank sets the "rank" field.
func (nu *NormUpdate) SetRank(i int) *NormUpdate {
	nu.mutation.ResetRank()
	nu.mutation.SetRank(i)
	return nu
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (nu *NormUpdate) SetNillableRank(i *int) *NormUpdate {
	if i != nil {
		nu.SetRank(*i)
	}
	return nu
}

// AddRank adds i to the "rank" field.
func (nu *NormUpdate) AddRank(i int) *NormUpdate {
	nu.mutation.AddRank(i)
	return nu
}

// SetMeta sets the "meta" field.
func (nu *NormUpdate) SetMeta(m map[string]interface{}) *NormUpdate {
	nu.mutation.SetMeta(m)
	return nu
}

// ClearMeta clears the value of the "meta" field.
func (nu *NormUpdate) ClearMeta() *NormUpdate {
	nu.mutation.ClearMeta()
	return nu
}

// SetSampleID sets the "sample" edge to the Sample entity by ID.
func (nu *NormUpdate) SetSampleID(id uuid.UUID) *NormUpdate {
	nu.mutation.SetSampleID(id)
	return nu
}

// SetSample sets the "sample" edge to the Sample entity.
func (nu *NormUpdate) SetSample(s *Sample) *NormUpdate {
	return nu.SetSampleID(s.ID)
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (nu *NormUpdate) SetScaleID(id uuid.UUID) *NormUpdate {
	nu.mutation.SetScaleID(id)
	return nu
}

// SetScale sets the "scale" edge to the Scale entity.
func (nu *NormUpdate) SetScale(s *Scale) *NormUpdate {
	return nu.SetScaleID(s.ID)
}

// Mutation returns the NormMutation object of the builder.
func (nu *NormUpdate) Mutation() *NormMutation {
	return nu.mutation
}

// ClearSample clears the "sample" edge to the Sample entity.
func (nu *NormUpdate) ClearSample() *NormUpdate {
	nu.mutation.ClearSample()
	return nu
}

// ClearScale clears the "scale" edge to the Scale entity.
func (nu *NormUpdate) ClearScale() *NormUpdate {
	nu.mutation.ClearScale()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NormUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	nu.defaults()
	if len(nu.hooks) == 0 {
		if err = nu.check(); err != nil {
			return 0, err
		}
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NormMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nu.check(); err != nil {
				return 0, err
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NormUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NormUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NormUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NormUpdate) defaults() {
	if _, ok := nu.mutation.UpdateTime(); !ok {
		v := norm.UpdateDefaultUpdateTime()
		nu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nu *NormUpdate) check() error {
	if v, ok := nu.mutation.Name(); ok {
		if err := norm.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Norm.name": %w`, err)}
		}
	}
	if v, ok := nu.mutation.Base(); ok {
		if err := norm.BaseValidator(v); err != nil {
			return &ValidationError{Name: "base", err: fmt.Errorf(`ent: validator failed for field "Norm.base": %w`, err)}
		}
	}
	if _, ok := nu.mutation.SampleID(); nu.mutation.SampleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Norm.sample"`)
	}
	if _, ok := nu.mutation.ScaleID(); nu.mutation.ScaleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Norm.scale"`)
	}
	return nil
}

func (nu *NormUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   norm.Table,
			Columns: norm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: norm.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: norm.FieldUpdateTime,
		})
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: norm.FieldName,
		})
	}
	if value, ok := nu.mutation.Base(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldBase,
		})
	}
	if value, ok := nu.mutation.AddedBase(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldBase,
		})
	}
	if value, ok := nu.mutation.Mean(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldMean,
		})
	}
	if value, ok := nu.mutation.AddedMean(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldMean,
		})
	}
	if value, ok := nu.mutation.Sigma(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldSigma,
		})
	}
	if value, ok := nu.mutation.AddedSigma(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldSigma,
		})
	}
	if value, ok := nu.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldRank,
		})
	}
	if value, ok := nu.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldRank,
		})
	}
	if value, ok := nu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: norm.FieldMeta,
		})
	}
	if nu.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: norm.FieldMeta,
		})
	}
	if nu.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.SampleTable,
			Columns: []string{norm.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: sample.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.SampleTable,
			Columns: []string{norm.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: sample.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.ScaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.ScaleTable,
			Columns: []string{norm.ScaleColumn},
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
	if nodes := nu.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.ScaleTable,
			Columns: []string{norm.ScaleColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{norm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// NormUpdateOne is the builder for updating a single Norm entity.
type NormUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NormMutation
}

// SetUpdateTime sets the "update_time" field.
func (nuo *NormUpdateOne) SetUpdateTime(t time.Time) *NormUpdateOne {
	nuo.mutation.SetUpdateTime(t)
	return nuo
}

// SetName sets the "name" field.
func (nuo *NormUpdateOne) SetName(s string) *NormUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetBase sets the "base" field.
func (nuo *NormUpdateOne) SetBase(i int) *NormUpdateOne {
	nuo.mutation.ResetBase()
	nuo.mutation.SetBase(i)
	return nuo
}

// SetNillableBase sets the "base" field if the given value is not nil.
func (nuo *NormUpdateOne) SetNillableBase(i *int) *NormUpdateOne {
	if i != nil {
		nuo.SetBase(*i)
	}
	return nuo
}

// AddBase adds i to the "base" field.
func (nuo *NormUpdateOne) AddBase(i int) *NormUpdateOne {
	nuo.mutation.AddBase(i)
	return nuo
}

// SetMean sets the "mean" field.
func (nuo *NormUpdateOne) SetMean(f float64) *NormUpdateOne {
	nuo.mutation.ResetMean()
	nuo.mutation.SetMean(f)
	return nuo
}

// AddMean adds f to the "mean" field.
func (nuo *NormUpdateOne) AddMean(f float64) *NormUpdateOne {
	nuo.mutation.AddMean(f)
	return nuo
}

// SetSigma sets the "sigma" field.
func (nuo *NormUpdateOne) SetSigma(f float64) *NormUpdateOne {
	nuo.mutation.ResetSigma()
	nuo.mutation.SetSigma(f)
	return nuo
}

// AddSigma adds f to the "sigma" field.
func (nuo *NormUpdateOne) AddSigma(f float64) *NormUpdateOne {
	nuo.mutation.AddSigma(f)
	return nuo
}

// SetRank sets the "rank" field.
func (nuo *NormUpdateOne) SetRank(i int) *NormUpdateOne {
	nuo.mutation.ResetRank()
	nuo.mutation.SetRank(i)
	return nuo
}

// SetNillableRank sets the "rank" field if the given value is not nil.
func (nuo *NormUpdateOne) SetNillableRank(i *int) *NormUpdateOne {
	if i != nil {
		nuo.SetRank(*i)
	}
	return nuo
}

// AddRank adds i to the "rank" field.
func (nuo *NormUpdateOne) AddRank(i int) *NormUpdateOne {
	nuo.mutation.AddRank(i)
	return nuo
}

// SetMeta sets the "meta" field.
func (nuo *NormUpdateOne) SetMeta(m map[string]interface{}) *NormUpdateOne {
	nuo.mutation.SetMeta(m)
	return nuo
}

// ClearMeta clears the value of the "meta" field.
func (nuo *NormUpdateOne) ClearMeta() *NormUpdateOne {
	nuo.mutation.ClearMeta()
	return nuo
}

// SetSampleID sets the "sample" edge to the Sample entity by ID.
func (nuo *NormUpdateOne) SetSampleID(id uuid.UUID) *NormUpdateOne {
	nuo.mutation.SetSampleID(id)
	return nuo
}

// SetSample sets the "sample" edge to the Sample entity.
func (nuo *NormUpdateOne) SetSample(s *Sample) *NormUpdateOne {
	return nuo.SetSampleID(s.ID)
}

// SetScaleID sets the "scale" edge to the Scale entity by ID.
func (nuo *NormUpdateOne) SetScaleID(id uuid.UUID) *NormUpdateOne {
	nuo.mutation.SetScaleID(id)
	return nuo
}

// SetScale sets the "scale" edge to the Scale entity.
func (nuo *NormUpdateOne) SetScale(s *Scale) *NormUpdateOne {
	return nuo.SetScaleID(s.ID)
}

// Mutation returns the NormMutation object of the builder.
func (nuo *NormUpdateOne) Mutation() *NormMutation {
	return nuo.mutation
}

// ClearSample clears the "sample" edge to the Sample entity.
func (nuo *NormUpdateOne) ClearSample() *NormUpdateOne {
	nuo.mutation.ClearSample()
	return nuo
}

// ClearScale clears the "scale" edge to the Scale entity.
func (nuo *NormUpdateOne) ClearScale() *NormUpdateOne {
	nuo.mutation.ClearScale()
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NormUpdateOne) Select(field string, fields ...string) *NormUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Norm entity.
func (nuo *NormUpdateOne) Save(ctx context.Context) (*Norm, error) {
	var (
		err  error
		node *Norm
	)
	nuo.defaults()
	if len(nuo.hooks) == 0 {
		if err = nuo.check(); err != nil {
			return nil, err
		}
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NormMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nuo.check(); err != nil {
				return nil, err
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Norm)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NormMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NormUpdateOne) SaveX(ctx context.Context) *Norm {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NormUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NormUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NormUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdateTime(); !ok {
		v := norm.UpdateDefaultUpdateTime()
		nuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NormUpdateOne) check() error {
	if v, ok := nuo.mutation.Name(); ok {
		if err := norm.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Norm.name": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.Base(); ok {
		if err := norm.BaseValidator(v); err != nil {
			return &ValidationError{Name: "base", err: fmt.Errorf(`ent: validator failed for field "Norm.base": %w`, err)}
		}
	}
	if _, ok := nuo.mutation.SampleID(); nuo.mutation.SampleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Norm.sample"`)
	}
	if _, ok := nuo.mutation.ScaleID(); nuo.mutation.ScaleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Norm.scale"`)
	}
	return nil
}

func (nuo *NormUpdateOne) sqlSave(ctx context.Context) (_node *Norm, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   norm.Table,
			Columns: norm.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: norm.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Norm.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, norm.FieldID)
		for _, f := range fields {
			if !norm.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != norm.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: norm.FieldUpdateTime,
		})
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: norm.FieldName,
		})
	}
	if value, ok := nuo.mutation.Base(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldBase,
		})
	}
	if value, ok := nuo.mutation.AddedBase(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldBase,
		})
	}
	if value, ok := nuo.mutation.Mean(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldMean,
		})
	}
	if value, ok := nuo.mutation.AddedMean(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldMean,
		})
	}
	if value, ok := nuo.mutation.Sigma(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldSigma,
		})
	}
	if value, ok := nuo.mutation.AddedSigma(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: norm.FieldSigma,
		})
	}
	if value, ok := nuo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldRank,
		})
	}
	if value, ok := nuo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: norm.FieldRank,
		})
	}
	if value, ok := nuo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: norm.FieldMeta,
		})
	}
	if nuo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: norm.FieldMeta,
		})
	}
	if nuo.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.SampleTable,
			Columns: []string{norm.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: sample.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.SampleTable,
			Columns: []string{norm.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: sample.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.ScaleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.ScaleTable,
			Columns: []string{norm.ScaleColumn},
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
	if nodes := nuo.mutation.ScaleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   norm.ScaleTable,
			Columns: []string{norm.ScaleColumn},
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
	_node = &Norm{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{norm.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}