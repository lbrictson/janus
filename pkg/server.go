package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/job"
	"github.com/lbrictson/janus/web"
	"github.com/markbates/goth/gothic"
)

var sessionName = "janus"

func RunServer(config *Config, db *ent.Client) {
	sessionName = config.SessionName
	ctx := context.Background()
	e := echo.New()
	e.HideBanner = true
	staticFS, err := fs.Sub(web.Assets, "static")
	if err != nil {
		panic(fmt.Sprintf("failed to load static assets: %v", err))
	}
	// Serve static files from the embedded filesystem
	e.StaticFS("/static", staticFS)
	registerRenderer(e, *config)
	authC, err := getAuthConfig(ctx, db)
	if err != nil {
		panic(fmt.Sprintf("failed to get auth config: %v", err))
	}
	unauthenticated := e.Group("")
	authenticatedRoutes := e.Group("")
	adminRequired := e.Group("")
	adminRequired.Use(middlewareMustBeLoggedIn, middlewareAdminRequired)
	authenticatedRoutes.Use(middlewareMustBeLoggedIn)
	e.Use(session.Middleware(sessions.NewCookieStore(authC.SessionKey)))
	// Login pages
	unauthenticated.GET("/login", loginPage(db, config))
	unauthenticated.POST("/login", loginForm(db, config))
	unauthenticated.GET("/auth/:provider/callback", completeSSOAuth(db))
	unauthenticated.GET("/auth/:provider", startSSOAuth())
	unauthenticated.GET("/logout", destroySession)
	// Dashboard pages
	authenticatedRoutes.GET("/", renderDashboard(db))
	adminRequired.GET("/projects/:id/delete", hookDeleteProject(db))
	// User pages
	adminRequired.GET("/users", renderUsersPage(db))
	adminRequired.GET("/users/:id/edit", renderEditUserPage(db))
	adminRequired.POST("/users/:id/password", formAdminEditUserSetNewPassword(db))
	adminRequired.POST("/users/:id/role", formAdminSetUserRole(db))
	adminRequired.GET("/users/new", renderCreateUserPage)
	adminRequired.POST("/users/new", formCreateNewUser(db))
	adminRequired.POST("/users/:id/delete", formDeleteUser(db))
	adminRequired.POST("/users/:id/permissions", formUpdateUserPermissions(db))
	// Notification pages
	adminRequired.GET("/notifications", renderNotificationPage(db))
	adminRequired.GET("/notifications/new", renderNewNotificationPage())
	adminRequired.POST("/notifications/new", formCreateNotificationChannel(db))
	adminRequired.GET("/notifications/:id/edit", renderNotificationChannelEditPage(db))
	adminRequired.POST("/hook/notifications/:id/status", hookNotificationToggleStatus(db))
	adminRequired.GET("/notifications/:id/delete", deleteNotificationChannel(db))
	adminRequired.POST("/notifications/:id/edit", formEditNotificationChannel(db))
	adminRequired.POST("/hook/notifications/:id/test", hookSendTestNotification(db, *config))
	// Admin pages
	adminRequired.GET("/admin", renderAdminPage(db, config))
	adminRequired.POST("/admin/data-retention", formAdminDataRetention(db))
	adminRequired.POST("/admin/smtp", formAdminSMTP(db))
	adminRequired.POST("/admin/job-settings", formJobSettings(db))
	adminRequired.POST("/admin/security", formAdminSecuritySettings(db, config))
	// Profile pages
	authenticatedRoutes.GET("/profile/password", renderChangePasswordPage())
	authenticatedRoutes.GET("/profile/api-key", renderAPIKeyViewPage(db))
	authenticatedRoutes.POST("/profile/api-key/regenerate", formRegenerateAPIKey(db))
	authenticatedRoutes.POST("/profile/password", formSelfSetNewPassword(db))
	// Project Pages
	adminRequired.GET("/project/new", renderNewProjectView)
	adminRequired.POST("/project/new", formNewProject(db))
	authenticatedRoutes.GET("/projects/:id", renderProjectViewPage(db))
	authenticatedRoutes.GET("/projects/:project_id/jobs/new", renderCreateJobView(db))
	authenticatedRoutes.POST("/projects/:project_id/jobs/new", formCreateJob(db))
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/edit", renderEditJobPage(db))
	authenticatedRoutes.POST("/projects/:project_id/jobs/:job_id/edit", formEditJob(db))
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/delete", hookDeleteJob(db))
	adminRequired.GET("/projects/:id/secrets", renderProjectSecretsView(db))
	adminRequired.GET("/projects/:id/secrets/new", renderProjectAddSecretsView(db))
	adminRequired.POST("/projects/:id/secrets/new", formProjectAddSecret(db))
	adminRequired.GET("/projects/:project_id/secrets/:secret_id/delete", hookDeleteSecret(db))
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/run", renderRunJobView(db))
	authenticatedRoutes.POST("/projects/:project_id/jobs/:job_id/run", formRunJob(db, *config))
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/run/:history_id", renderJobHistorySingleItemView(db))
	authenticatedRoutes.GET("/htmx/job/history/:history_id/output", htmxJobHistoryOutput())
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/history", renderJobHistoryView(db))
	adminRequired.GET("/projects/:project_id/edit", renderEditProjectPage(db))
	adminRequired.POST("/projects/:project_id/edit", formEditProject(db))
	// Schedule pages
	authenticatedRoutes.GET("/schedule", renderSchedulePage(db))
	// Webhook pages
	adminRequired.GET("/webhooks", renderWebooksView(db, config))
	adminRequired.GET("/webhooks/new", renderNewWebhookView(db))
	adminRequired.POST("/webhooks/new", formNewWebhook(db))
	adminRequired.GET("/webhooks/:id/delete", hookDeleteWebhook(db))
	unauthenticated.GET("/inbound-webhook/:key", handleIncomingJobWebhookTrigger(db, config))
	unauthenticated.POST("/inbound-webhook/:key", handleIncomingJobWebhookTrigger(db, config))
	// Version pages
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/versions", renderJobVersionsPage(db))
	authenticatedRoutes.GET("/projects/:project_id/jobs/:job_id/versions/:version_id", renderSingleVersionView(db))

	// /projects/1/jobs/1/versions/6/restore
	authenticatedRoutes.POST("/projects/:project_id/jobs/:job_id/versions/:version_id/restore", formRestoreVersionOfJob(db))
	// Load API Keys
	err = reloadAPIKeys(db)
	if err != nil {
		slog.Error("failed to load API keys", "error", err)
	}
	// API Router
	apiV1 := e.Group("/api/v1")
	apiV1.Use(middlewareAPIAuthRequired)
	// Project API
	apiV1.GET("/project", apiGetProjects(db))
	apiV1.GET("/project/:id", apiGetProject(db))
	apiV1.POST("/project", apiCreateProject(db))
	apiV1.DELETE("/project/:id", apiDeleteProject(db))
	apiV1.PUT("/project/:id", apiUpdateProject(db))
	// Configure SSO
	key := authC.SessionKey
	maxAge := 8640 * 3 // 3 days
	os.MkdirAll("tmp/sessions", 0755)
	store := sessions.NewFilesystemStore("tmp/sessions", key)
	store.MaxLength(8192) // 8Kb is now the maximum size of the session
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = true
	gothic.Store = store
	if authC.EnableSSO {
		err = wireSSOConnection(db, config)
		if err != nil {
			slog.Error("error wiring SSO connection", "error", err)
		}
	}
	// Load all cron jobs
	slog.Info("loading cron jobs")
	cronJobs, err := db.Job.Query().Where(job.CronScheduleNEQ("")).All(ctx)
	if err != nil {
		slog.Error("failed to get cron jobs", "error", err)
	}
	for _, j := range cronJobs {
		if j.ScheduleEnabled {
			addCronJob(db, j)
		}
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

func registerRenderer(e *echo.Echo, config Config) {
	renderer := Renderer{
		config: config,
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"json": func(v interface{}) template.JS {
				b, err := json.Marshal(v)
				if err != nil {
					return template.JS("[]")
				}
				return template.JS(b)
			},
		}).ParseFS(web.Assets, "templates/*.tmpl")),
	}
	e.Renderer = &renderer
}

type Renderer struct {
	templates *template.Template
	config    Config
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if data == nil {
		data = make(map[string]any)
	}
	templateData := data.(map[string]any)
	templateData["userID"] = c.Get("userID")
	templateData["email"] = c.Get("email")
	templateData["BrandName"] = r.config.BrandName
	role := c.Get("globalRole")
	isAdmin := false
	if role != nil {
		if role.(string) == "admin" {
			isAdmin = true
		}
	}
	templateData["isAdmin"] = isAdmin
	err := r.templates.ExecuteTemplate(w, name, templateData)
	if err != nil {
		slog.Error("error rendering template", "error", err)
	}
	return err
}

func renderErrorPage(c echo.Context, errorMessage string, errorCode int) error {
	return c.Render(errorCode, "error", map[string]any{"ErrorMessage": errorMessage})
}
