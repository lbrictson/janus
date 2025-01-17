// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/janus/ent/audit"
)

// Audit is the model entity for the Audit schema.
type Audit struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Audit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case audit.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Audit fields.
func (a *Audit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case audit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Audit.
// This includes values selected through modifiers, order, etc.
func (a *Audit) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Audit.
// Note that you need to call Audit.Unwrap() before calling this method if this Audit
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Audit) Update() *AuditUpdateOne {
	return NewAuditClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Audit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Audit) Unwrap() *Audit {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Audit is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Audit) String() string {
	var builder strings.Builder
	builder.WriteString("Audit(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Audits is a parsable slice of Audit.
type Audits []*Audit
