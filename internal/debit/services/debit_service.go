package services

import (
	"kasikorn-line-api/internal/debit/models"
	"kasikorn-line-api/internal/debit/repositories"
	"kasikorn-line-api/pkg/utils"
)

type DebitService interface {
	GetDebitCardDetailsByCardID(req models.GetDebitCardDetailsByCardIDRequest) (*models.GetDebitCardDetailsByCardIDResponse, error)
	GetDebitCardsByUserID(req models.GetDebitCardsByUserIDRequest) (*models.GetDebitCardsByUserIDResponse, error)
}

type debitService struct {
	repo repositories.DebitRepository
}

func NewDebitService(repo repositories.DebitRepository) DebitService {
	return &debitService{repo: repo}
}

func (s *debitService) GetDebitCardsByUserID(req models.GetDebitCardsByUserIDRequest) (*models.GetDebitCardsByUserIDResponse, error) {

	debitCards, errResponse := s.repo.GetDebitCardsByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	var cardIDs []string

	for _, card := range debitCards {
		cardIDs = append(cardIDs, card.CardID)
	}

	return &models.GetDebitCardsByUserIDResponse{
		CardIDs: cardIDs,
	}, nil
}

func (s *debitService) GetDebitCardDetailsByCardID(req models.GetDebitCardDetailsByCardIDRequest) (*models.GetDebitCardDetailsByCardIDResponse, error) {
	debitCard, errResponse := s.repo.GetDebitCardByCardID(req.CardID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardDesign, errResponse := s.repo.GetDebitCardDesignByCardID(req.CardID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardDetails, errResponse := s.repo.GetDebitCardDetailsByCardID(req.CardID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardStatus, errResponse := s.repo.GetDebitCardStatusByCardID(req.CardID)
	if errResponse != nil {
		return nil, errResponse
	}

	debitCardNumber := debitCardDetails.Number
	if req.MaskDebitCardNumber {
		maskedNumber := utils.MaskDebitCardNumber(debitCardNumber)
		debitCardNumber = &maskedNumber
	}

	

	debitCardDetailsResponse := &models.GetDebitCardDetailsByCardIDResponse{
		Name:        debitCard.Name,
		Color:       debitCardDesign.Color,
		BorderColor: debitCardDesign.BorderColor,
		Issuer:      debitCardDetails.Issuer,
		Number:      debitCardNumber,
		Status:      debitCardStatus.Status,
	}

	return debitCardDetailsResponse, nil
}
