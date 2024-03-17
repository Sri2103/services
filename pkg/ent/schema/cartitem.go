package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CartItem holds the schema definition for the CartItem entity.
type CartItem struct {
	ent.Schema
}

// Fields of the CartItem.
func (CartItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Int("quantity"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the CartItem.
func (CartItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cart", Cart.Type).
			Ref("items").
			Unique(),
		edge.To("product", Product.Type),
	}
}
