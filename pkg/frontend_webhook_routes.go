package pkg

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/inboundwebhook"
	"github.com/lbrictson/janus/ent/job"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func renderWebooksView(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FEWebhook struct {
			ID                int
			Key               string
			CreatedBy         string
			URL               string
			CreatedAtFriendly string
			JobLink           string
			JobName           string
			ProjectName       string
		}
		webhooks, err := db.InboundWebhook.Query().WithJob().All(c.Request().Context())
		if err != nil {
			slog.Error("failed to get webhooks from database", "error", err)
			return renderErrorPage(c, "Error getting webhooks from database", http.StatusInternalServerError)
		}
		feWebhooks := make([]FEWebhook, len(webhooks))
		for i, w := range webhooks {
			j, err := db.Job.Query().WithProject().Where(job.ID(w.Edges.Job.ID)).Only(c.Request().Context())
			if err != nil {
				slog.Error("failed to get job from database", "error", err)
				return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
			}
			feWebhooks[i] = FEWebhook{
				ID:                w.ID,
				Key:               w.Key,
				CreatedBy:         w.CreatedBy,
				URL:               fmt.Sprintf("%v/inbound-webhook/%v", config.ServerURL, w.Key),
				CreatedAtFriendly: w.CreatedAt.Format("2006-01-02 15:04:05"),
				JobLink:           fmt.Sprintf("/projects/%v/jobs/%v/history", j.Edges.Project.ID, j.ID),
				JobName:           j.Name,
				ProjectName:       j.Edges.Project.Name,
			}
		}
		return c.Render(200, "webhooks", map[string]interface{}{
			"Webhooks": feWebhooks,
		})
	}
}

func renderNewWebhookView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		jobs, err := db.Job.Query().WithProject().All(c.Request().Context())
		if err != nil {
			slog.Error("failed to get jobs from database", "error", err)
			return renderErrorPage(c, "Error getting jobs from database", http.StatusInternalServerError)
		}
		return c.Render(200, "new-webhook", map[string]interface{}{
			"Jobs": jobs,
		})
	}
}

func formNewWebhook(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		j := c.FormValue("job_id")
		jID, err := strconv.Atoi(j)
		if err != nil {
			slog.Error("failed to convert job ID to int", "error", err)
			return renderErrorPage(c, "Error converting job ID to int", http.StatusInternalServerError)
		}
		job, err := db.Job.Query().Where(job.ID(jID)).Only(c.Request().Context())
		if err != nil {
			slog.Error("failed to get job from database", "error", err)
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		key := generateLongString() + generateLongString()
		_, err = db.InboundWebhook.Create().
			SetKey(key).
			SetCreatedBy(c.Get("email").(string)).
			SetJob(job).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("failed to create webhook", "error", err)
			return renderErrorPage(c, "Error creating webhook", http.StatusInternalServerError)
		}
		return c.Redirect(302, "/webhooks")
	}
}

func hookDeleteWebhook(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		i, err := strconv.Atoi(id)
		if err != nil {
			slog.Error("failed to convert webhook ID to int", "error", err)
			return renderErrorPage(c, "Error converting webhook ID to int", http.StatusInternalServerError)
		}
		err = db.InboundWebhook.DeleteOneID(i).Exec(c.Request().Context())
		if err != nil {
			slog.Error("failed to delete webhook", "error", err)
			return renderErrorPage(c, "Error deleting webhook", http.StatusInternalServerError)
		}
		return c.Redirect(302, "/webhooks")
	}
}

func handleIncomingJobWebhookTrigger(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Param("key")
		webhook, err := db.InboundWebhook.Query().WithJob().Where(inboundwebhook.Key(key)).Only(c.Request().Context())
		if err != nil {
			slog.Error("failed to get webhook from database", "error", err)
			return c.String(500, "invalid webhook parameter")
		}
		j, err := db.Job.Query().WithProject().Where(job.ID(webhook.Edges.Job.ID)).Only(c.Request().Context())
		if err != nil {
			slog.Error("failed to get job from database", "error", err)
			return c.String(500, "invalid webhook parameter")
		}
		if j.RequiresFileUpload {
			db.JobHistory.Create().
				SetJob(j).
				SetProject(j.Edges.Project).
				SetTrigger("Inbound Webhook").
				SetStatus("failed").
				SetCreatedAt(time.Now()).
				SetWasSuccessful(false).
				SetDurationMs(0).
				SetExitCode(0).
				SetTriggeredByEmail(webhook.CreatedBy).
				SetTriggeredByID(0).
				SetOutput("job requires file upload which cannot work with webhook triggers").
				Save(context.Background())
			return c.String(500, "job requires file upload")
		}
		doesItHaveAnArgumentWithoutAdefaultValue := false
		var argValues []JobRuntimeArg
		if c.Request().Method == "POST" {
			// Get the body of the request
			body, err := io.ReadAll(c.Request().Body)
			if err != nil {
				slog.Error("failed to read body of request", "error", err)
				return c.String(500, "error reading body of request")
			}
			s := string(body)
			argValues = append(argValues, JobRuntimeArg{
				Name:  "WEBHOOK_PAYLOAD",
				Value: s,
			})
		}

		for _, arg := range j.Arguments {
			if arg.DefaultValue == "" {
				doesItHaveAnArgumentWithoutAdefaultValue = true
			} else {
				argValues = append(argValues, JobRuntimeArg{
					Name:  arg.Name,
					Value: arg.DefaultValue,
				})
			}
		}
		if doesItHaveAnArgumentWithoutAdefaultValue {
			db.JobHistory.Create().
				SetJob(j).
				SetProject(j.Edges.Project).
				SetTrigger("Inbound Webhook").
				SetStatus("failed").
				SetCreatedAt(time.Now()).
				SetWasSuccessful(false).
				SetDurationMs(0).
				SetExitCode(0).
				SetTriggeredByEmail(webhook.CreatedBy).
				SetTriggeredByID(0).
				SetOutput("job requires arguments to be passed in").
				Save(context.Background())
			return c.String(500, "job requires arguments that have no default value which is not supported by webhook triggers")
		}
		history, err := db.JobHistory.Create().
			SetJob(j).
			SetProject(j.Edges.Project).
			SetTrigger("Inbound Webhook").
			SetStatus("running").
			SetCreatedAt(time.Now()).
			SetWasSuccessful(false).
			SetDurationMs(0).
			SetExitCode(0).
			SetTriggeredByEmail(webhook.CreatedBy).
			SetTriggeredByID(0).
			Save(context.Background())
		if err != nil {
			slog.Error("error creating job history", "error", err)
			return renderErrorPage(c, "Error creating job history", http.StatusInternalServerError)
		}
		go runJob(db, j, history, argValues, nil, *config)
		return c.String(http.StatusOK, fmt.Sprintf("%v/projects/%v/jobs/%v/run/%v", config.ServerURL, j.Edges.Project.ID, j.ID, history.ID))
	}
}
