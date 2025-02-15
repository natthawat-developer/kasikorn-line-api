// Code generated by MockGen. DO NOT EDIT.
// Source: internal/debit/repositories/debit_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	models "kasikorn-line-api/internal/debit/repositories/models"
	error "kasikorn-line-api/pkg/error"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDebitRepository is a mock of DebitRepository interface.
type MockDebitRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDebitRepositoryMockRecorder
}

// MockDebitRepositoryMockRecorder is the mock recorder for MockDebitRepository.
type MockDebitRepositoryMockRecorder struct {
	mock *MockDebitRepository
}

// NewMockDebitRepository creates a new mock instance.
func NewMockDebitRepository(ctrl *gomock.Controller) *MockDebitRepository {
	mock := &MockDebitRepository{ctrl: ctrl}
	mock.recorder = &MockDebitRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDebitRepository) EXPECT() *MockDebitRepositoryMockRecorder {
	return m.recorder
}

// GetDebitCardByCardID mocks base method.
func (m *MockDebitRepository) GetDebitCardByCardID(CardID string) (*models.DebitCard, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebitCardByCardID", CardID)
	ret0, _ := ret[0].(*models.DebitCard)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetDebitCardByCardID indicates an expected call of GetDebitCardByCardID.
func (mr *MockDebitRepositoryMockRecorder) GetDebitCardByCardID(CardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebitCardByCardID", reflect.TypeOf((*MockDebitRepository)(nil).GetDebitCardByCardID), CardID)
}

// GetDebitCardDesignByCardID mocks base method.
func (m *MockDebitRepository) GetDebitCardDesignByCardID(CardID string) (*models.DebitCardDesign, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebitCardDesignByCardID", CardID)
	ret0, _ := ret[0].(*models.DebitCardDesign)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetDebitCardDesignByCardID indicates an expected call of GetDebitCardDesignByCardID.
func (mr *MockDebitRepositoryMockRecorder) GetDebitCardDesignByCardID(CardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebitCardDesignByCardID", reflect.TypeOf((*MockDebitRepository)(nil).GetDebitCardDesignByCardID), CardID)
}

// GetDebitCardDetailsByCardID mocks base method.
func (m *MockDebitRepository) GetDebitCardDetailsByCardID(CardID string) (*models.DebitCardDetails, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebitCardDetailsByCardID", CardID)
	ret0, _ := ret[0].(*models.DebitCardDetails)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetDebitCardDetailsByCardID indicates an expected call of GetDebitCardDetailsByCardID.
func (mr *MockDebitRepositoryMockRecorder) GetDebitCardDetailsByCardID(CardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebitCardDetailsByCardID", reflect.TypeOf((*MockDebitRepository)(nil).GetDebitCardDetailsByCardID), CardID)
}

// GetDebitCardStatusByCardID mocks base method.
func (m *MockDebitRepository) GetDebitCardStatusByCardID(CardID string) (*models.DebitCardStatus, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebitCardStatusByCardID", CardID)
	ret0, _ := ret[0].(*models.DebitCardStatus)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetDebitCardStatusByCardID indicates an expected call of GetDebitCardStatusByCardID.
func (mr *MockDebitRepositoryMockRecorder) GetDebitCardStatusByCardID(CardID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebitCardStatusByCardID", reflect.TypeOf((*MockDebitRepository)(nil).GetDebitCardStatusByCardID), CardID)
}

// GetDebitCardsByUserID mocks base method.
func (m *MockDebitRepository) GetDebitCardsByUserID(userID string) ([]models.DebitCard, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebitCardsByUserID", userID)
	ret0, _ := ret[0].([]models.DebitCard)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetDebitCardsByUserID indicates an expected call of GetDebitCardsByUserID.
func (mr *MockDebitRepositoryMockRecorder) GetDebitCardsByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebitCardsByUserID", reflect.TypeOf((*MockDebitRepository)(nil).GetDebitCardsByUserID), userID)
}
