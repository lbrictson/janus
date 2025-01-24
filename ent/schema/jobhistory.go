package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// Parameter represents a key-value pair for job execution parameters
type Parameter struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Sensitive bool   `json:"sensitive"`
}

// JobHistory holds the schema definition for the JobHistory entity.
type JobHistory struct {
	ent.Schema
}

// Fields of the JobHistory.
func (JobHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("was_successful"),
		field.Int64("duration_ms"),
		field.JSON("parameters", []Parameter{}).
			Optional(),
		field.Text("output").
			Default(""),
		field.Int("exit_code"),
		field.String("triggered_by_email"),
		field.Int("triggered_by_id"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.String("status").Default("running"),
		field.String("trigger").Default("manual"),
	}
}

// Edges of the JobHistory.
func (JobHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("history").
			Unique().
			Required(),
		edge.From("job", Job.Type).
			Ref("history").
			Unique().
			Required(),
	}
}

// Indexes of the JobHistory.
func (JobHistory) Indexes() []ent.Index {
	return []ent.Index{
		// Index for querying history by project and job
		index.Fields("created_at").
			Edges("project", "job"),
		// Index for querying by status
		index.Fields("status", "created_at"),
	}
}
