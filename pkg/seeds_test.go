package pkg

import (
	"context"
	"testing"

	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

func TestExecuteSeeds_FreshDatabase(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Execute seeds on fresh database
	err := ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("ExecuteSeeds failed: %v", err)
	}

	// Verify User was created
	users, err := client.User.Query().All(ctx)
	if err != nil {
		t.Fatalf("Failed to query users: %v", err)
	}
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
	if len(users) > 0 {
		user := users[0]
		if user.Email != "admin@localhost" {
			t.Errorf("Expected email 'admin@localhost', got '%s'", user.Email)
		}
		if !user.Admin {
			t.Error("Expected user to be admin")
		}
		if user.APIKey == "" {
			t.Error("Expected API key to be generated")
		}
		// Verify password hash was created
		if len(user.EncryptedPassword) == 0 {
			t.Error("Expected encrypted password to be set")
		}
		// Verify password is correct
		err = compareHashAndPassword(user.EncryptedPassword, "ChangeMeBeforeUse1234!")
		if err != nil {
			t.Errorf("Password verification failed: %v", err)
		}
	}

	// Verify AuthConfig was created
	authConfigs, err := client.AuthConfig.Query().All(ctx)
	if err != nil {
		t.Fatalf("Failed to query auth configs: %v", err)
	}
	if len(authConfigs) != 1 {
		t.Errorf("Expected 1 auth config, got %d", len(authConfigs))
	}
	if len(authConfigs) > 0 {
		config := authConfigs[0]
		if config.EnableSSO {
			t.Error("Expected SSO to be disabled by default")
		}
		if len(config.SessionKey) == 0 {
			t.Error("Expected session key to be generated")
		}
	}

	// Verify DataConfig was created
	dataConfigs, err := client.DataConfig.Query().All(ctx)
	if err != nil {
		t.Fatalf("Failed to query data configs: %v", err)
	}
	if len(dataConfigs) != 1 {
		t.Errorf("Expected 1 data config, got %d", len(dataConfigs))
	}
	if len(dataConfigs) > 0 {
		config := dataConfigs[0]
		if config.DaysToKeep != 180 {
			t.Errorf("Expected DaysToKeep to be 180, got %d", config.DaysToKeep)
		}
	}

	// Verify SMTPConfig was created
	smtpConfigs, err := client.SMTPConfig.Query().All(ctx)
	if err != nil {
		t.Fatalf("Failed to query SMTP configs: %v", err)
	}
	if len(smtpConfigs) != 1 {
		t.Errorf("Expected 1 SMTP config, got %d", len(smtpConfigs))
	}
	if len(smtpConfigs) > 0 {
		config := smtpConfigs[0]
		if config.SMTPServer != "localhost" {
			t.Errorf("Expected SMTP server 'localhost', got '%s'", config.SMTPServer)
		}
		if config.SMTPPort != 1025 {
			t.Errorf("Expected SMTP port 1025, got %d", config.SMTPPort)
		}
		if config.SMTPUsername != "admin@localhost" {
			t.Errorf("Expected SMTP username 'admin@localhost', got '%s'", config.SMTPUsername)
		}
		if config.SMTPSender != "admin@localhost" {
			t.Errorf("Expected SMTP sender 'admin@localhost', got '%s'", config.SMTPSender)
		}
		if config.SMTPTLS {
			t.Error("Expected SMTP TLS to be disabled")
		}
	}

	// Verify JobConfig was created
	jobConfigs, err := client.JobConfig.Query().All(ctx)
	if err != nil {
		t.Fatalf("Failed to query job configs: %v", err)
	}
	if len(jobConfigs) != 1 {
		t.Errorf("Expected 1 job config, got %d", len(jobConfigs))
	}
	if len(jobConfigs) > 0 {
		config := jobConfigs[0]
		if config.MaxConcurrentJobs != 100 {
			t.Errorf("Expected MaxConcurrentJobs to be 100, got %d", config.MaxConcurrentJobs)
		}
		if config.DefaultTimeoutSeconds != 3600 {
			t.Errorf("Expected DefaultTimeoutSeconds to be 3600, got %d", config.DefaultTimeoutSeconds)
		}
	}
}

