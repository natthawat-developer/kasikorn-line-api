package repositories

import (
	"errors"
	"kasikorn-line-api/internal/account/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type AccountRepository interface {
	GetAccountByUserID(userID string) ([]models.Account, *error.ErrorResponse)
	GetAccountByID(accountID string) (*models.Account, *error.ErrorResponse)
	GetAccountBalance(accountID string) (*models.AccountBalance, *error.ErrorResponse)
	GetAccountDetail(accountID string) (*models.AccountDetail, *error.ErrorResponse)
	GetAccountFlags(accountID string) ([]models.AccountFlag, *error.ErrorResponse)
}

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) AccountRepository {
	return &accountRepository{DB: DB}
}

func (r *accountRepository) GetAccountByUserID(userID string) ([]models.Account, *error.ErrorResponse) {
	var accounts []models.Account
	err := r.DB.Where(&models.Account{UserID: &userID}).Find(&accounts).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return accounts, nil
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

func (r *accountRepository) GetAccountBalance(accountID string) (*models.AccountBalance, *error.ErrorResponse) {
	var balance models.AccountBalance
	err := r.DB.Where(&models.AccountBalance{AccountID: accountID}).First(&balance).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &balance, nil
}

func (r *accountRepository) GetAccountDetail(accountID string) (*models.AccountDetail, *error.ErrorResponse) {
	var details models.AccountDetail
	err := r.DB.Where(&models.AccountDetail{AccountID: accountID}).First(&details).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &details, nil
}

func (r *accountRepository) GetAccountFlags(accountID string) ([]models.AccountFlag, *error.ErrorResponse) {
	var flags []models.AccountFlag
	err := r.DB.Where(&models.AccountFlag{AccountID: accountID}).Find(&flags).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return flags, nil

}
