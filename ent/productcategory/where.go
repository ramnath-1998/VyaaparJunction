// Code generated by ent, DO NOT EDIT.

package productcategory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLTE(FieldID, id))
}

// CategoryName applies equality check predicate on the "categoryName" field. It's identical to CategoryNameEQ.
func CategoryName(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldCategoryName, v))
}

// Identifier applies equality check predicate on the "Identifier" field. It's identical to IdentifierEQ.
func Identifier(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldIdentifier, v))
}

// CreatedOn applies equality check predicate on the "createdOn" field. It's identical to CreatedOnEQ.
func CreatedOn(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldCreatedOn, v))
}

// CategoryNameEQ applies the EQ predicate on the "categoryName" field.
func CategoryNameEQ(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldCategoryName, v))
}

// CategoryNameNEQ applies the NEQ predicate on the "categoryName" field.
func CategoryNameNEQ(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNEQ(FieldCategoryName, v))
}

// CategoryNameIn applies the In predicate on the "categoryName" field.
func CategoryNameIn(vs ...string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldIn(FieldCategoryName, vs...))
}

// CategoryNameNotIn applies the NotIn predicate on the "categoryName" field.
func CategoryNameNotIn(vs ...string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNotIn(FieldCategoryName, vs...))
}

// CategoryNameGT applies the GT predicate on the "categoryName" field.
func CategoryNameGT(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGT(FieldCategoryName, v))
}

// CategoryNameGTE applies the GTE predicate on the "categoryName" field.
func CategoryNameGTE(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGTE(FieldCategoryName, v))
}

// CategoryNameLT applies the LT predicate on the "categoryName" field.
func CategoryNameLT(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLT(FieldCategoryName, v))
}

// CategoryNameLTE applies the LTE predicate on the "categoryName" field.
func CategoryNameLTE(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLTE(FieldCategoryName, v))
}

// CategoryNameContains applies the Contains predicate on the "categoryName" field.
func CategoryNameContains(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldContains(FieldCategoryName, v))
}

// CategoryNameHasPrefix applies the HasPrefix predicate on the "categoryName" field.
func CategoryNameHasPrefix(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldHasPrefix(FieldCategoryName, v))
}

// CategoryNameHasSuffix applies the HasSuffix predicate on the "categoryName" field.
func CategoryNameHasSuffix(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldHasSuffix(FieldCategoryName, v))
}

// CategoryNameEqualFold applies the EqualFold predicate on the "categoryName" field.
func CategoryNameEqualFold(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEqualFold(FieldCategoryName, v))
}

// CategoryNameContainsFold applies the ContainsFold predicate on the "categoryName" field.
func CategoryNameContainsFold(v string) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldContainsFold(FieldCategoryName, v))
}

// IdentifierEQ applies the EQ predicate on the "Identifier" field.
func IdentifierEQ(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "Identifier" field.
func IdentifierNEQ(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "Identifier" field.
func IdentifierIn(vs ...uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "Identifier" field.
func IdentifierNotIn(vs ...uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "Identifier" field.
func IdentifierGT(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "Identifier" field.
func IdentifierGTE(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "Identifier" field.
func IdentifierLT(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "Identifier" field.
func IdentifierLTE(v uuid.UUID) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLTE(FieldIdentifier, v))
}

// CreatedOnEQ applies the EQ predicate on the "createdOn" field.
func CreatedOnEQ(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldEQ(FieldCreatedOn, v))
}

// CreatedOnNEQ applies the NEQ predicate on the "createdOn" field.
func CreatedOnNEQ(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNEQ(FieldCreatedOn, v))
}

// CreatedOnIn applies the In predicate on the "createdOn" field.
func CreatedOnIn(vs ...time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldIn(FieldCreatedOn, vs...))
}

// CreatedOnNotIn applies the NotIn predicate on the "createdOn" field.
func CreatedOnNotIn(vs ...time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldNotIn(FieldCreatedOn, vs...))
}

// CreatedOnGT applies the GT predicate on the "createdOn" field.
func CreatedOnGT(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGT(FieldCreatedOn, v))
}

// CreatedOnGTE applies the GTE predicate on the "createdOn" field.
func CreatedOnGTE(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldGTE(FieldCreatedOn, v))
}

// CreatedOnLT applies the LT predicate on the "createdOn" field.
func CreatedOnLT(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLT(FieldCreatedOn, v))
}

// CreatedOnLTE applies the LTE predicate on the "createdOn" field.
func CreatedOnLTE(v time.Time) predicate.ProductCategory {
	return predicate.ProductCategory(sql.FieldLTE(FieldCreatedOn, v))
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.ProductCategory {
	return predicate.ProductCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.ProductCategory {
	return predicate.ProductCategory(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductCategory) predicate.ProductCategory {
	return predicate.ProductCategory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductCategory) predicate.ProductCategory {
	return predicate.ProductCategory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductCategory) predicate.ProductCategory {
	return predicate.ProductCategory(sql.NotPredicates(p))
}
