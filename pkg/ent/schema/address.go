package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("street_address"),
		field.String("city"),
		field.String("state"),
		field.String("country"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("addresses").
			Unique(),
	}
}
