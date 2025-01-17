package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"net/http"
)

func renderAdminPage(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "admin", map[string]any{
			"Config": config,
		})
	}
}
