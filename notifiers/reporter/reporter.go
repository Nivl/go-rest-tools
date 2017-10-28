package reporter

import "github.com/Nivl/go-rest-tools/security/auth"

// Reporter describes a struct used to report errors
//go:generate mockgen -destination implementations/mockreporter/reporter.go -package mockreporter github.com/Nivl/go-rest-tools/notifiers/reporter Reporter
type Reporter interface {
	SetUser(u *auth.User)
	AddTag(key, value string)
	AddTags(tags map[string]string)
	ReportError(err error)
	ReportErrorAndWait(err error)
}
