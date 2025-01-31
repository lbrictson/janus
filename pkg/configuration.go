package pkg

import (
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Port            int    `envconfig:"JANUS_PORT" default:"8080"`
	DBType          string `envconfig:"JANUS_DB_TYPE" default:"sqlite"`
	DBPath          string `envconfig:"JANUS_DB_PATH" default:"janus.db"`
	DBHostname      string `envconfig:"JANUS_DB_HOSTNAME" default:"localhost"`
	DBPort          int    `envconfig:"JANUS_DB_PORT" default:"5432"`
	DBName          string `envconfig:"JANUS_DB_NAME" default:"postgres"`
	DBUser          string `envconfig:"JANUS_DB_USER" default:"postgres"`
	DBPassword      string `envconfig:"JANUS_DB_PASSWORD" default:"postgres"`
	DBSSLMode       string `envconfig:"JANUS_DB_SSL_MODE" default:"disable"`
	ServerURL       string `envconfig:"JANUS_URL" default:"http://localhost:8080"`
	DevelopmentMode bool   `envconfig:"JANUS_DEVELOPMENT_MODE" default:"false"`
	DisableMetrics  bool   `envconfig:"JANUS_DISABLE_METRICS" default:"false"`
	MetricsPort     int    `envconfig:"JANUS_METRICS_PORT" default:"8081"`
	SessionName     string `envconfig:"JANUS_SESSION_NAME" default:"janus"`
	BrandName       string `envconfig:"JANUS_BRAND_NAME" default:"Janus"`
}

func LoadConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("janus", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
