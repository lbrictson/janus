package schema

import "entgo.io/ent"

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

// Fields of the Audit.
func (Audit) Fields() []ent.Field {
	return nil
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return nil
}