func TestExecuteSeeds_Idempotency(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Execute seeds first time
	err := ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("First ExecuteSeeds failed: %v", err)
	}

	// Count entities after first run
	userCount1, _ := client.User.Query().Count(ctx)
	authConfigCount1, _ := client.AuthConfig.Query().Count(ctx)
	dataConfigCount1, _ := client.DataConfig.Query().Count(ctx)
	smtpConfigCount1, _ := client.SMTPConfig.Query().Count(ctx)
	jobConfigCount1, _ := client.JobConfig.Query().Count(ctx)

	// Execute seeds second time
	err = ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("Second ExecuteSeeds failed: %v", err)
	}

	// Count entities after second run
	userCount2, _ := client.User.Query().Count(ctx)
	authConfigCount2, _ := client.AuthConfig.Query().Count(ctx)
	dataConfigCount2, _ := client.DataConfig.Query().Count(ctx)
	smtpConfigCount2, _ := client.SMTPConfig.Query().Count(ctx)
	jobConfigCount2, _ := client.JobConfig.Query().Count(ctx)

	// Verify no duplicates were created
	if userCount1 != userCount2 {
		t.Errorf("User count changed: %d -> %d", userCount1, userCount2)
	}
	if authConfigCount1 != authConfigCount2 {
		t.Errorf("AuthConfig count changed: %d -> %d", authConfigCount1, authConfigCount2)
	}
	if dataConfigCount1 != dataConfigCount2 {
		t.Errorf("DataConfig count changed: %d -> %d", dataConfigCount1, dataConfigCount2)
	}
	if smtpConfigCount1 != smtpConfigCount2 {
		t.Errorf("SMTPConfig count changed: %d -> %d", smtpConfigCount1, smtpConfigCount2)
	}
	if jobConfigCount1 != jobConfigCount2 {
		t.Errorf("JobConfig count changed: %d -> %d", jobConfigCount1, jobConfigCount2)
	}

	// Verify we still have exactly one of each
	if userCount2 != 1 {
		t.Errorf("Expected 1 user after idempotent run, got %d", userCount2)
	}
	if authConfigCount2 != 1 {
		t.Errorf("Expected 1 auth config after idempotent run, got %d", authConfigCount2)
	}
	if dataConfigCount2 != 1 {
		t.Errorf("Expected 1 data config after idempotent run, got %d", dataConfigCount2)
	}
	if smtpConfigCount2 != 1 {
		t.Errorf("Expected 1 SMTP config after idempotent run, got %d", smtpConfigCount2)
	}
	if jobConfigCount2 != 1 {
		t.Errorf("Expected 1 job config after idempotent run, got %d", jobConfigCount2)
	}
}

