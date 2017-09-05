package auth

// Code generated; DO NOT EDIT.

import (
	"testing"

		"github.com/stretchr/testify/assert"

		"github.com/satori/go.uuid"

		"github.com/Nivl/go-rest-tools/storage/db/mockdb"

	"github.com/Nivl/go-rest-tools/types/datetime"
)






func TestSessionDoCreate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsert("*auth.Session")

	s := &Session{}
	err := s.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, s.ID, "ID should have been set")
	assert.NotNil(t, s.CreatedAt, "CreatedAt should have been set")
	assert.NotNil(t, s.UpdatedAt, "UpdatedAt should have been set")
}

func TestSessionDoCreateWithDate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsert("*auth.Session")

	createdAt := datetime.Now().AddDate(0, 0, 1)
	s := &Session{CreatedAt: createdAt}
	err := s.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, s.ID, "ID should have been set")
	assert.True(t, s.CreatedAt.Equal(createdAt), "CreatedAt should not have been updated")
	assert.NotNil(t, s.UpdatedAt, "UpdatedAt should have been set")
}

func TestSessionDoCreateFail(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsertError("*auth.Session")

	s := &Session{}
	err := s.doCreate(mockDB)

	assert.Error(t, err, "doCreate() should have fail")
	mockDB.AssertExpectations(t)
}







func TestSessionDelete(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectDeletion()

	s := &Session{}
	s.ID = uuid.NewV4().String()
	err := s.Delete(mockDB)

	assert.NoError(t, err, "Delete() should not have fail")
	mockDB.AssertExpectations(t)
}

func TestSessionDeleteWithoutID(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	s := &Session{}
	err := s.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
	mockDB.AssertExpectations(t)
}

func TestSessionDeleteError(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectDeletionError()

	s := &Session{}
	s.ID = uuid.NewV4().String()
	err := s.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
	mockDB.AssertExpectations(t)
}

func TestSessionGetID(t *testing.T) {
	s := &Session{}
	s.ID = uuid.NewV4().String()
	assert.Equal(t, s.ID, s.GetID(), "GetID() did not return the right ID")
}

func TestSessionSetID(t *testing.T) {
	s := &Session{}
	s.SetID(uuid.NewV4().String())
	assert.NotEmpty(t, s.ID, "SetID() did not set the ID")
}

func TestSessionIsZero(t *testing.T) {
	empty := &Session{}
	assert.True(t, empty.IsZero(), "IsZero() should return true for empty struct")

	var nilStruct *Session
	assert.True(t, nilStruct.IsZero(), "IsZero() should return true for nil struct")

	valid := &Session{ID: uuid.NewV4().String()}
	assert.False(t, valid.IsZero(), "IsZero() should return false for valid struct")
}