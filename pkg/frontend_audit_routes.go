package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"net/http"
)

func renderAuditPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type AuditItem struct {
			ID      int
			Time    string
			Actor   string
			Project string
			Job     string
		}
		audits := []AuditItem{
			{ID: 1, Time: "2021-01-01", Actor: "lbrictson", Project: "IT Admin", Job: "Archive Mailbox"},
			{ID: 2, Time: "2021-01-01", Actor: "lbrictson", Project: "eGov Toolbox", Job: "Init new FNB client"},
			{ID: 3, Time: "2021-01-01", Actor: "lbrictson", Project: "IT Admin", Job: "Onboard new employee"},
		}
		return c.Render(http.StatusOK, "audit", map[string]any{
			"AuditLogs": audits,
		})
	}
}
