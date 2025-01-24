package pkg

import (
	"context"
	"fmt"
	"github.com/lbrictson/janus/ent"
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
