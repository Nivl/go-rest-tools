package auth

// Code generated; DO NOT EDIT.

import (
	

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/satori/go.uuid"

	"github.com/Nivl/go-sqldb/implementations/mocksqldb"

	gomock "github.com/golang/mock/gomock"

	"github.com/Nivl/go-types/datetime"
)











func TestSessionDoCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertSuccess(&Session{})

	s := &Session{}
	err := s.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	assert.NotEmpty(t, s.ID, "ID should have been set")
	assert.NotNil(t, s.CreatedAt, "CreatedAt should have been set")
	assert.NotNil(t, s.UpdatedAt, "UpdatedAt should have been set")
}

func TestSessionDoCreateWithDate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertSuccess(&Session{})

	createdAt := datetime.Now().AddDate(0, 0, 1)
	s := &Session{CreatedAt: createdAt}
	err := s.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	assert.NotEmpty(t, s.ID, "ID should have been set")
	assert.True(t, s.CreatedAt.Equal(createdAt), "CreatedAt should not have been updated")
	assert.NotNil(t, s.UpdatedAt, "UpdatedAt should have been set")
}

func TestSessionDoCreateFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertError(&Session{}, errors.New("sql error"))

	s := &Session{}
	err := s.doCreate(mockDB)

	assert.Error(t, err, "doCreate() should have fail")
}







func TestSessionDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().DeletionSuccess()

	s := &Session{}
	s.ID = uuid.NewV4().String()
	err := s.Delete(mockDB)

	assert.NoError(t, err, "Delete() should not have fail")
}

func TestSessionDeleteWithoutID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	s := &Session{}
	err := s.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
}

func TestSessionDeleteError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().DeletionError(errors.New("sql error"))

	s := &Session{}
	s.ID = uuid.NewV4().String()
	err := s.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
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