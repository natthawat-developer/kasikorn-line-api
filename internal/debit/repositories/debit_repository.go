package repositories

import (
	"errors"
	"kasikorn-line-api/internal/debit/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type DebitRepository interface {
	GetDebitCardByUserID(userID string) (*models.DebitCard, *error.ErrorResponse)
	GetDebitCardDesignByID(userID string) (*models.DebitCardDesign, *error.ErrorResponse)
	GetDebitCardDetailsByUserID(userID string) (*models.DebitCardDetails, *error.ErrorResponse)
	GetDebitCardStatusByUserID(userID string) (*models.DebitCardStatus, *error.ErrorResponse)
}

type debitRepository struct {
	DB *gorm.DB
}


func NewDebitRepository(DB *gorm.DB) DebitRepository {
	return &debitRepository{DB: DB}
}

func (r *debitRepository) GetDebitCardByUserID(userID string) (*models.DebitCard, *error.ErrorResponse) {
	var debitCard models.DebitCard
	err := r.DB.Where(&models.DebitCard{UserID: &userID}).Find(&debitCard).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &debitCard, nil
}

func (r *debitRepository) GetDebitCardDesignByID(userID string) (*models.DebitCardDesign, *error.ErrorResponse) {
	var debitCardDesign models.DebitCardDesign
	err := r.DB.Where(&models.DebitCardDesign{UserID: &userID}).First(&debitCardDesign).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return &debitCardDesign, nil
}

func (r *debitRepository) GetDebitCardDetailsByUserID(userID string) (*models.DebitCardDetails, *error.ErrorResponse) {
	var debitCardDetail models.DebitCardDetails
	err := r.DB.Where(&models.DebitCardDetails{UserID: &userID}).Find(&debitCardDetail).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &debitCardDetail, nil
}

func (r *debitRepository) GetDebitCardStatusByUserID(userID string) (*models.DebitCardStatus, *error.ErrorResponse) {
	var debitCardStatus models.DebitCardStatus
	err := r.DB.Where(&models.DebitCardStatus{UserID: &userID}).First(&debitCardStatus).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &debitCardStatus, nil
}