package auth

// Code auto-generated; DO NOT EDIT

import (
	"errors"


	"github.com/Nivl/go-rest-tools/primitives/apierror"
	"github.com/Nivl/go-rest-tools/storage/db"
	uuid "github.com/satori/go.uuid"
)





// Exists checks if a user exists for a specific ID
func Exists(q db.DB, id string) (bool, error) {
	exists := false
	stmt := "SELECT exists(SELECT 1 FROM users WHERE id=$1 and deleted_at IS NULL)"
	err := db.Get(q, &exists, stmt, id)
	return exists, err
}

// Save creates or updates the article depending on the value of the id using
// a transaction
func (u *User) Save(q db.DB) error {
	if u.ID == "" {
		return u.Create(q)
	}

	return u.Update(q)
}

// Create persists a user in the database
func (u *User) Create(q db.DB) error {
	if u.ID != "" {
		return errors.New("cannot persist a user that already has an ID")
	}

	return u.doCreate(q)
}

// doCreate persists a user in the database using a Node
func (u *User) doCreate(q db.DB) error {
	if u == nil {
		return errors.New("user not instanced")
	}

	u.ID = uuid.NewV4().String()
	u.UpdatedAt = db.Now()
	if u.CreatedAt == nil {
		u.CreatedAt = db.Now()
	}

	stmt := "INSERT INTO users (id, created_at, updated_at, deleted_at, name, email, password, is_admin) VALUES (:id, :created_at, :updated_at, :deleted_at, :name, :email, :password, :is_admin)"
	_, err := q.NamedExec(stmt, u)

  return apierror.NewFromSQL(err)
}

// Update updates most of the fields of a persisted user using a transaction
// Excluded fields are id, created_at, deleted_at, etc.
func (u *User) Update(q db.DB) error {
	if u.ID == "" {
		return errors.New("cannot update a non-persisted user")
	}

	return u.doUpdate(q)
}

// doUpdate updates a user in the database using an optional transaction
func (u *User) doUpdate(q db.DB) error {
	if u.ID == "" {
		return errors.New("cannot update a non-persisted user")
	}

	u.UpdatedAt = db.Now()

	stmt := "UPDATE users SET id=:id, created_at=:created_at, updated_at=:updated_at, deleted_at=:deleted_at, name=:name, email=:email, password=:password, is_admin=:is_admin WHERE id=:id"
	_, err := q.NamedExec(stmt, u)

	return apierror.NewFromSQL(err)
}

// Delete removes a user from the database using a transaction
func (u *User) Delete(q db.DB) error {
	if u == nil {
		return errors.New("user not instanced")
	}

	if u.ID == "" {
		return errors.New("user has not been saved")
	}

	stmt := "DELETE FROM users WHERE id=$1"
	_, err := q.Exec(stmt, u.ID)

	return err
}

// Trash soft delete a user using a transaction
func (u *User) Trash(q db.DB) error {
	return u.doTrash(q)
}

// doTrash performs a soft delete operation on a user using an optional transaction
func (u *User) doTrash(q db.DB) error {
	if u.ID == "" {
		return errors.New("cannot trash a non-persisted user")
	}

	u.DeletedAt = db.Now()

	stmt := "UPDATE users SET deleted_at = $2 WHERE id=$1"
	_, err := q.Exec(stmt, u.ID, u.DeletedAt)
	return err
}

// IsZero checks if the object is either nil or don't have an ID
func (u *User) IsZero() bool {
	return u == nil || u.ID == ""
}