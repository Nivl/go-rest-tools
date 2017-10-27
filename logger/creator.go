package logger

// Creator describes a struct used to create loggers
//go:generate mockgen -destination implementations/mocklogger/creator.go -package mocklogger github.com/Nivl/go-rest-tools/logger Creator
type Creator interface {
	New() (Logger, error)
}
