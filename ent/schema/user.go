package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("name").NotEmpty(),
		field.UUID("identifier", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Time("createdOn").Default(time.Now().UTC).Immutable(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type),
	}
}
