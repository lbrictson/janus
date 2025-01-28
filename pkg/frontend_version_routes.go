package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobversion"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func renderJobVersionsPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FrontendVersion struct {
			*ent.JobVersion
			FriendlyCreatedTime string
		}
		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			slog.Warn("invalid job ID", "error", err)
			return renderErrorPage(c, "Invalid job ID", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.ID(jobIDInt)).WithProject().WithVersions().Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting job", "error", err)
			return renderErrorPage(c, "Error getting job", http.StatusInternalServerError)
		}
		vers, err := db.JobVersion.Query().Where(jobversion.HasJobWith(job.ID(jobIDInt))).Order(ent.Desc(jobversion.FieldCreatedAt)).All(c.Request().Context())
		versions := make([]*FrontendVersion, 0)
		for _, v := range vers {
			versions = append(versions, &FrontendVersion{
				JobVersion:          v,
				FriendlyCreatedTime: v.CreatedAt.Format(time.RFC1123),
			})
		}
		self, _ := getSelf(c, db)
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to view this version", http.StatusForbidden)
		}
		return c.Render(200, "job-versions", map[string]any{
			"Job":      j,
			"Versions": versions,
			"Project":  j.Edges.Project,
		})
	}
}

func renderSingleVersionView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FrontendVersion struct {
			*ent.JobVersion
			FriendlyCreatedTime string
		}
		versionID := c.Param("version_id")
		versionIDInt, err := strconv.Atoi(versionID)
		if err != nil {
			slog.Warn("invalid version ID", "error", err)
			return renderErrorPage(c, "Invalid version ID", http.StatusBadRequest)
		}
		v, err := db.JobVersion.Query().WithJob().Where(jobversion.ID(versionIDInt)).WithJob().Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting version", "error", err)
			return renderErrorPage(c, "Error getting version", http.StatusInternalServerError)
		}
		j, err := db.Job.Query().WithProject().Where(job.ID(v.Edges.Job.ID)).Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting job", "error", err)
			return renderErrorPage(c, "Error getting job", http.StatusInternalServerError)
		}
		vFE := &FrontendVersion{
			JobVersion:          v,
			FriendlyCreatedTime: v.CreatedAt.Format(time.RFC1123),
		}
		self, _ := getSelf(c, db)
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to view this version", http.StatusForbidden)
		}
		splitted := strings.Split(v.Script, "\n")
		return c.Render(200, "job-version-view", map[string]any{
			"Version":     vFE,
			"Job":         j,
			"Project":     j.Edges.Project,
			"ScriptLines": len(splitted),
		})
	}
}

func formRestoreVersionOfJob(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		versionID := c.Param("version_id")
		versionIDInt, err := strconv.Atoi(versionID)
		if err != nil {
			slog.Warn("invalid version ID", "error", err)
			return renderErrorPage(c, "Invalid version ID", http.StatusBadRequest)
		}
		v, err := db.JobVersion.Query().WithJob().Where(jobversion.ID(versionIDInt)).Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting version", "error", err)
			return renderErrorPage(c, "Error getting version", http.StatusInternalServerError)
		}
		j, err := db.Job.Query().WithProject().Where(job.ID(v.Edges.Job.ID)).Only(c.Request().Context())
		if err != nil {
			slog.Error("error getting job", "error", err)
			return renderErrorPage(c, "Error getting job", http.StatusInternalServerError)
		}
		self, _ := getSelf(c, db)
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to restore this version", http.StatusForbidden)
		}
		// Swap fields that are version controlled
		updatedJob, err := db.Job.UpdateOne(j).
			SetScript(v.Script).
			SetName(v.Name).
			SetDescription(v.Description).
			SetArguments(v.Arguments).
			SetCronSchedule(v.CronSchedule).
			SetScheduleEnabled(v.ScheduleEnabled).
			SetAllowConcurrentRuns(v.AllowConcurrentRuns).
			SetRequiresFileUpload(v.RequiresFileUpload).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error restoring version", "error", err)
			return renderErrorPage(c, "Error restoring version", http.StatusInternalServerError)
		}
		// Schedule cron if need be, if it was already scheduled that's fine
		if updatedJob.ScheduleEnabled {
			addCronJob(db, updatedJob)
		} else {
			removeCronJob(updatedJob)
		}
		slog.Info("restored version", "version", v.ID, "job", j.ID, "user", self.ID)
		// Create a new version
		_, err = db.JobVersion.Create().
			SetJobID(updatedJob.ID).
			SetScript(updatedJob.Script).
			SetName(updatedJob.Name).
			SetDescription(updatedJob.Description).
			SetArguments(updatedJob.Arguments).
			SetCronSchedule(updatedJob.CronSchedule).
			SetScheduleEnabled(updatedJob.ScheduleEnabled).
			SetAllowConcurrentRuns(updatedJob.AllowConcurrentRuns).
			SetRequiresFileUpload(updatedJob.RequiresFileUpload).
			SetChangedByEmail(self.Email).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error creating new version", "error", err)
			return renderErrorPage(c, "Error creating new version", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusFound, fmt.Sprintf("/projects/%v/jobs/%v/versions", j.Edges.Project.ID, j.ID))
	}
}
