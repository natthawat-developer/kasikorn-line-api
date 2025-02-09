package services

import (
	"kasikorn-line-api/internal/account/models"
	"kasikorn-line-api/internal/account/repositories"
)

type AccountService interface {
	GetAccountByUserID(req models.GetAccountByUserIDRequest) (*models.GetAccountByUserIDResponse, error)
	GetAccountDetail(req models.GetAccountDetailRequest) (*models.GetAccountDetailResponse, error)
	GetMainAccountByUserID(req models.GetMainAccountByUserIDRequest) (*models.GetMainAccountByUserIDResponse, error)
}

type accountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(repo repositories.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (s *accountService) GetAccountByUserID(req models.GetAccountByUserIDRequest) (*models.GetAccountByUserIDResponse, error) {
	accounts, errResponse := s.repo.GetAccountByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	var accountIDs []string
	for _, account := range accounts {
		accountIDs = append(accountIDs, account.AccountID)
	}

	accountResponse := &models.GetAccountByUserIDResponse{
		AccountIDs: accountIDs,
	}

	return accountResponse, nil
}

func (s *accountService) GetAccountDetail(req models.GetAccountDetailRequest) (*models.GetAccountDetailResponse, error) {
	account, errResponse := s.repo.GetAccountByID(req.AccountID)
	if errResponse != nil {
		return nil, errResponse
	}

	accountBalance, errResponse := s.repo.GetAccountBalance(req.AccountID)
	if errResponse != nil {
		return nil, errResponse
	}

	accountDetail, errResponse := s.repo.GetAccountDetail(req.AccountID)
	if errResponse != nil {
		return nil, errResponse
	}

	accountFlags, errResponse := s.repo.GetAccountFlags(req.AccountID)
	if errResponse != nil {
		return nil, errResponse
	}

	var mappedFlags []models.AccountFlag
	for _, flag := range accountFlags {
		mappedFlags = append(mappedFlags, models.AccountFlag{
			FlagType:  flag.FlagType,
			FlagValue: flag.FlagValue,
			CreatedAt: flag.CreatedAt,
			UpdatedAt: flag.UpdatedAt,
		})
	}

	accountResponse := &models.GetAccountDetailResponse{
		Type:          account.Type,
		Currency:      account.Currency,
		AccountNumber: account.AccountNumber,
		Issuer:        account.Issuer,
		Amount:        accountBalance.Amount,
		Color:         accountDetail.Color,
		IsMainAccount: accountDetail.IsMainAccount,
		Progress:      accountDetail.Progress,
		Flags:         mappedFlags,
	}

	return accountResponse, nil
}

func (s *accountService) GetMainAccountByUserID(req models.GetMainAccountByUserIDRequest) (*models.GetMainAccountByUserIDResponse, error) {
	mainAccount, errResponse := s.repo.GetMainAccountByUserID(req.UserID)
	if errResponse != nil {
		return nil, errResponse
	}

	if mainAccount == nil {
		return nil, nil
	}

	accountResponse := &models.GetMainAccountByUserIDResponse{
		AccountID: mainAccount.AccountID,
	}
	return accountResponse, nil
}
