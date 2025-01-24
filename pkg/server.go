package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/web"
	"html/template"
	"io"
	"io/fs"
	"log/slog"
)

func RunServer(config *Config, db *ent.Client) {
	ctx := context.Background()
	e := echo.New()
	e.HideBanner = true
	staticFS, err := fs.Sub(web.Assets, "static")
	if err != nil {
		panic(fmt.Sprintf("failed to load static assets: %v", err))
	}
	// Serve static files from the embedded filesystem
	e.StaticFS("/static", staticFS)
	renderer := Renderer{
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
	adminRequired.GET("/notifications/new", renderNewNotificationPage(db))
	adminRequired.POST("/notifications/new", formCreateNotificationChannel(db))
	adminRequired.GET("/notifications/:id/edit", renderNotificationChannelEditPage(db))
	adminRequired.POST("/hook/notifications/:id/status", hookNotificationToggleStatus(db))
	adminRequired.GET("/notifications/:id/delete", deleteNotificationChannel(db))
	adminRequired.POST("/notifications/:id/edit", formEditNotificationChannel(db))
	adminRequired.POST("/hook/notifications/:id/test", hookSendTestNotification(db, *config))
	// Admin pages
	adminRequired.GET("/admin", renderAdminPage(db, config))
	adminRequired.POST("/admin/data-retention", formAdminDataRetention(db))
	// Profile pages
	authenticatedRoutes.GET("/profile/password", renderChangePasswordPage(db))
	authenticatedRoutes.GET("/profile/api-key", renderAPIKeyViewPage(db))
	authenticatedRoutes.POST("/profile/api-key/regenerate", formRegenerateAPIKey(db))
	authenticatedRoutes.POST("/profile/password", formSelfSetNewPassword(db))
	// Project Pages
	adminRequired.GET("/project/new", renderNewProjectView)
	adminRequired.POST("/project/new", formNewProject(db))
	authenticatedRoutes.GET("/projects/:id", renderProjectViewPage(db))
	authenticatedRoutes.GET("/projects/:project_id/jobs/new", renderCreateJobView(db))
	authenticatedRoutes.POST("/projects/:project_id/jobs/new", formCreateJob(db, config))
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
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if data == nil {
		data = make(map[string]any)
	}
	templateData := data.(map[string]any)
	templateData["userID"] = c.Get("userID")
	templateData["email"] = c.Get("email")
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
