package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{

		field.Int("userId"),
		field.UUID("identifier", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("createdOn").Default(time.Now().UTC).Immutable(),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("order").
			Field("userId").
			Unique().
			Required(),
	}
}
