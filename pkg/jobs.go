package pkg

import (
	"context"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/jobhistory"
	"log/slog"
	"time"
)

func RunStaleJobCleaner(db *ent.Client, config Config) {
	// Get all jobs that are running
	for {
		runningJobs, err := db.JobHistory.Query().WithJob().Where(jobhistory.StatusEQ("running")).All(context.Background())
		if err != nil {
			slog.Error("failed to get running jobs", "error", err)
			time.Sleep(5 * time.Second)
			continue
		}
		for _, job := range runningJobs {
			expires := job.CreatedAt.Add(time.Duration(job.Edges.Job.TimeoutSeconds) * time.Second)
			// If the job has been running for more than the timeout, mark it as failed
			if time.Now().After(expires) {
				slog.Warn("marking job as failed due to timeout", "job_id", job.ID)
				o := job.Output
				o = o + "\n\nJob failed due to timeout"
				_, updateErr := db.JobHistory.UpdateOne(job).SetStatus("failed").
					SetOutput(o).
					SetWasSuccessful(false).
					Save(context.Background())
				if updateErr != nil {
					slog.Error("failed to update job status", "error", updateErr)
				}
				totalJobFailures.Inc()
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func RunJobCleaner(db *ent.Client, config Config) {
	for {
		c, err := db.DataConfig.Query().First(context.Background())
		if err != nil {
			slog.Error("failed to get data config", "error", err)
			time.Sleep(1 * time.Hour)
			continue
		}
		// Get all jobs histories that are older than the configured number of days
		expirationTime := time.Now().Add(-time.Duration(c.DaysToKeep) * 24 * time.Hour)
		oldJobs, err := db.JobHistory.Query().Where(jobhistory.CreatedAtLT(expirationTime)).All(context.Background())
		if err != nil {
			slog.Error("failed to get old jobs", "error", err)
			time.Sleep(1 * time.Hour)
			continue
		}
		for _, job := range oldJobs {
			err = db.JobHistory.DeleteOne(job).Exec(context.Background())
			if err != nil {
				slog.Error("failed to delete job", "error", err)
			}
		}
		time.Sleep(1 * time.Hour)
	}
}
