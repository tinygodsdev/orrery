// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/challenge"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/prediction"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/proof"
	"github.com/DanielTitkov/lentils/internal/repository/entgo/ent/user"
	"github.com/google/uuid"
)

// ChallengeCreate is the builder for creating a Challenge entity.
type ChallengeCreate struct {
	config
	mutation *ChallengeMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *ChallengeCreate) SetCreateTime(t time.Time) *ChallengeCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableCreateTime(t *time.Time) *ChallengeCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *ChallengeCreate) SetUpdateTime(t time.Time) *ChallengeCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableUpdateTime(t *time.Time) *ChallengeCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *ChallengeCreate) SetContent(s string) *ChallengeCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *ChallengeCreate) SetDescription(s string) *ChallengeCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableDescription(s *string) *ChallengeCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// SetOutcome sets the "outcome" field.
func (cc *ChallengeCreate) SetOutcome(b bool) *ChallengeCreate {
	cc.mutation.SetOutcome(b)
	return cc
}

// SetNillableOutcome sets the "outcome" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableOutcome(b *bool) *ChallengeCreate {
	if b != nil {
		cc.SetOutcome(*b)
	}
	return cc
}

// SetPublished sets the "published" field.
func (cc *ChallengeCreate) SetPublished(b bool) *ChallengeCreate {
	cc.mutation.SetPublished(b)
	return cc
}

// SetNillablePublished sets the "published" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillablePublished(b *bool) *ChallengeCreate {
	if b != nil {
		cc.SetPublished(*b)
	}
	return cc
}

// SetStartTime sets the "start_time" field.
func (cc *ChallengeCreate) SetStartTime(t time.Time) *ChallengeCreate {
	cc.mutation.SetStartTime(t)
	return cc
}

// SetEndTime sets the "end_time" field.
func (cc *ChallengeCreate) SetEndTime(t time.Time) *ChallengeCreate {
	cc.mutation.SetEndTime(t)
	return cc
}

// SetType sets the "type" field.
func (cc *ChallengeCreate) SetType(c challenge.Type) *ChallengeCreate {
	cc.mutation.SetType(c)
	return cc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableType(c *challenge.Type) *ChallengeCreate {
	if c != nil {
		cc.SetType(*c)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ChallengeCreate) SetID(u uuid.UUID) *ChallengeCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableID(u *uuid.UUID) *ChallengeCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddPredictionIDs adds the "predictions" edge to the Prediction entity by IDs.
func (cc *ChallengeCreate) AddPredictionIDs(ids ...uuid.UUID) *ChallengeCreate {
	cc.mutation.AddPredictionIDs(ids...)
	return cc
}

// AddPredictions adds the "predictions" edges to the Prediction entity.
func (cc *ChallengeCreate) AddPredictions(p ...*Prediction) *ChallengeCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddPredictionIDs(ids...)
}

// AddProofIDs adds the "proofs" edge to the Proof entity by IDs.
func (cc *ChallengeCreate) AddProofIDs(ids ...uuid.UUID) *ChallengeCreate {
	cc.mutation.AddProofIDs(ids...)
	return cc
}

// AddProofs adds the "proofs" edges to the Proof entity.
func (cc *ChallengeCreate) AddProofs(p ...*Proof) *ChallengeCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProofIDs(ids...)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (cc *ChallengeCreate) SetAuthorID(id uuid.UUID) *ChallengeCreate {
	cc.mutation.SetAuthorID(id)
	return cc
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (cc *ChallengeCreate) SetNillableAuthorID(id *uuid.UUID) *ChallengeCreate {
	if id != nil {
		cc = cc.SetAuthorID(*id)
	}
	return cc
}

// SetAuthor sets the "author" edge to the User entity.
func (cc *ChallengeCreate) SetAuthor(u *User) *ChallengeCreate {
	return cc.SetAuthorID(u.ID)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cc *ChallengeCreate) Mutation() *ChallengeMutation {
	return cc.mutation
}

// Save creates the Challenge in the database.
func (cc *ChallengeCreate) Save(ctx context.Context) (*Challenge, error) {
	var (
		err  error
		node *Challenge
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChallengeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChallengeCreate) SaveX(ctx context.Context) *Challenge {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChallengeCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChallengeCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChallengeCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := challenge.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := challenge.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.Published(); !ok {
		v := challenge.DefaultPublished
		cc.mutation.SetPublished(v)
	}
	if _, ok := cc.mutation.GetType(); !ok {
		v := challenge.DefaultType
		cc.mutation.SetType(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := challenge.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChallengeCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Challenge.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Challenge.update_time"`)}
	}
	if _, ok := cc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Challenge.content"`)}
	}
	if v, ok := cc.mutation.Content(); ok {
		if err := challenge.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Challenge.content": %w`, err)}
		}
	}
	if v, ok := cc.mutation.Description(); ok {
		if err := challenge.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Challenge.description": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Published(); !ok {
		return &ValidationError{Name: "published", err: errors.New(`ent: missing required field "Challenge.published"`)}
	}
	if _, ok := cc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "Challenge.start_time"`)}
	}
	if _, ok := cc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "end_time", err: errors.New(`ent: missing required field "Challenge.end_time"`)}
	}
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Challenge.type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := challenge.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Challenge.type": %w`, err)}
		}
	}
	return nil
}

func (cc *ChallengeCreate) sqlSave(ctx context.Context) (*Challenge, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ChallengeCreate) createSpec() (*Challenge, *sqlgraph.CreateSpec) {
	var (
		_node = &Challenge{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: challenge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: challenge.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: challenge.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: challenge.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := cc.mutation.Outcome(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: challenge.FieldOutcome,
		})
		_node.Outcome = &value
	}
	if value, ok := cc.mutation.Published(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: challenge.FieldPublished,
		})
		_node.Published = value
	}
	if value, ok := cc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := cc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: challenge.FieldType,
		})
		_node.Type = value
	}
	if nodes := cc.mutation.PredictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ProofsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   challenge.AuthorTable,
			Columns: []string{challenge.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_challenges = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChallengeCreateBulk is the builder for creating many Challenge entities in bulk.
type ChallengeCreateBulk struct {
	config
	builders []*ChallengeCreate
}

// Save creates the Challenge entities in the database.
func (ccb *ChallengeCreateBulk) Save(ctx context.Context) ([]*Challenge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Challenge, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChallengeMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChallengeCreateBulk) SaveX(ctx context.Context) []*Challenge {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChallengeCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChallengeCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
