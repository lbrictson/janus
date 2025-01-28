package pkg

import (
	"context"
	"fmt"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/user"
	"time"
)

// ExecuteSeeds runs the seed data for the application, it is required prior to a first start of the web server
func ExecuteSeeds(ctx context.Context, db *ent.Client, config *Config) error {
	// Add a user if there is none
	existingUsers, _ := db.User.Query().Count(ctx)
	if existingUsers == 0 {
		pw, err := hashAndSaltPassword("ChangeMeBeforeUse1234!")
		if err != nil {
			return fmt.Errorf("failed to hash password for seed user: %v", err)
		}
		_, err = db.User.Create().
			SetEmail("admin@localhost").
			SetEncryptedPassword(pw).
			SetAPIKey(generateLongString()).
			SetAdmin(true).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed user: %v", err)
		}
	}
	// Seed the default auth config
	existingAuthConfigs, _ := db.AuthConfig.Query().Count(ctx)
	if existingAuthConfigs == 0 {
		_, err := db.AuthConfig.Create().
			SetEnableSSO(false).
			SetSessionKey([]byte(generateLongString())).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed auth config: %v", err)
		}
	}
	// Seed the default data config
	existingDataConfigs, _ := db.DataConfig.Query().Count(ctx)
	if existingDataConfigs == 0 {
		_, err := db.DataConfig.Create().
			SetDaysToKeep(180).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed data config: %v", err)
		}
	}
	// Seed the default smtp config
	existingSMTPConfigs, _ := db.SMTPConfig.Query().Count(ctx)
	if existingSMTPConfigs == 0 {
		_, err := db.SMTPConfig.Create().
			SetSMTPServer("localhost").
			SetSMTPPort(1025).
			SetSMTPUsername("admin@localhost").
			SetSMTPPassword("Changeme!!!!!!").
			SetSMTPSender("admin@localhost").
			SetSMTPTLS(false).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed smtp config: %v", err)
		}
	}
	// Seed default job data
	existingJobConfig, _ := db.JobConfig.Query().Count(ctx)
	if existingJobConfig == 0 {
		_, err := db.JobConfig.Create().
			SetMaxConcurrentJobs(100).
			SetDefaultTimeoutSeconds(3600).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed job config: %v", err)
		}
	}
	return nil
}

func seedTestDatabase(db *ent.Client) error {
	ctx := context.Background()
	// Seed the default auth config
	existingAuthConfigs, _ := db.AuthConfig.Query().Count(ctx)
	if existingAuthConfigs == 0 {
		_, err := db.AuthConfig.Create().
			SetEnableSSO(false).
			SetSessionKey([]byte(generateLongString())).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed auth config: %v", err)
		}
	}
	// Seed the default data config
	existingDataConfigs, _ := db.DataConfig.Query().Count(ctx)
	if existingDataConfigs == 0 {
		_, err := db.DataConfig.Create().
			SetDaysToKeep(180).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed data config: %v", err)
		}
	}
	// Seed the default smtp config
	existingSMTPConfigs, _ := db.SMTPConfig.Query().Count(ctx)
	if existingSMTPConfigs == 0 {
		_, err := db.SMTPConfig.Create().
			SetSMTPServer("localhost").
			SetSMTPPort(1025).
			SetSMTPUsername("admin@localhost").
			SetSMTPPassword("Changeme!!!!!!").
			SetSMTPSender("admin@localhost").
			SetSMTPTLS(false).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed smtp config: %v", err)
		}
	}
	// Seed default job data
	existingJobConfig, _ := db.JobConfig.Query().Count(ctx)
	if existingJobConfig == 0 {
		_, err := db.JobConfig.Create().
			SetMaxConcurrentJobs(100).
			SetDefaultTimeoutSeconds(3600).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed job config: %v", err)
		}
	}
	usersToMake := []string{"admin@localhost", "user@localhost"}
	for _, email := range usersToMake {
		pw, err := hashAndSaltPassword("ChangeMeBeforeUse1234!")
		if err != nil {
			return fmt.Errorf("failed to hash password for seed user: %v", err)
		}
		_, err = db.User.Create().
			SetEmail(email).
			SetEncryptedPassword(pw).
			SetAPIKey(generateLongString()).
			SetAdmin(email == "admin@localhost").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed user: %v", err)
		}
	}
	projectsToMake := []string{"project1", "project2"}
	for _, name := range projectsToMake {
		_, err := db.Project.Create().
			SetName(name).
			SetDescription(name + " stuff").
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed project: %v", err)
		}
	}
	jobsToMake := []string{"job1", "job2"}
	for _, name := range jobsToMake {
		_, err := db.Job.Create().
			SetName(name).
			SetScheduleEnabled(false).
			SetCronSchedule("").
			SetAllowConcurrentRuns(false).
			SetScript("echo 'hello world'").
			SetCreatedAt(time.Now()).
			SetTimeoutSeconds(3600).
			SetRequiresFileUpload(false).
			SetLastRunTime(time.Now()).
			SetProject(db.Project.Query().Where(project.NameEQ("project1")).OnlyX(ctx)).
			SetLastEditTime(time.Now()).
			SetDescription("").
			SetNextCronRunTime(time.Now()).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed job: %v", err)
		}
	}
	// Number of job histories to make per job
	historiesToMake := 5
	for _, name := range jobsToMake {
		job, err := db.Job.Query().WithProject().Where(job.NameEQ(name)).Only(ctx)
		if err != nil {
			return fmt.Errorf("failed to get seed job: %v", err)
		}
		for i := 0; i < historiesToMake; i++ {
			_, err := db.JobHistory.Create().
				SetWasSuccessful(true).
				SetCreatedAt(time.Now()).
				SetDurationMs(100).
				SetJob(job).
				SetExitCode(0).
				SetOutput("").
				SetJob(job).
				SetTriggeredByID(1).
				SetTrigger("UI").
				SetTriggeredByEmail("admin@localhost").
				SetStatus("Success").
				SetProject(job.Edges.Project).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create seed job history: %v", err)
			}
		}
	}
	// Permissions
	permissionsToMake := []string{"admin@localhost", "user@localhost"}
	for _, email := range permissionsToMake {
		user, err := db.User.Query().Where(user.EmailEQ(email)).Only(ctx)
		if err != nil {
			return fmt.Errorf("failed to get seed user: %v", err)
		}
		project, err := db.Project.Query().Where(project.NameEQ("project1")).Only(ctx)
		if err != nil {
			return fmt.Errorf("failed to get seed project: %v", err)
		}
		_, err = db.ProjectUser.Create().
			SetUser(user).
			SetProject(project).
			SetCanEdit(true).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create seed project user: %v", err)
		}
	}
	// Create secrets
	sceretsToMake := map[string]string{
		"secret1": "value1",
		"secret2": "value2",
	}
	for _, p := range projectsToMake {
		project, err := db.Project.Query().Where(project.NameEQ(p)).Only(ctx)
		if err != nil {
			return fmt.Errorf("failed to get seed project: %v", err)
		}
		for key, value := range sceretsToMake {
			_, err := db.Secret.Create().
				SetName(key).
				SetValue(value).
				SetProject(project).
				SetUpdatedAt(time.Now()).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create seed secret: %v", err)
			}
		}
	}
	return nil
}
