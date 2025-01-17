package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"log/slog"
	"net/http"
)

func renderChangePasswordPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(200, "change-password", nil)
	}
}

func renderAPIKeyViewPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, err := getSelf(c, db)
		if err != nil {
			return renderErrorPage(c, "failed to get user from session", http.StatusBadRequest)
		}
		return c.Render(200, "api-key", map[string]any{
			"APIKey": self.APIKey,
		})
	}
}

func formRegenerateAPIKey(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		self, err := getSelf(c, db)
		if err != nil {
			slog.Error("failed to get user from session", "error", err)
			return renderErrorPage(c, "failed to get user from session", http.StatusBadRequest)
		}
		newKey := generateLongString()
		slog.Info("self service API key regeneration completed", "user", self.Email)
		self.Update().SetAPIKey(newKey).SaveX(c.Request().Context())
		return c.Render(200, "api-key", map[string]any{
			"APIKey":  newKey,
			"Success": "Successfully regenerated API key",
		})
	}
}

func formSelfSetNewPassword(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Form struct {
			NewPassword     string `form:"new_password"`
			CurrentPassword string `form:"current_password"`
			ConfirmPassword string `form:"confirm_password"`
		}
		f := new(Form)
		if err := c.Bind(f); err != nil {
			slog.Error("failed to bind form for password change", "error", err)
			return renderErrorPage(c, "failed to bind form", http.StatusBadRequest)
		}
		self, err := getSelf(c, db)
		if err != nil {
			slog.Error("failed to get user from session", "error", err)
			return renderErrorPage(c, "failed to get user from session", http.StatusBadRequest)
		}
		if f.NewPassword != f.ConfirmPassword {
			return c.Render(http.StatusBadRequest, "change-password", map[string]any{
				"Error": "New password and confirm password do not match",
			})
		}
		if err := compareHashAndPassword(self.EncryptedPassword, f.CurrentPassword); err != nil {
			return c.Render(http.StatusBadRequest, "change-password", map[string]any{
				"Error": "Current password is incorrect",
			})
		}
		if err := validatePassword(f.NewPassword); err != nil {
			return c.Render(http.StatusBadRequest, "change-password", map[string]any{
				"Error": err.Error(),
			})
		}
		h, err := hashAndSaltPassword(f.NewPassword)
		if err != nil {
			return renderErrorPage(c, "failed to hash new password", http.StatusInternalServerError)
		}
		slog.Info("self service password change completed", "user", self.Email)
		self.Update().SetEncryptedPassword(h).SaveX(c.Request().Context())
		return c.Render(http.StatusOK, "change-password", map[string]any{
			"Success": "Successfully changed password",
		})
	}
}
