package pkg

import (
	"context"
	"errors"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/user"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/azureadv2"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/openidConnect"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

func loginPage(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		authConfig, err := getAuthConfig(c.Request().Context(), db)
		if err != nil {
			slog.Error("error getting auth config", "error", err)
			return renderErrorPage(c, "Internal server error", http.StatusInternalServerError)
		}
		ssoPath := ""
		ssoName := ""
		if authConfig.EnableSSO {
			switch authConfig.SSOProvider {
			case "google":
				ssoName = "Google"
				ssoPath = "google"
			case "azureadv2":
				ssoName = "Microsoft"
				ssoPath = "azureadv2"
			case "oidc":
				ssoName = "Single-Sign-On"
				ssoPath = "openid-connect"
			default:
				ssoName = "Single-Sign-On"
				ssoPath = "openid-connect"
			}
		}
		return c.Render(http.StatusOK, "login", map[string]any{
			"EnableSSO":            authConfig.EnableSSO,
			"DisablePasswordLogin": authConfig.DisablePasswordLogin,
			"SSOName":              ssoName,
			"SSOPath":              ssoPath,
		})
	}
}

func loginForm(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Form struct {
			Email    string `form:"email"`
			Password string `form:"password"`
		}
		var form Form
		if err := c.Bind(&form); err != nil {
			return renderErrorPage(c, "Invalid form data, email and password are required", http.StatusUnauthorized)
		}
		form.Email = strings.TrimSpace(form.Email)
		form.Email = strings.ToLower(form.Email)
		u, err := db.User.Query().Where(user.EmailEQ(form.Email)).Only(c.Request().Context())
		if err != nil {
			loginFailures.Inc()
			slog.Warn("unknown user attempted to login", "email", form.Email)
			return renderErrorPage(c, "Invalid email or password", http.StatusUnauthorized)
		}
		if compareHashAndPassword(u.EncryptedPassword, form.Password) != nil {
			loginFailures.Inc()
			slog.Warn("user entered invalid password", "email", form.Email)
			return renderErrorPage(c, "Invalid email or password", http.StatusUnauthorized)
		}
		slog.Info("user logged in", "email", form.Email)
		loginSuccesses.Inc()
		return createNewSessionAndRedirect(c, "/", u)
	}
}

func getSession(c echo.Context) (*sessions.Session, error) {
	return session.Get(sessionName, c)
}

func destroySession(c echo.Context) error {
	s, err := getSession(c)
	if err != nil {
		slog.Error("error getting session", "error", err)
		return renderErrorPage(c, "Error getting session", http.StatusInternalServerError)
	}
	s.Options.MaxAge = -1
	s.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/login")
}

func createNewSessionAndRedirect(c echo.Context, path string, u *ent.User) error {
	s, err := getSession(c)
	if err != nil {
		slog.Error("error getting session when creating new session", "error", err)
		return renderErrorPage(c, "Error getting session", http.StatusInternalServerError)
	}
	globalRole := "user"
	if u.Admin == true {
		globalRole = "admin"
	}
	s.Options.MaxAge = 60 * 60 * 24 * 3
	s.Options.Path = "/"
	s.Options.Secure = true
	s.Options.HttpOnly = true
	s.Options.SameSite = http.SameSiteStrictMode
	s.Options.Domain = c.Request().Host
	s.Values["userID"] = u.ID
	s.Values["email"] = u.Email
	s.Values["globalRole"] = globalRole
	s.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, path)
}

func startSSOAuth() echo.HandlerFunc {
	return func(c echo.Context) error {
		gothic.BeginAuthHandler(c.Response(), c.Request())
		return nil
	}
}

func completeSSOAuth(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		u, err := gothic.CompleteUserAuth(c.Response(), c.Request())
		if err != nil {
			slog.Error("error completing SSO auth", "error", err)
			return renderErrorPage(c, "Error completing SSO auth", http.StatusInternalServerError)
		}
		existingUser, err := db.User.Query().WithProjectUsers().Where(user.EmailEqualFold(u.Email)).Only(ctx)
		if err != nil {
			// User does not exist, create a new user
			// Check if any users exist at all, if there are none this one will be an admin
			users, _ := db.User.Query().Count(ctx)
			isAdmin := false
			if users == 0 {
				isAdmin = true
			}
			usr, err := db.User.Create().
				SetEmail(strings.ToLower(u.Email)).
				SetAdmin(isAdmin).
				SetAPIKey(generateLongString()).
				SetCreatedAt(time.Now()).
				SetEncryptedPassword([]byte(generateLongString())).
				SetMustChangePassword(false).
				SetUpdatedAt(time.Now()).
				SetIsSSO(true).
				Save(ctx)
			if err != nil {
				slog.Error("failed to create user", "error", err)
				return renderErrorPage(c, "Failed to create user", http.StatusInternalServerError)
			}
			return createNewSessionAndRedirect(c, "/", usr)
		}
		// Save the user to the session
		db.User.Update().Where(user.IDEQ(existingUser.ID)).SetIsSSO(true).Save(ctx)
		return createNewSessionAndRedirect(c, "/", existingUser)
	}
}

func wireSSOConnection(db *ent.Client, config *Config) error {
	authConfig, err := getAuthConfig(context.Background(), db)
	if err != nil {
		slog.Error("error getting auth config", "error", err)
		return err
	}
	if authConfig.EnableSSO {
		switch authConfig.SSOProvider {
		case "google":
			goth.UseProviders(google.New(authConfig.SSOClientID, authConfig.SSOClientSecret, config.ServerURL+"/auth/google/callback", "email", "profile"))
			return nil
		case "azureadv2":
			goth.UseProviders(azureadv2.New(authConfig.SSOClientID, authConfig.SSOClientSecret, config.ServerURL+"/auth/azureadv2/callback", azureadv2.ProviderOptions{
				Scopes: []azureadv2.ScopeType{"User.Read"},
				Tenant: azureadv2.TenantType(authConfig.EntraTenantID),
			}))
			return nil
		case "oidc":
			open, err := openidConnect.New(authConfig.SSOClientID, authConfig.SSOClientSecret, config.ServerURL+"/auth/openid-connect/callback", authConfig.SSOAuthorizationURL, "email", "profile")
			if err != nil {
				return err
			}
			goth.UseProviders(open)
			return nil
		default:
			slog.Warn("unknown SSO provider", "provider", authConfig.SSOProvider)
			return errors.New("unknown SSO provider")
		}
	}
	return nil
}
