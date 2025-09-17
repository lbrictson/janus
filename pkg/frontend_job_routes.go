package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/jobhistory"
	"github.com/lbrictson/janus/ent/schema"
	"github.com/robfig/cron/v3"
)

func renderCreateJobView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		jobConfig, _ := getJobConfig(c.Request().Context(), db)
		projectID := c.Param("project_id")
		projectIDInt, err := strconv.Atoi(projectID)
		if err != nil {
			return renderErrorPage(c, "Error converting project ID to integer", http.StatusBadRequest)
		}
		self, _ := getSelf(c, db)
		if canUserEditProject(db, self.ID, projectIDInt) == false {
			return renderErrorPage(c, "You do not have permission to edit this project", http.StatusForbidden)
		}
		p, err := db.Project.Get(c.Request().Context(), projectIDInt)
		if err != nil {
			return renderErrorPage(c, "Error getting project from database", http.StatusInternalServerError)
		}
		n, err := db.NotificationChannel.Query().All(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting notification channels", http.StatusInternalServerError)
		}
		return c.Render(http.StatusOK, "create-job", map[string]any{
			"Project":              p,
			"NotificationChannels": n,
			"DefaultTimeout":       jobConfig.DefaultTimeoutSeconds,
		})
	}
}

func formCreateJob(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		// Define our form structure to use Echo's binding
		type JobArgForm struct {
			Name          string `form:"arg_names[]"`
			DefaultValue  string `form:"arg_defaults[]"`
			AllowedValues string `form:"arg_allowed_values[]"`
			Sensitive     string `form:"arg_sensitive[]"`
		}

		type Form struct {
			Name                      string `form:"name"`
			Description               string `form:"description"`
			TimeoutSeconds            int    `form:"timeout_seconds"`
			CronSchedule              string `form:"cron_schedule"`
			ScheduleEnabled           string `form:"schedule_enabled"`
			AllowConcurrentRuns       string `form:"allow_concurrent_runs"`
			RequiresFileUpload        string `form:"requires_file_upload"`
			NotifyOnStartChannelIDs   []int  `form:"notify_on_start_channel_ids"`
			NotifyOnSuccessChannelIDs []int  `form:"notify_on_success_channel_ids"`
			NotifyOnFailureChannelIDs []int  `form:"notify_on_failure_channel_ids"`
			Arguments                 []JobArgForm
			Script                    string `form:"script"`
		}

		projectID := c.Param("project_id")
		scheduleEnabled := c.FormValue("schedule_enabled") == "on"
		allowConcurrentRuns := c.FormValue("allow_concurrent_runs") == "on"
		requiresFileUpload := c.FormValue("requires_file_upload") == "on"
		// Parse the form
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("error binding form data", "error", err)
			return renderErrorPage(c, "Invalid form data", http.StatusBadRequest)
		}
		// Default timeout if not set
		if form.TimeoutSeconds == 0 {
			form.TimeoutSeconds = 3600 // Default 1 hour
		}

		// Get project
		projectIDInt, err := strconv.Atoi(projectID)
		if err != nil {
			return renderErrorPage(c, "Error converting project ID to integer", http.StatusBadRequest)
		}
		if !canUserEditProject(db, self.ID, projectIDInt) {
			return renderErrorPage(c, "You do not have permission to create jobs", http.StatusForbidden)
		}
		project, err := db.Project.Get(c.Request().Context(), projectIDInt)
		if err != nil {
			slog.Error("error getting project", "error", err)
			return renderErrorPage(c, "Error getting project", http.StatusInternalServerError)
		}

		// Process arguments from form data
		argNames := c.Request().Form["arg_names[]"]
		argTypes := c.Request().Form["arg_types[]"]
		argDefaults := c.Request().Form["arg_defaults[]"]
		argAllowedValues := c.Request().Form["arg_allowed_values[]"]
		argSensitive := c.Request().Form["arg_sensitive[]"]
		// Create a map of sensitive indexes
		sensitiveMap := make(map[int]bool)
		for _, v := range argSensitive {
			// The value will be the string representation of the index
			if idx, err := strconv.Atoi(v); err == nil {
				sensitiveMap[idx] = true
			}
		}
		doesJobHaveArgumentWithoutDefaultValue := false
		needToScheduleCron := false
		arguments := make([]schema.JobArgument, 0)
		for i := 0; i < len(argNames); i++ {
			if argNames[i] == "" {
				continue
			}

			var allowedValues []string
			if argAllowedValues[i] != "" {
				allowedValues = strings.Split(argAllowedValues[i], ",")
				for j := range allowedValues {
					allowedValues[j] = strings.TrimSpace(allowedValues[j])
				}
			}

			if len(allowedValues) > 0 {
				if argDefaults[i] != "" {
					defaultValueValid := false
					for _, v := range allowedValues {
						if v == argDefaults[i] {
							defaultValueValid = true
							break
						}
					}
					if !defaultValueValid {
						return renderErrorPage(c, "Default value must be one of the allowed values for argument "+argNames[i], http.StatusBadRequest)
					}
				} else {
					return renderErrorPage(c, "Default value is required for argument "+argNames[i], http.StatusBadRequest)
				}
			}
			if argDefaults[i] == "" {
				doesJobHaveArgumentWithoutDefaultValue = true
			}
			argType := "string" // Default type
			if i < len(argTypes) && argTypes[i] != "" {
				argType = argTypes[i]
			}
			arguments = append(arguments, schema.JobArgument{
				Name:          argNames[i],
				Type:          argType,
				DefaultValue:  argDefaults[i],
				AllowedValues: allowedValues,
				Sensitive:     sensitiveMap[i],
			})
		}
		cronNextRunTime := time.Now()
		if scheduleEnabled {
			if doesJobHaveArgumentWithoutDefaultValue {
				return renderErrorPage(c, "All arguments must have a default value when scheduling is enabled", http.StatusBadRequest)
			}
			if form.CronSchedule != "" {
				n, cronParseError := cron.ParseStandard(form.CronSchedule)
				if cronParseError != nil {
					slog.Error("error parsing cron schedule", "error", cronParseError)
					return renderErrorPage(c, "Error parsing cron schedule "+cronParseError.Error(), http.StatusBadRequest)
				}
				cronNextRunTime = n.Next(time.Now())
			} else {
				return renderErrorPage(c, "Cron schedule is required when schedule is enabled", http.StatusBadRequest)
			}
			needToScheduleCron = true
		}
		var startChannels []int
		var successChannels []int
		var failureChannels []int

		// Note the [] after the field name - this is important for multiple values
		startChannelIDs := c.Request().Form["notify_on_start_channel_ids[]"]
		successChannelIDs := c.Request().Form["notify_on_success_channel_ids[]"]
		failureChannelIDs := c.Request().Form["notify_on_failure_channel_ids[]"]

		// Process start channels
		for _, id := range startChannelIDs {
			if channelID, err := strconv.Atoi(id); err == nil {
				startChannels = append(startChannels, channelID)
			}
		}

		// Process success channels
		for _, id := range successChannelIDs {
			if channelID, err := strconv.Atoi(id); err == nil {
				successChannels = append(successChannels, channelID)
			}
		}

		// Process failure channels
		for _, id := range failureChannelIDs {
			if channelID, err := strconv.Atoi(id); err == nil {
				failureChannels = append(failureChannels, channelID)
			}
		}
		// Create the job
		job, err := db.Job.Create().
			SetName(form.Name).
			SetDescription(form.Description).
			SetTimeoutSeconds(form.TimeoutSeconds).
			SetCronSchedule(form.CronSchedule).
			SetScheduleEnabled(scheduleEnabled).
			SetAllowConcurrentRuns(allowConcurrentRuns).
			SetArguments(arguments).
			SetRequiresFileUpload(requiresFileUpload).
			SetNotifyOnStartChannelIds(startChannels).
			SetNotifyOnSuccessChannelIds(successChannels).
			SetNotifyOnFailureChannelIds(failureChannels).
			SetProject(project).
			SetScript(form.Script).
			SetLastEditTime(time.Now()).
			SetLastRunTime(time.Now()).
			SetLastRunSuccess(true).
			SetNextCronRunTime(cronNextRunTime).
			Save(c.Request().Context())

		if err != nil {
			slog.Error("failed to create job",
				"error", err,
				"project_id", projectID,
				"job_name", form.Name,
			)
			return renderErrorPage(c, "Error creating job", http.StatusInternalServerError)
		}

		slog.Info("job created successfully",
			"job_id", job.ID,
			"project_id", projectID,
			"job_name", form.Name,
		)
		if needToScheduleCron {
			addCronJob(db, job)
		}
		// Create the version
		_, err = db.JobVersion.Create().
			SetJob(job).
			SetName(job.Name).
			SetDescription(job.Description).
			SetCronSchedule(job.CronSchedule).
			SetScheduleEnabled(job.ScheduleEnabled).
			SetScript(job.Script).
			SetAllowConcurrentRuns(job.AllowConcurrentRuns).
			SetArguments(arguments).
			SetRequiresFileUpload(job.RequiresFileUpload).
			SetChangedByEmail(self.Email).
			SetCreatedAt(time.Now()).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error creating job version", "error", err)
			return renderErrorPage(c, "Error creating job version", http.StatusInternalServerError)
		}
		// Redirect to the project view page
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%s", projectID))
	}
}

func renderEditJobPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FrontendNotificationChannel struct {
			ent.NotificationChannel
			Selected bool
		}
		self, _ := getSelf(c, db)
		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			return renderErrorPage(c, "Error converting job ID to integer", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.IDEQ(jobIDInt)).WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to edit this job", http.StatusForbidden)
		}
		allNotificationChannels, err := db.NotificationChannel.Query().All(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting notification channels", http.StatusInternalServerError)
		}
		failureChannels := make([]FrontendNotificationChannel, 0)
		successChannels := make([]FrontendNotificationChannel, 0)
		startChannels := make([]FrontendNotificationChannel, 0)
		for _, nc := range allNotificationChannels {
			failureChannels = append(failureChannels, FrontendNotificationChannel{
				NotificationChannel: *nc,
				Selected:            false,
			})
			successChannels = append(successChannels, FrontendNotificationChannel{
				NotificationChannel: *nc,
				Selected:            false,
			})
			startChannels = append(startChannels, FrontendNotificationChannel{
				NotificationChannel: *nc,
				Selected:            false,
			})
		}
		for _, nc := range j.NotifyOnFailureChannelIds {
			for i := range failureChannels {
				if failureChannels[i].ID == nc {
					failureChannels[i].Selected = true
				}
			}
		}
		for _, nc := range j.NotifyOnSuccessChannelIds {
			for i := range successChannels {
				if successChannels[i].ID == nc {
					successChannels[i].Selected = true
				}
			}
		}
		for _, nc := range j.NotifyOnStartChannelIds {
			for i := range startChannels {
				if startChannels[i].ID == nc {
					startChannels[i].Selected = true
				}
			}
		}
		lines := strings.Split(j.Script, "\n")
		totalLinesInScript := len(lines)
		if totalLinesInScript < 6 {
			totalLinesInScript = 6
		}
		existingArgs, err := json.Marshal(j.Arguments)
		if err != nil {
			return renderErrorPage(c, "Error marshalling job arguments", http.StatusInternalServerError)
		}
		return c.Render(http.StatusOK, "edit-job", map[string]any{
			"Job":             j,
			"Project":         j.Edges.Project,
			"FailureChannels": failureChannels,
			"SuccessChannels": successChannels,
			"StartChannels":   startChannels,
			"ExistingArgs":    template.HTML(existingArgs),
			"ScriptLines":     totalLinesInScript,
		})
	}
}

