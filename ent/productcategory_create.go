// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/ent/product"
	"github.com/ramnath.1998/vyaaparjunction/ent/productcategory"
)

// ProductCategoryCreate is the builder for creating a ProductCategory entity.
type ProductCategoryCreate struct {
	config
	mutation *ProductCategoryMutation
	hooks    []Hook
}

// SetCategoryName sets the "categoryName" field.
func (pcc *ProductCategoryCreate) SetCategoryName(s string) *ProductCategoryCreate {
	pcc.mutation.SetCategoryName(s)
	return pcc
}

// SetIdentifier sets the "Identifier" field.
func (pcc *ProductCategoryCreate) SetIdentifier(u uuid.UUID) *ProductCategoryCreate {
	pcc.mutation.SetIdentifier(u)
	return pcc
}

// SetNillableIdentifier sets the "Identifier" field if the given value is not nil.
func (pcc *ProductCategoryCreate) SetNillableIdentifier(u *uuid.UUID) *ProductCategoryCreate {
	if u != nil {
		pcc.SetIdentifier(*u)
	}
	return pcc
}

// SetCreatedOn sets the "createdOn" field.
func (pcc *ProductCategoryCreate) SetCreatedOn(t time.Time) *ProductCategoryCreate {
	pcc.mutation.SetCreatedOn(t)
	return pcc
}

// SetNillableCreatedOn sets the "createdOn" field if the given value is not nil.
func (pcc *ProductCategoryCreate) SetNillableCreatedOn(t *time.Time) *ProductCategoryCreate {
	if t != nil {
		pcc.SetCreatedOn(*t)
	}
	return pcc
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (pcc *ProductCategoryCreate) AddProductIDs(ids ...int) *ProductCategoryCreate {
	pcc.mutation.AddProductIDs(ids...)
	return pcc
}

// AddProduct adds the "product" edges to the Product entity.
func (pcc *ProductCategoryCreate) AddProduct(p ...*Product) *ProductCategoryCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcc.AddProductIDs(ids...)
}

// Mutation returns the ProductCategoryMutation object of the builder.
func (pcc *ProductCategoryCreate) Mutation() *ProductCategoryMutation {
	return pcc.mutation
}

// Save creates the ProductCategory in the database.
func (pcc *ProductCategoryCreate) Save(ctx context.Context) (*ProductCategory, error) {
	pcc.defaults()
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *ProductCategoryCreate) SaveX(ctx context.Context) *ProductCategory {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *ProductCategoryCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *ProductCategoryCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *ProductCategoryCreate) defaults() {
	if _, ok := pcc.mutation.Identifier(); !ok {
		v := productcategory.DefaultIdentifier()
		pcc.mutation.SetIdentifier(v)
	}
	if _, ok := pcc.mutation.CreatedOn(); !ok {
		v := productcategory.DefaultCreatedOn()
		pcc.mutation.SetCreatedOn(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *ProductCategoryCreate) check() error {
	if _, ok := pcc.mutation.CategoryName(); !ok {
		return &ValidationError{Name: "categoryName", err: errors.New(`ent: missing required field "ProductCategory.categoryName"`)}
	}
	if v, ok := pcc.mutation.CategoryName(); ok {
		if err := productcategory.CategoryNameValidator(v); err != nil {
			return &ValidationError{Name: "categoryName", err: fmt.Errorf(`ent: validator failed for field "ProductCategory.categoryName": %w`, err)}
		}
	}
	if _, ok := pcc.mutation.Identifier(); !ok {
		return &ValidationError{Name: "Identifier", err: errors.New(`ent: missing required field "ProductCategory.Identifier"`)}
	}
	if _, ok := pcc.mutation.CreatedOn(); !ok {
		return &ValidationError{Name: "createdOn", err: errors.New(`ent: missing required field "ProductCategory.createdOn"`)}
	}
	return nil
}

func (pcc *ProductCategoryCreate) sqlSave(ctx context.Context) (*ProductCategory, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *ProductCategoryCreate) createSpec() (*ProductCategory, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductCategory{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(productcategory.Table, sqlgraph.NewFieldSpec(productcategory.FieldID, field.TypeInt))
	)
	if value, ok := pcc.mutation.CategoryName(); ok {
		_spec.SetField(productcategory.FieldCategoryName, field.TypeString, value)
		_node.CategoryName = value
	}
	if value, ok := pcc.mutation.Identifier(); ok {
		_spec.SetField(productcategory.FieldIdentifier, field.TypeUUID, value)
		_node.Identifier = value
	}
	if value, ok := pcc.mutation.CreatedOn(); ok {
		_spec.SetField(productcategory.FieldCreatedOn, field.TypeTime, value)
		_node.CreatedOn = value
	}
	if nodes := pcc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductTable,
			Columns: []string{productcategory.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCategoryCreateBulk is the builder for creating many ProductCategory entities in bulk.
type ProductCategoryCreateBulk struct {
	config
	err      error
	builders []*ProductCategoryCreate
}

// Save creates the ProductCategory entities in the database.
func (pccb *ProductCategoryCreateBulk) Save(ctx context.Context) ([]*ProductCategory, error) {
	if pccb.err != nil {
		return nil, pccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*ProductCategory, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductCategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *ProductCategoryCreateBulk) SaveX(ctx context.Context) []*ProductCategory {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *ProductCategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *ProductCategoryCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
