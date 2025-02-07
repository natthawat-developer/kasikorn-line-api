package repositories

import (
	"errors"
	"kasikorn-line-api/internal/debit/repositories/models"
	"kasikorn-line-api/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type DebitRepository interface {
	GetDebitCardsByUserID(userID string) ([]models.DebitCard, *error.ErrorResponse)
	GetDebitCardByCardID(CardID string) (*models.DebitCard, *error.ErrorResponse)
	GetDebitCardDesignByCardID(CardID string) (*models.DebitCardDesign, *error.ErrorResponse)
	GetDebitCardDetailsByCardID(CardID string) (*models.DebitCardDetails, *error.ErrorResponse)
	GetDebitCardStatusByCardID(CardID string) (*models.DebitCardStatus, *error.ErrorResponse)
}

type debitRepository struct {
	DB *gorm.DB
}


func NewDebitRepository(DB *gorm.DB) DebitRepository {
	return &debitRepository{DB: DB}
}

func (r *debitRepository) GetDebitCardsByUserID(userID string) ([]models.DebitCard, *error.ErrorResponse) {
	var debitCards []models.DebitCard
	err := r.DB.Where(&models.DebitCard{UserID: &userID}).Find(&debitCards).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return debitCards, nil
}


func (r *debitRepository) GetDebitCardByCardID(CardID string) (*models.DebitCard, *error.ErrorResponse) {
	var debitCard models.DebitCard
	err := r.DB.Where(&models.DebitCard{CardID: CardID}).First(&debitCard).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &debitCard, nil
}

func (r *debitRepository) GetDebitCardDesignByCardID(CardID string) (*models.DebitCardDesign, *error.ErrorResponse) {
	var debitCardDesign models.DebitCardDesign
	err := r.DB.Where(&models.DebitCardDesign{CardID: CardID}).First(&debitCardDesign).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}

	return &debitCardDesign, nil
}

func (r *debitRepository) GetDebitCardDetailsByCardID(CardID string) (*models.DebitCardDetails, *error.ErrorResponse) {
	var debitCardDetail models.DebitCardDetails
	err := r.DB.Where(&models.DebitCardDetails{CardID: CardID}).First(&debitCardDetail).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &debitCardDetail, nil
}

func (r *debitRepository) GetDebitCardStatusByCardID(CardID string) (*models.DebitCardStatus, *error.ErrorResponse) {
	var debitCardStatus models.DebitCardStatus
	err := r.DB.Where(&models.DebitCardStatus{CardID: CardID}).First(&debitCardStatus).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, error.NewErrorResponse(http.StatusNotFound, err.Error())
		}
		return nil, error.NewErrorResponse(http.StatusInternalServerError, err.Error())
	}
	return &debitCardStatus, nil
}