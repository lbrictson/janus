// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/janus/ent/project"
)

// Project is the model entity for the Project schema.
type Project struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProjectQuery when eager-loading is set.
	Edges        ProjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProjectEdges holds the relations/edges for other nodes in the graph.
type ProjectEdges struct {
	// ProjectUsers holds the value of the projectUsers edge.
	ProjectUsers []*ProjectUser `json:"projectUsers,omitempty"`
	// Jobs holds the value of the jobs edge.
	Jobs []*Job `json:"jobs,omitempty"`
	// History holds the value of the history edge.
	History []*JobHistory `json:"history,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectUsersOrErr returns the ProjectUsers value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) ProjectUsersOrErr() ([]*ProjectUser, error) {
	if e.loadedTypes[0] {
		return e.ProjectUsers, nil
	}
	return nil, &NotLoadedError{edge: "projectUsers"}
}

// JobsOrErr returns the Jobs value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) JobsOrErr() ([]*Job, error) {
	if e.loadedTypes[1] {
		return e.Jobs, nil
	}
	return nil, &NotLoadedError{edge: "jobs"}
}

// HistoryOrErr returns the History value or an error if the edge
// was not loaded in eager-loading.
func (e ProjectEdges) HistoryOrErr() ([]*JobHistory, error) {
	if e.loadedTypes[2] {
		return e.History, nil
	}
	return nil, &NotLoadedError{edge: "history"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Project) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			values[i] = new(sql.NullInt64)
		case project.FieldName, project.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Project fields.
func (pr *Project) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case project.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case project.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case project.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pr.Description = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Project.
// This includes values selected through modifiers, order, etc.
func (pr *Project) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryProjectUsers queries the "projectUsers" edge of the Project entity.
func (pr *Project) QueryProjectUsers() *ProjectUserQuery {
	return NewProjectClient(pr.config).QueryProjectUsers(pr)
}

// QueryJobs queries the "jobs" edge of the Project entity.
func (pr *Project) QueryJobs() *JobQuery {
	return NewProjectClient(pr.config).QueryJobs(pr)
}

// QueryHistory queries the "history" edge of the Project entity.
func (pr *Project) QueryHistory() *JobHistoryQuery {
	return NewProjectClient(pr.config).QueryHistory(pr)
}

// Update returns a builder for updating this Project.
// Note that you need to call Project.Unwrap() before calling this method if this Project
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Project) Update() *ProjectUpdateOne {
	return NewProjectClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Project entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Project) Unwrap() *Project {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Project is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Project) String() string {
	var builder strings.Builder
	builder.WriteString("Project(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(pr.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Projects is a parsable slice of Project.
type Projects []*Project
