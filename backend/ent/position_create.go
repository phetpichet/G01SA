// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet/app/ent/position"
	"github.com/Piichet/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
}

// SetPosition sets the position field.
func (pc *PositionCreate) SetPosition(s string) *PositionCreate {
	pc.mutation.SetPosition(s)
	return pc
}

// AddUserIDs adds the users edge to User by ids.
func (pc *PositionCreate) AddUserIDs(ids ...int) *PositionCreate {
	pc.mutation.AddUserIDs(ids...)
	return pc
}

// AddUsers adds the users edges to User.
func (pc *PositionCreate) AddUsers(u ...*User) *PositionCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	if _, ok := pc.mutation.Position(); !ok {
		return nil, &ValidationError{Name: "position", err: errors.New("ent: missing required field \"position\"")}
	}
	if v, ok := pc.mutation.Position(); ok {
		if err := position.PositionValidator(v); err != nil {
			return nil, &ValidationError{Name: "position", err: fmt.Errorf("ent: validator failed for field \"position\": %w", err)}
		}
	}
	var (
		err  error
		node *Position
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
	po, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	po.ID = int(id)
	return po, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		po    = &Position{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: position.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: position.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Position(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldPosition,
		})
		po.Position = value
	}
	if nodes := pc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.UsersTable,
			Columns: []string{position.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return po, _spec
}
