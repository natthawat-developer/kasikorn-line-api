package repositories

import (
	"errors"
	"kasikorn-line-api/internal/user/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(userID string) (*models.User, *error.ErrorResponse)
	GetUserGreetingByUserID(userID string) (*models.UserGreeting, *error.ErrorResponse)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}

func (r *userRepository) GetUserByID(userID string) (*models.User, *error.ErrorResponse) {
	var user models.User
	err := r.DB.Where(&models.User{UserID: userID}).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, "User not found")
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, "Database error: "+err.Error())
	}
	return &user, nil
}

func (r *userRepository) GetUserGreetingByUserID(userID string) (*models.UserGreeting, *error.ErrorResponse) {
	var userGreeting models.UserGreeting
	err := r.DB.Where(&models.UserGreeting{UserID: userID}).First(&userGreeting).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, "User greeting not found")
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, "Database error: "+err.Error())
	}
	return &userGreeting, nil

}
