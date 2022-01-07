// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"orse/ent/predicate"
	"orse/ent/productattributekey"
	"orse/ent/productattributevalue"
	"orse/ent/productspecsitem"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductAttributeValueUpdate is the builder for updating ProductAttributeValue entities.
type ProductAttributeValueUpdate struct {
	config
	hooks    []Hook
	mutation *ProductAttributeValueMutation
}

// Where appends a list predicates to the ProductAttributeValueUpdate builder.
func (pavu *ProductAttributeValueUpdate) Where(ps ...predicate.ProductAttributeValue) *ProductAttributeValueUpdate {
	pavu.mutation.Where(ps...)
	return pavu
}

// SetCreatedAt sets the "created_at" field.
func (pavu *ProductAttributeValueUpdate) SetCreatedAt(t time.Time) *ProductAttributeValueUpdate {
	pavu.mutation.SetCreatedAt(t)
	return pavu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pavu *ProductAttributeValueUpdate) SetNillableCreatedAt(t *time.Time) *ProductAttributeValueUpdate {
	if t != nil {
		pavu.SetCreatedAt(*t)
	}
	return pavu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (pavu *ProductAttributeValueUpdate) ClearCreatedAt() *ProductAttributeValueUpdate {
	pavu.mutation.ClearCreatedAt()
	return pavu
}

// SetUpdatedAt sets the "updated_at" field.
func (pavu *ProductAttributeValueUpdate) SetUpdatedAt(t time.Time) *ProductAttributeValueUpdate {
	pavu.mutation.SetUpdatedAt(t)
	return pavu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pavu *ProductAttributeValueUpdate) ClearUpdatedAt() *ProductAttributeValueUpdate {
	pavu.mutation.ClearUpdatedAt()
	return pavu
}

// SetKeyID sets the "key_id" field.
func (pavu *ProductAttributeValueUpdate) SetKeyID(i int) *ProductAttributeValueUpdate {
	pavu.mutation.SetKeyID(i)
	return pavu
}

// SetNillableKeyID sets the "key_id" field if the given value is not nil.
func (pavu *ProductAttributeValueUpdate) SetNillableKeyID(i *int) *ProductAttributeValueUpdate {
	if i != nil {
		pavu.SetKeyID(*i)
	}
	return pavu
}

// ClearKeyID clears the value of the "key_id" field.
func (pavu *ProductAttributeValueUpdate) ClearKeyID() *ProductAttributeValueUpdate {
	pavu.mutation.ClearKeyID()
	return pavu
}

// SetValue sets the "value" field.
func (pavu *ProductAttributeValueUpdate) SetValue(s string) *ProductAttributeValueUpdate {
	pavu.mutation.SetValue(s)
	return pavu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (pavu *ProductAttributeValueUpdate) SetNillableValue(s *string) *ProductAttributeValueUpdate {
	if s != nil {
		pavu.SetValue(*s)
	}
	return pavu
}

// ClearValue clears the value of the "value" field.
func (pavu *ProductAttributeValueUpdate) ClearValue() *ProductAttributeValueUpdate {
	pavu.mutation.ClearValue()
	return pavu
}

// SetSort sets the "sort" field.
func (pavu *ProductAttributeValueUpdate) SetSort(i int) *ProductAttributeValueUpdate {
	pavu.mutation.ResetSort()
	pavu.mutation.SetSort(i)
	return pavu
}

// AddSort adds i to the "sort" field.
func (pavu *ProductAttributeValueUpdate) AddSort(i int) *ProductAttributeValueUpdate {
	pavu.mutation.AddSort(i)
	return pavu
}

// SetKey sets the "key" edge to the ProductAttributeKey entity.
func (pavu *ProductAttributeValueUpdate) SetKey(p *ProductAttributeKey) *ProductAttributeValueUpdate {
	return pavu.SetKeyID(p.ID)
}

// AddItemIDs adds the "items" edge to the ProductSpecsItem entity by IDs.
func (pavu *ProductAttributeValueUpdate) AddItemIDs(ids ...int) *ProductAttributeValueUpdate {
	pavu.mutation.AddItemIDs(ids...)
	return pavu
}

// AddItems adds the "items" edges to the ProductSpecsItem entity.
func (pavu *ProductAttributeValueUpdate) AddItems(p ...*ProductSpecsItem) *ProductAttributeValueUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pavu.AddItemIDs(ids...)
}

// Mutation returns the ProductAttributeValueMutation object of the builder.
func (pavu *ProductAttributeValueUpdate) Mutation() *ProductAttributeValueMutation {
	return pavu.mutation
}

// ClearKey clears the "key" edge to the ProductAttributeKey entity.
func (pavu *ProductAttributeValueUpdate) ClearKey() *ProductAttributeValueUpdate {
	pavu.mutation.ClearKey()
	return pavu
}

// ClearItems clears all "items" edges to the ProductSpecsItem entity.
func (pavu *ProductAttributeValueUpdate) ClearItems() *ProductAttributeValueUpdate {
	pavu.mutation.ClearItems()
	return pavu
}

// RemoveItemIDs removes the "items" edge to ProductSpecsItem entities by IDs.
func (pavu *ProductAttributeValueUpdate) RemoveItemIDs(ids ...int) *ProductAttributeValueUpdate {
	pavu.mutation.RemoveItemIDs(ids...)
	return pavu
}

// RemoveItems removes "items" edges to ProductSpecsItem entities.
func (pavu *ProductAttributeValueUpdate) RemoveItems(p ...*ProductSpecsItem) *ProductAttributeValueUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pavu.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pavu *ProductAttributeValueUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pavu.defaults()
	if len(pavu.hooks) == 0 {
		affected, err = pavu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductAttributeValueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pavu.mutation = mutation
			affected, err = pavu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pavu.hooks) - 1; i >= 0; i-- {
			if pavu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pavu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pavu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pavu *ProductAttributeValueUpdate) SaveX(ctx context.Context) int {
	affected, err := pavu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pavu *ProductAttributeValueUpdate) Exec(ctx context.Context) error {
	_, err := pavu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pavu *ProductAttributeValueUpdate) ExecX(ctx context.Context) {
	if err := pavu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pavu *ProductAttributeValueUpdate) defaults() {
	if _, ok := pavu.mutation.UpdatedAt(); !ok && !pavu.mutation.UpdatedAtCleared() {
		v := productattributevalue.UpdateDefaultUpdatedAt()
		pavu.mutation.SetUpdatedAt(v)
	}
}

func (pavu *ProductAttributeValueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productattributevalue.Table,
			Columns: productattributevalue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productattributevalue.FieldID,
			},
		},
	}
	if ps := pavu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pavu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productattributevalue.FieldCreatedAt,
		})
	}
	if pavu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productattributevalue.FieldCreatedAt,
		})
	}
	if value, ok := pavu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productattributevalue.FieldUpdatedAt,
		})
	}
	if pavu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productattributevalue.FieldUpdatedAt,
		})
	}
	if value, ok := pavu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productattributevalue.FieldValue,
		})
	}
	if pavu.mutation.ValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productattributevalue.FieldValue,
		})
	}
	if value, ok := pavu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productattributevalue.FieldSort,
		})
	}
	if value, ok := pavu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productattributevalue.FieldSort,
		})
	}
	if pavu.mutation.KeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productattributevalue.KeyTable,
			Columns: []string{productattributevalue.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productattributekey.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pavu.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productattributevalue.KeyTable,
			Columns: []string{productattributevalue.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productattributekey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pavu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productattributevalue.ItemsTable,
			Columns: productattributevalue.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productspecsitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pavu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !pavu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productattributevalue.ItemsTable,
			Columns: productattributevalue.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productspecsitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pavu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productattributevalue.ItemsTable,
			Columns: productattributevalue.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productspecsitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pavu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productattributevalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProductAttributeValueUpdateOne is the builder for updating a single ProductAttributeValue entity.
type ProductAttributeValueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductAttributeValueMutation
}

// SetCreatedAt sets the "created_at" field.
func (pavuo *ProductAttributeValueUpdateOne) SetCreatedAt(t time.Time) *ProductAttributeValueUpdateOne {
	pavuo.mutation.SetCreatedAt(t)
	return pavuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pavuo *ProductAttributeValueUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductAttributeValueUpdateOne {
	if t != nil {
		pavuo.SetCreatedAt(*t)
	}
	return pavuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (pavuo *ProductAttributeValueUpdateOne) ClearCreatedAt() *ProductAttributeValueUpdateOne {
	pavuo.mutation.ClearCreatedAt()
	return pavuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pavuo *ProductAttributeValueUpdateOne) SetUpdatedAt(t time.Time) *ProductAttributeValueUpdateOne {
	pavuo.mutation.SetUpdatedAt(t)
	return pavuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pavuo *ProductAttributeValueUpdateOne) ClearUpdatedAt() *ProductAttributeValueUpdateOne {
	pavuo.mutation.ClearUpdatedAt()
	return pavuo
}

// SetKeyID sets the "key_id" field.
func (pavuo *ProductAttributeValueUpdateOne) SetKeyID(i int) *ProductAttributeValueUpdateOne {
	pavuo.mutation.SetKeyID(i)
	return pavuo
}

// SetNillableKeyID sets the "key_id" field if the given value is not nil.
func (pavuo *ProductAttributeValueUpdateOne) SetNillableKeyID(i *int) *ProductAttributeValueUpdateOne {
	if i != nil {
		pavuo.SetKeyID(*i)
	}
	return pavuo
}

// ClearKeyID clears the value of the "key_id" field.
func (pavuo *ProductAttributeValueUpdateOne) ClearKeyID() *ProductAttributeValueUpdateOne {
	pavuo.mutation.ClearKeyID()
	return pavuo
}

// SetValue sets the "value" field.
func (pavuo *ProductAttributeValueUpdateOne) SetValue(s string) *ProductAttributeValueUpdateOne {
	pavuo.mutation.SetValue(s)
	return pavuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (pavuo *ProductAttributeValueUpdateOne) SetNillableValue(s *string) *ProductAttributeValueUpdateOne {
	if s != nil {
		pavuo.SetValue(*s)
	}
	return pavuo
}

// ClearValue clears the value of the "value" field.
func (pavuo *ProductAttributeValueUpdateOne) ClearValue() *ProductAttributeValueUpdateOne {
	pavuo.mutation.ClearValue()
	return pavuo
}

// SetSort sets the "sort" field.
func (pavuo *ProductAttributeValueUpdateOne) SetSort(i int) *ProductAttributeValueUpdateOne {
	pavuo.mutation.ResetSort()
	pavuo.mutation.SetSort(i)
	return pavuo
}

// AddSort adds i to the "sort" field.
func (pavuo *ProductAttributeValueUpdateOne) AddSort(i int) *ProductAttributeValueUpdateOne {
	pavuo.mutation.AddSort(i)
	return pavuo
}

// SetKey sets the "key" edge to the ProductAttributeKey entity.
func (pavuo *ProductAttributeValueUpdateOne) SetKey(p *ProductAttributeKey) *ProductAttributeValueUpdateOne {
	return pavuo.SetKeyID(p.ID)
}

// AddItemIDs adds the "items" edge to the ProductSpecsItem entity by IDs.
func (pavuo *ProductAttributeValueUpdateOne) AddItemIDs(ids ...int) *ProductAttributeValueUpdateOne {
	pavuo.mutation.AddItemIDs(ids...)
	return pavuo
}

// AddItems adds the "items" edges to the ProductSpecsItem entity.
func (pavuo *ProductAttributeValueUpdateOne) AddItems(p ...*ProductSpecsItem) *ProductAttributeValueUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pavuo.AddItemIDs(ids...)
}

// Mutation returns the ProductAttributeValueMutation object of the builder.
func (pavuo *ProductAttributeValueUpdateOne) Mutation() *ProductAttributeValueMutation {
	return pavuo.mutation
}

// ClearKey clears the "key" edge to the ProductAttributeKey entity.
func (pavuo *ProductAttributeValueUpdateOne) ClearKey() *ProductAttributeValueUpdateOne {
	pavuo.mutation.ClearKey()
	return pavuo
}

// ClearItems clears all "items" edges to the ProductSpecsItem entity.
func (pavuo *ProductAttributeValueUpdateOne) ClearItems() *ProductAttributeValueUpdateOne {
	pavuo.mutation.ClearItems()
	return pavuo
}

// RemoveItemIDs removes the "items" edge to ProductSpecsItem entities by IDs.
func (pavuo *ProductAttributeValueUpdateOne) RemoveItemIDs(ids ...int) *ProductAttributeValueUpdateOne {
	pavuo.mutation.RemoveItemIDs(ids...)
	return pavuo
}

// RemoveItems removes "items" edges to ProductSpecsItem entities.
func (pavuo *ProductAttributeValueUpdateOne) RemoveItems(p ...*ProductSpecsItem) *ProductAttributeValueUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pavuo.RemoveItemIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pavuo *ProductAttributeValueUpdateOne) Select(field string, fields ...string) *ProductAttributeValueUpdateOne {
	pavuo.fields = append([]string{field}, fields...)
	return pavuo
}

// Save executes the query and returns the updated ProductAttributeValue entity.
func (pavuo *ProductAttributeValueUpdateOne) Save(ctx context.Context) (*ProductAttributeValue, error) {
	var (
		err  error
		node *ProductAttributeValue
	)
	pavuo.defaults()
	if len(pavuo.hooks) == 0 {
		node, err = pavuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductAttributeValueMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pavuo.mutation = mutation
			node, err = pavuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pavuo.hooks) - 1; i >= 0; i-- {
			if pavuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pavuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pavuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pavuo *ProductAttributeValueUpdateOne) SaveX(ctx context.Context) *ProductAttributeValue {
	node, err := pavuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pavuo *ProductAttributeValueUpdateOne) Exec(ctx context.Context) error {
	_, err := pavuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pavuo *ProductAttributeValueUpdateOne) ExecX(ctx context.Context) {
	if err := pavuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pavuo *ProductAttributeValueUpdateOne) defaults() {
	if _, ok := pavuo.mutation.UpdatedAt(); !ok && !pavuo.mutation.UpdatedAtCleared() {
		v := productattributevalue.UpdateDefaultUpdatedAt()
		pavuo.mutation.SetUpdatedAt(v)
	}
}

func (pavuo *ProductAttributeValueUpdateOne) sqlSave(ctx context.Context) (_node *ProductAttributeValue, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productattributevalue.Table,
			Columns: productattributevalue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productattributevalue.FieldID,
			},
		},
	}
	id, ok := pavuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProductAttributeValue.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := pavuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productattributevalue.FieldID)
		for _, f := range fields {
			if !productattributevalue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productattributevalue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pavuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pavuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productattributevalue.FieldCreatedAt,
		})
	}
	if pavuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productattributevalue.FieldCreatedAt,
		})
	}
	if value, ok := pavuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productattributevalue.FieldUpdatedAt,
		})
	}
	if pavuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productattributevalue.FieldUpdatedAt,
		})
	}
	if value, ok := pavuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productattributevalue.FieldValue,
		})
	}
	if pavuo.mutation.ValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: productattributevalue.FieldValue,
		})
	}
	if value, ok := pavuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productattributevalue.FieldSort,
		})
	}
	if value, ok := pavuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productattributevalue.FieldSort,
		})
	}
	if pavuo.mutation.KeyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productattributevalue.KeyTable,
			Columns: []string{productattributevalue.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productattributekey.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pavuo.mutation.KeyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productattributevalue.KeyTable,
			Columns: []string{productattributevalue.KeyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productattributekey.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pavuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productattributevalue.ItemsTable,
			Columns: productattributevalue.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productspecsitem.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pavuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !pavuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productattributevalue.ItemsTable,
			Columns: productattributevalue.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productspecsitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pavuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productattributevalue.ItemsTable,
			Columns: productattributevalue.ItemsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: productspecsitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProductAttributeValue{config: pavuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pavuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productattributevalue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}