func formEditJob(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		// Define our form structure to use Echo's binding
		type JobArgForm struct {
			Name          string `form:"arg_names[]"`
			DefaultValue  string `form:"arg_defaults[]"`
			AllowedValues string `form:"arg_allowed_values[]"`
			Sensitive     string `form:"arg_sensitive[]"`
		}

		type Form struct {
			Name                      string `form:"name"`
			Description               string `form:"description"`
			TimeoutSeconds            int    `form:"timeout_seconds"`
			CronSchedule              string `form:"cron_schedule"`
			ScheduleEnabled           string `form:"schedule_enabled"`
			AllowConcurrentRuns       string `form:"allow_concurrent_runs"`
			RequiresFileUpload        string `form:"requires_file_upload"`
			NotifyOnStartChannelIDs   []int  `form:"notify_on_start_channel_ids"`
			NotifyOnSuccessChannelIDs []int  `form:"notify_on_success_channel_ids"`
			NotifyOnFailureChannelIDs []int  `form:"notify_on_failure_channel_ids"`
			Arguments                 []JobArgForm
			Script                    string `form:"script"`
		}

		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			return renderErrorPage(c, "Error converting job ID to integer", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.IDEQ(jobIDInt)).WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to edit this job", http.StatusForbidden)
		}
		// Parse the form
		var form Form
		if err := c.Bind(&form); err != nil {
			slog.Error("error binding form data", "error", err)
			return renderErrorPage(c, "Invalid form data", http.StatusBadRequest)
		}
		// Default timeout if not set
		if form.TimeoutSeconds == 0 {
			form.TimeoutSeconds = 3600 // Default 1 hour
		}
		scheduleEnabled := c.FormValue("schedule_enabled") == "on"
		allowConcurrentRuns := c.FormValue("allow_concurrent_runs") == "on"
		requiresFileUpload := c.FormValue("requires_file_upload") == "on"
		argNames := c.Request().Form["arg_names[]"]
		argTypes := c.Request().Form["arg_types[]"]
		argDefaults := c.Request().Form["arg_defaults[]"]
		argAllowedValues := c.Request().Form["arg_allowed_values[]"]
		argSensitive := c.Request().Form["arg_sensitive[]"]
		sensitiveMap := make(map[int]bool)
		for _, v := range argSensitive {
			// The value will be the string representation of the index
			if idx, err := strconv.Atoi(v); err == nil {
				sensitiveMap[idx] = true
			}
		}
		needToScheduleCron := false
		doesJobHaveArgumentWithoutDefaultValue := false
		arguments := make([]schema.JobArgument, 0)
		for i := 0; i < len(argNames); i++ {
			if argNames[i] == "" {
				continue
			}

			var allowedValues []string
			if argAllowedValues[i] != "" {
				allowedValues = strings.Split(argAllowedValues[i], ",")
				for j := range allowedValues {
					allowedValues[j] = strings.TrimSpace(allowedValues[j])
				}
			}

			if len(allowedValues) > 0 {
				if argDefaults[i] != "" {
					defaultValueValid := false
					for _, v := range allowedValues {
						if v == argDefaults[i] {
							defaultValueValid = true
							break
						}
					}
					if !defaultValueValid {
						return renderErrorPage(c, "Default value must be one of the allowed values for argument "+argNames[i], http.StatusBadRequest)
					}
				} else {
					return renderErrorPage(c, "Default value is required for argument "+argNames[i], http.StatusBadRequest)
				}
			}
			if argDefaults[i] == "" {
				doesJobHaveArgumentWithoutDefaultValue = true
			}
			argType := "string" // Default type
			if i < len(argTypes) && argTypes[i] != "" {
				argType = argTypes[i]
			}
			arguments = append(arguments, schema.JobArgument{
				Name:          argNames[i],
				Type:          argType,
				DefaultValue:  argDefaults[i],
				AllowedValues: allowedValues,
				Sensitive:     sensitiveMap[i],
			})
		}
		cronNextRunTime := time.Now()
		if scheduleEnabled {
			if doesJobHaveArgumentWithoutDefaultValue {
				return renderErrorPage(c, "All arguments must have a default value when scheduling is enabled", http.StatusBadRequest)
			}
			if form.CronSchedule != "" {
				n, cronParseError := cron.ParseStandard(form.CronSchedule)
				if cronParseError != nil {
					slog.Error("error parsing cron schedule", "error", cronParseError)
					return renderErrorPage(c, "Error parsing cron schedule "+cronParseError.Error(), http.StatusBadRequest)
				}
				cronNextRunTime = n.Next(time.Now())
			} else {
				return renderErrorPage(c, "Cron schedule is required when schedule is enabled", http.StatusBadRequest)
			}
			needToScheduleCron = true
		}
		var startChannels []int
		var successChannels []int
		var failureChannels []int

		// Note the [] after the field name - this is important for multiple values
		startChannelIDs := c.Request().Form["notify_on_start_channel_ids[]"]
		successChannelIDs := c.Request().Form["notify_on_success_channel_ids[]"]
		failureChannelIDs := c.Request().Form["notify_on_failure_channel_ids[]"]

		// Process start channels
		for _, id := range startChannelIDs {
			if channelID, err := strconv.Atoi(id); err == nil {
				startChannels = append(startChannels, channelID)
			}
		}

		// Process success channels
		for _, id := range successChannelIDs {
			if channelID, err := strconv.Atoi(id); err == nil {
				successChannels = append(successChannels, channelID)
			}
		}

		// Process failure channels
		for _, id := range failureChannelIDs {
			if channelID, err := strconv.Atoi(id); err == nil {
				failureChannels = append(failureChannels, channelID)
			}
		}
		_, err = db.Job.UpdateOneID(jobIDInt).
			SetName(form.Name).
			SetDescription(form.Description).
			SetTimeoutSeconds(form.TimeoutSeconds).
			SetCronSchedule(form.CronSchedule).
			SetScheduleEnabled(scheduleEnabled).
			SetAllowConcurrentRuns(allowConcurrentRuns).
			SetArguments(arguments).
			SetRequiresFileUpload(requiresFileUpload).
			SetNotifyOnStartChannelIds(startChannels).
			SetNotifyOnSuccessChannelIds(successChannels).
			SetNotifyOnFailureChannelIds(failureChannels).
			SetScript(form.Script).
			SetLastEditTime(time.Now()).
			SetNextCronRunTime(cronNextRunTime).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error updating job", "error", err)
			return renderErrorPage(c, "Error updating job", http.StatusInternalServerError)
		}
		if needToScheduleCron {
			addCronJob(db, j)
		}
		// Create the version
		_, err = db.JobVersion.Create().
			SetJob(j).
			SetName(form.Name).
			SetDescription(form.Description).
			SetCronSchedule(form.CronSchedule).
			SetScheduleEnabled(scheduleEnabled).
			SetAllowConcurrentRuns(allowConcurrentRuns).
			SetArguments(arguments).
			SetRequiresFileUpload(requiresFileUpload).
			SetScript(form.Script).
			SetChangedByEmail(self.Email).
			SetCreatedAt(time.Now()).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("error creating job version", "error", err)
			return renderErrorPage(c, "Error creating job version", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%d", j.Edges.Project.ID))
	}
}

