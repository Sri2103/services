package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Immutable().StorageKey("_id"),
		field.String("name"),
		field.String("username"),
		field.String("email").Unique(),
		field.Bytes("password").Sensitive(),
		field.Time("created_at").Immutable().Default(time.Now()),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("carts", Cart.Type),
		edge.To("orders", Order.Type),
		edge.To("addresses", Address.Type),
	}
}
