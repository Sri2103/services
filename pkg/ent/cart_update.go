// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Sri2103/services/pkg/ent/cart"
	"github.com/Sri2103/services/pkg/ent/cartitem"
	"github.com/Sri2103/services/pkg/ent/predicate"
	"github.com/Sri2103/services/pkg/ent/user"
	"github.com/google/uuid"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// Where appends a list predicates to the CartUpdate builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CartUpdate) SetCreatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CartUpdate) SetNillableCreatedAt(t *time.Time) *CartUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CartUpdate) SetUpdatedAt(t time.Time) *CartUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// AddItemIDs adds the "items" edge to the CartItem entity by IDs.
func (cu *CartUpdate) AddItemIDs(ids ...uuid.UUID) *CartUpdate {
	cu.mutation.AddItemIDs(ids...)
	return cu
}

// AddItems adds the "items" edges to the CartItem entity.
func (cu *CartUpdate) AddItems(c ...*CartItem) *CartUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddItemIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *CartUpdate) SetUserID(id uuid.UUID) *CartUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cu *CartUpdate) SetNillableUserID(id *uuid.UUID) *CartUpdate {
	if id != nil {
		cu = cu.SetUserID(*id)
	}
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *CartUpdate) SetUser(u *User) *CartUpdate {
	return cu.SetUserID(u.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// ClearItems clears all "items" edges to the CartItem entity.
func (cu *CartUpdate) ClearItems() *CartUpdate {
	cu.mutation.ClearItems()
	return cu
}

// RemoveItemIDs removes the "items" edge to CartItem entities by IDs.
func (cu *CartUpdate) RemoveItemIDs(ids ...uuid.UUID) *CartUpdate {
	cu.mutation.RemoveItemIDs(ids...)
	return cu
}

// RemoveItems removes "items" edges to CartItem entities.
func (cu *CartUpdate) RemoveItems(c ...*CartItem) *CartUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveItemIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (cu *CartUpdate) ClearUser() *CartUpdate {
	cu.mutation.ClearUser()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CartUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(cart.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.ItemsTable,
			Columns: []string{cart.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !cu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.ItemsTable,
			Columns: []string{cart.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.ItemsTable,
			Columns: []string{cart.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartMutation
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CartUpdateOne) SetCreatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableCreatedAt(t *time.Time) *CartUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CartUpdateOne) SetUpdatedAt(t time.Time) *CartUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// AddItemIDs adds the "items" edge to the CartItem entity by IDs.
func (cuo *CartUpdateOne) AddItemIDs(ids ...uuid.UUID) *CartUpdateOne {
	cuo.mutation.AddItemIDs(ids...)
	return cuo
}

// AddItems adds the "items" edges to the CartItem entity.
func (cuo *CartUpdateOne) AddItems(c ...*CartItem) *CartUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddItemIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *CartUpdateOne) SetUserID(id uuid.UUID) *CartUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (cuo *CartUpdateOne) SetNillableUserID(id *uuid.UUID) *CartUpdateOne {
	if id != nil {
		cuo = cuo.SetUserID(*id)
	}
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *CartUpdateOne) SetUser(u *User) *CartUpdateOne {
	return cuo.SetUserID(u.ID)
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// ClearItems clears all "items" edges to the CartItem entity.
func (cuo *CartUpdateOne) ClearItems() *CartUpdateOne {
	cuo.mutation.ClearItems()
	return cuo
}

// RemoveItemIDs removes the "items" edge to CartItem entities by IDs.
func (cuo *CartUpdateOne) RemoveItemIDs(ids ...uuid.UUID) *CartUpdateOne {
	cuo.mutation.RemoveItemIDs(ids...)
	return cuo
}

// RemoveItems removes "items" edges to CartItem entities.
func (cuo *CartUpdateOne) RemoveItems(c ...*CartItem) *CartUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveItemIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *CartUpdateOne) ClearUser() *CartUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// Where appends a list predicates to the CartUpdate builder.
func (cuo *CartUpdateOne) Where(ps ...predicate.Cart) *CartUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CartUpdateOne) Select(field string, fields ...string) *CartUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cart entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CartUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := cart.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (_node *Cart, err error) {
	_spec := sqlgraph.NewUpdateSpec(cart.Table, cart.Columns, sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cart.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cart.FieldID)
		for _, f := range fields {
			if !cart.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(cart.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cart.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.ItemsTable,
			Columns: []string{cart.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !cuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.ItemsTable,
			Columns: []string{cart.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.ItemsTable,
			Columns: []string{cart.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cart.UserTable,
			Columns: []string{cart.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cart{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