func hookDeleteJob(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			return renderErrorPage(c, "Error converting job ID to integer", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.IDEQ(jobIDInt)).WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to delete this job", http.StatusForbidden)
		}
		// Remove it from cron if it has a schedule enabled
		if j.ScheduleEnabled {
			removeCronJob(j)
		}
		err = db.Job.DeleteOneID(jobIDInt).Exec(c.Request().Context())
		if err != nil {
			slog.Error("error deleting job", "error", err)
			return renderErrorPage(c, "Error deleting job", http.StatusInternalServerError)
		}
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%d", j.Edges.Project.ID))
	}
}

func renderRunJobView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			return renderErrorPage(c, "Error converting job ID to integer", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.IDEQ(jobIDInt)).WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to run this job", http.StatusForbidden)
		}
		return c.Render(http.StatusOK, "run-job", map[string]any{
			"Job":     j,
			"Project": j.Edges.Project,
		})
	}
}

func formRunJob(db *ent.Client, config Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			return renderErrorPage(c, "Error converting job ID to integer", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.IDEQ(jobIDInt)).WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		if !canUserEditProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to run this job", http.StatusForbidden)
		}
		var argValues []JobRuntimeArg
		for _, arg := range j.Arguments {
			argValue := c.FormValue("arg_" + arg.Name)
			argValues = append(argValues, JobRuntimeArg{
				Name:      arg.Name,
				Value:     argValue,
				Sensitive: arg.Sensitive,
			})
		}
		fileBytes := []byte{}
		if j.RequiresFileUpload {
			file, err := c.FormFile("file")
			if err != nil {
				return renderErrorPage(c, "Error getting file from form", http.StatusBadRequest)
			}
			src, err := file.Open()
			if err != nil {
				return renderErrorPage(c, "Error opening file", http.StatusBadRequest)
			}
			defer src.Close()
			// Read all into []byte
			b, err := io.ReadAll(src)
			if err != nil {
				return renderErrorPage(c, "Error reading file", http.StatusBadRequest)
			}
			fileBytes = b
		}
		// Create the history
		history, err := db.JobHistory.Create().
			SetJob(j).
			SetProject(j.Edges.Project).
			SetTrigger("Web UI").
			SetStatus("running").
			SetCreatedAt(time.Now()).
			SetWasSuccessful(false).
			SetDurationMs(0).
			SetExitCode(0).
			SetTriggeredByEmail(self.Email).
			SetTriggeredByID(self.ID).
			Save(context.Background())
		if err != nil {
			slog.Error("error creating job history", "error", err)
			return renderErrorPage(c, "Error creating job history", http.StatusInternalServerError)
		}
		go runJob(db, j, history, argValues, fileBytes, config)
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/projects/%d/jobs/%d/run/%d", j.Edges.Project.ID, j.ID, history.ID))
	}
}

