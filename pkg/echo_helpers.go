package pkg

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/user"
)

func getSelf(c echo.Context, db *ent.Client) (*ent.User, error) {
	i, ok := c.Get("userID").(int)
	if !ok {
		return nil, errors.New("failed to get userID from context")
	}
	u, err := db.User.Query().WithProjectUsers().Where(user.IDEQ(i)).Only(c.Request().Context())
	if err != nil {
		return nil, err
	}
	return u, nil
}
