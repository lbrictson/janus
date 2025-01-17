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
	return nil
}
