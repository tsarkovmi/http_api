package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	httpapi "github.com/tsarkovmi/http_api"
	"github.com/tsarkovmi/http_api/pkg/service"
	mock_service "github.com/tsarkovmi/http_api/pkg/service/mocks"
	"go.uber.org/mock/gomock"
)

func TestHandler_PostWorkers(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCRUD, worker httpapi.Worker)

	testTable := []struct {
		name                string
		inputBody           string
		inputWorker         httpapi.Worker
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","age":25,"salary":15000.50,"occupation":"work"}`,
			inputWorker: httpapi.Worker{
				Name:       "Test",
				Age:        25,
				Salary:     15000.50,
				Occupation: "work",
			},
			mockBehavior: func(s *mock_service.MockCRUD, worker httpapi.Worker) {
				s.EXPECT().CreateWorker(worker).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			post := mock_service.NewMockCRUD(c)
			testCase.mockBehavior(post, testCase.inputWorker)

			services := &service.Service{CRUD: post}
			handler := Newhandler(services)

			// Test

			r := gin.New()
			r.POST("/workers", handler.PostWorkers)

			// Test Request
			// с помощью либы httptest

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/workers",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// пакет assert, проверка того, что полученное совпадает с ожидаемым
			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
