package services

import (
	"net/http"
	"testing"

	"kasikorn-line-api/internal/user/models"
	"kasikorn-line-api/internal/user/repositories/mock"
	repoModel "kasikorn-line-api/internal/user/repositories/models"
	coreError "kasikorn-line-api/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserDetails_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repositories.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	req := models.UserRequest{UserID: "123"}
	repoUser := &repoModel.User{UserID: "123", Name: "John Doe"}
	repoGreeting := &repoModel.UserGreeting{UserID: "123", Greeting: "Hello!"}

	mockRepo.EXPECT().GetUserByID("123").Return(repoUser, nil)
	mockRepo.EXPECT().GetUserGreetingByUserID("123").Return(repoGreeting, nil)

	resp, err := userService.GetUserDetails(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John Doe", resp.Name)
	assert.Equal(t, "Hello!", resp.Greeting)
}

func TestGetUserDetails_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repositories.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	req := models.UserRequest{UserID: "123"}

	mockRepo.EXPECT().GetUserByID("123").Return(nil, coreError.NewErrorResponse(http.StatusNotFound, "user not found"))

	resp, err := userService.GetUserDetails(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusNotFound, err.(*coreError.ErrorResponse).Code)
}

func TestGetUserDetails_GreetingNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repositories.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	req := models.UserRequest{UserID: "123"}
	repoUser := &repoModel.User{UserID: "123", Name: "John Doe"}

	mockRepo.EXPECT().GetUserByID("123").Return(repoUser, nil)
	mockRepo.EXPECT().GetUserGreetingByUserID("123").Return(nil, coreError.NewErrorResponse(http.StatusNotFound, "greeting not found"))

	resp, err := userService.GetUserDetails(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusNotFound, err.(*coreError.ErrorResponse).Code)
}

func TestGetUserDetails_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repositories.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	req := models.UserRequest{UserID: "123"}

	mockRepo.EXPECT().GetUserByID("123").Return(nil, coreError.NewErrorResponse(http.StatusInternalServerError, "database error"))

	resp, err := userService.GetUserDetails(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, err.(*coreError.ErrorResponse).Code)
}
