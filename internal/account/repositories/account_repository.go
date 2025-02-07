package repositories

import (
	"errors"
	"kasikorn-line-api/internal/account/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountByID(accountID string) (*models.Account, *error.ErrorResponse)
}

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) AccountRepository {
	return &accountRepository{DB: DB}
}

func (r *accountRepository) GetAccountByID(accountID string) (*models.Account, *error.ErrorResponse) {
	var account models.Account
	err := r.DB.Where(&models.Account{AccountID: accountID}).First(&account).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &account, nil
}