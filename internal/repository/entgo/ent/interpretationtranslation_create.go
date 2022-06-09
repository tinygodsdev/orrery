// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretation"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/interpretationtranslation"
	"github.com/google/uuid"
)

// InterpretationTranslationCreate is the builder for creating a InterpretationTranslation entity.
type InterpretationTranslationCreate struct {
	config
	mutation *InterpretationTranslationMutation
	hooks    []Hook
}

// SetLocale sets the "locale" field.
func (itc *InterpretationTranslationCreate) SetLocale(i interpretationtranslation.Locale) *InterpretationTranslationCreate {
	itc.mutation.SetLocale(i)
	return itc
}

// SetContent sets the "content" field.
func (itc *InterpretationTranslationCreate) SetContent(s string) *InterpretationTranslationCreate {
	itc.mutation.SetContent(s)
	return itc
}

// SetID sets the "id" field.
func (itc *InterpretationTranslationCreate) SetID(u uuid.UUID) *InterpretationTranslationCreate {
	itc.mutation.SetID(u)
	return itc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (itc *InterpretationTranslationCreate) SetNillableID(u *uuid.UUID) *InterpretationTranslationCreate {
	if u != nil {
		itc.SetID(*u)
	}
	return itc
}

// SetInterpretationID sets the "interpretation" edge to the Interpretation entity by ID.
func (itc *InterpretationTranslationCreate) SetInterpretationID(id uuid.UUID) *InterpretationTranslationCreate {
	itc.mutation.SetInterpretationID(id)
	return itc
}

// SetNillableInterpretationID sets the "interpretation" edge to the Interpretation entity by ID if the given value is not nil.
func (itc *InterpretationTranslationCreate) SetNillableInterpretationID(id *uuid.UUID) *InterpretationTranslationCreate {
	if id != nil {
		itc = itc.SetInterpretationID(*id)
	}
	return itc
}

// SetInterpretation sets the "interpretation" edge to the Interpretation entity.
func (itc *InterpretationTranslationCreate) SetInterpretation(i *Interpretation) *InterpretationTranslationCreate {
	return itc.SetInterpretationID(i.ID)
}

// Mutation returns the InterpretationTranslationMutation object of the builder.
func (itc *InterpretationTranslationCreate) Mutation() *InterpretationTranslationMutation {
	return itc.mutation
}

// Save creates the InterpretationTranslation in the database.
func (itc *InterpretationTranslationCreate) Save(ctx context.Context) (*InterpretationTranslation, error) {
	var (
		err  error
		node *InterpretationTranslation
	)
	itc.defaults()
	if len(itc.hooks) == 0 {
		if err = itc.check(); err != nil {
			return nil, err
		}
		node, err = itc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InterpretationTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = itc.check(); err != nil {
				return nil, err
			}
			itc.mutation = mutation
			if node, err = itc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(itc.hooks) - 1; i >= 0; i-- {
			if itc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = itc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, itc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*InterpretationTranslation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InterpretationTranslationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (itc *InterpretationTranslationCreate) SaveX(ctx context.Context) *InterpretationTranslation {
	v, err := itc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itc *InterpretationTranslationCreate) Exec(ctx context.Context) error {
	_, err := itc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itc *InterpretationTranslationCreate) ExecX(ctx context.Context) {
	if err := itc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (itc *InterpretationTranslationCreate) defaults() {
	if _, ok := itc.mutation.ID(); !ok {
		v := interpretationtranslation.DefaultID()
		itc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itc *InterpretationTranslationCreate) check() error {
	if _, ok := itc.mutation.Locale(); !ok {
		return &ValidationError{Name: "locale", err: errors.New(`ent: missing required field "InterpretationTranslation.locale"`)}
	}
	if v, ok := itc.mutation.Locale(); ok {
		if err := interpretationtranslation.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "InterpretationTranslation.locale": %w`, err)}
		}
	}
	if _, ok := itc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "InterpretationTranslation.content"`)}
	}
	if v, ok := itc.mutation.Content(); ok {
		if err := interpretationtranslation.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "InterpretationTranslation.content": %w`, err)}
		}
	}
	return nil
}

func (itc *InterpretationTranslationCreate) sqlSave(ctx context.Context) (*InterpretationTranslation, error) {
	_node, _spec := itc.createSpec()
	if err := sqlgraph.CreateNode(ctx, itc.driver, _spec); err != nil {
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

func (itc *InterpretationTranslationCreate) createSpec() (*InterpretationTranslation, *sqlgraph.CreateSpec) {
	var (
		_node = &InterpretationTranslation{config: itc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: interpretationtranslation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: interpretationtranslation.FieldID,
			},
		}
	)
	if id, ok := itc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := itc.mutation.Locale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: interpretationtranslation.FieldLocale,
		})
		_node.Locale = value
	}
	if value, ok := itc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: interpretationtranslation.FieldContent,
		})
		_node.Content = value
	}
	if nodes := itc.mutation.InterpretationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   interpretationtranslation.InterpretationTable,
			Columns: []string{interpretationtranslation.InterpretationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: interpretation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.interpretation_translations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InterpretationTranslationCreateBulk is the builder for creating many InterpretationTranslation entities in bulk.
type InterpretationTranslationCreateBulk struct {
	config
	builders []*InterpretationTranslationCreate
}

// Save creates the InterpretationTranslation entities in the database.
func (itcb *InterpretationTranslationCreateBulk) Save(ctx context.Context) ([]*InterpretationTranslation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(itcb.builders))
	nodes := make([]*InterpretationTranslation, len(itcb.builders))
	mutators := make([]Mutator, len(itcb.builders))
	for i := range itcb.builders {
		func(i int, root context.Context) {
			builder := itcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InterpretationTranslationMutation)
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
					_, err = mutators[i+1].Mutate(root, itcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, itcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, itcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (itcb *InterpretationTranslationCreateBulk) SaveX(ctx context.Context) []*InterpretationTranslation {
	v, err := itcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itcb *InterpretationTranslationCreateBulk) Exec(ctx context.Context) error {
	_, err := itcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itcb *InterpretationTranslationCreateBulk) ExecX(ctx context.Context) {
	if err := itcb.Exec(ctx); err != nil {
		panic(err)
	}
}
