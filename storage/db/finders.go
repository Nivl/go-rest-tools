package db

// Get is the same as sqlx.Get() but do not returns an error on empty results
// func Get(q DB, dest interface{}, query string, args ...interface{}) error {
// 	err := q.Get(dest, query, args...)
// 	if IsNotFound(err) {
// 		return nil
// 	}
// 	return err
// }
