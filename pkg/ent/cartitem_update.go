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
	"github.com/Sri2103/services/pkg/ent/product"
	"github.com/google/uuid"
)

// CartItemUpdate is the builder for updating CartItem entities.
type CartItemUpdate struct {
	config
	hooks    []Hook
	mutation *CartItemMutation
}

// Where appends a list predicates to the CartItemUpdate builder.
func (ciu *CartItemUpdate) Where(ps ...predicate.CartItem) *CartItemUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetQuantity sets the "quantity" field.
func (ciu *CartItemUpdate) SetQuantity(i int) *CartItemUpdate {
	ciu.mutation.ResetQuantity()
	ciu.mutation.SetQuantity(i)
	return ciu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (ciu *CartItemUpdate) SetNillableQuantity(i *int) *CartItemUpdate {
	if i != nil {
		ciu.SetQuantity(*i)
	}
	return ciu
}

// AddQuantity adds i to the "quantity" field.
func (ciu *CartItemUpdate) AddQuantity(i int) *CartItemUpdate {
	ciu.mutation.AddQuantity(i)
	return ciu
}

// SetCreatedAt sets the "created_at" field.
func (ciu *CartItemUpdate) SetCreatedAt(t time.Time) *CartItemUpdate {
	ciu.mutation.SetCreatedAt(t)
	return ciu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciu *CartItemUpdate) SetNillableCreatedAt(t *time.Time) *CartItemUpdate {
	if t != nil {
		ciu.SetCreatedAt(*t)
	}
	return ciu
}

// SetUpdatedAt sets the "updated_at" field.
func (ciu *CartItemUpdate) SetUpdatedAt(t time.Time) *CartItemUpdate {
	ciu.mutation.SetUpdatedAt(t)
	return ciu
}

// SetCartID sets the "cart" edge to the Cart entity by ID.
func (ciu *CartItemUpdate) SetCartID(id uuid.UUID) *CartItemUpdate {
	ciu.mutation.SetCartID(id)
	return ciu
}

// SetNillableCartID sets the "cart" edge to the Cart entity by ID if the given value is not nil.
func (ciu *CartItemUpdate) SetNillableCartID(id *uuid.UUID) *CartItemUpdate {
	if id != nil {
		ciu = ciu.SetCartID(*id)
	}
	return ciu
}

// SetCart sets the "cart" edge to the Cart entity.
func (ciu *CartItemUpdate) SetCart(c *Cart) *CartItemUpdate {
	return ciu.SetCartID(c.ID)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (ciu *CartItemUpdate) AddProductIDs(ids ...uuid.UUID) *CartItemUpdate {
	ciu.mutation.AddProductIDs(ids...)
	return ciu
}

// AddProduct adds the "product" edges to the Product entity.
func (ciu *CartItemUpdate) AddProduct(p ...*Product) *CartItemUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ciu.AddProductIDs(ids...)
}

// Mutation returns the CartItemMutation object of the builder.
func (ciu *CartItemUpdate) Mutation() *CartItemMutation {
	return ciu.mutation
}

// ClearCart clears the "cart" edge to the Cart entity.
func (ciu *CartItemUpdate) ClearCart() *CartItemUpdate {
	ciu.mutation.ClearCart()
	return ciu
}

// ClearProduct clears all "product" edges to the Product entity.
func (ciu *CartItemUpdate) ClearProduct() *CartItemUpdate {
	ciu.mutation.ClearProduct()
	return ciu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (ciu *CartItemUpdate) RemoveProductIDs(ids ...uuid.UUID) *CartItemUpdate {
	ciu.mutation.RemoveProductIDs(ids...)
	return ciu
}

// RemoveProduct removes "product" edges to Product entities.
func (ciu *CartItemUpdate) RemoveProduct(p ...*Product) *CartItemUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ciu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *CartItemUpdate) Save(ctx context.Context) (int, error) {
	ciu.defaults()
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *CartItemUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *CartItemUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *CartItemUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciu *CartItemUpdate) defaults() {
	if _, ok := ciu.mutation.UpdatedAt(); !ok {
		v := cartitem.UpdateDefaultUpdatedAt()
		ciu.mutation.SetUpdatedAt(v)
	}
}

func (ciu *CartItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.Quantity(); ok {
		_spec.SetField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := ciu.mutation.AddedQuantity(); ok {
		_spec.AddField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := ciu.mutation.CreatedAt(); ok {
		_spec.SetField(cartitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ciu.mutation.UpdatedAt(); ok {
		_spec.SetField(cartitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if ciu.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: []string{cartitem.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: []string{cartitem.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.RemovedProductIDs(); len(nodes) > 0 && !ciu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cartitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// CartItemUpdateOne is the builder for updating a single CartItem entity.
type CartItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CartItemMutation
}

// SetQuantity sets the "quantity" field.
func (ciuo *CartItemUpdateOne) SetQuantity(i int) *CartItemUpdateOne {
	ciuo.mutation.ResetQuantity()
	ciuo.mutation.SetQuantity(i)
	return ciuo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (ciuo *CartItemUpdateOne) SetNillableQuantity(i *int) *CartItemUpdateOne {
	if i != nil {
		ciuo.SetQuantity(*i)
	}
	return ciuo
}

// AddQuantity adds i to the "quantity" field.
func (ciuo *CartItemUpdateOne) AddQuantity(i int) *CartItemUpdateOne {
	ciuo.mutation.AddQuantity(i)
	return ciuo
}

// SetCreatedAt sets the "created_at" field.
func (ciuo *CartItemUpdateOne) SetCreatedAt(t time.Time) *CartItemUpdateOne {
	ciuo.mutation.SetCreatedAt(t)
	return ciuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ciuo *CartItemUpdateOne) SetNillableCreatedAt(t *time.Time) *CartItemUpdateOne {
	if t != nil {
		ciuo.SetCreatedAt(*t)
	}
	return ciuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ciuo *CartItemUpdateOne) SetUpdatedAt(t time.Time) *CartItemUpdateOne {
	ciuo.mutation.SetUpdatedAt(t)
	return ciuo
}

// SetCartID sets the "cart" edge to the Cart entity by ID.
func (ciuo *CartItemUpdateOne) SetCartID(id uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.SetCartID(id)
	return ciuo
}

// SetNillableCartID sets the "cart" edge to the Cart entity by ID if the given value is not nil.
func (ciuo *CartItemUpdateOne) SetNillableCartID(id *uuid.UUID) *CartItemUpdateOne {
	if id != nil {
		ciuo = ciuo.SetCartID(*id)
	}
	return ciuo
}

// SetCart sets the "cart" edge to the Cart entity.
func (ciuo *CartItemUpdateOne) SetCart(c *Cart) *CartItemUpdateOne {
	return ciuo.SetCartID(c.ID)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (ciuo *CartItemUpdateOne) AddProductIDs(ids ...uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.AddProductIDs(ids...)
	return ciuo
}

// AddProduct adds the "product" edges to the Product entity.
func (ciuo *CartItemUpdateOne) AddProduct(p ...*Product) *CartItemUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ciuo.AddProductIDs(ids...)
}

// Mutation returns the CartItemMutation object of the builder.
func (ciuo *CartItemUpdateOne) Mutation() *CartItemMutation {
	return ciuo.mutation
}

// ClearCart clears the "cart" edge to the Cart entity.
func (ciuo *CartItemUpdateOne) ClearCart() *CartItemUpdateOne {
	ciuo.mutation.ClearCart()
	return ciuo
}

// ClearProduct clears all "product" edges to the Product entity.
func (ciuo *CartItemUpdateOne) ClearProduct() *CartItemUpdateOne {
	ciuo.mutation.ClearProduct()
	return ciuo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (ciuo *CartItemUpdateOne) RemoveProductIDs(ids ...uuid.UUID) *CartItemUpdateOne {
	ciuo.mutation.RemoveProductIDs(ids...)
	return ciuo
}

// RemoveProduct removes "product" edges to Product entities.
func (ciuo *CartItemUpdateOne) RemoveProduct(p ...*Product) *CartItemUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ciuo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the CartItemUpdate builder.
func (ciuo *CartItemUpdateOne) Where(ps ...predicate.CartItem) *CartItemUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *CartItemUpdateOne) Select(field string, fields ...string) *CartItemUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated CartItem entity.
func (ciuo *CartItemUpdateOne) Save(ctx context.Context) (*CartItem, error) {
	ciuo.defaults()
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *CartItemUpdateOne) SaveX(ctx context.Context) *CartItem {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *CartItemUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *CartItemUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ciuo *CartItemUpdateOne) defaults() {
	if _, ok := ciuo.mutation.UpdatedAt(); !ok {
		v := cartitem.UpdateDefaultUpdatedAt()
		ciuo.mutation.SetUpdatedAt(v)
	}
}

func (ciuo *CartItemUpdateOne) sqlSave(ctx context.Context) (_node *CartItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(cartitem.Table, cartitem.Columns, sqlgraph.NewFieldSpec(cartitem.FieldID, field.TypeUUID))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CartItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cartitem.FieldID)
		for _, f := range fields {
			if !cartitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cartitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.Quantity(); ok {
		_spec.SetField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := ciuo.mutation.AddedQuantity(); ok {
		_spec.AddField(cartitem.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := ciuo.mutation.CreatedAt(); ok {
		_spec.SetField(cartitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := ciuo.mutation.UpdatedAt(); ok {
		_spec.SetField(cartitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if ciuo.mutation.CartCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: []string{cartitem.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cartitem.CartTable,
			Columns: []string{cartitem.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(cart.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ciuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.RemovedProductIDs(); len(nodes) > 0 && !ciuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ciuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cartitem.ProductTable,
			Columns: []string{cartitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CartItem{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cartitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
