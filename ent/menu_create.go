// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"orse/ent/menu"
	"orse/ent/property"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetParentID sets the "parent_id" field.
func (mc *MenuCreate) SetParentID(i int) *MenuCreate {
	mc.mutation.SetParentID(i)
	return mc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableParentID(i *int) *MenuCreate {
	if i != nil {
		mc.SetParentID(*i)
	}
	return mc
}

// SetTree sets the "tree" field.
func (mc *MenuCreate) SetTree(s string) *MenuCreate {
	mc.mutation.SetTree(s)
	return mc
}

// SetNillableTree sets the "tree" field if the given value is not nil.
func (mc *MenuCreate) SetNillableTree(s *string) *MenuCreate {
	if s != nil {
		mc.SetTree(*s)
	}
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *MenuCreate) SetCreatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MenuCreate) SetUpdatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetTitle sets the "title" field.
func (mc *MenuCreate) SetTitle(s string) *MenuCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mc *MenuCreate) SetNillableName(s *string) *MenuCreate {
	if s != nil {
		mc.SetName(*s)
	}
	return mc
}

// SetURL sets the "url" field.
func (mc *MenuCreate) SetURL(s string) *MenuCreate {
	mc.mutation.SetURL(s)
	return mc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (mc *MenuCreate) SetNillableURL(s *string) *MenuCreate {
	if s != nil {
		mc.SetURL(*s)
	}
	return mc
}

// SetLevel sets the "level" field.
func (mc *MenuCreate) SetLevel(i int) *MenuCreate {
	mc.mutation.SetLevel(i)
	return mc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (mc *MenuCreate) SetNillableLevel(i *int) *MenuCreate {
	if i != nil {
		mc.SetLevel(*i)
	}
	return mc
}

// SetSort sets the "sort" field.
func (mc *MenuCreate) SetSort(i int) *MenuCreate {
	mc.mutation.SetSort(i)
	return mc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mc *MenuCreate) SetNillableSort(i *int) *MenuCreate {
	if i != nil {
		mc.SetSort(*i)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MenuCreate) SetStatus(pr property.Status) *MenuCreate {
	mc.mutation.SetStatus(pr)
	return mc
}

// SetIcon sets the "icon" field.
func (mc *MenuCreate) SetIcon(s string) *MenuCreate {
	mc.mutation.SetIcon(s)
	return mc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIcon(s *string) *MenuCreate {
	if s != nil {
		mc.SetIcon(*s)
	}
	return mc
}

// SetHidden sets the "hidden" field.
func (mc *MenuCreate) SetHidden(pr property.Status) *MenuCreate {
	mc.mutation.SetHidden(pr)
	return mc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (mc *MenuCreate) SetNillableHidden(pr *property.Status) *MenuCreate {
	if pr != nil {
		mc.SetHidden(*pr)
	}
	return mc
}

// SetDesc sets the "desc" field.
func (mc *MenuCreate) SetDesc(s string) *MenuCreate {
	mc.mutation.SetDesc(s)
	return mc
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDesc(s *string) *MenuCreate {
	if s != nil {
		mc.SetDesc(*s)
	}
	return mc
}

// SetParent sets the "parent" edge to the Menu entity.
func (mc *MenuCreate) SetParent(m *Menu) *MenuCreate {
	return mc.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mc *MenuCreate) AddChildIDs(ids ...int) *MenuCreate {
	mc.mutation.AddChildIDs(ids...)
	return mc
}

// AddChildren adds the "children" edges to the Menu entity.
func (mc *MenuCreate) AddChildren(m ...*Menu) *MenuCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddChildIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	var (
		err  error
		node *Menu
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := menu.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := menu.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Level(); !ok {
		v := menu.DefaultLevel
		mc.mutation.SetLevel(v)
	}
	if _, ok := mc.mutation.Sort(); !ok {
		v := menu.DefaultSort
		mc.mutation.SetSort(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	if _, ok := mc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if _, ok := mc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "level"`)}
	}
	if _, ok := mc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "sort"`)}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := mc.mutation.Status(); ok {
		if err := menu.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if v, ok := mc.mutation.Hidden(); ok {
		if err := menu.HiddenValidator(v); err != nil {
			return &ValidationError{Name: "hidden", err: fmt.Errorf(`ent: validator failed for field "hidden": %w`, err)}
		}
	}
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: menu.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Tree(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTree,
		})
		_node.Tree = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menu.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: menu.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
		_node.Name = value
	}
	if value, ok := mc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := mc.mutation.Level(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldLevel,
		})
		_node.Level = value
	}
	if value, ok := mc.mutation.Sort(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: menu.FieldSort,
		})
		_node.Sort = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: menu.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := mc.mutation.Icon(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldIcon,
		})
		_node.Icon = value
	}
	if value, ok := mc.mutation.Hidden(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: menu.FieldHidden,
		})
		_node.Hidden = value
	}
	if value, ok := mc.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldDesc,
		})
		_node.Desc = value
	}
	if nodes := mc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
