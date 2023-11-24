package handlers

import (
	"context"
	"job-port-api/internal/auth"
	"job-port-api/internal/middleware"
	"job-port-api/internal/models"
	"job-port-api/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func Test_handler_AddCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "missing jwt claims",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},
		{
			name: "Success",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"name":"Tek system","location":"mysore"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Ctxkey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := service.NewMockUserService(mc)
				ms.EXPECT().FetchAllCompanies().Return(models.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","location":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := handler{
				service: ms,
			}

			h.AddCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
