package auth

// Code generated; DO NOT EDIT.

import (
	"testing"

		"github.com/stretchr/testify/assert"

		"github.com/satori/go.uuid"

		"github.com/Nivl/go-rest-tools/storage/db/mockdb"

	"github.com/Nivl/go-types/datetime"
)


func TestUserSaveNew(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsert("*auth.User")

	u := &User{}
	err := u.Save(mockDB)

	assert.NoError(t, err, "Save() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, u.ID, "ID should have been set")
}

func TestUserSaveExisting(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectUpdate("*auth.User")

	u := &User{}
	id := uuid.NewV4().String()
	u.ID = id
	err := u.Save(mockDB)

	assert.NoError(t, err, "Save() should not have fail")
	mockDB.AssertExpectations(t)
	assert.Equal(t, id, u.ID, "ID should not have changed")
}

func TestUserCreate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsert("*auth.User")

	u := &User{}
	err := u.Create(mockDB)

	assert.NoError(t, err, "Create() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.CreatedAt, "CreatedAt should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserCreateWithID(t *testing.T) {
	mockDB := &mockdb.Queryable{}

	u := &User{}
	u.ID = uuid.NewV4().String()

	err := u.Create(mockDB)
	assert.Error(t, err, "Create() should have fail")
	mockDB.AssertExpectations(t)
}

func TestUserDoCreate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsert("*auth.User")

	u := &User{}
	err := u.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.CreatedAt, "CreatedAt should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserDoCreateWithDate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsert("*auth.User")

	createdAt := datetime.Now().AddDate(0, 0, 1)
	u := &User{CreatedAt: createdAt}
	err := u.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.True(t, u.CreatedAt.Equal(createdAt), "CreatedAt should not have been updated")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserDoCreateFail(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectInsertError("*auth.User")

	u := &User{}
	err := u.doCreate(mockDB)

	assert.Error(t, err, "doCreate() should have fail")
	mockDB.AssertExpectations(t)
}


func TestUserUpdate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectUpdate("*auth.User")

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.Update(mockDB)

	assert.NoError(t, err, "Update() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserUpdateWithoutID(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	u := &User{}
	err := u.Update(mockDB)

	assert.Error(t, err, "Update() should not have fail")
	mockDB.AssertExpectations(t)
}


func TestUserDoUpdate(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectUpdate("*auth.User")

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.doUpdate(mockDB)

	assert.NoError(t, err, "doUpdate() should not have fail")
	mockDB.AssertExpectations(t)
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserDoUpdateWithoutID(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	u := &User{}
	err := u.doUpdate(mockDB)

	assert.Error(t, err, "doUpdate() should not have fail")
	mockDB.AssertExpectations(t)
}

func TestUserDoUpdateFail(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectUpdateError("*auth.User")

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.doUpdate(mockDB)

	assert.Error(t, err, "doUpdate() should have fail")
	mockDB.AssertExpectations(t)
}

func TestUserDelete(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectDeletion()

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.Delete(mockDB)

	assert.NoError(t, err, "Delete() should not have fail")
	mockDB.AssertExpectations(t)
}

func TestUserDeleteWithoutID(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	u := &User{}
	err := u.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
	mockDB.AssertExpectations(t)
}

func TestUserDeleteError(t *testing.T) {
	mockDB := &mockdb.Queryable{}
	mockDB.ExpectDeletionError()

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
	mockDB.AssertExpectations(t)
}

func TestUserGetID(t *testing.T) {
	u := &User{}
	u.ID = uuid.NewV4().String()
	assert.Equal(t, u.ID, u.GetID(), "GetID() did not return the right ID")
}

func TestUserSetID(t *testing.T) {
	u := &User{}
	u.SetID(uuid.NewV4().String())
	assert.NotEmpty(t, u.ID, "SetID() did not set the ID")
}

func TestUserIsZero(t *testing.T) {
	empty := &User{}
	assert.True(t, empty.IsZero(), "IsZero() should return true for empty struct")

	var nilStruct *User
	assert.True(t, nilStruct.IsZero(), "IsZero() should return true for nil struct")

	valid := &User{ID: uuid.NewV4().String()}
	assert.False(t, valid.IsZero(), "IsZero() should return false for valid struct")
}