// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// ProductName applies equality check predicate on the "productName" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductName, v))
}

// Identifier applies equality check predicate on the "Identifier" field. It's identical to IdentifierEQ.
func Identifier(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldIdentifier, v))
}

// CreatedOn applies equality check predicate on the "createdOn" field. It's identical to CreatedOnEQ.
func CreatedOn(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedOn, v))
}

// CategoryId applies equality check predicate on the "categoryId" field. It's identical to CategoryIdEQ.
func CategoryId(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCategoryId, v))
}

// ProductNameEQ applies the EQ predicate on the "productName" field.
func ProductNameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldProductName, v))
}

// ProductNameNEQ applies the NEQ predicate on the "productName" field.
func ProductNameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldProductName, v))
}

// ProductNameIn applies the In predicate on the "productName" field.
func ProductNameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldProductName, vs...))
}

// ProductNameNotIn applies the NotIn predicate on the "productName" field.
func ProductNameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldProductName, vs...))
}

// ProductNameGT applies the GT predicate on the "productName" field.
func ProductNameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldProductName, v))
}

// ProductNameGTE applies the GTE predicate on the "productName" field.
func ProductNameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldProductName, v))
}

// ProductNameLT applies the LT predicate on the "productName" field.
func ProductNameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldProductName, v))
}

// ProductNameLTE applies the LTE predicate on the "productName" field.
func ProductNameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldProductName, v))
}

// ProductNameContains applies the Contains predicate on the "productName" field.
func ProductNameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldProductName, v))
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "productName" field.
func ProductNameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldProductName, v))
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "productName" field.
func ProductNameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldProductName, v))
}

// ProductNameEqualFold applies the EqualFold predicate on the "productName" field.
func ProductNameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldProductName, v))
}

// ProductNameContainsFold applies the ContainsFold predicate on the "productName" field.
func ProductNameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldProductName, v))
}

// IdentifierEQ applies the EQ predicate on the "Identifier" field.
func IdentifierEQ(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "Identifier" field.
func IdentifierNEQ(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "Identifier" field.
func IdentifierIn(vs ...uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "Identifier" field.
func IdentifierNotIn(vs ...uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "Identifier" field.
func IdentifierGT(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "Identifier" field.
func IdentifierGTE(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "Identifier" field.
func IdentifierLT(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "Identifier" field.
func IdentifierLTE(v uuid.UUID) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldIdentifier, v))
}

// CreatedOnEQ applies the EQ predicate on the "createdOn" field.
func CreatedOnEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreatedOn, v))
}

// CreatedOnNEQ applies the NEQ predicate on the "createdOn" field.
func CreatedOnNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreatedOn, v))
}

// CreatedOnIn applies the In predicate on the "createdOn" field.
func CreatedOnIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreatedOn, vs...))
}

// CreatedOnNotIn applies the NotIn predicate on the "createdOn" field.
func CreatedOnNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreatedOn, vs...))
}

// CreatedOnGT applies the GT predicate on the "createdOn" field.
func CreatedOnGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreatedOn, v))
}

// CreatedOnGTE applies the GTE predicate on the "createdOn" field.
func CreatedOnGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreatedOn, v))
}

// CreatedOnLT applies the LT predicate on the "createdOn" field.
func CreatedOnLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreatedOn, v))
}

// CreatedOnLTE applies the LTE predicate on the "createdOn" field.
func CreatedOnLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreatedOn, v))
}

// CategoryIdEQ applies the EQ predicate on the "categoryId" field.
func CategoryIdEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCategoryId, v))
}

// CategoryIdNEQ applies the NEQ predicate on the "categoryId" field.
func CategoryIdNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCategoryId, v))
}

// CategoryIdIn applies the In predicate on the "categoryId" field.
func CategoryIdIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCategoryId, vs...))
}

// CategoryIdNotIn applies the NotIn predicate on the "categoryId" field.
func CategoryIdNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCategoryId, vs...))
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.ProductCategory) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
