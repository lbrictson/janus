package schema

import "entgo.io/ent"

// NotificationChannel holds the schema definition for the NotificationChannel entity.
type NotificationChannel struct {
	ent.Schema
}

// Fields of the NotificationChannel.
func (NotificationChannel) Fields() []ent.Field {
	return nil
}

// Edges of the NotificationChannel.
func (NotificationChannel) Edges() []ent.Edge {
	return nil
}
