package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProjectUser holds the schema definition for the ProjectUser entity.
type ProjectUser struct {
	ent.Schema
}

// Fields of the ProjectUser.
func (ProjectUser) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("can_edit").Default(false),
	}
}

// Edges of the ProjectUser.
func (ProjectUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Required().Ref("projectUsers").Unique(),
		edge.From("user", User.Type).Required().Ref("projectUsers").Unique(),
	}
}
