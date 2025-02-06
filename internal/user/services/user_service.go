package services

import (
	"net/http"

	"kasikorn-line-api/internal/user/models"
	"kasikorn-line-api/internal/user/repositories"

	coreError "kasikorn-line-api/pkg/error"

	"github.com/jinzhu/copier"
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

	var userResponse models.UserResponse
	err := copier.Copy(&userResponse, repoUser)
	if err != nil {
		return nil, coreError.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return &userResponse, nil
}
