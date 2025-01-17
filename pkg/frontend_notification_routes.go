package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"net/http"
)

func renderNotificationPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Channels struct {
			ID       int
			Name     string
			Type     string
			LastUsed string
		}
		channels := []Channels{
			{ID: 1, Name: "Email DevOps", Type: "Email", LastUsed: "2021-01-01"},
			{ID: 2, Name: "#team-devops slack channel", Type: "Slack", LastUsed: "2021-01-01"},
			{ID: 3, Name: "#p1 teams room", Type: "Teams", LastUsed: "2021-01-01"},
		}
		return c.Render(http.StatusOK, "notifications", map[string]any{
			"Channels": channels,
		})
	}
}
