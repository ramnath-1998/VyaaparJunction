// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/ent/productcategory"
)

// ProductCategory is the model entity for the ProductCategory schema.
type ProductCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CategoryName holds the value of the "categoryName" field.
	CategoryName string `json:"categoryName,omitempty"`
	// Identifier holds the value of the "Identifier" field.
	Identifier uuid.UUID `json:"Identifier,omitempty"`
	// CreatedOn holds the value of the "createdOn" field.
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductCategoryQuery when eager-loading is set.
	Edges        ProductCategoryEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductCategoryEdges holds the relations/edges for other nodes in the graph.
type ProductCategoryEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e ProductCategoryEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productcategory.FieldID:
			values[i] = new(sql.NullInt64)
		case productcategory.FieldCategoryName:
			values[i] = new(sql.NullString)
		case productcategory.FieldCreatedOn:
			values[i] = new(sql.NullTime)
		case productcategory.FieldIdentifier:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductCategory fields.
func (pc *ProductCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		case productcategory.FieldCategoryName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field categoryName", values[i])
			} else if value.Valid {
				pc.CategoryName = value.String
			}
		case productcategory.FieldIdentifier:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field Identifier", values[i])
			} else if value != nil {
				pc.Identifier = *value
			}
		case productcategory.FieldCreatedOn:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdOn", values[i])
			} else if value.Valid {
				pc.CreatedOn = value.Time
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductCategory.
// This includes values selected through modifiers, order, etc.
func (pc *ProductCategory) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the ProductCategory entity.
func (pc *ProductCategory) QueryProduct() *ProductQuery {
	return NewProductCategoryClient(pc.config).QueryProduct(pc)
}

// Update returns a builder for updating this ProductCategory.
// Note that you need to call ProductCategory.Unwrap() before calling this method if this ProductCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *ProductCategory) Update() *ProductCategoryUpdateOne {
	return NewProductCategoryClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the ProductCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *ProductCategory) Unwrap() *ProductCategory {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductCategory is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *ProductCategory) String() string {
	var builder strings.Builder
	builder.WriteString("ProductCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("categoryName=")
	builder.WriteString(pc.CategoryName)
	builder.WriteString(", ")
	builder.WriteString("Identifier=")
	builder.WriteString(fmt.Sprintf("%v", pc.Identifier))
	builder.WriteString(", ")
	builder.WriteString("createdOn=")
	builder.WriteString(pc.CreatedOn.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ProductCategories is a parsable slice of ProductCategory.
type ProductCategories []*ProductCategory