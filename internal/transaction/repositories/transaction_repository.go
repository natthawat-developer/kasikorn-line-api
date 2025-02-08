package repositories

import (
	"errors"
	"kasikorn-line-api/internal/transaction/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionByUserID(userID string) ([]models.Transaction, *error.ErrorResponse)
	GetTransactionByTransactionID(transactionID string) (*models.Transaction, *error.ErrorResponse)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(DB *gorm.DB) TransactionRepository {
	return &transactionRepository{DB: DB}
}

func (r *transactionRepository) GetTransactionByUserID(userID string) ([]models.Transaction, *error.ErrorResponse) {
	var transactions []models.Transaction
	err := r.DB.Where(&models.Transaction{UserID: &userID}).Find(&transactions).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return transactions, nil
}

func (r *transactionRepository) GetTransactionByTransactionID(transactionID string) (*models.Transaction, *error.ErrorResponse) {
	var transaction models.Transaction
	err := r.DB.Where(&models.Transaction{TransactionID: transactionID}).First(&transaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &transaction, nil
}
