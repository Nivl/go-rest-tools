package auth

// Code generated; DO NOT EDIT.

import (
	"strings"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/satori/go.uuid"

	"github.com/Nivl/go-sqldb/implementations/mocksqldb"

	gomock "github.com/golang/mock/gomock"

	"github.com/Nivl/go-types/datetime"
)

func TestJoinUserSQL(t *testing.T) {
	fields := []string{ "id", "created_at", "updated_at", "deleted_at", "name", "email", "password", "is_admin" }
	totalFields := len(fields)
	output := JoinUserSQL("tofind")

	assert.Equal(t, totalFields*2, strings.Count(output, "tofind."), "wrong number of fields returned")
	assert.True(t, strings.HasSuffix(output, "\""), "JoinSQL() output should end with a \"")
}

func TestUserSaveNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertSuccess(&User{})

	u := &User{}
	err := u.Save(mockDB)

	assert.NoError(t, err, "Save() should not have fail")
	assert.NotEmpty(t, u.ID, "ID should have been set")
}

func TestUserSaveExisting(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().UpdateSuccess(&User{})

	u := &User{}
	id := uuid.NewV4().String()
	u.ID = id
	err := u.Save(mockDB)

	assert.NoError(t, err, "Save() should not have fail")
	assert.Equal(t, id, u.ID, "ID should not have changed")
}

func TestUserCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertSuccess(&User{})

	u := &User{}
	err := u.Create(mockDB)

	assert.NoError(t, err, "Create() should not have fail")
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.CreatedAt, "CreatedAt should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserCreateWithID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)

	u := &User{}
	u.ID = uuid.NewV4().String()

	err := u.Create(mockDB)
	assert.Error(t, err, "Create() should have fail")
}

func TestUserDoCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertSuccess(&User{})

	u := &User{}
	err := u.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.CreatedAt, "CreatedAt should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserDoCreateWithDate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertSuccess(&User{})

	createdAt := datetime.Now().AddDate(0, 0, 1)
	u := &User{CreatedAt: createdAt}
	err := u.doCreate(mockDB)

	assert.NoError(t, err, "doCreate() should not have fail")
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.True(t, u.CreatedAt.Equal(createdAt), "CreatedAt should not have been updated")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserDoCreateFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().InsertError(&User{}, errors.New("sql error"))

	u := &User{}
	err := u.doCreate(mockDB)

	assert.Error(t, err, "doCreate() should have fail")
}


func TestUserUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().UpdateSuccess(&User{})

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.Update(mockDB)

	assert.NoError(t, err, "Update() should not have fail")
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserUpdateWithoutID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	u := &User{}
	err := u.Update(mockDB)

	assert.Error(t, err, "Update() should not have fail")
}


func TestUserDoUpdate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().UpdateSuccess(&User{})

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.doUpdate(mockDB)

	assert.NoError(t, err, "doUpdate() should not have fail")
	assert.NotEmpty(t, u.ID, "ID should have been set")
	assert.NotNil(t, u.UpdatedAt, "UpdatedAt should have been set")
}

func TestUserDoUpdateWithoutID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	u := &User{}
	err := u.doUpdate(mockDB)

	assert.Error(t, err, "doUpdate() should not have fail")
}

func TestUserDoUpdateFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().UpdateError(&User{}, errors.New("sql error"))

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.doUpdate(mockDB)

	assert.Error(t, err, "doUpdate() should have fail")
}

func TestUserDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().DeletionSuccess()

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.Delete(mockDB)

	assert.NoError(t, err, "Delete() should not have fail")
}

func TestUserDeleteWithoutID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	u := &User{}
	err := u.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
}

func TestUserDeleteError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDB := mocksqldb.NewMockQueryable(mockCtrl)
	mockDB.EXPECT().DeletionError(errors.New("sql error"))

	u := &User{}
	u.ID = uuid.NewV4().String()
	err := u.Delete(mockDB)

	assert.Error(t, err, "Delete() should have fail")
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