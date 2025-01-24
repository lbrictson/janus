package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// DataConfig holds the schema definition for the DataConfig entity.
type DataConfig struct {
	ent.Schema
}

// Fields of the DataConfig.
func (DataConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int("days_to_keep").Default(180),
	}
}

// Edges of the DataConfig.
func (DataConfig) Edges() []ent.Edge {
	return nil
}
