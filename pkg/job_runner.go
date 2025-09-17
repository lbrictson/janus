package pkg

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/jobhistory"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/schema"
	"github.com/lbrictson/janus/ent/secret"
	"github.com/lbrictson/janus/pkg/notification_sender"
)

var runningJobOutputs = make(map[int][]string)

type JobRuntimeArg struct {
	Name      string
	Value     string
	Sensitive bool
}

func runJob(db *ent.Client, job *ent.Job, history *ent.JobHistory, argValues []JobRuntimeArg, fileBytes []byte, config Config) (int, error) {
	// Before we run any job we need to check for max concurrency
	runningJobs, _ := db.JobHistory.Query().WithJob().Where(jobhistory.StatusEQ("running")).All(context.Background())
	runningJobsCount := len(runningJobs)
	jConfig, _ := getJobConfig(context.Background(), db)
	if runningJobsCount > jConfig.MaxConcurrentJobs {
		totalJobFailures.Inc()
		// Save the history
		_, err := db.JobHistory.UpdateOne(history).
			SetStatus("failed").
			SetOutput("Max concurrent jobs reached.  Aborting.").
			SetWasSuccessful(false).
			SetDurationMs(1).
			SetExitCode(1).
			Save(context.Background())
		if err != nil {
			slog.Error("failed to update job history", "error", err)
		}
		failChannels := []*ent.NotificationChannel{}
		for _, i := range job.NotifyOnFailureChannelIds {
			nc, err := db.NotificationChannel.Get(context.Background(), i)
			if err != nil {
				totalJobFailures.Inc()
				return 0, err
			}
			failChannels = append(failChannels, nc)
		}
		for _, n := range failChannels {
			err = sendNotification(db, n, notification_sender.NewNotificationInput{
				JobName:     job.Name,
				ProjectName: job.Edges.Project.Name,
				JobStatus:   notification_sender.FAILURE,
				JobDuration: "0ms",
				CallbackURL: fmt.Sprintf("%s/projects/%d/jobs/%d/run/%d", config.ServerURL, job.Edges.Project.ID, job.ID, history.ID),
			})
			if err != nil {
				notificationsFailure.Inc()
			} else {
				notificationsSent.Inc()
			}
		}
		return 0, fmt.Errorf("max concurrent jobs reached (%v/%v).  Aborting.", runningJobsCount, jConfig.MaxConcurrentJobs)
	}

	// Before we run any job we need to check if it allows concurrent runs, if it doesn't we abort if it is already running
	if !job.AllowConcurrentRuns {
		// This job doesn't allow concurrent runs, we need to check if it is already running
		runningJobs, err := db.JobHistory.Query().WithJob().Where(jobhistory.StatusEQ("running")).All(context.Background())
		if err != nil {
			totalJobFailures.Inc()
			return 0, err
		}
		for _, j := range runningJobs {
			if j.Edges.Job.ID == job.ID {
				if j.ID == history.ID {
					continue
				}
				totalJobFailures.Inc()
				// Save the history
				_, err = db.JobHistory.UpdateOne(history).
					SetStatus("failed").
					SetOutput("Job is already running and does not allow concurrent runs.  Aborting.").
					SetWasSuccessful(false).
					SetDurationMs(1).
					SetExitCode(1).
					Save(context.Background())
				if err != nil {
					slog.Error("failed to update job history", "error", err)
				}
				failChannels := []*ent.NotificationChannel{}
				for _, i := range job.NotifyOnFailureChannelIds {
					nc, err := db.NotificationChannel.Get(context.Background(), i)
					if err != nil {
						totalJobFailures.Inc()
						return 0, err
					}
					failChannels = append(failChannels, nc)
				}
				for _, n := range failChannels {
					err = sendNotification(db, n, notification_sender.NewNotificationInput{
						JobName:     job.Name,
						ProjectName: job.Edges.Project.Name,
						JobStatus:   notification_sender.FAILURE,
						JobDuration: "0ms",
						CallbackURL: fmt.Sprintf("%s/projects/%d/jobs/%d/run/%d", config.ServerURL, job.Edges.Project.ID, job.ID, history.ID),
					})
					if err != nil {
						notificationsFailure.Inc()
					} else {
						notificationsSent.Inc()
					}
				}
				return 0, fmt.Errorf("job is already running and does not allow concurrent runs.  Aborting.")
			}
		}
	}
	totalJobRuns.Inc()
	if runningJobOutputs == nil {
		runningJobOutputs = make(map[int][]string)
	}
	// We need to gather any notification channels this job needs
	startChannels := []*ent.NotificationChannel{}
	failChannels := []*ent.NotificationChannel{}
	successChannels := []*ent.NotificationChannel{}
	for _, i := range job.NotifyOnStartChannelIds {
		nc, err := db.NotificationChannel.Get(context.Background(), i)
		if err != nil {
			totalJobFailures.Inc()
			return 0, err
		}
		startChannels = append(startChannels, nc)
	}
	for _, i := range job.NotifyOnFailureChannelIds {
		nc, err := db.NotificationChannel.Get(context.Background(), i)
		if err != nil {
			totalJobFailures.Inc()
			return 0, err
		}
		failChannels = append(failChannels, nc)
	}
	for _, i := range job.NotifyOnSuccessChannelIds {
		nc, err := db.NotificationChannel.Get(context.Background(), i)
		if err != nil {
			totalJobFailures.Inc()
			return 0, err
		}
		successChannels = append(successChannels, nc)
	}
	// Get all the secrets from the project the job is in
	availableSecrets, err := db.Secret.Query().Where(secret.HasProjectWith(project.IDEQ(job.Edges.Project.ID))).All(context.Background())
	if err != nil {
		totalJobFailures.Inc()
		return 0, err
	}
	params := []schema.Parameter{}
	for _, arg := range argValues {
		params = append(params, schema.Parameter{
			Name:      arg.Name,
			Value:     arg.Value,
			Sensitive: arg.Sensitive,
		})
	}
	db.JobHistory.Update().Where(jobhistory.IDEQ(history.ID)).
		SetParameters(params).
		Save(context.Background())
	historyID := history.ID
	runningJobOutputs[historyID] = []string{}
	script := injectSecretsIntoScript(job.Script, availableSecrets)
	args := make(map[string]string)
	for _, arg := range argValues {
		args[arg.Name] = arg.Value
	}
	script = injectArgsIntoScript(script, args)
	script = strings.Replace(script, "\r\n", "\n", -1)
	// Replace windows line endings with unix line endings
	script = strings.Replace(script, "\r", "\n", -1)
	// Execute the script
	ctx := context.Background()
	scriptCtx, cancelFunc := context.WithDeadline(ctx, time.Now().Add(time.Duration(job.TimeoutSeconds)*time.Second))
	defer cancelFunc()
	envVars := []string{}
	// Add PATH to env vars
	envVars = append(envVars, "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin")
	hiddenValues := []string{}
	for _, s := range availableSecrets {
		hiddenValues = append(hiddenValues, s.Value)
		envVars = append(envVars, fmt.Sprintf("%s=%s", formatSecretNameForJob(s.Name), s.Value))
	}
	for _, s := range argValues {
		if s.Sensitive {
			hiddenValues = append(hiddenValues, s.Value)
		}
		envVars = append(envVars, fmt.Sprintf("%s=%s", formatArgNameForJob(s.Name), s.Value))
	}
	for _, n := range startChannels {
		sendNotification(db, n, notification_sender.NewNotificationInput{
			JobName:     job.Name,
			ProjectName: job.Edges.Project.Name,
			JobStatus:   notification_sender.STARTING,
			JobDuration: "0ms",
			CallbackURL: fmt.Sprintf("%s/projects/%d/jobs/%d/run/%d", config.ServerURL, job.Edges.Project.ID, job.ID, historyID),
		})
	}
	start := time.Now()
	statusCode, err := executeScript(scriptCtx, script, historyID, envVars, hiddenValues, fileBytes)
	duration := time.Since(start)
	if err != nil {
		o := strings.Join(runningJobOutputs[historyID], "\n")
		// Set the history as failed
		if err == context.DeadlineExceeded {
			err = fmt.Errorf("timeout after %d seconds", job.TimeoutSeconds)
		}
		if err.Error() == "command failed: signal: killed" {
			err = fmt.Errorf("timeout after %d seconds", job.TimeoutSeconds)
		}
		o = fmt.Sprintf("%s\n\nError: %s", o, err.Error())
		db.JobHistory.Update().Where(jobhistory.IDEQ(historyID)).
			SetStatus("failed").
			SetDurationMs(duration.Milliseconds()).
			SetOutput(o).
			SetExitCode(statusCode).
			SetWasSuccessful(false).
			Save(context.Background())
		db.Job.UpdateOne(job).SetLastRunSuccess(false).
			SetLastRunTime(time.Now()).
			Save(context.Background())
		totalJobFailures.Inc()
		// Send notification for failure
		for _, n := range failChannels {
			sendNotification(db, n, notification_sender.NewNotificationInput{
				JobName:     job.Name,
				ProjectName: job.Edges.Project.Name,
				JobStatus:   notification_sender.FAILURE,
				JobDuration: duration.String(),
				CallbackURL: fmt.Sprintf("%s/projects/%d/jobs/%d/run/%d", config.ServerURL, job.Edges.Project.ID, job.ID, historyID),
			})
		}
	} else {
		o := strings.Join(runningJobOutputs[historyID], "\n")
		// Set the history as successful
		db.JobHistory.Update().Where(jobhistory.IDEQ(historyID)).
			SetStatus("success").
			SetDurationMs(duration.Milliseconds()).
			SetOutput(o).
			SetExitCode(statusCode).
			SetWasSuccessful(true).
			Save(context.Background())
		db.Job.UpdateOne(job).
			SetLastRunTime(time.Now()).
			SetLastRunSuccess(true).Save(context.Background())
		totalJobSuccesses.Inc()
		// Send notification for success
		for _, n := range successChannels {
			sendNotification(db, n, notification_sender.NewNotificationInput{
				JobName:     job.Name,
				ProjectName: job.Edges.Project.Name,
				JobStatus:   notification_sender.SUCCESS,
				JobDuration: duration.String(),
				CallbackURL: fmt.Sprintf("%s/projects/%d/jobs/%d/run/%d", config.ServerURL, job.Edges.Project.ID, job.ID, historyID),
			})
		}
	}
	// Remove job from output
	delete(runningJobOutputs, historyID)
	return statusCode, err
}

