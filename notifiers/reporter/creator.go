package reporter

// Creator describes a struct used to create reporters
//go:generate mockgen -destination implementations/mockreporter/creator.go -package mockreporter github.com/Nivl/go-rest-tools/notifiers/reporter Creator
type Creator interface {
	New() (Reporter, error)
}
