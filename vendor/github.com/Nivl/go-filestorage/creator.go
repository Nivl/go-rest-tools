package filestorage

// Creator creates new FileStorage
//go:generate mockgen -destination implementations/mockfilestorage/creator.go -package mockfilestorage github.com/Nivl/go-filestorage Creator
type Creator interface {
	New() (FileStorage, error)
}
