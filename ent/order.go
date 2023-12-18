// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/ramnath.1998/vyaaparjunction/ent/order"
	"github.com/ramnath.1998/vyaaparjunction/ent/user"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId int `json:"userId,omitempty"`
	// Identifier holds the value of the "identifier" field.
	Identifier uuid.UUID `json:"identifier,omitempty"`
	// CreatedOn holds the value of the "createdOn" field.
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges        OrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldUserId:
			values[i] = new(sql.NullInt64)
		case order.FieldCreatedOn:
			values[i] = new(sql.NullTime)
		case order.FieldIdentifier:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldUserId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value.Valid {
				o.UserId = int(value.Int64)
			}
		case order.FieldIdentifier:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field identifier", values[i])
			} else if value != nil {
				o.Identifier = *value
			}
		case order.FieldCreatedOn:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdOn", values[i])
			} else if value.Valid {
				o.CreatedOn = value.Time
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Order entity.
func (o *Order) QueryUser() *UserQuery {
	return NewOrderClient(o.config).QueryUser(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("userId=")
	builder.WriteString(fmt.Sprintf("%v", o.UserId))
	builder.WriteString(", ")
	builder.WriteString("identifier=")
	builder.WriteString(fmt.Sprintf("%v", o.Identifier))
	builder.WriteString(", ")
	builder.WriteString("createdOn=")
	builder.WriteString(o.CreatedOn.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order