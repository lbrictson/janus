// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuditsColumns holds the columns for the "audits" table.
	AuditsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// AuditsTable holds the schema information for the "audits" table.
	AuditsTable = &schema.Table{
		Name:       "audits",
		Columns:    AuditsColumns,
		PrimaryKey: []*schema.Column{AuditsColumns[0]},
	}
	// AuthConfigsColumns holds the columns for the "auth_configs" table.
	AuthConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "enable_sso", Type: field.TypeBool, Default: false},
		{Name: "disable_password_login", Type: field.TypeBool, Default: false},
		{Name: "sso_provider", Type: field.TypeString, Default: ""},
		{Name: "sso_client_id", Type: field.TypeString, Default: ""},
		{Name: "sso_client_secret", Type: field.TypeString, Default: ""},
		{Name: "sso_redirect_uri", Type: field.TypeString, Default: ""},
		{Name: "sso_authorization_url", Type: field.TypeString, Default: ""},
		{Name: "sso_token_url", Type: field.TypeString, Default: ""},
		{Name: "sso_user_info_url", Type: field.TypeString, Default: ""},
		{Name: "entra_tenant_id", Type: field.TypeString, Default: ""},
		{Name: "google_allowed_domains", Type: field.TypeString, Default: ""},
		{Name: "session_key", Type: field.TypeBytes},
	}
	// AuthConfigsTable holds the schema information for the "auth_configs" table.
	AuthConfigsTable = &schema.Table{
		Name:       "auth_configs",
		Columns:    AuthConfigsColumns,
		PrimaryKey: []*schema.Column{AuthConfigsColumns[0]},
	}
	// JobsColumns holds the columns for the "jobs" table.
	JobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "cron_schedule", Type: field.TypeString, Nullable: true},
		{Name: "schedule_enabled", Type: field.TypeBool, Default: false},
		{Name: "allow_concurrent_runs", Type: field.TypeBool, Default: false},
		{Name: "arguments", Type: field.TypeJSON, Nullable: true},
		{Name: "requires_file_upload", Type: field.TypeBool, Default: false},
		{Name: "average_duration_ms", Type: field.TypeInt64, Nullable: true, Default: 0},
		{Name: "timeout_seconds", Type: field.TypeInt, Nullable: true, Default: 3600},
		{Name: "last_edit_time", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "notify_on_start_channel_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "notify_on_success_channel_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "notify_on_failure_channel_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "last_run_time", Type: field.TypeTime},
		{Name: "next_cron_run_time", Type: field.TypeTime},
		{Name: "last_run_success", Type: field.TypeBool, Default: true},
		{Name: "project_jobs", Type: field.TypeInt},
	}
	// JobsTable holds the schema information for the "jobs" table.
	JobsTable = &schema.Table{
		Name:       "jobs",
		Columns:    JobsColumns,
		PrimaryKey: []*schema.Column{JobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "jobs_projects_jobs",
				Columns:    []*schema.Column{JobsColumns[18]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "job_name_project_jobs",
				Unique:  true,
				Columns: []*schema.Column{JobsColumns[1], JobsColumns[18]},
			},
		},
	}
	// JobHistoriesColumns holds the columns for the "job_histories" table.
	JobHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "was_successful", Type: field.TypeBool},
		{Name: "duration_ms", Type: field.TypeInt64},
		{Name: "parameters", Type: field.TypeJSON, Nullable: true},
		{Name: "output", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "exit_code", Type: field.TypeInt},
		{Name: "triggered_by_email", Type: field.TypeString},
		{Name: "triggered_by_id", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString, Default: "running"},
		{Name: "trigger", Type: field.TypeString, Default: "manual"},
		{Name: "job_history", Type: field.TypeInt},
		{Name: "project_history", Type: field.TypeInt},
	}
	// JobHistoriesTable holds the schema information for the "job_histories" table.
	JobHistoriesTable = &schema.Table{
		Name:       "job_histories",
		Columns:    JobHistoriesColumns,
		PrimaryKey: []*schema.Column{JobHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "job_histories_jobs_history",
				Columns:    []*schema.Column{JobHistoriesColumns[11]},
				RefColumns: []*schema.Column{JobsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "job_histories_projects_history",
				Columns:    []*schema.Column{JobHistoriesColumns[12]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "jobhistory_created_at_project_history_job_history",
				Unique:  false,
				Columns: []*schema.Column{JobHistoriesColumns[8], JobHistoriesColumns[12], JobHistoriesColumns[11]},
			},
			{
				Name:    "jobhistory_status_created_at",
				Unique:  false,
				Columns: []*schema.Column{JobHistoriesColumns[9], JobHistoriesColumns[8]},
			},
		},
	}
	// NotificationChannelsColumns holds the columns for the "notification_channels" table.
	NotificationChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// NotificationChannelsTable holds the schema information for the "notification_channels" table.
	NotificationChannelsTable = &schema.Table{
		Name:       "notification_channels",
		Columns:    NotificationChannelsColumns,
		PrimaryKey: []*schema.Column{NotificationChannelsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
	}
	// ProjectUsersColumns holds the columns for the "project_users" table.
	ProjectUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "can_edit", Type: field.TypeBool, Default: false},
		{Name: "project_project_users", Type: field.TypeInt},
		{Name: "user_project_users", Type: field.TypeInt},
	}
	// ProjectUsersTable holds the schema information for the "project_users" table.
	ProjectUsersTable = &schema.Table{
		Name:       "project_users",
		Columns:    ProjectUsersColumns,
		PrimaryKey: []*schema.Column{ProjectUsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_users_projects_projectUsers",
				Columns:    []*schema.Column{ProjectUsersColumns[2]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "project_users_users_projectUsers",
				Columns:    []*schema.Column{ProjectUsersColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "encrypted_password", Type: field.TypeBytes},
		{Name: "admin", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "api_key", Type: field.TypeString, Unique: true},
		{Name: "must_change_password", Type: field.TypeBool, Default: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuditsTable,
		AuthConfigsTable,
		JobsTable,
		JobHistoriesTable,
		NotificationChannelsTable,
		ProjectsTable,
		ProjectUsersTable,
		UsersTable,
	}
)

func init() {
	JobsTable.ForeignKeys[0].RefTable = ProjectsTable
	JobHistoriesTable.ForeignKeys[0].RefTable = JobsTable
	JobHistoriesTable.ForeignKeys[1].RefTable = ProjectsTable
	ProjectUsersTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectUsersTable.ForeignKeys[1].RefTable = UsersTable
}
