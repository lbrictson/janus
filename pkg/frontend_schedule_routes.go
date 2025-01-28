package pkg

import (
	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"log/slog"
	"net/http"
)

func renderSchedulePage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type ScheduledJob struct {
			JobName      string
			ProjectName  string
			ProjectID    int
			CronSchedule string
			NextRun      string
			JobID        int
		}
		var scheduledJobs []ScheduledJob
		jobs, err := db.Job.Query().Where(job.ScheduleEnabledEQ(true)).WithProject().All(c.Request().Context())
		if err != nil {
			slog.Error("failed to get scheduled jobs", "error", err)
			return renderErrorPage(c, "Failed to get scheduled jobs", http.StatusInternalServerError)
		}
		for _, j := range jobs {
			next := ""
			// Check if job is in the crons map
			cronJobMapLock.Lock()
			entryID, ok := cronJobs[j.ID]
			cronJobMapLock.Unlock()
			if !ok {
				continue
			} else {
				entry := cronService.Entry(entryID)
				next = humanize.Time(entry.Next)
			}
			scheduledJobs = append(scheduledJobs, ScheduledJob{
				JobName:      j.Name,
				ProjectName:  j.Edges.Project.Name,
				ProjectID:    j.Edges.Project.ID,
				CronSchedule: j.CronSchedule,
				NextRun:      next,
				JobID:        j.ID,
			})
		}
		return c.Render(200, "scheduled-jobs", map[string]interface{}{
			"ScheduledJobs": scheduledJobs,
		})
	}
}
