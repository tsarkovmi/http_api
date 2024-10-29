package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
		{
			name:                "Wrong input",
			inputBody:           `{"name":"Test"}`,
			inputWorker:         httpapi.Worker{},
			mockBehavior:        func(s *mock_service.MockCRUD, worker httpapi.Worker) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"name":"Test","age":25,"salary":15000.50,"occupation":"work"}`,
			inputWorker: httpapi.Worker{
				Name:       "Test",
				Age:        25,
				Salary:     15000.50,
				Occupation: "work",
			},
			mockBehavior: func(s *mock_service.MockCRUD, worker httpapi.Worker) {
				s.EXPECT().CreateWorker(worker).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"something went wrong"}`,
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

func TestHandler_GetWorkers(t *testing.T) {
	// Тип для поведения мока
	type mockBehavior func(s *mock_service.MockCRUD)

	// Таблица тестов
	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockCRUD) {
				s.EXPECT().GetAllWorkers().Return([]httpapi.Worker{
					{ID: 1, Name: "John Doe", Age: 30, Salary: 3000.0, Occupation: "Engineer"},
					{ID: 2, Name: "Jane Doe", Age: 28, Salary: 3500.0, Occupation: "Manager"},
				}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
				"JSON": "[{\"id\":1,\"name\":\"John Doe\",\"age\":30,\"salary\":3000,\"occupation\":\"Engineer\"},{\"id\":2,\"name\":\"Jane Doe\",\"age\":28,\"salary\":3500,\"occupation\":\"Manager\"}]",
				"XML": "<Worker><id>1</id><name>John Doe</name><age>30</age><salary>3000</salary><occupation>Engineer</occupation></Worker><Worker><id>2</id><name>Jane Doe</name><age>28</age><salary>3500</salary><occupation>Manager</occupation></Worker>",
				"TOML": "[[]]\nid = 1\nname = 'John Doe'\nage = 30\nsalary = 3000.0\noccupation = 'Engineer'\n\n[[]]\nid = 2\nname = 'Jane Doe'\nage = 28\nsalary = 3500.0\noccupation = 'Manager'\n"
			}`,
		},
		{
			name: "Worker not found",
			mockBehavior: func(s *mock_service.MockCRUD) {
				s.EXPECT().GetAllWorkers().Return([]httpapi.Worker{}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
			   "JSON": "[]",
			   "XML": "",
			   "TOML": "[]"
			}`,
		},
		{
			name: "Internal Server Error",
			mockBehavior: func(s *mock_service.MockCRUD) {
				s.EXPECT().GetAllWorkers().Return(nil, errors.New("internal server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"internal server error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockCRUD := mock_service.NewMockCRUD(c)
			testCase.mockBehavior(mockCRUD)

			services := &service.Service{CRUD: mockCRUD}
			handler := Newhandler(services)

			r := gin.New()
			r.GET("/workers", handler.GetWorkers)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/workers", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_GetWorkerByID(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCRUD, id int)

	testTable := []struct {
		name                 string
		idParam              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "OK",
			idParam: "1",
			mockBehavior: func(s *mock_service.MockCRUD, id int) {
				s.EXPECT().FindWorkerByID(id).Return(httpapi.Worker{
					ID:         1,
					Name:       "John Doe",
					Age:        30,
					Salary:     3000.0,
					Occupation: "Engineer",
				}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `{
				"JSON": "{\"id\":1,\"name\":\"John Doe\",\"age\":30,\"salary\":3000,\"occupation\":\"Engineer\"}",
				"XML": "<Worker><id>1</id><name>John Doe</name><age>30</age><salary>3000</salary><occupation>Engineer</occupation></Worker>",
				"TOML": "id = 1\nname = 'John Doe'\nage = 30\nsalary = 3000.0\noccupation = 'Engineer'\n"
			}`,
		},
		{
			name:    "Worker Not Found",
			idParam: "1",
			mockBehavior: func(s *mock_service.MockCRUD, id int) {
				s.EXPECT().FindWorkerByID(id).Return(httpapi.Worker{
					ID:         1,
					Name:       "John Doe",
					Age:        30,
					Salary:     3000.0,
					Occupation: "Engineer",
				}, errors.New("worker not found"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"worker not found"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			mockCRUD := mock_service.NewMockCRUD(c)
			testCase.mockBehavior(mockCRUD, 1)

			services := &service.Service{CRUD: mockCRUD}
			handler := Newhandler(services)

			r := gin.New()
			r.GET("/workers/:id", handler.GetWorkerByID)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/workers/"+testCase.idParam, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
