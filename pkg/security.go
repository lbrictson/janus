package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/projectuser"
	"github.com/lbrictson/janus/ent/user"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"strings"
	"sync"
)

func hashAndSaltPassword(password string) ([]byte, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return h, err
}

func compareHashAndPassword(hashedPassword []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New(fmt.Sprintf("password must be at least %v characters long", 8))
	}
	if len(password) > 128 {
		return errors.New(fmt.Sprintf("password must be at most %v characters long", 128))
	}
	return nil
}

func generateLongString() string {
	return strings.Replace(fmt.Sprintf("%v%v%v%v", uuid.NewString(), uuid.NewString(), uuid.NewString(), uuid.NewString()), "-", "", -1)
}

var apiKeys = make(map[string]*ent.User)
var apiKeysLock = &sync.RWMutex{}

func reloadAPIKeys(db *ent.Client) error {
	apiKeysLock.Lock()
	apiKeys = make(map[string]*ent.User)
	allUsers, err := db.User.Query().WithProjectUsers().Where(user.APIKeyNEQ("")).All(context.Background())
	if err != nil {
		apiKeysLock.Unlock()
		return fmt.Errorf("failed to get users: %v", err)
	}
	for _, u := range allUsers {
		apiKeys[u.APIKey] = u
	}
	apiKeysLock.Unlock()
	return nil
}

func middlewareAPIAuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-Key")
		if apiKey == "" {
			return c.JSON(http.StatusUnauthorized, APIError{ErrorMessage: "API key required"})
		}
		// Get the API key from the cache
		apiKeysLock.RLock()
		u, ok := apiKeys[apiKey]
		apiKeysLock.RUnlock()
		if !ok {
			return c.JSON(http.StatusUnauthorized, APIError{ErrorMessage: "Invalid API key"})
		}
		c.Set("userID", u.ID)
		c.Set("email", u.Email)
		if u.Admin {
			c.Set("globalRole", "admin")
		} else {
			c.Set("globalRole", "user")
		}
		return next(c)
	}
}

func middlewareMustBeLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s, err := getSession(c)
		if err != nil {
			slog.Error("error getting session", "error", err)
			return c.Redirect(http.StatusFound, "/login")
		}
		userID, ok := s.Values["userID"]
		if !ok {
			return c.Redirect(http.StatusFound, "/login")
		}
		if userID == nil {
			return c.Redirect(http.StatusFound, "/login")
		}
		c.Set("userID", userID)
		c.Set("email", s.Values["email"])
		c.Set("globalRole", s.Values["globalRole"])
		return next(c)
	}
}

func canUserEditProject(db *ent.Client, userID int, projectID int) bool {
	p, err := db.ProjectUser.Query().WithUser().WithProject().Where(projectuser.HasProjectWith(project.IDEQ(projectID)), projectuser.HasUserWith(user.IDEQ(userID))).All(context.Background())
	if err != nil {
		slog.Error("error getting project user", "error", err)
		return false
	}
	if len(p) == 0 {
		return false
	}
	if p[0].CanEdit {
		return true
	}
	return false
}

func canUserViewProject(db *ent.Client, userID int, projectID int) bool {
	p, err := db.ProjectUser.Query().WithUser().WithProject().Where(projectuser.HasProjectWith(project.IDEQ(projectID)), projectuser.HasUserWith(user.IDEQ(userID))).All(context.Background())
	if err != nil {
		slog.Error("error getting project user", "error", err)
		return false
	}
	if len(p) == 0 {
		return false
	}
	return true
}

func middlewareAdminRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, ok := c.Get("globalRole").(string)
		if !ok {
			return renderErrorPage(c, "Internal error, unable to determine user role", http.StatusInternalServerError)
		}
		if role != "admin" {
			return renderErrorPage(c, "You must be an admin to access this page", http.StatusForbidden)
		}
		return next(c)
	}
}
