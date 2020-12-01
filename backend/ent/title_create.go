// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Piichet/app/ent/title"
	"github.com/Piichet/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// TitleCreate is the builder for creating a Title entity.
type TitleCreate struct {
	config
	mutation *TitleMutation
	hooks    []Hook
}

// SetTitle sets the title field.
func (tc *TitleCreate) SetTitle(s string) *TitleCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// AddUserIDs adds the users edge to User by ids.
func (tc *TitleCreate) AddUserIDs(ids ...int) *TitleCreate {
	tc.mutation.AddUserIDs(ids...)
	return tc
}

// AddUsers adds the users edges to User.
func (tc *TitleCreate) AddUsers(u ...*User) *TitleCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddUserIDs(ids...)
}

// Mutation returns the TitleMutation object of the builder.
func (tc *TitleCreate) Mutation() *TitleMutation {
	return tc.mutation
}

// Save creates the Title in the database.
func (tc *TitleCreate) Save(ctx context.Context) (*Title, error) {
	if _, ok := tc.mutation.Title(); !ok {
		return nil, &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := tc.mutation.Title(); ok {
		if err := title.TitleValidator(v); err != nil {
			return nil, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	var (
		err  error
		node *Title
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TitleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TitleCreate) SaveX(ctx context.Context) *Title {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TitleCreate) sqlSave(ctx context.Context) (*Title, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TitleCreate) createSpec() (*Title, *sqlgraph.CreateSpec) {
	var (
		t     = &Title{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: title.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: title.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldTitle,
		})
		t.Title = value
	}
	if nodes := tc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   title.UsersTable,
			Columns: []string{title.UsersColumn},
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
	return t, _spec
}
