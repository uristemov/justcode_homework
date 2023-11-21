package service

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/murat96k/kitaptar.kz/api"
	"homeworks/hw15/entity"
	"homeworks/hw15/handler"
	mock_service "homeworks/hw15/service/mock"

	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestHandler_createUser(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, user entity.User)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"
	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"firstname":"Test_user", "lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: entity.User{
				FirstName: "Test_user",
				LastName:  "Test_user",
				Password:  "password",
				Email:     "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, user entity.User) {
				s.EXPECT().CreateUser(gomock.Any(), &user).Return(userID, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"e79e360e-cb68-40a1-911e-a8a75068ef79"}`,
		},
		{
			name:                 "Wrong input (Missing firstname)",
			inputBody:            `{"lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser:            entity.User{},
			mockBehavior:         func(s *mock_service.MockService, user entity.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"firstname":"Test_user", "lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: entity.User{
				FirstName: "Test_user",
				LastName:  "Test_user",
				Password:  "password",
				Email:     "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, user entity.User) {
				s.EXPECT().CreateUser(gomock.Any(), &user).Return("", errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputUser)

			mockHandler := handler.New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/register", bytes.NewBufferString(testCase.inputBody))

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_updateUser(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, req api.UpdateUserRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            api.UpdateUserRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Update all data",
			inputBody: `{"firstname":"Test_user", "lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: api.UpdateUserRequest{
				FirstName: "Test_user",
				LastName:  "Test_user",
				Password:  "password",
				Email:     "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateUser(gomock.Any(), userID, &req).Return(nil)
			},

			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"User data updated!"}`,
		},
		{
			name:      "Missing firstname input",
			inputBody: `{"lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: api.UpdateUserRequest{
				LastName: "Test_user",
				Password: "password",
				Email:    "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateUser(gomock.Any(), userID, &req).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"User data updated!"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"firstname":"Test_user", "lastname":"Test_user", "email":"test_mockuser@gmail.com", "password":"password"}`,
			inputUser: api.UpdateUserRequest{
				FirstName: "Test_user",
				LastName:  "Test_user",
				Password:  "password",
				Email:     "test_mockuser@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateUser(gomock.Any(), userID, &req).Return(errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:      "Empty response field",
			inputBody: `{}`,
			inputUser: api.UpdateUserRequest{},
			mockBehavior: func(s *mock_service.MockService, req api.UpdateUserRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {

			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputUser)

			mockHandler := handler.New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", "/user/update", bytes.NewBufferString(testCase.inputBody))

			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}
