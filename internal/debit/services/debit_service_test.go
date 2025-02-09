package services

import (
	"net/http"
	"testing"

	"kasikorn-line-api/internal/debit/models"
	"kasikorn-line-api/internal/debit/repositories/mock"
	repoModel "kasikorn-line-api/internal/debit/repositories/models"
	coreError "kasikorn-line-api/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetDebitCardsByUserID_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockDebitRepository(ctrl)
	debitService := NewDebitService(mockRepo)

	req := models.GetDebitCardsByUserIDRequest{UserID: "user123"}

	mockRepo.EXPECT().GetDebitCardsByUserID("user123").Return(nil, coreError.NewErrorResponse(http.StatusInternalServerError, "database error"))

	resp, err := debitService.GetDebitCardsByUserID(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusInternalServerError, err.(*coreError.ErrorResponse).Code)
}

func TestGetDebitCardDetailsByCardID_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockDebitRepository(ctrl)
	debitService := NewDebitService(mockRepo)

	req := models.GetDebitCardDetailsByCardIDRequest{CardID: "card_001", UnmaskDebitCardNumber: false}
	repoDebitCard := &repoModel.DebitCard{Name: strPtr("John Doe")}
	repoDebitCardDesign := &repoModel.DebitCardDesign{Color: strPtr("blue"), BorderColor: strPtr("red")}
	repoDebitCardDetails := &repoModel.DebitCardDetails{Issuer: strPtr("Visa"), Number: strPtr("6821 5668 7876 2379")}
	repoDebitCardStatus := &repoModel.DebitCardStatus{Status: strPtr("Active")}

	mockRepo.EXPECT().GetDebitCardByCardID("card_001").Return(repoDebitCard, nil)
	mockRepo.EXPECT().GetDebitCardDesignByCardID("card_001").Return(repoDebitCardDesign, nil)
	mockRepo.EXPECT().GetDebitCardDetailsByCardID("card_001").Return(repoDebitCardDetails, nil)
	mockRepo.EXPECT().GetDebitCardStatusByCardID("card_001").Return(repoDebitCardStatus, nil)

	resp, err := debitService.GetDebitCardDetailsByCardID(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "John Doe", *resp.Name)
	assert.Equal(t, "blue", *resp.Color)
	assert.Equal(t, "red", *resp.BorderColor)
	assert.Equal(t, "Visa", *resp.Issuer)
	assert.Equal(t, "6821 56•• •••• 2379", *resp.Number)
	assert.Equal(t, "Active", *resp.Status)
}

func TestGetDebitCardDetailsByCardID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockDebitRepository(ctrl)
	debitService := NewDebitService(mockRepo)

	req := models.GetDebitCardDetailsByCardIDRequest{CardID: "card_001"}

	mockRepo.EXPECT().GetDebitCardByCardID("card_001").Return(nil, coreError.NewErrorResponse(http.StatusNotFound, "debit card not found"))

	resp, err := debitService.GetDebitCardDetailsByCardID(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, http.StatusNotFound, err.(*coreError.ErrorResponse).Code)
}

// Helper function to create pointer
func strPtr(s string) *string {
	return &s
}
