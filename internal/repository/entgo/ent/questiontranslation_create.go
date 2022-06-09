// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/questiontranslation"
	"github.com/google/uuid"
)

// QuestionTranslationCreate is the builder for creating a QuestionTranslation entity.
type QuestionTranslationCreate struct {
	config
	mutation *QuestionTranslationMutation
	hooks    []Hook
}

// SetLocale sets the "locale" field.
func (qtc *QuestionTranslationCreate) SetLocale(q questiontranslation.Locale) *QuestionTranslationCreate {
	qtc.mutation.SetLocale(q)
	return qtc
}

// SetContent sets the "content" field.
func (qtc *QuestionTranslationCreate) SetContent(s string) *QuestionTranslationCreate {
	qtc.mutation.SetContent(s)
	return qtc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (qtc *QuestionTranslationCreate) SetNillableContent(s *string) *QuestionTranslationCreate {
	if s != nil {
		qtc.SetContent(*s)
	}
	return qtc
}

// SetHeaderContent sets the "header_content" field.
func (qtc *QuestionTranslationCreate) SetHeaderContent(s string) *QuestionTranslationCreate {
	qtc.mutation.SetHeaderContent(s)
	return qtc
}

// SetNillableHeaderContent sets the "header_content" field if the given value is not nil.
func (qtc *QuestionTranslationCreate) SetNillableHeaderContent(s *string) *QuestionTranslationCreate {
	if s != nil {
		qtc.SetHeaderContent(*s)
	}
	return qtc
}

// SetFooterContent sets the "footer_content" field.
func (qtc *QuestionTranslationCreate) SetFooterContent(s string) *QuestionTranslationCreate {
	qtc.mutation.SetFooterContent(s)
	return qtc
}

// SetNillableFooterContent sets the "footer_content" field if the given value is not nil.
func (qtc *QuestionTranslationCreate) SetNillableFooterContent(s *string) *QuestionTranslationCreate {
	if s != nil {
		qtc.SetFooterContent(*s)
	}
	return qtc
}

// SetID sets the "id" field.
func (qtc *QuestionTranslationCreate) SetID(u uuid.UUID) *QuestionTranslationCreate {
	qtc.mutation.SetID(u)
	return qtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (qtc *QuestionTranslationCreate) SetNillableID(u *uuid.UUID) *QuestionTranslationCreate {
	if u != nil {
		qtc.SetID(*u)
	}
	return qtc
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (qtc *QuestionTranslationCreate) SetQuestionID(id uuid.UUID) *QuestionTranslationCreate {
	qtc.mutation.SetQuestionID(id)
	return qtc
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (qtc *QuestionTranslationCreate) SetNillableQuestionID(id *uuid.UUID) *QuestionTranslationCreate {
	if id != nil {
		qtc = qtc.SetQuestionID(*id)
	}
	return qtc
}

// SetQuestion sets the "question" edge to the Question entity.
func (qtc *QuestionTranslationCreate) SetQuestion(q *Question) *QuestionTranslationCreate {
	return qtc.SetQuestionID(q.ID)
}

// Mutation returns the QuestionTranslationMutation object of the builder.
func (qtc *QuestionTranslationCreate) Mutation() *QuestionTranslationMutation {
	return qtc.mutation
}

// Save creates the QuestionTranslation in the database.
func (qtc *QuestionTranslationCreate) Save(ctx context.Context) (*QuestionTranslation, error) {
	var (
		err  error
		node *QuestionTranslation
	)
	qtc.defaults()
	if len(qtc.hooks) == 0 {
		if err = qtc.check(); err != nil {
			return nil, err
		}
		node, err = qtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionTranslationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qtc.check(); err != nil {
				return nil, err
			}
			qtc.mutation = mutation
			if node, err = qtc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(qtc.hooks) - 1; i >= 0; i-- {
			if qtc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = qtc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, qtc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*QuestionTranslation)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from QuestionTranslationMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qtc *QuestionTranslationCreate) SaveX(ctx context.Context) *QuestionTranslation {
	v, err := qtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qtc *QuestionTranslationCreate) Exec(ctx context.Context) error {
	_, err := qtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtc *QuestionTranslationCreate) ExecX(ctx context.Context) {
	if err := qtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (qtc *QuestionTranslationCreate) defaults() {
	if _, ok := qtc.mutation.ID(); !ok {
		v := questiontranslation.DefaultID()
		qtc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (qtc *QuestionTranslationCreate) check() error {
	if _, ok := qtc.mutation.Locale(); !ok {
		return &ValidationError{Name: "locale", err: errors.New(`ent: missing required field "QuestionTranslation.locale"`)}
	}
	if v, ok := qtc.mutation.Locale(); ok {
		if err := questiontranslation.LocaleValidator(v); err != nil {
			return &ValidationError{Name: "locale", err: fmt.Errorf(`ent: validator failed for field "QuestionTranslation.locale": %w`, err)}
		}
	}
	return nil
}

func (qtc *QuestionTranslationCreate) sqlSave(ctx context.Context) (*QuestionTranslation, error) {
	_node, _spec := qtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qtc.driver, _spec); err != nil {
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

func (qtc *QuestionTranslationCreate) createSpec() (*QuestionTranslation, *sqlgraph.CreateSpec) {
	var (
		_node = &QuestionTranslation{config: qtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: questiontranslation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: questiontranslation.FieldID,
			},
		}
	)
	if id, ok := qtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := qtc.mutation.Locale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: questiontranslation.FieldLocale,
		})
		_node.Locale = value
	}
	if value, ok := qtc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := qtc.mutation.HeaderContent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldHeaderContent,
		})
		_node.HeaderContent = value
	}
	if value, ok := qtc.mutation.FooterContent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questiontranslation.FieldFooterContent,
		})
		_node.FooterContent = value
	}
	if nodes := qtc.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questiontranslation.QuestionTable,
			Columns: []string{questiontranslation.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.question_translations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QuestionTranslationCreateBulk is the builder for creating many QuestionTranslation entities in bulk.
type QuestionTranslationCreateBulk struct {
	config
	builders []*QuestionTranslationCreate
}

// Save creates the QuestionTranslation entities in the database.
func (qtcb *QuestionTranslationCreateBulk) Save(ctx context.Context) ([]*QuestionTranslation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qtcb.builders))
	nodes := make([]*QuestionTranslation, len(qtcb.builders))
	mutators := make([]Mutator, len(qtcb.builders))
	for i := range qtcb.builders {
		func(i int, root context.Context) {
			builder := qtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuestionTranslationMutation)
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
					_, err = mutators[i+1].Mutate(root, qtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, qtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qtcb *QuestionTranslationCreateBulk) SaveX(ctx context.Context) []*QuestionTranslation {
	v, err := qtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (qtcb *QuestionTranslationCreateBulk) Exec(ctx context.Context) error {
	_, err := qtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qtcb *QuestionTranslationCreateBulk) ExecX(ctx context.Context) {
	if err := qtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
