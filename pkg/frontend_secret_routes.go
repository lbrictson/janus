package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/secret"
	"log/slog"
	"net/http"
	"strconv"
)

func renderProjectSecretsView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Secret struct {
			ent.Secret
			CreatedAtFriendly string
			UpdatedAtFriendly string
		}
		projectID := c.Param("id")
		projectIDInt, err := strconv.Atoi(projectID)
		if err != nil {
			return renderErrorPage(c, "Invalid project ID", http.StatusBadRequest)
		}
		p, err := db.Project.Query().Where(project.IDEQ(projectIDInt)).Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting project", "error", err)
			return renderErrorPage(c, "Project not found", http.StatusNotFound)
		}
		s, err := db.Secret.Query().Where(secret.HasProjectWith(project.IDEQ(projectIDInt))).Order(ent.Desc(secret.FieldName)).All(c.Request().Context())
		if err != nil {
			slog.Error("error getting secrets", "error", err)
			return renderErrorPage(c, "Error getting secrets", http.StatusInternalServerError)
		}
		secrets := make([]Secret, 0)
		for _, s := range s {
			secrets = append(secrets, Secret{
				Secret:            *s,
				CreatedAtFriendly: s.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAtFriendly: s.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		return c.Render(http.StatusOK, "project-secrets", map[string]any{
			"Project": p,
			"Secrets": secrets,
		})
	}
}

func renderProjectAddSecretsView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.Param("id")
		projectIDInt, err := strconv.Atoi(projectID)
		if err != nil {
			return renderErrorPage(c, "Invalid project ID", http.StatusBadRequest)
		}
		p, err := db.Project.Query().Where(project.IDEQ(projectIDInt)).Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting project", "error", err)
			return renderErrorPage(c, "Project not found", http.StatusNotFound)
		}
		return c.Render(http.StatusOK, "add-secret", map[string]any{
			"Project": p,
		})
	}
}

func formProjectAddSecret(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.Param("id")
		projectIDInt, err := strconv.Atoi(projectID)
		if err != nil {
			return renderErrorPage(c, "Invalid project ID", http.StatusBadRequest)
		}
		p, err := db.Project.Query().Where(project.IDEQ(projectIDInt)).Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting project", "error", err)
			return renderErrorPage(c, "Project not found", http.StatusNotFound)
		}
		name := c.FormValue("name")
		value := c.FormValue("value")
		err = isJobOrSecretNameValid(name)
		if err != nil {
			return renderErrorPage(c, err.Error(), http.StatusBadRequest)
		}
		_, err = db.Secret.Create().
			SetName(name).
			SetValue(value).
			SetProject(p).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error creating secret", "error", err)
			return renderErrorPage(c, "Error creating secret", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%v/secrets", projectID))
	}
}

func hookDeleteSecret(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		secretID := c.Param("secret_id")
		secretIDInt, err := strconv.Atoi(secretID)
		if err != nil {
			return renderErrorPage(c, "Invalid secret ID", http.StatusBadRequest)
		}
		_, err = db.Secret.Delete().Where(secret.IDEQ(secretIDInt)).Exec(c.Request().Context())
		if err != nil {
			slog.Error("error deleting secret", "error", err)
			return renderErrorPage(c, "Error deleting secret", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%v/secrets", c.Param("project_id")))
	}
}
