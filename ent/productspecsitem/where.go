// Code generated by entc, DO NOT EDIT.

package productspecsitem

import (
	"orse/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ProductSpecsID applies equality check predicate on the "product_specs_id" field. It's identical to ProductSpecsIDEQ.
func ProductSpecsID(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductSpecsID), v))
	})
}

// ValueID applies equality check predicate on the "value_id" field. It's identical to ValueIDEQ.
func ValueID(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValueID), v))
	})
}

// Sort applies equality check predicate on the "sort" field. It's identical to SortEQ.
func Sort(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// ProductSpecsIDEQ applies the EQ predicate on the "product_specs_id" field.
func ProductSpecsIDEQ(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProductSpecsID), v))
	})
}

// ProductSpecsIDNEQ applies the NEQ predicate on the "product_specs_id" field.
func ProductSpecsIDNEQ(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProductSpecsID), v))
	})
}

// ProductSpecsIDIn applies the In predicate on the "product_specs_id" field.
func ProductSpecsIDIn(vs ...int) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProductSpecsID), v...))
	})
}

// ProductSpecsIDNotIn applies the NotIn predicate on the "product_specs_id" field.
func ProductSpecsIDNotIn(vs ...int) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProductSpecsID), v...))
	})
}

// ProductSpecsIDIsNil applies the IsNil predicate on the "product_specs_id" field.
func ProductSpecsIDIsNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProductSpecsID)))
	})
}

// ProductSpecsIDNotNil applies the NotNil predicate on the "product_specs_id" field.
func ProductSpecsIDNotNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProductSpecsID)))
	})
}

// ValueIDEQ applies the EQ predicate on the "value_id" field.
func ValueIDEQ(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValueID), v))
	})
}

// ValueIDNEQ applies the NEQ predicate on the "value_id" field.
func ValueIDNEQ(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValueID), v))
	})
}

// ValueIDIn applies the In predicate on the "value_id" field.
func ValueIDIn(vs ...int) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValueID), v...))
	})
}

// ValueIDNotIn applies the NotIn predicate on the "value_id" field.
func ValueIDNotIn(vs ...int) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValueID), v...))
	})
}

// ValueIDGT applies the GT predicate on the "value_id" field.
func ValueIDGT(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValueID), v))
	})
}

// ValueIDGTE applies the GTE predicate on the "value_id" field.
func ValueIDGTE(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValueID), v))
	})
}

// ValueIDLT applies the LT predicate on the "value_id" field.
func ValueIDLT(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValueID), v))
	})
}

// ValueIDLTE applies the LTE predicate on the "value_id" field.
func ValueIDLTE(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValueID), v))
	})
}

// ValueIDIsNil applies the IsNil predicate on the "value_id" field.
func ValueIDIsNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValueID)))
	})
}

// ValueIDNotNil applies the NotNil predicate on the "value_id" field.
func ValueIDNotNil() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValueID)))
	})
}

// SortEQ applies the EQ predicate on the "sort" field.
func SortEQ(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSort), v))
	})
}

// SortNEQ applies the NEQ predicate on the "sort" field.
func SortNEQ(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSort), v))
	})
}

// SortIn applies the In predicate on the "sort" field.
func SortIn(vs ...int) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSort), v...))
	})
}

// SortNotIn applies the NotIn predicate on the "sort" field.
func SortNotIn(vs ...int) predicate.ProductSpecsItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSort), v...))
	})
}

// SortGT applies the GT predicate on the "sort" field.
func SortGT(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSort), v))
	})
}

// SortGTE applies the GTE predicate on the "sort" field.
func SortGTE(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSort), v))
	})
}

// SortLT applies the LT predicate on the "sort" field.
func SortLT(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSort), v))
	})
}

// SortLTE applies the LTE predicate on the "sort" field.
func SortLTE(v int) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSort), v))
	})
}

// HasSpecs applies the HasEdge predicate on the "specs" edge.
func HasSpecs() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpecsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpecsTable, SpecsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpecsWith applies the HasEdge predicate on the "specs" edge with a given conditions (other predicates).
func HasSpecsWith(preds ...predicate.ProductSpecs) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SpecsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpecsTable, SpecsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasValues applies the HasEdge predicate on the "values" edge.
func HasValues() predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ValuesTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ValuesTable, ValuesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasValuesWith applies the HasEdge predicate on the "values" edge with a given conditions (other predicates).
func HasValuesWith(preds ...predicate.ProductAttributeValue) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ValuesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ValuesTable, ValuesPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductSpecsItem) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductSpecsItem) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductSpecsItem) predicate.ProductSpecsItem {
	return predicate.ProductSpecsItem(func(s *sql.Selector) {
		p(s.Not())
	})
}