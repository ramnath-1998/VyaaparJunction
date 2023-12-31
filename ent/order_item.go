// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ramnath.1998/vyaaparjunction/ent/order_item"
)

// Order_Item is the model entity for the Order_Item schema.
type Order_Item struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order_Item) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order_item.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order_Item fields.
func (oi *Order_Item) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order_item.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order_Item.
// This includes values selected through modifiers, order, etc.
func (oi *Order_Item) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// Update returns a builder for updating this Order_Item.
// Note that you need to call Order_Item.Unwrap() before calling this method if this Order_Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *Order_Item) Update() *OrderItemUpdateOne {
	return NewOrderItemClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the Order_Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *Order_Item) Unwrap() *Order_Item {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order_Item is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *Order_Item) String() string {
	var builder strings.Builder
	builder.WriteString("Order_Item(")
	builder.WriteString(fmt.Sprintf("id=%v", oi.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Order_Items is a parsable slice of Order_Item.
type Order_Items []*Order_Item
