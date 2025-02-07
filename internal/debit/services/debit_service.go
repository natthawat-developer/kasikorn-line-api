package services

import (
	"kasikorn-line-api/internal/debit/models"
	"kasikorn-line-api/internal/debit/repositories"
)

type DebitService interface {
	GetDebitCardDetailsByUserID(req models.GetDebitCardDetailsByUserIDRequest) (*models.GetDebitCardDetailsByUserIDResponse, error)
}

type debitService struct {
	repo repositories.DebitRepository
}

func NewDebitService(repo repositories.DebitRepository) DebitService {
	return &debitService{repo: repo}
}

func (s *debitService) GetDebitCardDetailsByUserID(req models.GetDebitCardDetailsByUserIDRequest) (*models.GetDebitCardDetailsByUserIDResponse, error) {
	debitCard, errResponse := s.repo.GetDebitCardByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardDesign, errResponse := s.repo.GetDebitCardDesignByID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardDetails, errResponse := s.repo.GetDebitCardDetailsByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardStatus, errResponse := s.repo.GetDebitCardStatusByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardDetailsResponse := &models.GetDebitCardDetailsByUserIDResponse{
		Name:        debitCard.Name,
		Color:       debitCardDesign.Color,
		BorderColor: debitCardDesign.BorderColor,
		Issuer:      debitCardDetails.Issuer,
		Number:      debitCardDetails.Number,
		Status:      debitCardStatus.Status,
	}

	return debitCardDetailsResponse, nil
}