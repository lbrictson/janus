package main

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/pkg"
	_ "github.com/lib/pq"
	"log/slog"
	_ "modernc.org/sqlite"
	"os"
	"strings"
)

func main() {
	// Configure json logger
	j := slog.NewJSONHandler(os.Stdout, nil)
	newLogger := slog.New(j)
	slog.SetDefault(newLogger)
	slog.Info("starting janus")
	c, err := pkg.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load configuration: %v", err))
	}
	db, err := connectDatabase(c)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	slog.Info("connected to database and ran migrations successfully")
	if err := pkg.ExecuteSeeds(context.Background(), db, c); err != nil {
		panic(fmt.Sprintf("failed to seed database with initial data: %v", err))
	}
	slog.Info("start webserver")
	pkg.RunServer(c, db)
}

func connectDatabase(config *pkg.Config) (*ent.Client, error) {
	switch strings.ToLower(config.DBType) {
	case "sqlite":
		return connectSqlite(config)
	case "postgres":
		return connectPostgres(config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.DBType)
	}
}

func connectSqlite(config *pkg.Config) (*ent.Client, error) {
	params := "_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)&_pragma=synchronous(NORMAL)&_pragma=journal_size_limit(100000000)"
	db, err := sql.Open("sqlite", fmt.Sprintf("%v?%v", config.DBPath, params))
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.SQLite, db.DB())
	conn := ent.NewClient(ent.Driver(drv))
	if err := conn.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return conn, nil
}

func connectPostgres(config *pkg.Config) (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v", config.DBHostname, config.DBPort, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return client, nil
}
