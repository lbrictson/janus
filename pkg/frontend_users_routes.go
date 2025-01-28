package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/projectuser"
	"github.com/lbrictson/janus/ent/user"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
)

func renderUsersPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type FrontendUserData struct {
			ID           int
			Email        string
			Role         string
			ProjectCount int
			IsSSO        bool
		}
		users, err := db.User.Query().Order(ent.Desc(user.FieldEmail)).All(c.Request().Context())
		if err != nil {
			slog.Error("failed to query users for users page", "error", err)
			return renderErrorPage(c, "failed to query users", http.StatusInternalServerError)
		}
		frontendUsers := make([]FrontendUserData, len(users))
		for i, u := range users {
			role := "User"
			if u.Admin {
				role = "Admin"
			}
			frontendUsers[i] = FrontendUserData{
				ID:           u.ID,
				Email:        u.Email,
				Role:         role,
				ProjectCount: 0,
				IsSSO:        u.IsSSO,
			}
		}
		wg := sync.WaitGroup{}
		for i := range frontendUsers {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				_, count, err := getUserProjectPermissions(db, frontendUsers[i].ID)
				if err != nil {
					slog.Error("failed to get user permissions", "error", err)
					return
				}
				frontendUsers[i].ProjectCount = count
			}(i)
		}
		wg.Wait()
		return c.Render(http.StatusOK, "users", map[string]any{
			"Users": frontendUsers,
		})
	}
}

func renderEditUserPage(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			return renderErrorPage(c, "failed to parse user id", http.StatusBadRequest)
		}
		u, err := db.User.Query().WithProjectUsers().Where(user.ID(id)).Only(c.Request().Context())
		if err != nil {
			slog.Error("failed to query user for edit page", "error", err)
			return renderErrorPage(c, "failed to query user", http.StatusInternalServerError)
		}
		type ProjectData struct {
			ID     int
			Name   string
			Access string
		}
		projects, err := db.Project.Query().Order(ent.Desc(project.FieldName)).All(c.Request().Context())
		if err != nil {
			slog.Error("failed to query projects for edit page", "error", err)
			return renderErrorPage(c, "failed to query projects", http.StatusInternalServerError)
		}
		projectData := make([]ProjectData, len(projects))
		for i, p := range projects {
			projectData[i] = ProjectData{
				ID:     p.ID,
				Name:   p.Name,
				Access: "None",
			}
		}
		userPermissions, _, err := getUserProjectPermissions(db, id)
		if err != nil {
			slog.Error("failed to get user permissions", "error", err)
			return renderErrorPage(c, "failed to get user permissions", http.StatusInternalServerError)
		}
		for i, p := range projectData {
			if access, ok := userPermissions[p.ID]; ok {
				projectData[i].Access = access
			}
		}
		return c.Render(http.StatusOK, "edit-user", map[string]any{
			"User":     u,
			"Projects": projectData,
		})
	}
}

func formAdminEditUserSetNewPassword(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			return renderErrorPage(c, "failed to parse user id", http.StatusBadRequest)
		}
		u, err := db.User.Query().Where(user.ID(id)).Only(c.Request().Context())
		if err != nil {
			slog.Error("failed to query user for password reset", "error", err)
			return renderErrorPage(c, "failed to query user", http.StatusInternalServerError)
		}
		type Form struct {
			Password        string `form:"new_password"`
			ConfirmPassword string `form:"confirm_password"`
		}
		form := new(Form)
		if err := c.Bind(form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "failed to bind form", http.StatusBadRequest)
		}
		if form.Password != form.ConfirmPassword {
			return c.Render(http.StatusBadRequest, "edit-user", map[string]any{
				"User":  u,
				"Error": "Passwords do not match",
			})
		}
		if validatePassword(form.Password) != nil {
			return c.Render(http.StatusBadRequest, "edit-user", map[string]any{
				"User":  u,
				"Error": "Password does not meet requirements",
			})
		}
		hashedPassword, err := hashAndSaltPassword(form.Password)
		if err != nil {
			slog.Error("failed to hash password", "error", err)
			return renderErrorPage(c, "failed to hash password", http.StatusInternalServerError)
		}
		_, err = db.User.Update().Where(user.ID(id)).SetEncryptedPassword(hashedPassword).Save(c.Request().Context())
		if err != nil {
			slog.Error("failed to update user password", "error", err)
			return renderErrorPage(c, "failed to update user password", http.StatusInternalServerError)
		}
		slog.Info("admin updated user password", "user_id", id, "admin_id", c.Get("userID"))
		return c.Render(http.StatusOK, "edit-user", map[string]any{
			"User":    u,
			"Success": "Password updated",
		})
	}
}

