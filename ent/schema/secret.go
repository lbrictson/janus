package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Secret holds the schema definition for the Secret entity.
type Secret struct {
	ent.Schema
}

// Fields of the Secret.
func (Secret) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("value"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Secret.
func (Secret) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).Unique(),
	}
}
