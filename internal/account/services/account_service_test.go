package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"kasikorn-line-api/internal/account/models"
	"kasikorn-line-api/internal/account/repositories/mock"
	repoModel "kasikorn-line-api/internal/account/repositories/models"
	coreError "kasikorn-line-api/pkg/error"
)

func TestGetAccountByUserID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetAccountByUserIDRequest{UserID: "user123"}
	repoAccounts := []repoModel.Account{
		{AccountID: "account1"},
		{AccountID: "account2"},
	}

	mockRepo.EXPECT().GetAccountByUserID("user123").Return(repoAccounts, nil)

	resp, err := accountService.GetAccountByUserID(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.AccountIDs, 2)
	assert.Equal(t, "account1", resp.AccountIDs[0])
	assert.Equal(t, "account2", resp.AccountIDs[1])
}

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func float64Ptr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}

func TestGetAccountDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetAccountDetailRequest{AccountID: "account1"}

	repoAccount := &repoModel.Account{
		AccountID:     "account1",
		Type:          strPtr("Saving"),
		Currency:      strPtr("USD"),
		AccountNumber: strPtr("123456"),
		Issuer:        strPtr("BankA"),
	}

	repoAccountBalance := &repoModel.AccountBalance{
		Amount: float64Ptr(1000.0),
	}

	repoAccountDetail := &repoModel.AccountDetail{
		Color:         strPtr("Blue"),
		IsMainAccount: boolPtr(true),
		Progress:      intPtr(80),
	}

	repoAccountFlags := []repoModel.AccountFlag{
		{FlagType: "Active", FlagValue: "true"},
	}

	mockRepo.EXPECT().GetAccountByID("account1").Return(repoAccount, nil)
	mockRepo.EXPECT().GetAccountBalance("account1").Return(repoAccountBalance, nil)
	mockRepo.EXPECT().GetAccountDetail("account1").Return(repoAccountDetail, nil)
	mockRepo.EXPECT().GetAccountFlags("account1").Return(repoAccountFlags, nil)

	resp, err := accountService.GetAccountDetail(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Saving", *resp.Type)
	assert.Equal(t, "USD", *resp.Currency)
	assert.Equal(t, "123456", *resp.AccountNumber)
	assert.Equal(t, "BankA", *resp.Issuer)
	assert.Equal(t, 1000.0, *resp.Amount)
	assert.Equal(t, "Blue", *resp.Color)
	assert.True(t, *resp.IsMainAccount)
	assert.Equal(t, 80, *resp.Progress)
	assert.Len(t, resp.Flags, 1)
	assert.Equal(t, "Active", resp.Flags[0].FlagType)
}

func TestGetAccountDetail_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetAccountDetailRequest{AccountID: "account1"}
	mockError := &coreError.ErrorResponse{Message: "database error"}

	mockRepo.EXPECT().GetAccountByID("account1").Return(nil, mockError)

	resp, err := accountService.GetAccountDetail(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetMainAccountByUserID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetMainAccountByUserIDRequest{UserID: "user123"}

	repoMainAccount := repoModel.AccountDetail{
		AccountID: "mainAccountID",
	}

	mockRepo.EXPECT().GetMainAccountByUserID("user123").Return(&repoMainAccount, nil)

	resp, err := accountService.GetMainAccountByUserID(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "mainAccountID", resp.AccountID)
}

func TestGetMainAccountByUserID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetMainAccountByUserIDRequest{UserID: "user123"}
	mockError := &coreError.ErrorResponse{Message: "database error"}

	mockRepo.EXPECT().GetMainAccountByUserID("user123").Return(nil, mockError)

	resp, err := accountService.GetMainAccountByUserID(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}
