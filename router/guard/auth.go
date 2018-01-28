package guard

import (
	"github.com/Nivl/go-rest-tools/security/auth"
	"github.com/Nivl/go-rest-tools/types/apperror"
)

// RouteAuth represents a middleware used to allow/block the access to an endpoint
type RouteAuth func(*auth.User) apperror.Error

// LoggedUserAccess is a auth middleware that filters out anonymous users
func LoggedUserAccess(u *auth.User) apperror.Error {
	if u == nil || u.ID == "" {
		return apperror.NewUnauthorized()
	}
	return nil
}

// AdminAccess is a auth middleware that filters out non admin users
func AdminAccess(u *auth.User) apperror.Error {
	if err := LoggedUserAccess(u); err != nil {
		return apperror.NewUnauthorized()
	}
	if !u.IsAdmin {
		return apperror.NewForbidden()
	}
	return nil
}
