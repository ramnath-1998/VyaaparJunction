package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("productName").NotEmpty(),
		field.UUID("Identifier", uuid.UUID{}).Default(uuid.New),
		field.Time("createdOn").Default(time.Now().UTC).Immutable(),
		field.Int("categoryId"),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", ProductCategory.Type).
			Ref("product").
			Field("categoryId").
			Unique().
			Required(),
	}
}
