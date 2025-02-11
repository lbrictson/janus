// Code generated by ent, DO NOT EDIT.

package inboundwebhook

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the inboundwebhook type in the database.
	Label = "inbound_webhook"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeJob holds the string denoting the job edge name in mutations.
	EdgeJob = "job"
	// Table holds the table name of the inboundwebhook in the database.
	Table = "inbound_webhooks"
	// JobTable is the table that holds the job relation/edge.
	JobTable = "inbound_webhooks"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "inbound_webhook_job"
)

// Columns holds all SQL columns for inboundwebhook fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldCreatedBy,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "inbound_webhooks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"inbound_webhook_job",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the InboundWebhook queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByJobField orders the results by job field.
func ByJobField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJobStep(), sql.OrderByField(field, opts...))
	}
}
func newJobStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JobInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, JobTable, JobColumn),
	)
}
