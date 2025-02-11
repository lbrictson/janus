// Code generated by ent, DO NOT EDIT.

package project

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeProjectUsers holds the string denoting the projectusers edge name in mutations.
	EdgeProjectUsers = "projectUsers"
	// EdgeJobs holds the string denoting the jobs edge name in mutations.
	EdgeJobs = "jobs"
	// EdgeHistory holds the string denoting the history edge name in mutations.
	EdgeHistory = "history"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// ProjectUsersTable is the table that holds the projectUsers relation/edge.
	ProjectUsersTable = "project_users"
	// ProjectUsersInverseTable is the table name for the ProjectUser entity.
	// It exists in this package in order to avoid circular dependency with the "projectuser" package.
	ProjectUsersInverseTable = "project_users"
	// ProjectUsersColumn is the table column denoting the projectUsers relation/edge.
	ProjectUsersColumn = "project_project_users"
	// JobsTable is the table that holds the jobs relation/edge.
	JobsTable = "jobs"
	// JobsInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobsInverseTable = "jobs"
	// JobsColumn is the table column denoting the jobs relation/edge.
	JobsColumn = "project_jobs"
	// HistoryTable is the table that holds the history relation/edge.
	HistoryTable = "job_histories"
	// HistoryInverseTable is the table name for the JobHistory entity.
	// It exists in this package in order to avoid circular dependency with the "jobhistory" package.
	HistoryInverseTable = "job_histories"
	// HistoryColumn is the table column denoting the history relation/edge.
	HistoryColumn = "project_history"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Project queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByProjectUsersCount orders the results by projectUsers count.
func ByProjectUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProjectUsersStep(), opts...)
	}
}

// ByProjectUsers orders the results by projectUsers terms.
func ByProjectUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByJobsCount orders the results by jobs count.
func ByJobsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newJobsStep(), opts...)
	}
}

// ByJobs orders the results by jobs terms.
func ByJobs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newJobsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByHistoryCount orders the results by history count.
func ByHistoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHistoryStep(), opts...)
	}
}

// ByHistory orders the results by history terms.
func ByHistory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHistoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProjectUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectUsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProjectUsersTable, ProjectUsersColumn),
	)
}
func newJobsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(JobsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, JobsTable, JobsColumn),
	)
}
func newHistoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HistoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HistoryTable, HistoryColumn),
	)
}
