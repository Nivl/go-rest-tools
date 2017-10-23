package hasher

// Hasher hashes and validate strings
//go:generate mockgen -destination mockhasher/hasher.go -package mockhasher github.com/Nivl/go-rest-tools/security/hasher Hasher
type Hasher interface {
	// Hash returns a hash for the provided string
	Hash(raw string) (string, error)

	// IsValid checks if a hash and a string match
	IsValid(hash string, raw string) bool
}
