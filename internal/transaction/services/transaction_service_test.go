package services

import (
	"net/http"
	"testing"

	"kasikorn-line-api/internal/transaction/models"
	"kasikorn-line-api/internal/transaction/repositories/mock"
	repoModel "kasikorn-line-api/internal/transaction/repositories/models"
	coreError "kasikorn-line-api/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetTransactionByUserID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTransactionRepository(ctrl)
	transactionService := NewTransactionService(mockRepo)

	req := models.GetTransactionByUserIDRequest{UserID: "user123"}
	repoTransactions := []repoModel.Transaction{
		{TransactionID: "txn_001"},
		{TransactionID: "txn_002"},
	}

	mockRepo.EXPECT().GetTransactionByUserID("user123").Return(repoTransactions, nil)

	resp, err := transactionService.GetTransactionByUserID(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.ElementsMatch(t, []string{"txn_001", "txn_002"}, resp.TransactionIDs)
}

func TestGetTransactionByUserID_NoTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTransactionRepository(ctrl)
	transactionService := NewTransactionService(mockRepo)

	req := models.GetTransactionByUserIDRequest{UserID: "user123"}

	mockRepo.EXPECT().GetTransactionByUserID("user123").Return([]repoModel.Transaction{}, nil)

	resp, err := transactionService.GetTransactionByUserID(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Empty(t, resp.TransactionIDs)
}

func TestGetTransactionByUserID_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTransactionRepository(ctrl)
	transactionService := NewTransactionService(mockRepo)

	req := models.GetTransactionByUserIDRequest{UserID: "user123"}

	mockRepo.EXPECT().GetTransactionByUserID("user123").Return(nil, coreError.NewErrorResponse(http.StatusInternalServerError, "database error"))

	resp, err := transactionService.GetTransactionByUserID(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, err.(*coreError.ErrorResponse).Code)
}

func TestGetTransactionDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTransactionRepository(ctrl)
	transactionService := NewTransactionService(mockRepo)

	req := models.GetTransactionDetailRequest{TransactionID: "txn_001"}
	repoTransaction := &repoModel.Transaction{
		Name:   strPtr("John Doe"),
		Image:  strPtr("image_url"),
		IsBank: boolPtr(true),
	}

	mockRepo.EXPECT().GetTransactionByTransactionID("txn_001").Return(repoTransaction, nil)

	resp, err := transactionService.GetTransactionDetail(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John Doe", *resp.Name)
	assert.Equal(t, "image_url", *resp.Image)
	assert.Equal(t, true, *resp.IsBank)
}

func TestGetTransactionDetail_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTransactionRepository(ctrl)
	transactionService := NewTransactionService(mockRepo)

	req := models.GetTransactionDetailRequest{TransactionID: "txn_001"}

	mockRepo.EXPECT().GetTransactionByTransactionID("txn_001").Return(nil, coreError.NewErrorResponse(http.StatusNotFound, "transaction not found"))

	resp, err := transactionService.GetTransactionDetail(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusNotFound, err.(*coreError.ErrorResponse).Code)
}

func TestGetTransactionDetail_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockTransactionRepository(ctrl)
	transactionService := NewTransactionService(mockRepo)

	req := models.GetTransactionDetailRequest{TransactionID: "txn_001"}

	mockRepo.EXPECT().GetTransactionByTransactionID("txn_001").Return(nil, coreError.NewErrorResponse(http.StatusInternalServerError, "database error"))

	resp, err := transactionService.GetTransactionDetail(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, err.(*coreError.ErrorResponse).Code)
}

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}
