package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// InboundWebhook holds the schema definition for the InboundWebhook entity.
type InboundWebhook struct {
	ent.Schema
}

// Fields of the InboundWebhook.
func (InboundWebhook) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("created_by"),
		field.Time("created_at").Default(time.Now),
		field.Bool("require_api_key").Default(false),
		field.String("api_key").Optional().Nillable(),
	}
}

// Edges of the InboundWebhook.
func (InboundWebhook) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("job", Job.Type).Unique(),
	}
}
