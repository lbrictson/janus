package pkg

import (
	"context"
	"github.com/lbrictson/janus/ent"
	"github.com/lbrictson/janus/ent/projectuser"
	"github.com/lbrictson/janus/ent/user"
)

func getUserProjectPermissions(db *ent.Client, userID int) (map[int]string, int, error) {
	projects, err := db.Project.Query().All(context.Background())
	if err != nil {
		return nil, 0, err
	}
	permissions := make(map[int]string)
	for _, p := range projects {
		permissions[p.ID] = "None"
	}
	projectsWithPermissions := 0
	userSpecificPermissions := db.ProjectUser.Query().WithProject().WithUser().Where(projectuser.HasUserWith(user.IDEQ(userID))).AllX(context.Background())
	for _, up := range userSpecificPermissions {
		if up.CanEdit {
			permissions[up.Edges.Project.ID] = "Edit"
		} else {
			permissions[up.Edges.Project.ID] = "View"
		}
		projectsWithPermissions++
	}
	return permissions, projectsWithPermissions, nil
}