func renderJobHistorySingleItemView(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		historyID := c.Param("history_id")
		historyIDInt, err := strconv.Atoi(historyID)
		if err != nil {
			return renderErrorPage(c, "Error converting history ID to integer", http.StatusBadRequest)

		}
		h, err := db.JobHistory.Query().Where(jobhistory.IDEQ(historyIDInt)).WithJob().WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job history from database", http.StatusInternalServerError)
		}
		if !canUserViewProject(db, self.ID, h.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to view this job history", http.StatusForbidden)
		}
		realTimeOutput := false
		if h.Status == "running" {
			realTimeOutput = true
		}
		// Check if output is in the runningJobOutputs map
		if runningJobOutputs == nil {
			runningJobOutputs = make(map[int][]string)
		}
		_, ok := runningJobOutputs[h.ID]
		output := ""
		if ok {
			output = strings.Join(runningJobOutputs[h.ID], "\n")
		}
		return c.Render(http.StatusOK, "job-history-single", map[string]any{
			"History":        h,
			"Output":         output,
			"RealTimeOutput": realTimeOutput,
			"Project":        h.Edges.Project,
		})
	}
}

func htmxJobHistoryOutput() echo.HandlerFunc {
	return func(c echo.Context) error {
		historyID := c.Param("history_id")
		historyIDInt, err := strconv.Atoi(historyID)
		if err != nil {
			return renderErrorPage(c, "Error converting history ID to integer", http.StatusBadRequest)
		}
		// Check if output is in the runningJobOutputs map
		if runningJobOutputs == nil {
			runningJobOutputs = make(map[int][]string)
		}
		_, ok := runningJobOutputs[historyIDInt]
		output := ""
		if ok {
			output = strings.Join(runningJobOutputs[historyIDInt], "\n")
		} else {
			c.Response().Header().Add("HX-Refresh", "true")
			return c.String(http.StatusOK, "Reload the page to see the output")
		}
		return c.Render(http.StatusOK, "output-slot", map[string]any{
			"Output":  output,
			"History": ent.JobHistory{ID: historyIDInt},
		})
	}
}

