package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SMTPConfig holds the schema definition for the SMTPConfig entity.
type SMTPConfig struct {
	ent.Schema
}

// Fields of the SMTPConfig.
func (SMTPConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("smtp_server").Default(""),
		field.Int("smtp_port").Default(0),
		field.String("smtp_username").Default(""),
		field.String("smtp_password").Default(""),
		field.String("smtp_sender").Default(""),
		field.Bool("smtp_tls").Default(true),
	}
}

// Edges of the SMTPConfig.
func (SMTPConfig) Edges() []ent.Edge {
	return nil
}
