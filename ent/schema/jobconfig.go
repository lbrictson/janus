package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// JobConfig holds the schema definition for the JobConfig entity.
type JobConfig struct {
	ent.Schema
}

// Fields of the JobConfig.
func (JobConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("max_concurrent_jobs").Default(100),
		field.Int("default_timeout_seconds").Default(600),
	}
}

// Edges of the JobConfig.
func (JobConfig) Edges() []ent.Edge {
	return nil
}
