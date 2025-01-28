package pkg

import (
	"context"
	"github.com/lbrictson/janus/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestGetAuthConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Create initial auth config
	initial, err := client.AuthConfig.Create().
		SetEnableSSO(false).
		SetDisablePasswordLogin(false).
		SetSSOProvider("none").
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating auth config: %v", err)
	}

	// Test getting config
	config, err := getAuthConfig(ctx, client)
	if err != nil {
		t.Errorf("getAuthConfig failed: %v", err)
	}
	if config.ID != initial.ID {
		t.Errorf("got wrong config ID")
	}

	// Test caching works
	lastAuthConfig = nil
	config1, _ := getAuthConfig(ctx, client)
	config2, _ := getAuthConfig(ctx, client)
	if config1 != config2 {
		t.Error("cache not working")
	}
}

func TestUpdateAuthConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Create initial config
	initial, err := client.AuthConfig.Create().
		SetEnableSSO(false).
		SetDisablePasswordLogin(false).
		SetSSOProvider("none").
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating auth config: %v", err)
	}

	// Update config
	initial.EnableSSO = true
	err = updateAuthConfig(ctx, client, initial)
	if err != nil {
		t.Errorf("updateAuthConfig failed: %v", err)
	}

	// Verify update
	updated, err := getAuthConfig(ctx, client)
	if err != nil || !updated.EnableSSO {
		t.Error("update not persisted")
	}
}

func TestGetDataConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	initial, err := client.DataConfig.Create().
		SetDaysToKeep(30).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating data config: %v", err)
	}

	config, err := getDataConfig(ctx, client)
	if err != nil {
		t.Errorf("getDataConfig failed: %v", err)
	}
	if config.ID != initial.ID {
		t.Error("got wrong config")
	}

	// Test cache
	lastDataConfig = nil
	c1, _ := getDataConfig(ctx, client)
	c2, _ := getDataConfig(ctx, client)
	if c1 != c2 {
		t.Error("cache not working")
	}
}

func TestUpdateDataConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	initial, err := client.DataConfig.Create().
		SetDaysToKeep(30).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating data config: %v", err)
	}

	initial.DaysToKeep = 60
	err = updateDataConfig(ctx, client, initial)
	if err != nil {
		t.Errorf("updateDataConfig failed: %v", err)
	}

	updated, _ := getDataConfig(ctx, client)
	if updated.DaysToKeep != 60 {
		t.Error("update not persisted")
	}
}

func TestGetJobConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	initial, err := client.JobConfig.Create().
		SetMaxConcurrentJobs(5).
		SetDefaultTimeoutSeconds(300).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating job config: %v", err)
	}

	config, err := getJobConfig(ctx, client)
	if err != nil {
		t.Errorf("getJobConfig failed: %v", err)
	}
	if config.ID != initial.ID {
		t.Error("got wrong config")
	}

	// Test cache
	lastJobConfig = nil
	c1, _ := getJobConfig(ctx, client)
	c2, _ := getJobConfig(ctx, client)
	if c1 != c2 {
		t.Error("cache not working")
	}
}

func TestUpdateJobConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	initial, err := client.JobConfig.Create().
		SetMaxConcurrentJobs(5).
		SetDefaultTimeoutSeconds(300).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating job config: %v", err)
	}

	initial.MaxConcurrentJobs = 10
	err = updateJobConfig(ctx, client, initial)
	if err != nil {
		t.Errorf("updateJobConfig failed: %v", err)
	}

	updated, _ := getJobConfig(ctx, client)
	if updated.MaxConcurrentJobs != 10 {
		t.Error("update not persisted")
	}
}

func TestGetSMTPConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	initial, err := client.SMTPConfig.Create().
		SetSMTPServer("smtp.test.com").
		SetSMTPPort(587).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating smtp config: %v", err)
	}

	config, err := getSMTPConfig(ctx, client)
	if err != nil {
		t.Errorf("getSMTPConfig failed: %v", err)
	}
	if config.ID != initial.ID {
		t.Error("got wrong config")
	}

	// Test cache
	lastSMTPConfig = nil
	c1, _ := getSMTPConfig(ctx, client)
	c2, _ := getSMTPConfig(ctx, client)
	if c1 != c2 {
		t.Error("cache not working")
	}
}

func TestUpdateSMTPConfig(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	initial, err := client.SMTPConfig.Create().
		SetSMTPServer("smtp.test.com").
		SetSMTPPort(587).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating smtp config: %v", err)
	}

	initial.SMTPPort = 25
	err = updateSMTPConfig(ctx, client, initial)
	if err != nil {
		t.Errorf("updateSMTPConfig failed: %v", err)
	}

	updated, _ := getSMTPConfig(ctx, client)
	if updated.SMTPPort != 25 {
		t.Error("update not persisted")
	}
}
