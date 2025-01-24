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
			slog.Error("failed to get data config: %v", err)
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
			slog.Error("failed to bind form: %v", err)
			return renderErrorPage(c, "Failed to bind form", http.StatusBadRequest)
		}
		fmt.Println(form.JobHistoryDays)
		data.DaysToKeep = form.JobHistoryDays
		if err := updateDataConfig(c.Request().Context(), db, data); err != nil {
			slog.Error("failed to update data config: %v", err)
			return renderErrorPage(c, "Failed to update data config", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, "/admin")
	}
}
