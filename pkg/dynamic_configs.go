package pkg

import (
	"context"
	"fmt"
	"github.com/lbrictson/janus/ent"
)

// Used for caching the auth config since it is gotten quite frequently on an authenticated page which opens the site up
// to a DoS attack if the DB is hit every time
var lastAuthConfig *ent.AuthConfig
var lastDataConfig *ent.DataConfig
var lastJobConfig *ent.JobConfig
var lastSMTPConfig *ent.SMTPConfig

func getAuthConfig(ctx context.Context, db *ent.Client) (*ent.AuthConfig, error) {
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
	n, err := db.AuthConfig.UpdateOne(newConfig).
		SetEnableSSO(newConfig.EnableSSO).
		SetDisablePasswordLogin(newConfig.DisablePasswordLogin).
		SetSSOProvider(newConfig.SSOProvider).
		SetSSOClientID(newConfig.SSOClientID).
		SetSSOClientSecret(newConfig.SSOClientSecret).
		SetSSORedirectURI(newConfig.SSORedirectURI).
		SetSSOAuthorizationURL(newConfig.SSOAuthorizationURL).
		SetSSOTokenURL(newConfig.SSOTokenURL).
		SetSSOUserInfoURL(newConfig.SSOUserInfoURL).
		SetEntraTenantID(newConfig.EntraTenantID).
		SetGoogleAllowedDomains(newConfig.GoogleAllowedDomains).
		SetSessionKey(newConfig.SessionKey).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update auth config: %v", err)
	}
	lastAuthConfig = n
	return nil
}

func getDataConfig(ctx context.Context, db *ent.Client) (*ent.DataConfig, error) {
	if lastDataConfig != nil {
		return lastDataConfig, nil
	}
	dConfig, err := db.DataConfig.Query().Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get data config: %v", err)
	}
	lastDataConfig = dConfig
	return dConfig, nil
}

func updateDataConfig(ctx context.Context, db *ent.Client, newConfig *ent.DataConfig) error {
	fmt.Println(newConfig.DaysToKeep)
	updateConfig, err := db.DataConfig.UpdateOne(newConfig).
		SetDaysToKeep(newConfig.DaysToKeep).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update data config: %v", err)
	}
	lastDataConfig = updateConfig
	return nil
}

func getJobConfig(ctx context.Context, db *ent.Client) (*ent.JobConfig, error) {
	if lastJobConfig != nil {
		return lastJobConfig, nil
	}
	jConfig, err := db.JobConfig.Query().Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get job config: %v", err)
	}
	lastJobConfig = jConfig
	return jConfig, nil
}

func updateJobConfig(ctx context.Context, db *ent.Client, newConfig *ent.JobConfig) error {
	_, err := db.JobConfig.UpdateOne(newConfig).
		SetMaxConcurrentJobs(newConfig.MaxConcurrentJobs).
		SetDefaultTimeoutSeconds(newConfig.DefaultTimeoutSeconds).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update job config: %v", err)
	}
	lastJobConfig = newConfig
	return nil
}

func getSMTPConfig(ctx context.Context, db *ent.Client) (*ent.SMTPConfig, error) {
	if lastSMTPConfig != nil {
		return lastSMTPConfig, nil
	}
	sConfig, err := db.SMTPConfig.Query().Only(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get smtp config: %v", err)
	}
	lastSMTPConfig = sConfig
	return sConfig, nil
}

func updateSMTPConfig(ctx context.Context, db *ent.Client, newConfig *ent.SMTPConfig) error {
	_, err := db.SMTPConfig.UpdateOne(newConfig).
		SetSMTPPassword(newConfig.SMTPPassword).
		SetSMTPPort(newConfig.SMTPPort).
		SetSMTPServer(newConfig.SMTPServer).
		SetSMTPUsername(newConfig.SMTPUsername).
		SetSMTPSender(newConfig.SMTPSender).
		SetSMTPTLS(newConfig.SMTPTLS).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update smtp config: %v", err)
	}
	lastSMTPConfig = newConfig
	return nil
}
