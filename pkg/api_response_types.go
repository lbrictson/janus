package pkg

import "github.com/lbrictson/janus/ent"

type APIError struct {
	ErrorMessage string `json:"error"`
}

type APIUser struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"has_global_admin_permissions"`
}

type APIProject struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	NumberOfJobs int       `json:"number_of_jobs"`
	Editors      []APIUser `json:"editors"`
	Users        []APIUser `json:"users"`
}

func APIUserFromEntUser(u *ent.User) APIUser {
	return APIUser{
		ID:      u.ID,
		Email:   u.Email,
		IsAdmin: u.Admin,
	}
}
