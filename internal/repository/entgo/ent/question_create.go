// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/questiontranslation"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/test"
	"github.com/google/uuid"
)

// QuestionCreate is the builder for creating a Question entity.
type QuestionCreate struct {
	config
	mutation *QuestionMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (qc *QuestionCreate) SetCreateTime(t time.Time) *QuestionCreate {
	qc.mutation.SetCreateTime(t)
	return qc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableCreateTime(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetCreateTime(*t)
	}
	return qc
}

// SetUpdateTime sets the "update_time" field.
func (qc *QuestionCreate) SetUpdateTime(t time.Time) *QuestionCreate {
	qc.mutation.SetUpdateTime(t)
	return qc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableUpdateTime(t *time.Time) *QuestionCreate {
	if t != nil {
		qc.SetUpdateTime(*t)
	}
	return qc
}

// SetOrder sets the "order" field.
func (qc *QuestionCreate) SetOrder(i int) *QuestionCreate {
	qc.mutation.SetOrder(i)
	return qc
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableOrder(i *int) *QuestionCreate {
	if i != nil {
		qc.SetOrder(*i)
	}
	return qc
}

// SetCode sets the "code" field.
func (qc *QuestionCreate) SetCode(s string) *QuestionCreate {
	qc.mutation.SetCode(s)
	return qc
}

// SetType sets the "type" field.
func (qc *QuestionCreate) SetType(q question.Type) *QuestionCreate {
	qc.mutation.SetType(q)
	return qc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableType(q *question.Type) *QuestionCreate {
	if q != nil {
		qc.SetType(*q)
	}
	return qc
}

// SetID sets the "id" field.
func (qc *QuestionCreate) SetID(u uuid.UUID) *QuestionCreate {
	qc.mutation.SetID(u)
	return qc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (qc *QuestionCreate) SetNillableID(u *uuid.UUID) *QuestionCreate {
	if u != nil {
		qc.SetID(*u)
	}
	return qc
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (qc *QuestionCreate) AddItemIDs(ids ...uuid.UUID) *QuestionCreate {
	qc.mutation.AddItemIDs(ids...)
	return qc
}

// AddItems adds the "items" edges to the Item entity.
func (qc *QuestionCreate) AddItems(i ...*Item) *QuestionCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return qc.AddItemIDs(ids...)
}

// AddTranslationIDs adds the "translations" edge to the QuestionTranslation entity by IDs.
func (qc *QuestionCreate) AddTranslationIDs(ids ...uuid.UUID) *QuestionCreate {
	qc.mutation.AddTranslationIDs(ids...)
	return qc
}

// AddTranslations adds the "translations" edges to the QuestionTranslation entity.
func (qc *QuestionCreate) AddTranslations(q ...*QuestionTranslation) *QuestionCreate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qc.AddTranslationIDs(ids...)
}

// AddTestIDs adds the "test" edge to the Test entity by IDs.
func (qc *QuestionCreate) AddTestIDs(ids ...uuid.UUID) *QuestionCreate {
	qc.mutation.AddTestIDs(ids...)
	return qc
}

// AddTest adds the "test" edges to the Test entity.
func (qc *QuestionCreate) AddTest(t ...*Test) *QuestionCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return qc.AddTestIDs(ids...)
}

// Mutation returns the QuestionMutation object of the builder.
func (qc *QuestionCreate) Mutation() *QuestionMutation {
	return qc.mutation
}

// Save creates the Question in the database.
func (qc *QuestionCreate) Save(ctx context.Context) (*Question, error) {
	var (
		err  error
		node *Question
	)
	qc.defaults()
	if len(qc.hooks) == 0 {
		if err = qc.check(); err != nil {
			return nil, err
		}
		node, err = qc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qc.check(); err != nil {
				return nil, err
			}
			qc.mutation = mutation
			if node, err = qc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(qc.hooks) - 1; i >= 0; i-- {
			if qc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, qc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Question)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from QuestionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QuestionCreate) SaveX(ctx context.Context) *Question {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qc *QuestionCreate) Exec(ctx context.Context) error {
	_, err := qc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qc *QuestionCreate) ExecX(ctx context.Context) {
	if err := qc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qc *QuestionCreate) defaults() {
	if _, ok := qc.mutation.CreateTime(); !ok {
		v := question.DefaultCreateTime()
		qc.mutation.SetCreateTime(v)
	}
	if _, ok := qc.mutation.UpdateTime(); !ok {
		v := question.DefaultUpdateTime()
		qc.mutation.SetUpdateTime(v)
	}
	if _, ok := qc.mutation.Order(); !ok {
		v := question.DefaultOrder
		qc.mutation.SetOrder(v)
	}
	if _, ok := qc.mutation.GetType(); !ok {
		v := question.DefaultType
		qc.mutation.SetType(v)
	}
	if _, ok := qc.mutation.ID(); !ok {
		v := question.DefaultID()
		qc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qc *QuestionCreate) check() error {
	if _, ok := qc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Question.create_time"`)}
	}
	if _, ok := qc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Question.update_time"`)}
	}
	if _, ok := qc.mutation.Order(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required field "Question.order"`)}
	}
	if _, ok := qc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Question.code"`)}
	}
	if v, ok := qc.mutation.Code(); ok {
		if err := question.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Question.code": %w`, err)}
		}
	}
	if _, ok := qc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Question.type"`)}
	}
	if v, ok := qc.mutation.GetType(); ok {
		if err := question.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Question.type": %w`, err)}
		}
	}
	return nil
}

func (qc *QuestionCreate) sqlSave(ctx context.Context) (*Question, error) {
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
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

func (qc *QuestionCreate) createSpec() (*Question, *sqlgraph.CreateSpec) {
	var (
		_node = &Question{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: question.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: question.FieldID,
			},
		}
	)
	if id, ok := qc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := qc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := qc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: question.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := qc.mutation.Order(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: question.FieldOrder,
		})
		_node.Order = value
	}
	if value, ok := qc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: question.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := qc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: question.FieldType,
		})
		_node.Type = value
	}
	if nodes := qc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   question.ItemsTable,
			Columns: question.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   question.TranslationsTable,
			Columns: []string{question.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: questiontranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.TestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   question.TestTable,
			Columns: question.TestPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: test.FieldID,
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

// QuestionCreateBulk is the builder for creating many Question entities in bulk.
type QuestionCreateBulk struct {
	config
	builders []*QuestionCreate
}

// Save creates the Question entities in the database.
func (qcb *QuestionCreateBulk) Save(ctx context.Context) ([]*Question, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Question, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuestionMutation)
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
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QuestionCreateBulk) SaveX(ctx context.Context) []*Question {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qcb *QuestionCreateBulk) Exec(ctx context.Context) error {
	_, err := qcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qcb *QuestionCreateBulk) ExecX(ctx context.Context) {
	if err := qcb.Exec(ctx); err != nil {
		panic(err)
	}
}
