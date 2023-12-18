package schema

import "entgo.io/ent"

type Order_Item struct {
	ent.Schema
}

func (Order_Item) Fields() []ent.Field {
	return nil
}

func (Order_Item) Edges() []ent.Edge {
	return nil
}
