package pkg

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
	"net/http/httptest"
	"testing"
)

func TestGetSelf(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()
	ctx := context.Background()

	// Create test user
	user, err := client.User.Create().
		SetEmail("test@example.com").
		SetEncryptedPassword([]byte("hash")).
		SetAPIKey(generateLongString()).
		Save(ctx)
	if err != nil {
		t.Fatalf("failed creating user: %v", err)
	}

	tests := []struct {
		name      string
		userID    interface{}
		wantError bool
	}{
		{
			name:      "valid user",
			userID:    user.ID,
			wantError: false,
		},
		{
			name:      "invalid user ID type",
			userID:    "not an int",
			wantError: true,
		},
		{
			name:      "non-existent user",
			userID:    999999,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup echo context
			e := echo.New()
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.Set("userID", tt.userID)

			// Test getSelf
			u, err := getSelf(c, client)
			if tt.wantError {
				if err == nil {
					t.Error("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if u.ID != user.ID {
					t.Error("got wrong user")
				}
			}
		})
	}
}
