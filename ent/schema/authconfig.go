package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthConfig holds the schema definition for the AuthConfig entity.
type AuthConfig struct {
	ent.Schema
}

// Fields of the AuthConfig.
func (AuthConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("enable_sso").Default(false),
		field.Bool("disable_password_login").Default(false),
		field.String("sso_provider").Default(""),
		field.String("sso_client_id").Default(""),
		field.String("sso_client_secret").Default(""),
		field.String("sso_redirect_uri").Default(""),
		field.String("sso_authorization_url").Default(""),
		field.String("sso_token_url").Default(""),
		field.String("sso_user_info_url").Default(""),
		field.String("entra_tenant_id").Default(""),
		field.String("google_allowed_domains").Default(""),
		field.Bytes("session_key").Default([]byte("")),
	}
}

// Edges of the AuthConfig.
func (AuthConfig) Edges() []ent.Edge {
	return nil
}
