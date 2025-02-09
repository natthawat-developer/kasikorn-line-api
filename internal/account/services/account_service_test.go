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

// TestGetAccountByUserID_Success tests successful retrieval of account by user ID.
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

// Helper function to create pointer to string
func strPtr(s string) *string {
	return &s
}

// Helper function to create pointer to bool
func boolPtr(b bool) *bool {
	return &b
}

// Helper function to create pointer to float64
func float64Ptr(f float64) *float64 {
	return &f
}

// Helper function to create pointer to int
func intPtr(i int) *int {
	return &i
}

// TestGetAccountDetail_Success tests successful retrieval of account details.
func TestGetAccountDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetAccountDetailRequest{AccountID: "account1"}

	// Create pointer to repoAccount
	repoAccount := &repoModel.Account{
		AccountID:    "account1",
		Type:         strPtr("Saving"),    // Pointer to string
		Currency:     strPtr("USD"),       // Pointer to string
		AccountNumber: strPtr("123456"),   // Pointer to string
		Issuer:       strPtr("BankA"),     // Pointer to string
	}

	repoAccountBalance := &repoModel.AccountBalance{
		Amount: float64Ptr(1000.0), // Pointer to float64
	}

	repoAccountDetail := &repoModel.AccountDetail{
		Color:         strPtr("Blue"),    // Pointer to string
		IsMainAccount: boolPtr(true),     // Pointer to bool
		Progress:      intPtr(80),        // Pointer to int
	}

	repoAccountFlags := []repoModel.AccountFlag{
		{FlagType: "Active", FlagValue: "true"},  // Regular strings, no pointers needed
	}

	// Mock repo method calls with pointers
	mockRepo.EXPECT().GetAccountByID("account1").Return(repoAccount, nil)
	mockRepo.EXPECT().GetAccountBalance("account1").Return(repoAccountBalance, nil)
	mockRepo.EXPECT().GetAccountDetail("account1").Return(repoAccountDetail, nil)
	mockRepo.EXPECT().GetAccountFlags("account1").Return(repoAccountFlags, nil)

	resp, err := accountService.GetAccountDetail(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Saving", *resp.Type)  // Dereference pointer for comparison
	assert.Equal(t, "USD", *resp.Currency)  // Dereference pointer for comparison
	assert.Equal(t, "123456", *resp.AccountNumber)  // Dereference pointer for comparison
	assert.Equal(t, "BankA", *resp.Issuer)  // Dereference pointer for comparison
	assert.Equal(t, 1000.0, *resp.Amount)  // Dereference pointer for comparison
	assert.Equal(t, "Blue", *resp.Color)  // Dereference pointer for comparison
	assert.True(t, *resp.IsMainAccount)  // Dereference pointer for comparison
	assert.Equal(t, 80, *resp.Progress)  // Dereference pointer for comparison
	assert.Len(t, resp.Flags, 1)
	assert.Equal(t, "Active", resp.Flags[0].FlagType)  // No dereferencing needed here
}


// TestGetAccountDetail_Error tests the error case for account details.
func TestGetAccountDetail_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetAccountDetailRequest{AccountID: "account1"}
	mockError := &coreError.ErrorResponse{Message: "database error"}

	// Simulate error for GetAccountByID method
	mockRepo.EXPECT().GetAccountByID("account1").Return(nil, mockError) // Return pointer to nil (no account)

	resp, err := accountService.GetAccountDetail(req)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, resp)
}


// TestGetMainAccountByUserID_Success tests successful retrieval of main account by user ID.
func TestGetMainAccountByUserID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetMainAccountByUserIDRequest{UserID: "user123"}

	// Change to AccountDetail if that's the expected return type
	repoMainAccount := repoModel.AccountDetail{
		AccountID: "mainAccountID",  // Assuming AccountDetail has AccountID field
	}

	// Update the mock to return *repoModel.AccountDetail
	mockRepo.EXPECT().GetMainAccountByUserID("user123").Return(&repoMainAccount, nil)

	resp, err := accountService.GetMainAccountByUserID(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "mainAccountID", resp.AccountID)
}


// TestGetMainAccountByUserID_Error tests error handling when retrieving the main account by user ID.
func TestGetMainAccountByUserID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockAccountRepository(ctrl)
	accountService := NewAccountService(mockRepo)

	req := models.GetMainAccountByUserIDRequest{UserID: "user123"}
	mockError := &coreError.ErrorResponse{Message: "database error"}

	// Simulate error for GetMainAccountByUserID method
	// Use AccountDetail here instead of Account if that's what the method expects
	mockRepo.EXPECT().GetMainAccountByUserID("user123").Return(nil, mockError)

	resp, err := accountService.GetMainAccountByUserID(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