func renderJobHistoryView(db *ent.Client) echo.HandlerFunc {
	type NiceHistory struct {
		*ent.JobHistory
		FriendlyTime     string
		FriendlyDuration string
	}
	return func(c echo.Context) error {
		self, _ := getSelf(c, db)
		jobID := c.Param("job_id")
		jobIDInt, err := strconv.Atoi(jobID)
		if err != nil {
			return renderErrorPage(c, "Error converting job ID to integer", http.StatusBadRequest)
		}
		j, err := db.Job.Query().Where(job.IDEQ(jobIDInt)).WithProject().Only(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job from database", http.StatusInternalServerError)
		}
		if !canUserViewProject(db, self.ID, j.Edges.Project.ID) {
			return renderErrorPage(c, "You do not have permission to view this job history", http.StatusForbidden)
		}
		histories, err := db.JobHistory.Query().Where(jobhistory.HasJobWith(job.IDEQ(jobIDInt))).WithProject().WithJob().Order(ent.Desc(jobhistory.FieldCreatedAt)).All(c.Request().Context())
		if err != nil {
			return renderErrorPage(c, "Error getting job histories from database", http.StatusInternalServerError)
		}
		niceHistories := make([]NiceHistory, 0)
		for _, h := range histories {
			niceHistories = append(niceHistories, NiceHistory{
				JobHistory:   h,
				FriendlyTime: humanize.Time(h.CreatedAt),
			})
		}
		return c.Render(http.StatusOK, "job-histories", map[string]any{
			"Job":     j,
			"History": niceHistories,
			"Project": j.Edges.Project,
		})
	}
}
