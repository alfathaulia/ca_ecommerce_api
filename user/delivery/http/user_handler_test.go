package http_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alfathaulia/ca_ecommerce_api/domain"
	"github.com/alfathaulia/ca_ecommerce_api/domain/mocks"
	userHttp "github.com/alfathaulia/ca_ecommerce_api/user/delivery/http"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	var mockUser domain.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUcase := new(mocks.UserUsecase)
	mockListUser := make([]domain.User, 0)
	mockListUser = append(mockListUser, mockUser)
	num := 1
	cursor := "2"
	mockUcase.On("Fetch", mock.Anything, cursor, int64(num)).Return(mockListUser, "10", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users?num=1&cursor="+cursor, strings.NewReader("/"))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c := e.NewContext(req, w)
	c.SetPath("users?num")
	c.QueryParam("num")

	c.QueryParam("cursor")

	handler := userHttp.UserHandler{
		UUsecase: mockUcase,
	}

	err = handler.FetchUser(c)
	require.NoError(t, err)
	responseCursor := w.Header().Get("X-Cursor")
	assert.Equal(t, "10", responseCursor)
	assert.Equal(t, http.StatusOK, w.Code)
	mockUcase.AssertExpectations(t)

}
