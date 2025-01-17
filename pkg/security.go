package pkg

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"net/http"
	"strings"
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
