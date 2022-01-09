package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type OrderSetting struct {
	ent.Schema
}

func (OrderSetting) Config() ent.Config {
	return ent.Config{
		Table: "order_pay",
	}
}

// Fields of the Order.
func (OrderSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_overtime ").
			Optional().
			Default(0).
			Comment("订单超时关闭时间(分)"),
		field.Int("confirm_overtime").
			Optional().
			Default(0).
			Comment("发货后自动确认收货时间（天）"),
		field.Int("finish_overtime").
			Optional().
			Default(0).
			Comment("自动完成交易时间，不能申请售后（天）"),
		field.Int("comment_overtime").
			Optional().
			Default(0).
			Comment("订单完成后自动好评时间（天）"),
	}
}

// Edges of the OrderPay.
func (OrderSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("pays").
			Unique().
			Field("order_id"),
	}
}

func (OrderSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
func (OrderSetting) Index() []ent.Index {
	return []ent.Index{

	}
}
