package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ProductCate struct {
	ent.Schema
}

func (ProductCate) Config() ent.Config {
	return ent.Config{
		Table: "product_cate",
	}
}

func (ProductCate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (ProductCate) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("sort"),
	}
}

func (ProductCate) Edges() []ent.Edge {
	return []ent.Edge{}
}
