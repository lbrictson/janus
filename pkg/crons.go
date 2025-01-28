package pkg

import (
	"context"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobhistory"
	"github.com/robfig/cron/v3"
	"log/slog"
	"sync"
	"time"
)

var cronJobs map[int]cron.EntryID
var cronJobMapLock = &sync.Mutex{}
var cronService *cron.Cron

func init() {
	cronService = cron.New()
	cronJobs = make(map[int]cron.EntryID)
	cronService.Start()
}

func cronJobWrapper(db *ent.Client, j ent.Job) func() {
	return func() {
		// Get the job to make sure it hasn't been deleted
		freshJob, err := db.Job.Query().WithProject().Where(job.IDEQ(j.ID)).Only(context.Background())
		if err != nil {
			slog.Error("failed to get job to run as cron, that is unexpected", "error", err)
			return
		}
		historyItem, err := db.JobHistory.Create().SetJob(freshJob).
			SetProject(freshJob.Edges.Project).
			SetTrigger("Scheduler").
			SetStatus("running").
			SetCreatedAt(time.Now()).
			SetWasSuccessful(false).
			SetDurationMs(0).
			SetExitCode(0).
			SetTriggeredByEmail("SYSTEM").
			SetTriggeredByID(0).
			Save(context.Background())
		if err != nil {
			slog.Error("failed to create history item", "error", err)
			return
		}
		args := []JobRuntimeArg{}
		for _, x := range freshJob.Arguments {
			if x.DefaultValue == "" {
				slog.Warn("cron job was scheduled to run but had no default value for arguments, cancelled run", "job_id", freshJob.ID, "job_name", freshJob.Name)
				db.JobHistory.Update().Where(jobhistory.IDEQ(historyItem.ID)).SetStatus("failed").SetWasSuccessful(false).SetOutput("Job failed due to having one or more arguments without a defautl value").Save(context.Background())
				return
			}
			args = append(args, JobRuntimeArg{
				Name:      x.Name,
				Value:     x.DefaultValue,
				Sensitive: x.Sensitive,
			})
		}
		// Run the job
		c, _ := LoadConfig()
		slog.Info("executing scheduled run of job", "job_id", freshJob.ID, "job_name", freshJob.Name, "project_id", freshJob.Edges.Project.ID, "project_name", freshJob.Edges.Project.Name)
		runJob(db, freshJob, historyItem, args, nil, *c)
	}
}

func addCronJob(db *ent.Client, j *ent.Job) {
	cronJobMapLock.Lock()
	// See if the job is already in the map
	entryID, ok := cronJobs[j.ID]
	if ok {
		// Remove it first so we can re-add it
		cronService.Remove(entryID)
		delete(cronJobs, j.ID)
	}
	entry, err := cronService.AddFunc(j.CronSchedule, cronJobWrapper(db, *j))
	if err != nil {
		slog.Error("failed to add cron job", "error", err)
	}
	cronJobs[j.ID] = entry
	cronJobMapLock.Unlock()
	slog.Info("added cron job", "job_id", j.ID)
}

func removeCronJob(j *ent.Job) {
	cronJobMapLock.Lock()
	entryID, ok := cronJobs[j.ID]
	if ok {
		cronService.Remove(entryID)
		delete(cronJobs, j.ID)
	}
	cronJobMapLock.Unlock()
}
