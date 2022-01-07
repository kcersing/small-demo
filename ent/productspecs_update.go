// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"orse/ent/predicate"
	"orse/ent/product"
	"orse/ent/productspecs"
	"orse/ent/productspecsitem"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductSpecsUpdate is the builder for updating ProductSpecs entities.
type ProductSpecsUpdate struct {
	config
	hooks    []Hook
	mutation *ProductSpecsMutation
}

// Where appends a list predicates to the ProductSpecsUpdate builder.
func (psu *ProductSpecsUpdate) Where(ps ...predicate.ProductSpecs) *ProductSpecsUpdate {
	psu.mutation.Where(ps...)
	return psu
}

// SetCreatedAt sets the "created_at" field.
func (psu *ProductSpecsUpdate) SetCreatedAt(t time.Time) *ProductSpecsUpdate {
	psu.mutation.SetCreatedAt(t)
	return psu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (psu *ProductSpecsUpdate) SetNillableCreatedAt(t *time.Time) *ProductSpecsUpdate {
	if t != nil {
		psu.SetCreatedAt(*t)
	}
	return psu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (psu *ProductSpecsUpdate) ClearCreatedAt() *ProductSpecsUpdate {
	psu.mutation.ClearCreatedAt()
	return psu
}

// SetUpdatedAt sets the "updated_at" field.
func (psu *ProductSpecsUpdate) SetUpdatedAt(t time.Time) *ProductSpecsUpdate {
	psu.mutation.SetUpdatedAt(t)
	return psu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (psu *ProductSpecsUpdate) ClearUpdatedAt() *ProductSpecsUpdate {
	psu.mutation.ClearUpdatedAt()
	return psu
}

// SetProductID sets the "product_id" field.
func (psu *ProductSpecsUpdate) SetProductID(i int) *ProductSpecsUpdate {
	psu.mutation.SetProductID(i)
	return psu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (psu *ProductSpecsUpdate) SetNillableProductID(i *int) *ProductSpecsUpdate {
	if i != nil {
		psu.SetProductID(*i)
	}
	return psu
}

// ClearProductID clears the value of the "product_id" field.
func (psu *ProductSpecsUpdate) ClearProductID() *ProductSpecsUpdate {
	psu.mutation.ClearProductID()
	return psu
}

// SetName sets the "name" field.
func (psu *ProductSpecsUpdate) SetName(s string) *ProductSpecsUpdate {
	psu.mutation.SetName(s)
	return psu
}

// SetSn sets the "sn" field.
func (psu *ProductSpecsUpdate) SetSn(s string) *ProductSpecsUpdate {
	psu.mutation.SetSn(s)
	return psu
}

// SetStock sets the "stock" field.
func (psu *ProductSpecsUpdate) SetStock(i int) *ProductSpecsUpdate {
	psu.mutation.ResetStock()
	psu.mutation.SetStock(i)
	return psu
}

// AddStock adds i to the "stock" field.
func (psu *ProductSpecsUpdate) AddStock(i int) *ProductSpecsUpdate {
	psu.mutation.AddStock(i)
	return psu
}

// SetSales sets the "sales" field.
func (psu *ProductSpecsUpdate) SetSales(i int) *ProductSpecsUpdate {
	psu.mutation.ResetSales()
	psu.mutation.SetSales(i)
	return psu
}

// AddSales adds i to the "sales" field.
func (psu *ProductSpecsUpdate) AddSales(i int) *ProductSpecsUpdate {
	psu.mutation.AddSales(i)
	return psu
}

// SetPrice sets the "price" field.
func (psu *ProductSpecsUpdate) SetPrice(f float64) *ProductSpecsUpdate {
	psu.mutation.ResetPrice()
	psu.mutation.SetPrice(f)
	return psu
}

// AddPrice adds f to the "price" field.
func (psu *ProductSpecsUpdate) AddPrice(f float64) *ProductSpecsUpdate {
	psu.mutation.AddPrice(f)
	return psu
}

// SetSalePrice sets the "sale_price" field.
func (psu *ProductSpecsUpdate) SetSalePrice(f float64) *ProductSpecsUpdate {
	psu.mutation.ResetSalePrice()
	psu.mutation.SetSalePrice(f)
	return psu
}

// AddSalePrice adds f to the "sale_price" field.
func (psu *ProductSpecsUpdate) AddSalePrice(f float64) *ProductSpecsUpdate {
	psu.mutation.AddSalePrice(f)
	return psu
}

// SetCreateID sets the "create_id" field.
func (psu *ProductSpecsUpdate) SetCreateID(i int) *ProductSpecsUpdate {
	psu.mutation.ResetCreateID()
	psu.mutation.SetCreateID(i)
	return psu
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (psu *ProductSpecsUpdate) SetNillableCreateID(i *int) *ProductSpecsUpdate {
	if i != nil {
		psu.SetCreateID(*i)
	}
	return psu
}

// AddCreateID adds i to the "create_id" field.
func (psu *ProductSpecsUpdate) AddCreateID(i int) *ProductSpecsUpdate {
	psu.mutation.AddCreateID(i)
	return psu
}

// ClearCreateID clears the value of the "create_id" field.
func (psu *ProductSpecsUpdate) ClearCreateID() *ProductSpecsUpdate {
	psu.mutation.ClearCreateID()
	return psu
}

// SetProduct sets the "product" edge to the Product entity.
func (psu *ProductSpecsUpdate) SetProduct(p *Product) *ProductSpecsUpdate {
	return psu.SetProductID(p.ID)
}

// AddItemIDs adds the "items" edge to the ProductSpecsItem entity by IDs.
func (psu *ProductSpecsUpdate) AddItemIDs(ids ...int) *ProductSpecsUpdate {
	psu.mutation.AddItemIDs(ids...)
	return psu
}

// AddItems adds the "items" edges to the ProductSpecsItem entity.
func (psu *ProductSpecsUpdate) AddItems(p ...*ProductSpecsItem) *ProductSpecsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psu.AddItemIDs(ids...)
}

// Mutation returns the ProductSpecsMutation object of the builder.
func (psu *ProductSpecsUpdate) Mutation() *ProductSpecsMutation {
	return psu.mutation
}

// ClearProduct clears the "product" edge to the Product entity.
func (psu *ProductSpecsUpdate) ClearProduct() *ProductSpecsUpdate {
	psu.mutation.ClearProduct()
	return psu
}

// ClearItems clears all "items" edges to the ProductSpecsItem entity.
func (psu *ProductSpecsUpdate) ClearItems() *ProductSpecsUpdate {
	psu.mutation.ClearItems()
	return psu
}

// RemoveItemIDs removes the "items" edge to ProductSpecsItem entities by IDs.
func (psu *ProductSpecsUpdate) RemoveItemIDs(ids ...int) *ProductSpecsUpdate {
	psu.mutation.RemoveItemIDs(ids...)
	return psu
}

// RemoveItems removes "items" edges to ProductSpecsItem entities.
func (psu *ProductSpecsUpdate) RemoveItems(p ...*ProductSpecsItem) *ProductSpecsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psu.RemoveItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (psu *ProductSpecsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	psu.defaults()
	if len(psu.hooks) == 0 {
		affected, err = psu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductSpecsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psu.mutation = mutation
			affected, err = psu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(psu.hooks) - 1; i >= 0; i-- {
			if psu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (psu *ProductSpecsUpdate) SaveX(ctx context.Context) int {
	affected, err := psu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (psu *ProductSpecsUpdate) Exec(ctx context.Context) error {
	_, err := psu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psu *ProductSpecsUpdate) ExecX(ctx context.Context) {
	if err := psu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psu *ProductSpecsUpdate) defaults() {
	if _, ok := psu.mutation.UpdatedAt(); !ok && !psu.mutation.UpdatedAtCleared() {
		v := productspecs.UpdateDefaultUpdatedAt()
		psu.mutation.SetUpdatedAt(v)
	}
}

func (psu *ProductSpecsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productspecs.Table,
			Columns: productspecs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productspecs.FieldID,
			},
		},
	}
	if ps := psu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productspecs.FieldCreatedAt,
		})
	}
	if psu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productspecs.FieldCreatedAt,
		})
	}
	if value, ok := psu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productspecs.FieldUpdatedAt,
		})
	}
	if psu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productspecs.FieldUpdatedAt,
		})
	}
	if value, ok := psu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productspecs.FieldName,
		})
	}
	if value, ok := psu.mutation.Sn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productspecs.FieldSn,
		})
	}
	if value, ok := psu.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldStock,
		})
	}
	if value, ok := psu.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldStock,
		})
	}
	if value, ok := psu.mutation.Sales(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldSales,
		})
	}
	if value, ok := psu.mutation.AddedSales(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldSales,
		})
	}
	if value, ok := psu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldPrice,
		})
	}
	if value, ok := psu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldPrice,
		})
	}
	if value, ok := psu.mutation.SalePrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldSalePrice,
		})
	}
	if value, ok := psu.mutation.AddedSalePrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldSalePrice,
		})
	}
	if value, ok := psu.mutation.CreateID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldCreateID,
		})
	}
	if value, ok := psu.mutation.AddedCreateID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldCreateID,
		})
	}
	if psu.mutation.CreateIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: productspecs.FieldCreateID,
		})
	}
	if psu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productspecs.ProductTable,
			Columns: []string{productspecs.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productspecs.ProductTable,
			Columns: []string{productspecs.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productspecs.ItemsTable,
			Columns: []string{productspecs.ItemsColumn},
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
	if nodes := psu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !psu.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productspecs.ItemsTable,
			Columns: []string{productspecs.ItemsColumn},
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
	if nodes := psu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productspecs.ItemsTable,
			Columns: []string{productspecs.ItemsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, psu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productspecs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProductSpecsUpdateOne is the builder for updating a single ProductSpecs entity.
type ProductSpecsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductSpecsMutation
}

// SetCreatedAt sets the "created_at" field.
func (psuo *ProductSpecsUpdateOne) SetCreatedAt(t time.Time) *ProductSpecsUpdateOne {
	psuo.mutation.SetCreatedAt(t)
	return psuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (psuo *ProductSpecsUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductSpecsUpdateOne {
	if t != nil {
		psuo.SetCreatedAt(*t)
	}
	return psuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (psuo *ProductSpecsUpdateOne) ClearCreatedAt() *ProductSpecsUpdateOne {
	psuo.mutation.ClearCreatedAt()
	return psuo
}

// SetUpdatedAt sets the "updated_at" field.
func (psuo *ProductSpecsUpdateOne) SetUpdatedAt(t time.Time) *ProductSpecsUpdateOne {
	psuo.mutation.SetUpdatedAt(t)
	return psuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (psuo *ProductSpecsUpdateOne) ClearUpdatedAt() *ProductSpecsUpdateOne {
	psuo.mutation.ClearUpdatedAt()
	return psuo
}

// SetProductID sets the "product_id" field.
func (psuo *ProductSpecsUpdateOne) SetProductID(i int) *ProductSpecsUpdateOne {
	psuo.mutation.SetProductID(i)
	return psuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (psuo *ProductSpecsUpdateOne) SetNillableProductID(i *int) *ProductSpecsUpdateOne {
	if i != nil {
		psuo.SetProductID(*i)
	}
	return psuo
}

// ClearProductID clears the value of the "product_id" field.
func (psuo *ProductSpecsUpdateOne) ClearProductID() *ProductSpecsUpdateOne {
	psuo.mutation.ClearProductID()
	return psuo
}

// SetName sets the "name" field.
func (psuo *ProductSpecsUpdateOne) SetName(s string) *ProductSpecsUpdateOne {
	psuo.mutation.SetName(s)
	return psuo
}

// SetSn sets the "sn" field.
func (psuo *ProductSpecsUpdateOne) SetSn(s string) *ProductSpecsUpdateOne {
	psuo.mutation.SetSn(s)
	return psuo
}

// SetStock sets the "stock" field.
func (psuo *ProductSpecsUpdateOne) SetStock(i int) *ProductSpecsUpdateOne {
	psuo.mutation.ResetStock()
	psuo.mutation.SetStock(i)
	return psuo
}

// AddStock adds i to the "stock" field.
func (psuo *ProductSpecsUpdateOne) AddStock(i int) *ProductSpecsUpdateOne {
	psuo.mutation.AddStock(i)
	return psuo
}

// SetSales sets the "sales" field.
func (psuo *ProductSpecsUpdateOne) SetSales(i int) *ProductSpecsUpdateOne {
	psuo.mutation.ResetSales()
	psuo.mutation.SetSales(i)
	return psuo
}

// AddSales adds i to the "sales" field.
func (psuo *ProductSpecsUpdateOne) AddSales(i int) *ProductSpecsUpdateOne {
	psuo.mutation.AddSales(i)
	return psuo
}

// SetPrice sets the "price" field.
func (psuo *ProductSpecsUpdateOne) SetPrice(f float64) *ProductSpecsUpdateOne {
	psuo.mutation.ResetPrice()
	psuo.mutation.SetPrice(f)
	return psuo
}

// AddPrice adds f to the "price" field.
func (psuo *ProductSpecsUpdateOne) AddPrice(f float64) *ProductSpecsUpdateOne {
	psuo.mutation.AddPrice(f)
	return psuo
}

// SetSalePrice sets the "sale_price" field.
func (psuo *ProductSpecsUpdateOne) SetSalePrice(f float64) *ProductSpecsUpdateOne {
	psuo.mutation.ResetSalePrice()
	psuo.mutation.SetSalePrice(f)
	return psuo
}

// AddSalePrice adds f to the "sale_price" field.
func (psuo *ProductSpecsUpdateOne) AddSalePrice(f float64) *ProductSpecsUpdateOne {
	psuo.mutation.AddSalePrice(f)
	return psuo
}

// SetCreateID sets the "create_id" field.
func (psuo *ProductSpecsUpdateOne) SetCreateID(i int) *ProductSpecsUpdateOne {
	psuo.mutation.ResetCreateID()
	psuo.mutation.SetCreateID(i)
	return psuo
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (psuo *ProductSpecsUpdateOne) SetNillableCreateID(i *int) *ProductSpecsUpdateOne {
	if i != nil {
		psuo.SetCreateID(*i)
	}
	return psuo
}

// AddCreateID adds i to the "create_id" field.
func (psuo *ProductSpecsUpdateOne) AddCreateID(i int) *ProductSpecsUpdateOne {
	psuo.mutation.AddCreateID(i)
	return psuo
}

// ClearCreateID clears the value of the "create_id" field.
func (psuo *ProductSpecsUpdateOne) ClearCreateID() *ProductSpecsUpdateOne {
	psuo.mutation.ClearCreateID()
	return psuo
}

// SetProduct sets the "product" edge to the Product entity.
func (psuo *ProductSpecsUpdateOne) SetProduct(p *Product) *ProductSpecsUpdateOne {
	return psuo.SetProductID(p.ID)
}

// AddItemIDs adds the "items" edge to the ProductSpecsItem entity by IDs.
func (psuo *ProductSpecsUpdateOne) AddItemIDs(ids ...int) *ProductSpecsUpdateOne {
	psuo.mutation.AddItemIDs(ids...)
	return psuo
}

// AddItems adds the "items" edges to the ProductSpecsItem entity.
func (psuo *ProductSpecsUpdateOne) AddItems(p ...*ProductSpecsItem) *ProductSpecsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psuo.AddItemIDs(ids...)
}

// Mutation returns the ProductSpecsMutation object of the builder.
func (psuo *ProductSpecsUpdateOne) Mutation() *ProductSpecsMutation {
	return psuo.mutation
}

// ClearProduct clears the "product" edge to the Product entity.
func (psuo *ProductSpecsUpdateOne) ClearProduct() *ProductSpecsUpdateOne {
	psuo.mutation.ClearProduct()
	return psuo
}

// ClearItems clears all "items" edges to the ProductSpecsItem entity.
func (psuo *ProductSpecsUpdateOne) ClearItems() *ProductSpecsUpdateOne {
	psuo.mutation.ClearItems()
	return psuo
}

// RemoveItemIDs removes the "items" edge to ProductSpecsItem entities by IDs.
func (psuo *ProductSpecsUpdateOne) RemoveItemIDs(ids ...int) *ProductSpecsUpdateOne {
	psuo.mutation.RemoveItemIDs(ids...)
	return psuo
}

// RemoveItems removes "items" edges to ProductSpecsItem entities.
func (psuo *ProductSpecsUpdateOne) RemoveItems(p ...*ProductSpecsItem) *ProductSpecsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psuo.RemoveItemIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (psuo *ProductSpecsUpdateOne) Select(field string, fields ...string) *ProductSpecsUpdateOne {
	psuo.fields = append([]string{field}, fields...)
	return psuo
}

// Save executes the query and returns the updated ProductSpecs entity.
func (psuo *ProductSpecsUpdateOne) Save(ctx context.Context) (*ProductSpecs, error) {
	var (
		err  error
		node *ProductSpecs
	)
	psuo.defaults()
	if len(psuo.hooks) == 0 {
		node, err = psuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductSpecsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			psuo.mutation = mutation
			node, err = psuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(psuo.hooks) - 1; i >= 0; i-- {
			if psuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (psuo *ProductSpecsUpdateOne) SaveX(ctx context.Context) *ProductSpecs {
	node, err := psuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (psuo *ProductSpecsUpdateOne) Exec(ctx context.Context) error {
	_, err := psuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psuo *ProductSpecsUpdateOne) ExecX(ctx context.Context) {
	if err := psuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psuo *ProductSpecsUpdateOne) defaults() {
	if _, ok := psuo.mutation.UpdatedAt(); !ok && !psuo.mutation.UpdatedAtCleared() {
		v := productspecs.UpdateDefaultUpdatedAt()
		psuo.mutation.SetUpdatedAt(v)
	}
}

func (psuo *ProductSpecsUpdateOne) sqlSave(ctx context.Context) (_node *ProductSpecs, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productspecs.Table,
			Columns: productspecs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: productspecs.FieldID,
			},
		},
	}
	id, ok := psuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ProductSpecs.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := psuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productspecs.FieldID)
		for _, f := range fields {
			if !productspecs.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productspecs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := psuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := psuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productspecs.FieldCreatedAt,
		})
	}
	if psuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productspecs.FieldCreatedAt,
		})
	}
	if value, ok := psuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: productspecs.FieldUpdatedAt,
		})
	}
	if psuo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: productspecs.FieldUpdatedAt,
		})
	}
	if value, ok := psuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productspecs.FieldName,
		})
	}
	if value, ok := psuo.mutation.Sn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: productspecs.FieldSn,
		})
	}
	if value, ok := psuo.mutation.Stock(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldStock,
		})
	}
	if value, ok := psuo.mutation.AddedStock(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldStock,
		})
	}
	if value, ok := psuo.mutation.Sales(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldSales,
		})
	}
	if value, ok := psuo.mutation.AddedSales(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldSales,
		})
	}
	if value, ok := psuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldPrice,
		})
	}
	if value, ok := psuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldPrice,
		})
	}
	if value, ok := psuo.mutation.SalePrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldSalePrice,
		})
	}
	if value, ok := psuo.mutation.AddedSalePrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: productspecs.FieldSalePrice,
		})
	}
	if value, ok := psuo.mutation.CreateID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldCreateID,
		})
	}
	if value, ok := psuo.mutation.AddedCreateID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: productspecs.FieldCreateID,
		})
	}
	if psuo.mutation.CreateIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: productspecs.FieldCreateID,
		})
	}
	if psuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productspecs.ProductTable,
			Columns: []string{productspecs.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := psuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productspecs.ProductTable,
			Columns: []string{productspecs.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if psuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productspecs.ItemsTable,
			Columns: []string{productspecs.ItemsColumn},
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
	if nodes := psuo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !psuo.mutation.ItemsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productspecs.ItemsTable,
			Columns: []string{productspecs.ItemsColumn},
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
	if nodes := psuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productspecs.ItemsTable,
			Columns: []string{productspecs.ItemsColumn},
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
	_node = &ProductSpecs{config: psuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, psuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productspecs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}