func formAdminSetUserRole(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			return renderErrorPage(c, "failed to parse user id", http.StatusBadRequest)
		}
		u, err := db.User.Query().Where(user.ID(id)).Only(c.Request().Context())
		if err != nil {
			slog.Error("failed to query user for role change", "error", err)
			return renderErrorPage(c, "failed to query user", http.StatusInternalServerError)
		}
		type Form struct {
			Role string `form:"role"`
		}
		form := new(Form)
		if err := c.Bind(form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "failed to bind form", http.StatusBadRequest)
		}
		role := false
		if form.Role == "admin" {
			role = true
		}
		_, err = db.User.Update().Where(user.ID(id)).SetAdmin(role).Save(c.Request().Context())
		if err != nil {
			slog.Error("failed to update user role", "error", err)
			return renderErrorPage(c, "failed to update user role", http.StatusInternalServerError)
		}
		slog.Info("admin updated user role", "user_id", id, "admin_id", c.Get("userID"))
		reloadAPIKeys(db)
		return c.Render(http.StatusOK, "edit-user", map[string]any{
			"User":    u,
			"Success": "Role updated",
		})
	}
}

func renderCreateUserPage(c echo.Context) error {
	return c.Render(http.StatusOK, "new-user", nil)
}

func formCreateNewUser(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Form struct {
			Email    string `form:"email"`
			Password string `form:"password"`
			Role     string `form:"role"`
		}
		form := new(Form)
		if err := c.Bind(form); err != nil {
			slog.Error("failed to bind form", "error", err)
			return renderErrorPage(c, "failed to bind form", http.StatusBadRequest)
		}
		if validatePassword(form.Password) != nil {
			return c.Render(http.StatusBadRequest, "new-user", map[string]any{
				"Error": "Password does not meet requirements",
			})
		}
		hashedPassword, err := hashAndSaltPassword(form.Password)
		if err != nil {
			slog.Error("failed to hash password", "error", err)
			return renderErrorPage(c, "failed to hash password", http.StatusInternalServerError)
		}
		isAdmin := false
		if form.Role == "admin" {
			isAdmin = true
		}
		_, err = db.User.Create().
			SetEmail(form.Email).
			SetEncryptedPassword(hashedPassword).
			SetAdmin(isAdmin).
			SetAPIKey(generateLongString()).
			Save(c.Request().Context())
		if err != nil {
			slog.Error("failed to create user", "error", err)
			return renderErrorPage(c, "failed to create user", http.StatusInternalServerError)
		}
		slog.Info("admin created new user", "email", form.Email, "admin_id", c.Get("userID"))
		reloadAPIKeys(db)
		return c.Redirect(http.StatusSeeOther, "/users")
	}
}

func formDeleteUser(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			return renderErrorPage(c, "failed to parse user id", http.StatusBadRequest)
		}
		_, err = db.User.Delete().Where(user.ID(id)).Exec(c.Request().Context())
		if err != nil {
			slog.Error("failed to delete user", "error", err)
			return renderErrorPage(c, "failed to delete user", http.StatusInternalServerError)
		}
		slog.Info("admin deleted user", "user_id", id, "admin_id", c.Get("userID"))
		reloadAPIKeys(db)
		return c.Redirect(http.StatusSeeOther, "/users")
	}
}

func formUpdateUserPermissions(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			return renderErrorPage(c, "failed to parse user id", http.StatusBadRequest)
		}
		allProjects, err := db.Project.Query().All(c.Request().Context())
		if err != nil {
			slog.Error("failed to query projects", "error", err)
			return renderErrorPage(c, "failed to query projects", http.StatusInternalServerError)
		}
		userPermissions, _, err := getUserProjectPermissions(db, id)
		if err != nil {
			slog.Error("failed to get user permissions", "error", err)
			return renderErrorPage(c, "failed to get user permissions", http.StatusInternalServerError)
		}
		for _, p := range allProjects {
			access := c.FormValue(strconv.Itoa(p.ID))
			if access == "None" {
				db.ProjectUser.Delete().Where(projectuser.HasUserWith(user.IDEQ(id)), projectuser.HasProjectWith(project.IDEQ(p.ID))).Exec(c.Request().Context())
				continue
			}
			if userPermissions[p.ID] != "None" {
				_, err = db.ProjectUser.Update().Where(projectuser.HasUserWith(user.IDEQ(id)), projectuser.HasProjectWith(project.IDEQ(p.ID))).SetCanEdit(access == "Edit").Save(c.Request().Context())
				if err != nil {
					slog.Error("failed to update project user", "error", err)
					return renderErrorPage(c, "failed to update project user", http.StatusInternalServerError)
				}
				continue
			}
			_, err = db.ProjectUser.Create().
				SetProjectID(p.ID).
				SetUserID(id).
				SetCanEdit(access == "Edit").
				Save(c.Request().Context())
			if err != nil {
				slog.Error("failed to create project user", "error", err)
				return renderErrorPage(c, "failed to create project user", http.StatusInternalServerError)
			}
		}
		reloadAPIKeys(db)
		slog.Info("admin updated user permissions", "user_id", id, "admin_id", c.Get("userID"))
		return c.Redirect(http.StatusSeeOther, "/users")
	}
}
