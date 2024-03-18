// Code generated by ent, DO NOT EDIT.

package cartitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cartitem type in the database.
	Label = "cart_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeCart holds the string denoting the cart edge name in mutations.
	EdgeCart = "cart"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// Table holds the table name of the cartitem in the database.
	Table = "cart_items"
	// CartTable is the table that holds the cart relation/edge.
	CartTable = "cart_items"
	// CartInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartInverseTable = "carts"
	// CartColumn is the table column denoting the cart relation/edge.
	CartColumn = "cart_items"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "products"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "cart_item_product"
)

// Columns holds all SQL columns for cartitem fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cart_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"cart_items",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the CartItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCartField orders the results by cart field.
func ByCartField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCartStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductCount orders the results by product count.
func ByProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductStep(), opts...)
	}
}

// ByProduct orders the results by product terms.
func ByProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCartStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CartInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CartTable, CartColumn),
	)
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
	)
}