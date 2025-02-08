package services

import (
	"kasikorn-line-api/internal/transaction/models"
	"kasikorn-line-api/internal/transaction/repositories"
)

type TransactionService interface {
	GetTransactionDetail(req models.GetTransactionDetailRequest) (*models.GetTransactionDetailResponse, error)
	GetTransactionByUserID(req models.GetTransactionByUserIDRequest) (*models.GetTransactionByUserIDResponse, error)
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) GetTransactionByUserID(req models.GetTransactionByUserIDRequest) (*models.GetTransactionByUserIDResponse, error) {

	transactions, errResponse := s.repo.GetTransactionByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}


	var transactionIDs []string
	for _, transaction := range transactions {
		transactionIDs = append(transactionIDs, transaction.TransactionID)
	}

	accountResponse := &models.GetTransactionByUserIDResponse{
		TransactionIDs: transactionIDs,
	}

	return accountResponse, nil
}

func (s *transactionService) GetTransactionDetail(req models.GetTransactionDetailRequest) (*models.GetTransactionDetailResponse, error) {
	transaction, errResponse := s.repo.GetTransactionByTransactionID(req.TransactionID)
	if errResponse != nil {
		return nil, errResponse
	}

	transactionCardDetailsResponse := &models.GetTransactionDetailResponse{
		Name:   transaction.Name,
		Image:  transaction.Image,
		IsBank: transaction.IsBank,
	}

	return transactionCardDetailsResponse, nil
}
