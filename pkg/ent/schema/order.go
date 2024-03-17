package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Float("total_amount"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("orders").
			Unique(),
		edge.To("items", OrderItem.Type),
	}
}