func TestExecuteSeeds_PartialSeeding(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Pre-create only user and auth config
	pw, err := hashAndSaltPassword("ExistingPassword123!")
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}
	
	existingUser, err := client.User.Create().
		SetEmail("existing@example.com").
		SetEncryptedPassword(pw).
		SetAPIKey(generateLongString()).
		SetAdmin(false).
		Save(ctx)
	if err != nil {
		t.Fatalf("Failed to create existing user: %v", err)
	}

	existingAuthConfig, err := client.AuthConfig.Create().
		SetEnableSSO(true).
		SetSessionKey([]byte("existing-session-key")).
		Save(ctx)
	if err != nil {
		t.Fatalf("Failed to create existing auth config: %v", err)
	}

	// Execute seeds with partial data
	err = ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("ExecuteSeeds with partial data failed: %v", err)
	}

	// Verify existing data wasn't modified
	users, _ := client.User.Query().All(ctx)
	if len(users) != 1 {
		t.Errorf("Expected 1 user (existing), got %d", len(users))
	}
	if len(users) > 0 {
		user := users[0]
		if user.ID != existingUser.ID {
			t.Error("User was replaced instead of preserved")
		}
		if user.Email != "existing@example.com" {
			t.Errorf("User email was modified: %s", user.Email)
		}
	}

	authConfigs, _ := client.AuthConfig.Query().All(ctx)
	if len(authConfigs) != 1 {
		t.Errorf("Expected 1 auth config (existing), got %d", len(authConfigs))
	}
	if len(authConfigs) > 0 {
		config := authConfigs[0]
		if config.ID != existingAuthConfig.ID {
			t.Error("AuthConfig was replaced instead of preserved")
		}
		if !config.EnableSSO {
			t.Error("AuthConfig EnableSSO was modified")
		}
	}

	// Verify missing configs were created
	dataConfigCount, _ := client.DataConfig.Query().Count(ctx)
	if dataConfigCount != 1 {
		t.Errorf("Expected DataConfig to be created, got %d", dataConfigCount)
	}

	smtpConfigCount, _ := client.SMTPConfig.Query().Count(ctx)
	if smtpConfigCount != 1 {
		t.Errorf("Expected SMTPConfig to be created, got %d", smtpConfigCount)
	}

	jobConfigCount, _ := client.JobConfig.Query().Count(ctx)
	if jobConfigCount != 1 {
		t.Errorf("Expected JobConfig to be created, got %d", jobConfigCount)
	}
}

func TestExecuteSeeds_EmptyDatabase(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Verify database is empty before seeding
	userCount, _ := client.User.Query().Count(ctx)
	if userCount != 0 {
		t.Fatalf("Expected empty database, found %d users", userCount)
	}

	// Execute seeds
	err := ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("ExecuteSeeds failed on empty database: %v", err)
	}

	// Verify all entities were created
	userCount, _ = client.User.Query().Count(ctx)
	authConfigCount, _ := client.AuthConfig.Query().Count(ctx)
	dataConfigCount, _ := client.DataConfig.Query().Count(ctx)
	smtpConfigCount, _ := client.SMTPConfig.Query().Count(ctx)
	jobConfigCount, _ := client.JobConfig.Query().Count(ctx)

	if userCount == 0 {
		t.Error("No users were created")
	}
	if authConfigCount == 0 {
		t.Error("No auth config was created")
	}
	if dataConfigCount == 0 {
		t.Error("No data config was created")
	}
	if smtpConfigCount == 0 {
		t.Error("No SMTP config was created")
	}
	if jobConfigCount == 0 {
		t.Error("No job config was created")
	}
}

