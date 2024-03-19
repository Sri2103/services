// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Sri2103/services/pkg/ent/address"
	"github.com/Sri2103/services/pkg/ent/cart"
	"github.com/Sri2103/services/pkg/ent/cartitem"
	"github.com/Sri2103/services/pkg/ent/category"
	"github.com/Sri2103/services/pkg/ent/order"
	"github.com/Sri2103/services/pkg/ent/orderitem"
	"github.com/Sri2103/services/pkg/ent/product"
	"github.com/Sri2103/services/pkg/ent/role"
	"github.com/Sri2103/services/pkg/ent/schema"
	"github.com/Sri2103/services/pkg/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressFields[5].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressFields[6].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// address.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	address.UpdateDefaultUpdatedAt = addressDescUpdatedAt.UpdateDefault.(func() time.Time)
	// addressDescID is the schema descriptor for id field.
	addressDescID := addressFields[0].Descriptor()
	// address.DefaultID holds the default value on creation for the id field.
	address.DefaultID = addressDescID.Default.(func() uuid.UUID)
	cartFields := schema.Cart{}.Fields()
	_ = cartFields
	// cartDescCreatedAt is the schema descriptor for created_at field.
	cartDescCreatedAt := cartFields[1].Descriptor()
	// cart.DefaultCreatedAt holds the default value on creation for the created_at field.
	cart.DefaultCreatedAt = cartDescCreatedAt.Default.(func() time.Time)
	// cartDescUpdatedAt is the schema descriptor for updated_at field.
	cartDescUpdatedAt := cartFields[2].Descriptor()
	// cart.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cart.DefaultUpdatedAt = cartDescUpdatedAt.Default.(func() time.Time)
	// cart.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cart.UpdateDefaultUpdatedAt = cartDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cartDescID is the schema descriptor for id field.
	cartDescID := cartFields[0].Descriptor()
	// cart.DefaultID holds the default value on creation for the id field.
	cart.DefaultID = cartDescID.Default.(func() uuid.UUID)
	cartitemFields := schema.CartItem{}.Fields()
	_ = cartitemFields
	// cartitemDescCreatedAt is the schema descriptor for created_at field.
	cartitemDescCreatedAt := cartitemFields[2].Descriptor()
	// cartitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	cartitem.DefaultCreatedAt = cartitemDescCreatedAt.Default.(func() time.Time)
	// cartitemDescUpdatedAt is the schema descriptor for updated_at field.
	cartitemDescUpdatedAt := cartitemFields[3].Descriptor()
	// cartitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cartitem.DefaultUpdatedAt = cartitemDescUpdatedAt.Default.(func() time.Time)
	// cartitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	cartitem.UpdateDefaultUpdatedAt = cartitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// cartitemDescID is the schema descriptor for id field.
	cartitemDescID := cartitemFields[0].Descriptor()
	// cartitem.DefaultID holds the default value on creation for the id field.
	cartitem.DefaultID = cartitemDescID.Default.(func() uuid.UUID)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreatedAt is the schema descriptor for created_at field.
	categoryDescCreatedAt := categoryFields[2].Descriptor()
	// category.DefaultCreatedAt holds the default value on creation for the created_at field.
	category.DefaultCreatedAt = categoryDescCreatedAt.Default.(func() time.Time)
	// categoryDescUpdatedAt is the schema descriptor for updated_at field.
	categoryDescUpdatedAt := categoryFields[3].Descriptor()
	// category.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	category.DefaultUpdatedAt = categoryDescUpdatedAt.Default.(func() time.Time)
	// category.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	category.UpdateDefaultUpdatedAt = categoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// categoryDescID is the schema descriptor for id field.
	categoryDescID := categoryFields[0].Descriptor()
	// category.DefaultID holds the default value on creation for the id field.
	category.DefaultID = categoryDescID.Default.(func() uuid.UUID)
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderFields[2].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderFields[3].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderDescID is the schema descriptor for id field.
	orderDescID := orderFields[0].Descriptor()
	// order.DefaultID holds the default value on creation for the id field.
	order.DefaultID = orderDescID.Default.(func() uuid.UUID)
	orderitemFields := schema.OrderItem{}.Fields()
	_ = orderitemFields
	// orderitemDescCreatedAt is the schema descriptor for created_at field.
	orderitemDescCreatedAt := orderitemFields[3].Descriptor()
	// orderitem.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderitem.DefaultCreatedAt = orderitemDescCreatedAt.Default.(func() time.Time)
	// orderitemDescUpdatedAt is the schema descriptor for updated_at field.
	orderitemDescUpdatedAt := orderitemFields[4].Descriptor()
	// orderitem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderitem.DefaultUpdatedAt = orderitemDescUpdatedAt.Default.(func() time.Time)
	// orderitem.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderitem.UpdateDefaultUpdatedAt = orderitemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderitemDescID is the schema descriptor for id field.
	orderitemDescID := orderitemFields[0].Descriptor()
	// orderitem.DefaultID holds the default value on creation for the id field.
	orderitem.DefaultID = orderitemDescID.Default.(func() uuid.UUID)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productFields[4].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productFields[5].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() uuid.UUID)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[2].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[3].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleFields[0].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[6].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
}
