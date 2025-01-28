package pkg

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/project"
	"github.com/lbrictson/janus/ent/projectuser"
	"net/http"
	"strconv"
)

func apiGetProjects(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		projects, err := db.Project.Query().WithJobs().All(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get projects: %v", err)})
		}
		permissions, err := db.ProjectUser.Query().WithUser().WithProject().Where(projectuser.HasProject()).All(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get permissions: %v", err)})
		}
		allUsers, err := db.User.Query().WithProjectUsers().All(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get users: %v", err)})
		}
		userMap := make(map[int]*ent.User)
		for _, u := range allUsers {
			userMap[u.ID] = u
		}
		resp := []APIProject{}
		for _, p := range projects {
			proj := APIProject{
				ID:           p.ID,
				Name:         p.Name,
				Description:  p.Description,
				NumberOfJobs: len(p.Edges.Jobs),
				Editors:      make([]APIUser, 0),
				Users:        make([]APIUser, 0),
			}
			for _, perm := range permissions {
				if perm.Edges.Project.ID == p.ID {
					if perm.CanEdit {
						proj.Editors = append(proj.Editors, APIUserFromEntUser(userMap[perm.Edges.User.ID]))
					} else {
						proj.Users = append(proj.Users, APIUserFromEntUser(userMap[perm.Edges.User.ID]))
					}
				}
			}
			resp = append(resp, proj)
		}
		return c.JSON(http.StatusOK, resp)
	}
}

func apiGetProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError{ErrorMessage: fmt.Sprintf("invalid project id: %v", err)})
		}
		project, err := db.Project.Query().Where(project.IDEQ(idInt)).WithJobs().Only(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get project: %v", err)})
		}
		p := APIProject{
			ID:           project.ID,
			Name:         project.Name,
			Description:  project.Description,
			NumberOfJobs: len(project.Edges.Jobs),
			Editors:      make([]APIUser, 0),
			Users:        make([]APIUser, 0),
		}
		permissions, err := db.ProjectUser.Query().WithUser().WithProject().Where(projectuser.HasProject()).All(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get permissions: %v", err)})
		}
		allUsers, err := db.User.Query().WithProjectUsers().All(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get users: %v", err)})
		}
		userMap := make(map[int]*ent.User)
		for _, u := range allUsers {
			userMap[u.ID] = u
		}
		for _, perm := range permissions {
			if perm.Edges.Project.ID == project.ID {
				if perm.CanEdit {
					p.Editors = append(p.Editors, APIUserFromEntUser(userMap[perm.Edges.User.ID]))
				} else {
					p.Users = append(p.Users, APIUserFromEntUser(userMap[perm.Edges.User.ID]))
				}
			}
		}
		return c.JSON(http.StatusOK, p)
	}
}

func apiCreateProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		type Form struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		f := new(Form)
		if err := c.Bind(f); err != nil {
			return c.JSON(http.StatusBadRequest, APIError{ErrorMessage: fmt.Sprintf("failed to bind form: %v", err)})
		}
		project, err := db.Project.Create().SetName(f.Name).SetDescription(f.Description).Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to create project: %v", err)})
		}
		return c.JSON(http.StatusOK, APIProject{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
		})
	}
}

func apiUpdateProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError{ErrorMessage: fmt.Sprintf("invalid project id: %v", err)})
		}
		type Form struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		f := new(Form)
		if err := c.Bind(f); err != nil {
			return c.JSON(http.StatusBadRequest, APIError{ErrorMessage: fmt.Sprintf("failed to bind form: %v", err)})
		}
		project, err := db.Project.Get(c.Request().Context(), idInt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to get project: %v", err)})
		}
		project, err = project.Update().SetName(f.Name).SetDescription(f.Description).Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to update project: %v", err)})
		}
		return apiGetProject(db)(c)
	}
}

func apiDeleteProject(db *ent.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, APIError{ErrorMessage: fmt.Sprintf("invalid project id: %v", err)})
		}
		err = db.Project.DeleteOneID(idInt).Exec(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, APIError{ErrorMessage: fmt.Sprintf("failed to delete project: %v", err)})
		}
		return c.JSON(http.StatusOK, APIError{ErrorMessage: "project deleted"})
	}
}
