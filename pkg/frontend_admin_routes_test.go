package pkg

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *ent.Client {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	ctx := context.Background()

	// Create required configs
	_, err := client.AuthConfig.Create().Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.DataConfig.Create().Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.JobConfig.Create().Save(ctx)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.SMTPConfig.Create().Save(ctx)
	if err != nil {
		t.Fatal(err)
	}

	return client
}

func TestRenderAdminPage(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	e := echo.New()
	registerRenderer(e, Config{
		BrandName: "Janus",
	})

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := renderAdminPage(client, &Config{})
	err := handler(c)

	if err != nil {
		t.Errorf("renderAdminPage returned error: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("expected status OK; got %v", rec.Code)
	}
}

func TestFormAdminDataRetention(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	e := echo.New()
	f := make(url.Values)
	f.Set("job_history_days", "30")
	req := httptest.NewRequest(http.MethodPost, "/admin/data", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := formAdminDataRetention(client)
	err := handler(c)

	if err != nil {
		t.Errorf("formAdminDataRetention returned error: %v", err)
	}
	if rec.Code != http.StatusSeeOther {
		t.Errorf("expected status SeeOther; got %v", rec.Code)
	}

	// Verify update
	ctx := context.Background()
	config, _ := client.DataConfig.Query().Only(ctx)
	if config.DaysToKeep != 30 {
		t.Errorf("expected DaysToKeep=30; got %v", config.DaysToKeep)
	}
}

func TestFormAdminSMTP(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	e := echo.New()
	f := make(url.Values)
	f.Set("hostname", "smtp.test.com")
	f.Set("port", "587")
	f.Set("username", "test")
	f.Set("password", "pass")
	f.Set("from_address", "test@test.com")

	req := httptest.NewRequest(http.MethodPost, "/admin/smtp", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := formAdminSMTP(client)
	err := handler(c)

	if err != nil {
		t.Errorf("formAdminSMTP returned error: %v", err)
	}

	// Verify update
	ctx := context.Background()
	config, _ := client.SMTPConfig.Query().Only(ctx)
	if config.SMTPServer != "smtp.test.com" {
		t.Error("SMTP settings not updated correctly")
	}
}

func TestFormJobSettings(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	e := echo.New()
	f := make(url.Values)
	f.Set("default_timeout", "300")
	f.Set("max_concurrent_jobs", "5")

	req := httptest.NewRequest(http.MethodPost, "/admin/jobs", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := formJobSettings(client)
	err := handler(c)

	if err != nil {
		t.Errorf("formJobSettings returned error: %v", err)
	}

	// Verify update
	ctx := context.Background()
	config, _ := client.JobConfig.Query().Only(ctx)
	if config.DefaultTimeoutSeconds != 300 || config.MaxConcurrentJobs != 5 {
		t.Error("Job settings not updated correctly")
	}
}

func TestFormAdminSecuritySettings(t *testing.T) {
	client := setupTestDB(t)
	defer client.Close()

	e := echo.New()
	f := make(url.Values)
	f.Set("disable_password_login", "on")
	f.Set("enable_sso", "on")
	f.Set("sso_provider", "google")
	f.Set("sso_client_id", "client123")
	f.Set("sso_client_secret", "secret123")

	req := httptest.NewRequest(http.MethodPost, "/admin/security", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := formAdminSecuritySettings(client, &Config{})
	err := handler(c)

	if err != nil {
		t.Errorf("formAdminSecuritySettings returned error: %v", err)
	}

	// Verify update
	ctx := context.Background()
	config, _ := client.AuthConfig.Query().Only(ctx)
	if !config.DisablePasswordLogin || !config.EnableSSO || config.SSOProvider != "google" {
		t.Error("Security settings not updated correctly")
	}
}
