package pkg

import (
	"context"
	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobhistory"
	"github.com/lbrictson/janus/ent/project"
	"net/http"
	"sync"
	"time"
)

func renderDashboard(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type ProjectItem struct {
			ID           int
			Name         string
			Description  string
			LastJobRun   string
			NumberOfJobs int
			CanAccess    bool
		}
		allProjects, err := db.Project.Query().WithJobs().Order(ent.Desc(project.FieldName)).All(c.Request().Context())
		if err != nil {
			return c.Render(http.StatusInternalServerError, "dashboard", map[string]any{
				"Error": "Error getting projects from database",
			})
		}
		projects := make([]ProjectItem, len(allProjects))
		// Create a wait group of all the projects to gather stats from
		// and then populate the projects slice with the data
		projectWG := sync.WaitGroup{}
		projectWG.Add(len(allProjects))
		for p := range allProjects {
			go func(p int) {
				projLastRunJob, _ := db.Job.Query().Where(job.HasProjectWith(project.ID(allProjects[p].ID))).Order(ent.Desc(job.FieldLastRunTime)).First(c.Request().Context())
				lastRun := "Never"
				if projLastRunJob != nil {
					lastRun = humanize.Time(projLastRunJob.LastRunTime)
				}
				projects[p] = ProjectItem{
					ID:           allProjects[p].ID,
					Name:         allProjects[p].Name,
					Description:  allProjects[p].Description,
					LastJobRun:   lastRun,
					NumberOfJobs: len(allProjects[p].Edges.Jobs),
				}
				history, _ := db.JobHistory.Query().Where(jobhistory.HasProjectWith(project.ID(allProjects[p].ID))).Order(ent.Desc(jobhistory.FieldCreatedAt)).First(c.Request().Context())
				if history == nil {
					projects[p].LastJobRun = "Never"
				} else {
					projects[p].LastJobRun = humanize.Time(history.CreatedAt)
				}
				projectWG.Done()
			}(p)
		}
		// Create a wait group to run the calculations concurrently
		wg := sync.WaitGroup{}
		wg.Add(5)
		var successfulJobsLast24Hours int
		var failedJobsLast24Hours int
		var runningJobs int
		var scheduledJobs int
		var userSpecificPermissions map[int]string
		go func() {
			successfulJobsLast24Hours = calculateSuccessfulRunsLast24Hours(db)
			wg.Done()
		}()
		go func() {
			failedJobsLast24Hours = calculateFailedRunsLast24Hours(db)
			wg.Done()
		}()
		go func() {
			runningJobs = calculateRunningJobs(db)
			wg.Done()
		}()
		go func() {
			self, _ := getSelf(c, db)
			userSpecificPermissions, _, _ = getUserProjectPermissions(db, self.ID)
			wg.Done()
		}()
		go func() {
			scheduledJobs = calculateScheduledJobs(db)
			wg.Done()
		}()
		wg.Wait()
		projectWG.Wait()
		for i := range projects {
			if userSpecificPermissions[projects[i].ID] != "None" {
				projects[i].CanAccess = true
			}
		}
		return c.Render(http.StatusOK, "dashboard", map[string]any{
			"Projects":       projects,
			"SuccessfulJobs": successfulJobsLast24Hours,
			"FailedJobs":     failedJobsLast24Hours,
			"RunningJobs":    runningJobs,
			"ScheduledJobs":  scheduledJobs,
		})
	}
}

func calculateSuccessfulRunsLast24Hours(db *ent.Client) int {
	jobsSuccessfulLast24Hours, err := db.JobHistory.Query().Where(jobhistory.WasSuccessful(true), jobhistory.CreatedAtGTE(time.Now().Add(-24*time.Hour)), jobhistory.DurationMsGTE(1)).Count(context.Background())
	if err != nil {
		return 0
	}
	return jobsSuccessfulLast24Hours
}

func calculateFailedRunsLast24Hours(db *ent.Client) int {
	jobFailedLast24Hours, err := db.JobHistory.Query().Where(jobhistory.WasSuccessful(false), jobhistory.CreatedAtGTE(time.Now().Add(-24*time.Hour)), jobhistory.DurationMsGTE(1)).Count(context.Background())
	if err != nil {
		return 0
	}
	return jobFailedLast24Hours
}

func calculateRunningJobs(db *ent.Client) int {
	runningJobs, err := db.JobHistory.Query().Where(jobhistory.Status("running")).Count(context.Background())
	if err != nil {
		return 0
	}
	return runningJobs
}

func calculateScheduledJobs(db *ent.Client) int {
	scheduledJobs, err := db.Job.Query().Where(job.ScheduleEnabledEQ(true)).Count(context.Background())
	if err != nil {
		return 0
	}
	return scheduledJobs
}
