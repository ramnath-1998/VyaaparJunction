package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type ProductCategory struct {
	ent.Schema
}

func (ProductCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("categoryName").NotEmpty(),
		field.UUID("Identifier", uuid.UUID{}).Default(uuid.New),
		field.Time("createdOn").Default(time.Now().UTC).Immutable(),
	}
}

func (ProductCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
	}
}
