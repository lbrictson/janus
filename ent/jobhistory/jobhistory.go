// Code generated by ent, DO NOT EDIT.

package jobhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the jobhistory type in the database.
	Label = "job_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWasSuccessful holds the string denoting the was_successful field in the database.
	FieldWasSuccessful = "was_successful"
	// FieldDurationMs holds the string denoting the duration_ms field in the database.
	FieldDurationMs = "duration_ms"
	// FieldParameters holds the string denoting the parameters field in the database.
	FieldParameters = "parameters"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldExitCode holds the string denoting the exit_code field in the database.
	FieldExitCode = "exit_code"
	// FieldTriggeredByEmail holds the string denoting the triggered_by_email field in the database.
	FieldTriggeredByEmail = "triggered_by_email"
	// FieldTriggeredByID holds the string denoting the triggered_by_id field in the database.
	FieldTriggeredByID = "triggered_by_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldTrigger holds the string denoting the trigger field in the database.
	FieldTrigger = "trigger"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeJob holds the string denoting the job edge name in mutations.
	EdgeJob = "job"
	// Table holds the table name of the jobhistory in the database.
	Table = "job_histories"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "job_histories"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_history"
	// JobTable is the table that holds the job relation/edge.
	JobTable = "job_histories"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "job_history"
)

// Columns holds all SQL columns for jobhistory fields.
var Columns = []string{
	FieldID,
	FieldWasSuccessful,
	FieldDurationMs,
	FieldParameters,
	FieldOutput,
	FieldExitCode,
	FieldTriggeredByEmail,
	FieldTriggeredByID,
	FieldCreatedAt,
	FieldStatus,
	FieldTrigger,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "job_histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"job_history",
	"project_history",
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
	// DefaultOutput holds the default value on creation for the "output" field.
	DefaultOutput string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultTrigger holds the default value on creation for the "trigger" field.
	DefaultTrigger string
)

// OrderOption defines the ordering options for the JobHistory queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByWasSuccessful orders the results by the was_successful field.
func ByWasSuccessful(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWasSuccessful, opts...).ToFunc()
}

// ByDurationMs orders the results by the duration_ms field.
func ByDurationMs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDurationMs, opts...).ToFunc()
}

// ByOutput orders the results by the output field.
func ByOutput(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutput, opts...).ToFunc()
}

// ByExitCode orders the results by the exit_code field.
func ByExitCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExitCode, opts...).ToFunc()
}

// ByTriggeredByEmail orders the results by the triggered_by_email field.
func ByTriggeredByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggeredByEmail, opts...).ToFunc()
}

// ByTriggeredByID orders the results by the triggered_by_id field.
func ByTriggeredByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTriggeredByID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByTrigger orders the results by the trigger field.
func ByTrigger(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrigger, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByJobField orders the results by job field.
func ByJobField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJobStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newJobStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JobInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, JobTable, JobColumn),
	)
}
