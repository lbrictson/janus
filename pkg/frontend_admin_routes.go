package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"log/slog"
	"net/http"
)

func renderAdminPage(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		sec, _ := getAuthConfig(c.Request().Context(), db)
		data, err := getDataConfig(c.Request().Context(), db)
		if err != nil {
			slog.Error("failed to get data config", "error", err)
			return renderErrorPage(c, "Failed to get data config", http.StatusInternalServerError)
		}
		jobConfig, _ := getJobConfig(c.Request().Context(), db)
		smtpConfig, _ := getSMTPConfig(c.Request().Context(), db)
		return c.Render(http.StatusOK, "admin", map[string]any{
			"Config":        config,
			"Security":      sec,
			"SMTP":          smtpConfig,
			"JobSettings":   jobConfig,
			"DataRetention": data,
		})
	}
}

func formAdminDataRetention(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, _ := getDataConfig(c.Request().Context(), db)
		type Form struct {
			JobHistoryDays int `form:"job_history_days"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "Failed to bind form", http.StatusBadRequest)
		}
		fmt.Println(form.JobHistoryDays)
		data.DaysToKeep = form.JobHistoryDays
		if err := updateDataConfig(c.Request().Context(), db, data); err != nil {
			slog.Error("failed to update data config", "error", err)
			return renderErrorPage(c, "Failed to update data config", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func formAdminSMTP(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		smtp, _ := getSMTPConfig(c.Request().Context(), db)
		type Form struct {
			Host        string `form:"hostname"`
			Port        int    `form:"port"`
			Username    string `form:"username"`
			Password    string `form:"password"`
			FromAddress string `form:"from_address"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "Failed to bind form", http.StatusBadRequest)
		}
		smtp.SMTPServer = form.Host
		smtp.SMTPPort = form.Port
		smtp.SMTPUsername = form.Username
		smtp.SMTPPassword = form.Password
		smtp.SMTPSender = form.FromAddress
		if err := updateSMTPConfig(c.Request().Context(), db, smtp); err != nil {
			slog.Error("failed to update smtp config", "error", err)
			return renderErrorPage(c, "Failed to update smtp config", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func formJobSettings(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		jobConfig, _ := getJobConfig(c.Request().Context(), db)
		type Form struct {
			MaxJobDuration int `form:"default_timeout"`
			MaxConcurrent  int `form:"max_concurrent_jobs"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "Failed to bind form", http.StatusBadRequest)
		}
		jobConfig.DefaultTimeoutSeconds = form.MaxJobDuration
		jobConfig.MaxConcurrentJobs = form.MaxConcurrent
		if err := updateJobConfig(c.Request().Context(), db, jobConfig); err != nil {
			slog.Error("failed to update job config", "error", err)
			return renderErrorPage(c, "Failed to update job config", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func formAdminSecuritySettings(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		sec, _ := getAuthConfig(c.Request().Context(), db)
		type Form struct {
			DisablePasswordLogin string `form:"disable_password_login"`
			EnableSSO            string `form:"enable_sso"`
			SSOProvider          string `form:"sso_provider"`
			SSOClientID          string `form:"sso_client_id"`
			SSOClientSecret      string `form:"sso_client_secret"`
			SSORedirectURL       string `form:"sso_redirect_url"`
			SSOAuthURL           string `form:"sso_auth_url"`
			SSOTokenURL          string `form:"sso_token_url"`
			SSOUserInfoURL       string `form:"sso_userinfo_url"`
			EntraTenantID        string `form:"entra_tenant_id"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "Failed to bind form", http.StatusBadRequest)
		}
		disablePasswordLogin := false
		if form.DisablePasswordLogin == "on" {
			disablePasswordLogin = true
		}
		enableSSO := false
		if form.EnableSSO == "on" {
			enableSSO = true
		}
		sec.SSORedirectURI = form.SSORedirectURL
		sec.SSOClientID = form.SSOClientID
		sec.SSOClientSecret = form.SSOClientSecret
		sec.SSOAuthorizationURL = form.SSOAuthURL
		sec.SSOTokenURL = form.SSOTokenURL
		sec.SSOUserInfoURL = form.SSOUserInfoURL
		sec.SSOProvider = form.SSOProvider
		sec.DisablePasswordLogin = disablePasswordLogin
		sec.EnableSSO = enableSSO
		sec.EntraTenantID = form.EntraTenantID
		if err := updateAuthConfig(c.Request().Context(), db, sec); err != nil {
			slog.Error("failed to update auth config", "error", err)
			return renderErrorPage(c, "Failed to update auth config", http.StatusInternalServerError)
		}
		wireSSOConnection(db, config)
		return c.Redirect(http.StatusSeeOther, "/admin")
	}
}
