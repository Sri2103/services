package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().StructTag(`json:"id"`),
		field.String("name").StructTag(`json:"name"`),
		field.String("description").StructTag(`json:"description,omitempty"`),
		field.Float32("price").StructTag(`json:"price"`),
		field.Time("created_at").Immutable().Default(time.Now()).StructTag(`json:"createdAt,omitempty"`),
		field.Time("updated_at").Default(nil),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
