package auth

// Code generated; DO NOT EDIT.

import (
	"errors"
	

	"github.com/Nivl/go-rest-tools/types/apperror"
	"github.com/Nivl/go-types/datetime"
	"github.com/Nivl/go-sqldb"
	uuid "github.com/satori/go.uuid"
)












// doCreate persists a session in the database using a Node
func (s *Session) doCreate(q sqldb.Queryable) error {
	s.ID = uuid.NewV4().String()
	s.UpdatedAt = datetime.Now()
	if s.CreatedAt == nil {
		s.CreatedAt = datetime.Now()
	}

	stmt := "INSERT INTO user_sessions (id, created_at, updated_at, deleted_at, user_id) VALUES (:id, :created_at, :updated_at, :deleted_at, :user_id)"
	_, err := q.NamedExec(stmt, s)

  return apperror.NewFromSQL(err)
}





// Delete removes a session from the database
func (s *Session) Delete(q sqldb.Queryable) error {
	if s.ID == "" {
		return errors.New("session has not been saved")
	}

	stmt := "DELETE FROM user_sessions WHERE id=$1"
	_, err := q.Exec(stmt, s.ID)

	return err
}

// IsZero checks if the object is either nil or don't have an ID
func (s *Session) IsZero() bool {
	return s == nil || s.ID == ""
}