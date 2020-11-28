// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/kallydev/privacy/ent/jdmodel"
	"github.com/kallydev/privacy/ent/predicate"
)

// JDModelDelete is the builder for deleting a JDModel entity.
type JDModelDelete struct {
	config
	hooks    []Hook
	mutation *JDModelMutation
}

// Where adds a new predicate to the delete builder.
func (jmd *JDModelDelete) Where(ps ...predicate.JDModel) *JDModelDelete {
	jmd.mutation.predicates = append(jmd.mutation.predicates, ps...)
	return jmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (jmd *JDModelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(jmd.hooks) == 0 {
		affected, err = jmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JDModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			jmd.mutation = mutation
			affected, err = jmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(jmd.hooks) - 1; i >= 0; i-- {
			mut = jmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, jmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (jmd *JDModelDelete) ExecX(ctx context.Context) int {
	n, err := jmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (jmd *JDModelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: jdmodel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: jdmodel.FieldID,
			},
		},
	}
	if ps := jmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, jmd.driver, _spec)
}

// JDModelDeleteOne is the builder for deleting a single JDModel entity.
type JDModelDeleteOne struct {
	jmd *JDModelDelete
}

// Exec executes the deletion query.
func (jmdo *JDModelDeleteOne) Exec(ctx context.Context) error {
	n, err := jmdo.jmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{jdmodel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (jmdo *JDModelDeleteOne) ExecX(ctx context.Context) {
	jmdo.jmd.ExecX(ctx)
}