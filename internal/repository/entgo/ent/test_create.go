// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/question"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/scale"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/tag"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/take"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/test"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/testdisplay"
	"github.com/DanielTitkov/orrery/internal/repository/entgo/ent/testtranslation"
	"github.com/google/uuid"
)

// TestCreate is the builder for creating a Test entity.
type TestCreate struct {
	config
	mutation *TestMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (tc *TestCreate) SetCreateTime(t time.Time) *TestCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TestCreate) SetNillableCreateTime(t *time.Time) *TestCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TestCreate) SetUpdateTime(t time.Time) *TestCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TestCreate) SetNillableUpdateTime(t *time.Time) *TestCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetCode sets the "code" field.
func (tc *TestCreate) SetCode(s string) *TestCreate {
	tc.mutation.SetCode(s)
	return tc
}

// SetPublished sets the "published" field.
func (tc *TestCreate) SetPublished(b bool) *TestCreate {
	tc.mutation.SetPublished(b)
	return tc
}

// SetNillablePublished sets the "published" field if the given value is not nil.
func (tc *TestCreate) SetNillablePublished(b *bool) *TestCreate {
	if b != nil {
		tc.SetPublished(*b)
	}
	return tc
}

// SetAvailableLocales sets the "available_locales" field.
func (tc *TestCreate) SetAvailableLocales(s []string) *TestCreate {
	tc.mutation.SetAvailableLocales(s)
	return tc
}

// SetID sets the "id" field.
func (tc *TestCreate) SetID(u uuid.UUID) *TestCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TestCreate) SetNillableID(u *uuid.UUID) *TestCreate {
	if u != nil {
		tc.SetID(*u)
	}
	return tc
}

// AddTakeIDs adds the "takes" edge to the Take entity by IDs.
func (tc *TestCreate) AddTakeIDs(ids ...uuid.UUID) *TestCreate {
	tc.mutation.AddTakeIDs(ids...)
	return tc
}

// AddTakes adds the "takes" edges to the Take entity.
func (tc *TestCreate) AddTakes(t ...*Take) *TestCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTakeIDs(ids...)
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (tc *TestCreate) AddQuestionIDs(ids ...uuid.UUID) *TestCreate {
	tc.mutation.AddQuestionIDs(ids...)
	return tc
}

// AddQuestions adds the "questions" edges to the Question entity.
func (tc *TestCreate) AddQuestions(q ...*Question) *TestCreate {
	ids := make([]uuid.UUID, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return tc.AddQuestionIDs(ids...)
}

// AddTranslationIDs adds the "translations" edge to the TestTranslation entity by IDs.
func (tc *TestCreate) AddTranslationIDs(ids ...uuid.UUID) *TestCreate {
	tc.mutation.AddTranslationIDs(ids...)
	return tc
}

// AddTranslations adds the "translations" edges to the TestTranslation entity.
func (tc *TestCreate) AddTranslations(t ...*TestTranslation) *TestCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTranslationIDs(ids...)
}

// AddScaleIDs adds the "scales" edge to the Scale entity by IDs.
func (tc *TestCreate) AddScaleIDs(ids ...uuid.UUID) *TestCreate {
	tc.mutation.AddScaleIDs(ids...)
	return tc
}

// AddScales adds the "scales" edges to the Scale entity.
func (tc *TestCreate) AddScales(s ...*Scale) *TestCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tc.AddScaleIDs(ids...)
}

// SetDisplayID sets the "display" edge to the TestDisplay entity by ID.
func (tc *TestCreate) SetDisplayID(id uuid.UUID) *TestCreate {
	tc.mutation.SetDisplayID(id)
	return tc
}

// SetNillableDisplayID sets the "display" edge to the TestDisplay entity by ID if the given value is not nil.
func (tc *TestCreate) SetNillableDisplayID(id *uuid.UUID) *TestCreate {
	if id != nil {
		tc = tc.SetDisplayID(*id)
	}
	return tc
}

// SetDisplay sets the "display" edge to the TestDisplay entity.
func (tc *TestCreate) SetDisplay(t *TestDisplay) *TestCreate {
	return tc.SetDisplayID(t.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tc *TestCreate) AddTagIDs(ids ...uuid.UUID) *TestCreate {
	tc.mutation.AddTagIDs(ids...)
	return tc
}

// AddTags adds the "tags" edges to the Tag entity.
func (tc *TestCreate) AddTags(t ...*Tag) *TestCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTagIDs(ids...)
}

// Mutation returns the TestMutation object of the builder.
func (tc *TestCreate) Mutation() *TestMutation {
	return tc.mutation
}

// Save creates the Test in the database.
func (tc *TestCreate) Save(ctx context.Context) (*Test, error) {
	var (
		err  error
		node *Test
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Test)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TestCreate) SaveX(ctx context.Context) *Test {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TestCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TestCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TestCreate) defaults() {
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := test.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := test.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.Published(); !ok {
		v := test.DefaultPublished
		tc.mutation.SetPublished(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := test.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TestCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Test.create_time"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Test.update_time"`)}
	}
	if _, ok := tc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Test.code"`)}
	}
	if v, ok := tc.mutation.Code(); ok {
		if err := test.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Test.code": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Published(); !ok {
		return &ValidationError{Name: "published", err: errors.New(`ent: missing required field "Test.published"`)}
	}
	return nil
}

func (tc *TestCreate) sqlSave(ctx context.Context) (*Test, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TestCreate) createSpec() (*Test, *sqlgraph.CreateSpec) {
	var (
		_node = &Test{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: test.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: test.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: test.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: test.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: test.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := tc.mutation.Published(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: test.FieldPublished,
		})
		_node.Published = value
	}
	if value, ok := tc.mutation.AvailableLocales(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: test.FieldAvailableLocales,
		})
		_node.AvailableLocales = value
	}
	if nodes := tc.mutation.TakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TakesTable,
			Columns: []string{test.TakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: take.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.QuestionsTable,
			Columns: test.QuestionsPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TranslationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   test.TranslationsTable,
			Columns: []string{test.TranslationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testtranslation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ScalesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.ScalesTable,
			Columns: test.ScalesPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DisplayIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   test.DisplayTable,
			Columns: []string{test.DisplayColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: testdisplay.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   test.TagsTable,
			Columns: test.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
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

// TestCreateBulk is the builder for creating many Test entities in bulk.
type TestCreateBulk struct {
	config
	builders []*TestCreate
}

// Save creates the Test entities in the database.
func (tcb *TestCreateBulk) Save(ctx context.Context) ([]*Test, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Test, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TestMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TestCreateBulk) SaveX(ctx context.Context) []*Test {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TestCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TestCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