func TestExecuteSeeds_DefaultValues(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	err := ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("ExecuteSeeds failed: %v", err)
	}

	// Test specific default values in detail
	tests := []struct {
		name     string
		validate func(t *testing.T, client *ent.Client, ctx context.Context)
	}{
		{
			name: "User defaults",
			validate: func(t *testing.T, client *ent.Client, ctx context.Context) {
				user, err := client.User.Query().First(ctx)
				if err != nil {
					t.Fatalf("Failed to get user: %v", err)
				}
				if user.Email != "admin@localhost" {
					t.Errorf("Wrong email: %s", user.Email)
				}
				if !user.Admin {
					t.Error("User should be admin")
				}
				if len(user.APIKey) < 32 {
					t.Error("API key too short")
				}
			},
		},
		{
			name: "AuthConfig defaults",
			validate: func(t *testing.T, client *ent.Client, ctx context.Context) {
				config, err := client.AuthConfig.Query().First(ctx)
				if err != nil {
					t.Fatalf("Failed to get auth config: %v", err)
				}
				if config.EnableSSO {
					t.Error("SSO should be disabled by default")
				}
				if len(config.SessionKey) < 32 {
					t.Error("Session key too short")
				}
			},
		},
		{
			name: "DataConfig defaults",
			validate: func(t *testing.T, client *ent.Client, ctx context.Context) {
				config, err := client.DataConfig.Query().First(ctx)
				if err != nil {
					t.Fatalf("Failed to get data config: %v", err)
				}
				if config.DaysToKeep != 180 {
					t.Errorf("Wrong DaysToKeep: %d", config.DaysToKeep)
				}
			},
		},
		{
			name: "SMTPConfig defaults",
			validate: func(t *testing.T, client *ent.Client, ctx context.Context) {
				config, err := client.SMTPConfig.Query().First(ctx)
				if err != nil {
					t.Fatalf("Failed to get SMTP config: %v", err)
				}
				if config.SMTPServer != "localhost" {
					t.Errorf("Wrong SMTP server: %s", config.SMTPServer)
				}
				if config.SMTPPort != 1025 {
					t.Errorf("Wrong SMTP port: %d", config.SMTPPort)
				}
				if config.SMTPTLS {
					t.Error("TLS should be disabled by default")
				}
			},
		},
		{
			name: "JobConfig defaults",
			validate: func(t *testing.T, client *ent.Client, ctx context.Context) {
				config, err := client.JobConfig.Query().First(ctx)
				if err != nil {
					t.Fatalf("Failed to get job config: %v", err)
				}
				if config.MaxConcurrentJobs != 100 {
					t.Errorf("Wrong MaxConcurrentJobs: %d", config.MaxConcurrentJobs)
				}
				if config.DefaultTimeoutSeconds != 3600 {
					t.Errorf("Wrong DefaultTimeoutSeconds: %d", config.DefaultTimeoutSeconds)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.validate(t, client, ctx)
		})
	}
}

func TestExecuteSeeds_UniqueAPIKeys(t *testing.T) {
	// This test verifies that if we were to create multiple users,
	// they would have unique API keys
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Execute seeds to create first user
	err := ExecuteSeeds(ctx, client)
	if err != nil {
		t.Fatalf("ExecuteSeeds failed: %v", err)
	}

	firstUser, err := client.User.Query().First(ctx)
	if err != nil {
		t.Fatalf("Failed to get first user: %v", err)
	}

	// Create a second user manually
	pw, _ := hashAndSaltPassword("TestPassword123!")
	secondUser, err := client.User.Create().
		SetEmail("test2@localhost").
		SetEncryptedPassword(pw).
		SetAPIKey(generateLongString()).
		SetAdmin(false).
		Save(ctx)
	if err != nil {
		t.Fatalf("Failed to create second user: %v", err)
	}

	// Verify API keys are different
	if firstUser.APIKey == secondUser.APIKey {
		t.Error("API keys should be unique")
	}

	// Verify both API keys have sufficient length
	if len(firstUser.APIKey) < 32 {
		t.Errorf("First user API key too short: %d chars", len(firstUser.APIKey))
	}
	if len(secondUser.APIKey) < 32 {
		t.Errorf("Second user API key too short: %d chars", len(secondUser.APIKey))
	}
}

func TestExecuteSeeds_SessionKeyUniqueness(t *testing.T) {
	// Test that session keys are unique across different runs
	client1 := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client1.Close()
	
	client2 := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client2.Close()
	
	ctx := context.Background()

	// Execute seeds on both databases
	err := ExecuteSeeds(ctx, client1)
	if err != nil {
		t.Fatalf("ExecuteSeeds on client1 failed: %v", err)
	}
	
	err = ExecuteSeeds(ctx, client2)
	if err != nil {
		t.Fatalf("ExecuteSeeds on client2 failed: %v", err)
	}

	// Get session keys from both
	config1, _ := client1.AuthConfig.Query().First(ctx)
	config2, _ := client2.AuthConfig.Query().First(ctx)

	// Verify session keys are different
	if string(config1.SessionKey) == string(config2.SessionKey) {
		t.Error("Session keys should be unique across different database instances")
	}
}