package handler

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/ohpasha/rest-api/pkg/service"
	mock_service "github.com/ohpasha/rest-api/pkg/service/mocks"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                string
		headerName          string
		headerValue         string
		token               string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedResposeCode string
	}{
		{
			name:        "ok",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedResposeCode: "1",
		},
		{
			name:                "No header",
			headerName:          "",
			headerValue:         "Bearer token",
			token:               "token",
			mockBehavior:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedResposeCode: `{"message":"empty Authorization header"}`,
		},
		{
			name:                "Invalid header",
			headerName:          "Authorization",
			headerValue:         "Bearr token",
			token:               "token",
			mockBehavior:        func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode:  401,
			expectedResposeCode: `{"message":"wrong Authorization header"}`,
		},
		{
			name:        "Service Failure",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(0, errors.New("failed to parse token"))
			},
			expectedStatusCode:  401,
			expectedResposeCode: `{"message":"failed to parse token"}`,
		},
	}

	for _, testCase := range testTable {
		c := gomock.NewController(t)
		defer c.Finish()

		auth := mock_service.NewMockAuthorization(c)
		testCase.mockBehavior(auth, testCase.token)

		services := &service.Service{
			Authorization: auth,
		}

		handler := NewHandler(services)

		r := gin.New()

		r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
			id, _ := c.Get(userContext)
			c.String(200, fmt.Sprintf("%d", id.(int)))
		})

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set(testCase.headerName, testCase.headerValue)

		r.ServeHTTP(w, req)

		assert.Equal(t, w.Code, testCase.expectedStatusCode)
		assert.Equal(t, w.Body.String(), testCase.expectedResposeCode)
	}
}
