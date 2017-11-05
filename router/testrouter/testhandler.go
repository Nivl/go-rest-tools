package testrouter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Nivl/go-rest-tools/router"
	"github.com/Nivl/go-rest-tools/router/mockrouter"
	"github.com/Nivl/go-rest-tools/types/apierror"
	mockdb "github.com/Nivl/go-sqldb/implementations/mocksqldb"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	// ErrDup contains the errcode of a unique constraint violation
	ErrDup = "23505"
)

var (
	// serverError represents a database connection error
	serverError = &pq.Error{
		Code:    "08006",
		Message: "error: connection failure",
		Detail:  "the connection to the database failed",
	}
)

func newConflictError(fieldName string) *pq.Error {
	return &pq.Error{
		Code:    ErrDup,
		Message: "error: duplicate field",
		Detail:  fmt.Sprintf("Key (%s)=(Google) already exists.", fieldName),
	}
}

// ConflictTestParams represents the params needed by ConflictTest
type ConflictTestParams struct {
	Handler           router.RouteHandler
	HandlerParams     interface{}
	StructConflicting string
	FieldConflicting  string
}

// ConflictInsertTest test an handler and expects a 409 on an insert statement
func ConflictInsertTest(t *testing.T, p *ConflictTestParams) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Mock the database & add expectations
	mockDB := mockdb.NewMockConnection(mockCtrl)
	mockDB.QEXPECT().InsertError(p.StructConflicting, newConflictError(p.FieldConflicting))

	// Mock the request & add expectations
	req := &mockrouter.HTTPRequest{}
	req.On("Params").Return(p.HandlerParams)

	// call the handler
	err := p.Handler(req, &router.Dependencies{DB: mockDB})

	// Assert everything
	assert.Error(t, err)
	req.AssertExpectations(t)

	apiError := apierror.Convert(err)
	assert.Equal(t, http.StatusConflict, apiError.HTTPStatus())
	assert.Equal(t, p.FieldConflicting, apiError.Field())
}

// ConflictUpdateTest test am handler and expects a 409 on an update statement
func ConflictUpdateTest(t *testing.T, p *ConflictTestParams) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Mock the database & add expectations
	mockDB := mockdb.NewMockConnection(mockCtrl)
	mockDB.QEXPECT().GetSuccess(p.StructConflicting, nil)
	mockDB.QEXPECT().UpdateError(p.StructConflicting, newConflictError(p.FieldConflicting))

	// Mock the request & add expectations
	req := new(mockrouter.HTTPRequest)
	req.On("Params").Return(p.HandlerParams)

	// call the handler
	err := p.Handler(req, &router.Dependencies{DB: mockDB})

	// Assert everything
	assert.Error(t, err)
	req.AssertExpectations(t)

	apiError := apierror.Convert(err)
	assert.Equal(t, http.StatusConflict, apiError.HTTPStatus())
	assert.Equal(t, p.FieldConflicting, apiError.Field())
}
