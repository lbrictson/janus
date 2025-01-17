package pkg

import (
	"context"
	"fmt"
	"github.com/lbrictson/janus/ent"
)

// Used for caching the auth config since it is gotten quite frequently on an authenticated page which opens the site up
// to a DoS attack if the DB is hit every time
var lastAuthConfig *ent.AuthConfig

func getAuthconfig(ctx context.Context, db *ent.Client) (*ent.AuthConfig, error) {
	if lastAuthConfig != nil {
		return lastAuthConfig, nil
	}
	aConfig, err := db.AuthConfig.Query().Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get auth config: %v", err)
	}
	lastAuthConfig = aConfig
	return aConfig, nil
}

func updateAuthConfig(ctx context.Context, db *ent.Client, newConfig *ent.AuthConfig) error {
	_, err := db.AuthConfig.UpdateOne(newConfig).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update auth config: %v", err)
	}
	lastAuthConfig = newConfig
	return nil
}
