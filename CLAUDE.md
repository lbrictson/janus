# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Janus is a Job (script) runner web application designed for non-technical users to execute scripts authored by technical users. It's built with Go, uses the Echo web framework, and supports both SQLite and PostgreSQL databases.

## Core Architecture

### Technology Stack
- **Backend**: Go with Echo v4 web framework
- **Database**: SQLite (default) or PostgreSQL, using Ent ORM
- **Frontend**: Server-side rendered templates with HTMX for dynamic updates
- **Authentication**: Local auth and SSO support via Goth
- **Background Jobs**: Built-in cron scheduler using robfig/cron
- **Notifications**: Multiple channels (Email, Slack, Discord, AWS SNS/EventBridge, PagerDuty, Twilio, Webhooks)

### Key Components
- **Server Entry**: `cmd/server/main.go` - Initializes database, runs migrations, starts web server
- **Route Handlers**: `pkg/frontend_*_routes.go` files contain all HTTP handlers
- **Job Execution**: `pkg/job_runner.go` - Handles script execution and output capture
- **Background Tasks**: `pkg/background_jobs.go` and `pkg/crons.go` - Manages scheduled and recurring jobs
- **Database Models**: `ent/schema/` directory contains all entity definitions (14 models including Job, Project, User, JobHistory, Secret, NotificationChannel)
- **Configuration**: `pkg/configuration.go` - Environment-based configuration using envconfig
- **Authentication**: `pkg/frontend_auth_routes.go` - Handles local and SSO authentication flows
- **Metrics**: `pkg/metrics.go` - Prometheus metrics collection and exposure

## Development Commands

### Running the Server
```bash
# Run with default SQLite database
go run cmd/server/main.go

# With hot reloading (requires Air)
air

# Run with PostgreSQL (set environment variables first)
JANUS_DB_TYPE=postgres JANUS_DB_HOSTNAME=localhost JANUS_DB_PORT=5432 JANUS_DB_NAME=janus JANUS_DB_USER=postgres JANUS_DB_PASSWORD=password go run cmd/server/main.go
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/...

# Run tests with verbose output
go test -v ./...

# Run a specific test function
go test -run TestFunctionName ./pkg/...
```

### Database Operations
```bash
# Generate Ent code after schema changes
go generate ./ent

# Add new database model
go run -mod=mod entgo.io/ent/cmd/ent new <ModelName>

# Run migrations (automatic on server start)
# Migrations are handled automatically by Ent when the server starts
```

### Building
```bash
# Build binary
go build -o janus cmd/server/main.go

# Build with CGO disabled (for static binary)
CGO_ENABLED=0 go build -o janus cmd/server/main.go

# Build Docker images
docker build -t janus:latest .
docker build -f slim.Dockerfile -t janus:slim .
```

### Linting and Formatting
```bash
# Format code
go fmt ./...

# Run go vet
go vet ./...

# Update dependencies
go mod tidy
```

## Known Issues

### Dependency Compatibility
- **tablewriter v0.0.5 pinned**: The project uses a replace directive to pin `github.com/olekukonko/tablewriter` to v0.0.5 due to compatibility issues with `github.com/jaytaylor/html2text` (used by the Hermes email template library). Newer versions of tablewriter removed constants that html2text depends on.
  - **Impact**: Cannot update tablewriter beyond v0.0.5
  - **Workaround**: Using `replace` directive in go.mod
  - **Long-term solution**: Consider replacing Hermes email library with Go's built-in `html/template` or alternative email templating solution

## Environment Variables

Key configuration is done through environment variables (all prefixed with `JANUS_`):

- `JANUS_PORT` - Web server port (default: 8080)
- `JANUS_DB_TYPE` - Database type: "sqlite" or "postgres" (default: sqlite)
- `JANUS_DB_PATH` - SQLite database file path (default: janus.db)
- `JANUS_DB_HOSTNAME`, `JANUS_DB_PORT`, `JANUS_DB_NAME`, `JANUS_DB_USER`, `JANUS_DB_PASSWORD` - PostgreSQL connection details
- `JANUS_URL` - Server base URL for callbacks (default: http://localhost:8080)
- `JANUS_DEVELOPMENT_MODE` - Enable development features (default: false)
- `JANUS_METRICS_PORT` - Prometheus metrics port (default: 8081)

## Project Structure

- **Web Routes**: All routes are defined in `pkg/server.go` with handlers in `pkg/frontend_*_routes.go`
- **Job Execution**: Jobs are bash scripts executed via `os/exec`, with output captured and stored in database
- **Permissions**: Project-based permission system defined in `pkg/permissions.go`
- **Templates**: Server-side templates in `web/templates/` using Go's html/template
- **Static Assets**: Embedded in binary from `web/static/` directory
- **API Endpoints**: REST API handlers in `pkg/api_*.go` files for programmatic access
- **Notification System**: Pluggable notification channels in `pkg/notification_sender/` directory

## Default Credentials

When running locally for development:
- Username: `admin@localhost`
- Password: `ChangeMeBeforeUse1234!`

## Docker Development

A docker-compose file is provided for testing email notifications with a local mail server:
```bash
docker-compose up -d
# Mail server UI: http://localhost:8025
# SMTP: localhost:1025
```

## CI/CD

The project uses GitHub Actions for continuous integration and deployment:
- **Release Workflow** (`.github/workflows/release.yml`): Automatically builds and publishes Docker images to Docker Hub on tagged releases
- **Documentation** (`.github/workflows/docs.yml`): Builds and deploys MkDocs documentation to GitHub Pages
- **Docker Publishing Script** (`cicd/publish_docker.sh`): Automated script for building and pushing Docker images