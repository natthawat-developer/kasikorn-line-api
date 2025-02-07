package services

import (
	"kasikorn-line-api/internal/user/models"
	"kasikorn-line-api/internal/user/repositories"
)

type UserService interface {
	GetUserDetails(req models.UserRequest) (*models.UserResponse, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUserDetails(req models.UserRequest) (*models.UserResponse, error) {
	repoUser, errResponse := s.repo.GetUserByID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	userGreeting, errResponse := s.repo.GetUserGreetingByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	userResponse := &models.UserResponse{
		UserID:   repoUser.UserID,
		Name:     repoUser.Name,
		Greeting: userGreeting.Greeting,
	}

	return userResponse, nil
}
