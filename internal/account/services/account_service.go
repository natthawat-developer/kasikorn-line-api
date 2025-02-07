package services

import (
	"kasikorn-line-api/internal/account/models"
	"kasikorn-line-api/internal/account/repositories"
)

type AccountService interface {
	GetAccountDetails(req models.AccountRequest) (*models.AccountResponse, error)
}

type accountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(repo repositories.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) GetAccountDetails(req models.AccountRequest) (*models.AccountResponse, error) {
	repoAccount, errResponse := s.repo.GetAccountByID(req.AccountID)
	if errResponse != nil {
		return nil, errResponse
	}

	accountDetails, errResponse := s.repo.GetAccountByID(req.AccountID)
	if errResponse != nil {
		return nil, errResponse
	}

	accountResponse := &models.AccountResponse{
		AccountID: repoAccount.AccountID,
		Type:      *repoAccount.Type,
		Currency:  *repoAccount.Currency,
		Issuer:    *accountDetails.Issuer,
	}

	return accountResponse, nil
}
