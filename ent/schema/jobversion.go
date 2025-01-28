package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// JobVersion holds the schema definition for the JobVersion entity.
type JobVersion struct {
	ent.Schema
}

func (JobVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.String("name"),
		field.String("description"),
		field.String("script"),
		field.String("cron_schedule").Optional(),
		field.Bool("schedule_enabled"),
		field.Bool("allow_concurrent_runs"),
		field.JSON("arguments", []JobArgument{}),
		field.Bool("requires_file_upload"),
		field.String("changed_by_email"),
	}
}

func (JobVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("job", Job.Type).
			Ref("versions").
			Unique().
			Required(),
	}
}
