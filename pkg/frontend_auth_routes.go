package pkg

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/user"
	"log/slog"
	"net/http"
	"strings"
)

func loginPage(db *ent.Client, config *Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		authConfig, err := getAuthconfig(c.Request().Context(), db)
		if err != nil {
			slog.Error("error getting auth config", "error", err)
			return renderErrorPage(c, "Internal server error", http.StatusInternalServerError)
		}
		ssoName := ""
		if authConfig.EnableSSO {
			switch authConfig.SSOProvider {
			case "google":
				ssoName = "Google"
			case "entra":
				ssoName = "Microsoft"
			case "saml":
				ssoName = "Single-Sign-On"
			default:
				ssoName = "Single-Sign-On"
			}
		}
		return c.Render(http.StatusOK, "login", map[string]any{
			"EnableSSO":            authConfig.EnableSSO,
			"DisablePasswordLogin": authConfig.DisablePasswordLogin,
			"SSOName":              ssoName,
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
			slog.Warn("unknown user attempted to login", "email", form.Email)
			return renderErrorPage(c, "Invalid email or password", http.StatusUnauthorized)
		}
		if compareHashAndPassword(u.EncryptedPassword, form.Password) != nil {
			slog.Warn("user entered invalid password", "email", form.Email)
			return renderErrorPage(c, "Invalid email or password", http.StatusUnauthorized)
		}
		slog.Info("user logged in", "email", form.Email)
		return createNewSessionAndRedirect(c, "/", u)
	}
}

func getSession(c echo.Context) (*sessions.Session, error) {
	return session.Get("janus", c)
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
		slog.Error("error getting session", "error", err)
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
