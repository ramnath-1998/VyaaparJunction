// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ramnath.1998/vyaaparjunction/ent/order_item"
)

// OrderItemCreate is the builder for creating a Order_Item entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the Order_Item in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*Order_Item, error) {
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *Order_Item {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*Order_Item, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*Order_Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Order_Item{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(order_item.Table, sqlgraph.NewFieldSpec(order_item.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// OrderItemCreateBulk is the builder for creating many Order_Item entities in bulk.
type OrderItemCreateBulk struct {
	config
	err      error
	builders []*OrderItemCreate
}

// Save creates the Order_Item entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*Order_Item, error) {
	if oicb.err != nil {
		return nil, oicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*Order_Item, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*Order_Item {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
