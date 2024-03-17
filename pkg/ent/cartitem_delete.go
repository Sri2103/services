// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sri2103/services/pkg/ent/cartitem"
	"github.com/Sri2103/services/pkg/ent/predicate"
)

// CartItemDelete is the builder for deleting a CartItem entity.
type CartItemDelete struct {
	config
	hooks    []Hook
	mutation *CartItemMutation
}

// Where appends a list predicates to the CartItemDelete builder.
func (cid *CartItemDelete) Where(ps ...predicate.CartItem) *CartItemDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *CartItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cid.sqlExec, cid.mutation, cid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *CartItemDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *CartItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(cartitem.Table, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID))
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cid.mutation.done = true
	return affected, err
}

// CartItemDeleteOne is the builder for deleting a single CartItem entity.
type CartItemDeleteOne struct {
	cid *CartItemDelete
}

// Where appends a list predicates to the CartItemDelete builder.
func (cido *CartItemDeleteOne) Where(ps ...predicate.CartItem) *CartItemDeleteOne {
	cido.cid.mutation.Where(ps...)
	return cido
}

// Exec executes the deletion query.
func (cido *CartItemDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{cartitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *CartItemDeleteOne) ExecX(ctx context.Context) {
	if err := cido.Exec(ctx); err != nil {
		panic(err)
	}
}