func injectSecretsIntoScript(script string, secrets []*ent.Secret) string {
	for _, s := range secrets {
		script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", formatSecretNameForJob(s.Name)), s.Value)
	}
	return script
}

func injectArgsIntoScript(script string, args map[string]string) string {
	for k, v := range args {
		script = strings.ReplaceAll(script, fmt.Sprintf("{{%s}}", formatArgNameForJob(k)), v)
	}
	return script
}

func formatSecretNameForJob(name string) string {
	return strings.ToUpper(fmt.Sprintf("JANUS_SECRET_%s", name))
}

func formatArgNameForJob(name string) string {
	return strings.ToUpper(fmt.Sprintf("JANUS_ARG_%s", name))
}

func executeScript(ctx context.Context, script string, runID int, environmentVariables []string, hiddenValues []string, fileBytes []byte) (int, error) {
	// Create directory to execute script in
	dir := fmt.Sprintf("tmp/janus/%d", runID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return 1, fmt.Errorf("failed to create directory: %w", err)
	}
	if fileBytes != nil {
		filePath := fmt.Sprintf("%s/file", dir)
		if err := os.WriteFile(filePath, fileBytes, 0644); err != nil {
			return 1, fmt.Errorf("failed to write script file: %w", err)
		}
	}
	// save the script file
	scriptPath := fmt.Sprintf("%s/script.sh", dir)
	if err := os.WriteFile(scriptPath, []byte(script), 0755); err != nil {
		return 1, fmt.Errorf("failed to write script file: %w", err)
	}
	// Create command with interpreter
	cmd := exec.CommandContext(ctx, scriptPath)
	// Set env variables
	cmd.Env = append(cmd.Env, environmentVariables...)
	// Create pipes for stdout and stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		slog.Error("Failed to create stdout pipe", "error", err)
		return 1, fmt.Errorf("failed to create stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		slog.Error("Failed to create stderr pipe", "error", err)
		return 1, fmt.Errorf("failed to create stderr pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		slog.Error("Failed to start command", "error", err)
		return 1, fmt.Errorf("failed to start command: %w", err)
	}

	// Create output channels
	stdoutChan := make(chan ScriptOutput, 100)
	stderrChan := make(chan ScriptOutput, 100)
	errorChan := make(chan error, 2)
	doneChan := make(chan struct{})
	readCompleteChan := make(chan struct{}, 2) // Signal when readers are done

	// Use WaitGroup to track all goroutines
	var wg sync.WaitGroup
	wg.Add(3) // 2 for stream readers, 1 for saveOutputToDB

	// Start goroutines to handle output streams
	go func() {
		defer wg.Done()
		streamOutput(stdout, runID, "stdout", stdoutChan, errorChan)
		close(stdoutChan)
		readCompleteChan <- struct{}{} // Signal reading complete
	}()
	go func() {
		defer wg.Done()
		streamOutput(stderr, runID, "stderr", stderrChan, errorChan)
		close(stderrChan)
		readCompleteChan <- struct{}{} // Signal reading complete
	}()
	go func() {
		defer wg.Done()
		saveOutputToDB(ctx, stdoutChan, stderrChan, hiddenValues, doneChan)
	}()

	// Wait for both readers to complete or timeout
	readersDone := make(chan struct{})
	go func() {
		<-readCompleteChan // Wait for stdout reader
		<-readCompleteChan // Wait for stderr reader
		close(readersDone)
	}()

	// Wait for readers to complete or context deadline
	select {
	case <-readersDone:
		// Readers completed successfully
	case <-ctx.Done():
		// Context deadline exceeded
		slog.Warn("Context deadline exceeded while waiting for readers")
	}

	// Now it's safe to call Wait
	cmdErr := cmd.Wait()

	// Signal saveOutputToDB to stop after command completes
	close(doneChan)

	// Wait for all goroutines to finish
	wg.Wait()

	// Close error channel after all producers are done
	close(errorChan)

	// Collect any streaming errors (channel is now closed, so this will not block)
	var streamingErrors []error
	for err := range errorChan {
		if err != nil {
			// Filter out "file already closed" errors as they're expected
			if !strings.Contains(err.Error(), "file already closed") {
				streamingErrors = append(streamingErrors, err)
			}
		}
	}

	// Clean up directory
	os.RemoveAll(dir)

	// Check if command failed
	if cmdErr != nil {
		slog.Error("Command failed", "error", cmdErr)
		return 1, fmt.Errorf("command failed: %w", cmdErr)
	}

	// Check for streaming errors
	if len(streamingErrors) > 0 {
		slog.Error("Streaming errors occurred", "errors", streamingErrors)
		return 1, fmt.Errorf("streaming error: %w", streamingErrors[0])
	}

	return 0, nil
}

type ScriptOutput struct {
	ID        int
	Line      string
	Timestamp time.Time
}

func streamOutput(
	reader io.Reader,
	runID int,
	streamType string,
	outputChan chan<- ScriptOutput,
	errorChan chan<- error,
) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		output := ScriptOutput{
			ID:        runID,
			Line:      scanner.Text(),
			Timestamp: time.Now(),
		}

		outputChan <- output
	}

	if err := scanner.Err(); err != nil {
		errorChan <- fmt.Errorf("%s scanning error: %w", streamType, err)
	}
}

