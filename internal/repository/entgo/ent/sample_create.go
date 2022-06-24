// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/domain"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/norm"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/sample"
	"github.com/google/uuid"
)

// SampleCreate is the builder for creating a Sample entity.
type SampleCreate struct {
	config
	mutation *SampleMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *SampleCreate) SetCreateTime(t time.Time) *SampleCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SampleCreate) SetNillableCreateTime(t *time.Time) *SampleCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SampleCreate) SetUpdateTime(t time.Time) *SampleCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SampleCreate) SetNillableUpdateTime(t *time.Time) *SampleCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetCode sets the "code" field.
func (sc *SampleCreate) SetCode(s string) *SampleCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetCriteria sets the "criteria" field.
func (sc *SampleCreate) SetCriteria(dc domain.SampleCriteria) *SampleCreate {
	sc.mutation.SetCriteria(dc)
	return sc
}

// SetID sets the "id" field.
func (sc *SampleCreate) SetID(u uuid.UUID) *SampleCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SampleCreate) SetNillableID(u *uuid.UUID) *SampleCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddNormIDs adds the "norms" edge to the Norm entity by IDs.
func (sc *SampleCreate) AddNormIDs(ids ...uuid.UUID) *SampleCreate {
	sc.mutation.AddNormIDs(ids...)
	return sc
}

// AddNorms adds the "norms" edges to the Norm entity.
func (sc *SampleCreate) AddNorms(n ...*Norm) *SampleCreate {
	ids := make([]uuid.UUID, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return sc.AddNormIDs(ids...)
}

// Mutation returns the SampleMutation object of the builder.
func (sc *SampleCreate) Mutation() *SampleMutation {
	return sc.mutation
}

// Save creates the Sample in the database.
func (sc *SampleCreate) Save(ctx context.Context) (*Sample, error) {
	var (
		err  error
		node *Sample
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SampleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Sample)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SampleMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SampleCreate) SaveX(ctx context.Context) *Sample {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SampleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SampleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SampleCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := sample.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := sample.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := sample.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SampleCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Sample.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Sample.update_time"`)}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Sample.code"`)}
	}
	if v, ok := sc.mutation.Code(); ok {
		if err := sample.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Sample.code": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Criteria(); !ok {
		return &ValidationError{Name: "criteria", err: errors.New(`ent: missing required field "Sample.criteria"`)}
	}
	return nil
}

func (sc *SampleCreate) sqlSave(ctx context.Context) (*Sample, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (sc *SampleCreate) createSpec() (*Sample, *sqlgraph.CreateSpec) {
	var (
		_node = &Sample{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: sample.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: sample.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sample.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sample.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sample.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := sc.mutation.Criteria(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: sample.FieldCriteria,
		})
		_node.Criteria = value
	}
	if nodes := sc.mutation.NormsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sample.NormsTable,
			Columns: []string{sample.NormsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: norm.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SampleCreateBulk is the builder for creating many Sample entities in bulk.
type SampleCreateBulk struct {
	config
	builders []*SampleCreate
}

// Save creates the Sample entities in the database.
func (scb *SampleCreateBulk) Save(ctx context.Context) ([]*Sample, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sample, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SampleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SampleCreateBulk) SaveX(ctx context.Context) []*Sample {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SampleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SampleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}