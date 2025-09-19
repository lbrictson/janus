package pkg

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobhistory"
	"github.com/lbrictson/janus/ent/project"
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
		type JobItemFrontend struct {
			ent.Job
			LastRunTimeFriendly string
			LastRunState        string
		}
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
		j, err := db.Job.Query().WithProject().WithHistory().Where(job.HasProjectWith(project.IDEQ(projIDInt))).Order(ent.Asc(job.FieldName)).All(c.Request().Context())
		if err != nil {
			slog.Error("error getting jobs for project", "error", err)
			return renderErrorPage(c, "Error getting jobs for project", http.StatusInternalServerError)
		}
		jobs := make([]JobItemFrontend, len(j))
		for i, jo := range j {
			jobs[i] = JobItemFrontend{
				Job: *jo,
			}
			// Get the history
			h, _ := db.JobHistory.Query().Where(jobhistory.HasJobWith(job.ID(jo.ID))).Order(ent.Desc(jobhistory.FieldCreatedAt)).First(c.Request().Context())
			if h != nil {
				jobs[i].LastRunTimeFriendly = humanize.Time(h.CreatedAt)
				jobs[i].LastRunState = h.Status
			} else {
				jobs[i].LastRunTimeFriendly = "Never"
				jobs[i].LastRunState = "No Data"
			}
		}
		return c.Render(http.StatusOK, "project-view", map[string]any{
			"Project": p,
			"Jobs":    jobs,
			"CanEdit": canEdit,
		})
	}
}

func hookDeleteProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projID := c.Param("id")
		projIDInt, err := strconv.Atoi(projID)
		if err != nil {
			slog.Error("error parsing project ID", "error", err)
			return renderErrorPage(c, "Error parsing project ID", http.StatusBadRequest)
		}
		err = db.Project.DeleteOneID(projIDInt).Exec(c.Request().Context())
		if err != nil {
			slog.Error("error deleting project", "error", err)
			return renderErrorPage(c, "Error deleting project", http.StatusInternalServerError)
		}
		slog.Info("deleted project", "project", projIDInt, "user", c.Get("userID"))
		return c.Redirect(http.StatusSeeOther, "/")
	}
}

func renderEditProjectPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projID := c.Param("project_id")
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
		return c.Render(http.StatusOK, "edit-project", map[string]any{
			"Project": p,
		})
	}
}

func formEditProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectID := c.Param("project_id")
		projectIDInt, err := strconv.Atoi(projectID)
		if err != nil {
			slog.Error("error parsing project ID", "error", err)
			return renderErrorPage(c, "Error parsing project ID", http.StatusBadRequest)
		}

		type Form struct {
			Name        string `form:"name"`
			Description string `form:"description"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("error binding form for edit project", "error", err)
			return c.Render(http.StatusBadRequest, "edit-project", map[string]any{
				"Error": "Error binding form",
			})
		}
		p, err := db.Project.Get(c.Request().Context(), projectIDInt)
		if err != nil {
			slog.Error("error getting project", "error", err)
			return renderErrorPage(c, "Error getting project", http.StatusInternalServerError)
		}
		p, err = p.Update().SetName(form.Name).SetDescription(form.Description).Save(c.Request().Context())
		if err != nil {
			slog.Error("error updating project", "error", err)
			return c.Render(http.StatusInternalServerError, "edit-project", map[string]any{
				"Error": "Error updating project",
			})
		}
		slog.Info("updated project", "project", p, "user", c.Get("userID"))
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%d", projectIDInt))
	}
}