func saveOutputToDB(
	ctx context.Context,
	stdoutChan, stderrChan <-chan ScriptOutput,
	hiddenValues []string,
	done <-chan struct{},
) {

	for {
		select {
		case output, ok := <-stdoutChan:
			if ok {
				if err := insertOutput(ctx, output, hiddenValues); err != nil {
					slog.Error("Failed to save stdout", "error", err)
				}
			}
		case output, ok := <-stderrChan:
			if ok {
				if err := insertOutput(ctx, output, hiddenValues); err != nil {
					slog.Error("Failed to save stderr", "error", err)
				}
			}
		case <-done:
			// Drain any remaining outputs
			for output := range stdoutChan {
				if err := insertOutput(ctx, output, hiddenValues); err != nil {
					slog.Error("Failed to save stdout", "error", err)
				}
			}
			for output := range stderrChan {
				if err := insertOutput(ctx, output, hiddenValues); err != nil {
					slog.Error("Failed to save stderr", "error", err)
				}
			}
			return
		case <-ctx.Done():
			return
		}
	}
}

func insertOutput(
	ctx context.Context,
	output ScriptOutput,
	hiddenValues []string,
) error {
	// Check if job id is in the map
	if _, ok := runningJobOutputs[output.ID]; !ok {
		runningJobOutputs[output.ID] = []string{}
	}
	cleanedString := output.Line
	for _, v := range hiddenValues {
		cleanedString = strings.ReplaceAll(cleanedString, v, "****")
	}
	runningJobOutputs[output.ID] = append(runningJobOutputs[output.ID], cleanedString)
	return nil
}
