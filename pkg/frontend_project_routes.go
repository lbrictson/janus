package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/project"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

func renderNewProjectView(c echo.Context) error {
	return c.Render(http.StatusOK, "new-project", nil)
}

func formNewProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Form struct {
			Name        string `form:"name"`
			Description string `form:"description"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("error binding form for create new project", "error", err)
			return c.Render(http.StatusBadRequest, "new-project", map[string]any{
				"Error": "Error binding form",
			})
		}
		allProjects, err := db.Project.Query().All(c.Request().Context())
		if err != nil {
			slog.Error("error querying all projects", "error", err)
			return c.Render(http.StatusInternalServerError, "new-project", map[string]any{
				"Error": "Error getting projects from database",
			})
		}
		for _, project := range allProjects {
			if strings.ToLower(project.Name) == strings.ToLower(form.Name) {
				return c.Render(http.StatusBadRequest, "new-project", map[string]any{
					"Error": "Project with that name already exists",
				})
			}
		}
		newProject, err := db.Project.Create().
			SetName(form.Name).
			SetDescription(form.Description).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error creating new project", "error", err)
			return c.Render(http.StatusInternalServerError, "new-project", map[string]any{
				"Error": "Error creating new project",
			})
		}
		// When creating a new project the creator should be an editor
		self, _ := getSelf(c, db)
		db.ProjectUser.Create().SetProject(newProject).SetUserID(self.ID).SetCanEdit(true).Save(c.Request().Context())
		slog.Info("created new project", "project", newProject, "user", c.Get("userID"))
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%d", newProject.ID))
	}
}

func renderProjectViewPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Determine if the user can view this page
		self, _ := getSelf(c, db)
		userPermissions, _, err := getUserProjectPermissions(db, self.ID)

		projID := c.Param("id")
		projIDInt, err := strconv.Atoi(projID)
		if err != nil {
			slog.Error("error parsing project ID", "error", err)
			return renderErrorPage(c, "Error parsing project ID", http.StatusBadRequest)
		}
		p, err := db.Project.Get(c.Request().Context(), projIDInt)
		if err != nil {
			slog.Error("error getting project", "error", err)
			return renderErrorPage(c, "Error getting project", http.StatusInternalServerError)
		}
		canEdit := false
		if userPermissions[projIDInt] == "None" {
			return renderErrorPage(c, "You do not have permission to view this project", http.StatusForbidden)
		}
		if userPermissions[projIDInt] == "Edit" {
			canEdit = true
		}
		j, err := db.Job.Query().WithProject().Where(job.HasProjectWith(project.IDEQ(projIDInt))).All(c.Request().Context())
		if err != nil {
			slog.Error("error getting jobs for project", "error", err)
			return renderErrorPage(c, "Error getting jobs for project", http.StatusInternalServerError)
		}
		return c.Render(http.StatusOK, "project-view", map[string]any{
			"Project": p,
			"Jobs":    j,
			"CanEdit": canEdit,
		})
	}
}
