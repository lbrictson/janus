package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// JobArgument represents a single argument configuration for a job
type JobArgument struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"` // "string", "number", "date", "datetime"
	AllowedValues []string `json:"allowed_values,omitempty"`
	DefaultValue  string   `json:"default_value"`
	Sensitive     bool     `json:"sensitive"`
}

// Job holds the schema definition for the Job entity.
type Job struct {
	ent.Schema
}

// Fields of the Job.
func (Job) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.String("cron_schedule").
			Optional(),
		field.Bool("schedule_enabled").
			Default(false),
		field.Bool("allow_concurrent_runs").
			Default(false),
		field.JSON("arguments", []JobArgument{}).
			Optional().
			Comment("List of arguments that can be passed to this job"),
		field.Bool("requires_file_upload").
			Default(false),
		field.Int64("average_duration_ms").
			Optional().
			Default(0),
		field.Int("timeout_seconds").
			Optional().
			Default(3600), // Default 1 hour
		field.Time("last_edit_time").
			UpdateDefault(time.Now),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Ints("notify_on_start_channel_ids").
			Optional(),
		field.Ints("notify_on_success_channel_ids").
			Optional(),
		field.Ints("notify_on_failure_channel_ids").
			Optional(),
		field.Time("last_run_time"),
		field.Time("next_cron_run_time"),
		field.Text("script"),
		field.Bool("last_run_success").Default(true),
		field.Bool("created_by_api").Default(false),
	}
}

// Edges of the Job.
func (Job) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("jobs").
			Unique().
			Required(),
		edge.To("history", JobHistory.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade}),
		edge.To("versions", JobVersion.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade}),
	}
}

// Indexes of the Job
func (Job) Indexes() []ent.Index {
	return []ent.Index{
		// Index for querying jobs by project efficiently
		index.Fields("name").
			Edges("project").
			Unique(),
	}
}
