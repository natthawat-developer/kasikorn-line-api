// Code generated by MockGen. DO NOT EDIT.
// Source: internal/account/repositories/account_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	models "kasikorn-line-api/internal/account/repositories/models"
	error "kasikorn-line-api/pkg/error"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// GetAccountBalance mocks base method.
func (m *MockAccountRepository) GetAccountBalance(accountID string) (*models.AccountBalance, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountBalance", accountID)
	ret0, _ := ret[0].(*models.AccountBalance)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetAccountBalance indicates an expected call of GetAccountBalance.
func (mr *MockAccountRepositoryMockRecorder) GetAccountBalance(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountBalance", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountBalance), accountID)
}

// GetAccountByID mocks base method.
func (m *MockAccountRepository) GetAccountByID(accountID string) (*models.Account, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", accountID)
	ret0, _ := ret[0].(*models.Account)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockAccountRepositoryMockRecorder) GetAccountByID(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountByID), accountID)
}

// GetAccountByUserID mocks base method.
func (m *MockAccountRepository) GetAccountByUserID(userID string) ([]models.Account, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByUserID", userID)
	ret0, _ := ret[0].([]models.Account)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetAccountByUserID indicates an expected call of GetAccountByUserID.
func (mr *MockAccountRepositoryMockRecorder) GetAccountByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByUserID", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountByUserID), userID)
}

// GetAccountDetail mocks base method.
func (m *MockAccountRepository) GetAccountDetail(accountID string) (*models.AccountDetail, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountDetail", accountID)
	ret0, _ := ret[0].(*models.AccountDetail)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetAccountDetail indicates an expected call of GetAccountDetail.
func (mr *MockAccountRepositoryMockRecorder) GetAccountDetail(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountDetail", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountDetail), accountID)
}

// GetAccountFlags mocks base method.
func (m *MockAccountRepository) GetAccountFlags(accountID string) ([]models.AccountFlag, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountFlags", accountID)
	ret0, _ := ret[0].([]models.AccountFlag)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetAccountFlags indicates an expected call of GetAccountFlags.
func (mr *MockAccountRepositoryMockRecorder) GetAccountFlags(accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountFlags", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountFlags), accountID)
}

// GetMainAccountByUserID mocks base method.
func (m *MockAccountRepository) GetMainAccountByUserID(userID string) (*models.AccountDetail, *error.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainAccountByUserID", userID)
	ret0, _ := ret[0].(*models.AccountDetail)
	ret1, _ := ret[1].(*error.ErrorResponse)
	return ret0, ret1
}

// GetMainAccountByUserID indicates an expected call of GetMainAccountByUserID.
func (mr *MockAccountRepositoryMockRecorder) GetMainAccountByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainAccountByUserID", reflect.TypeOf((*MockAccountRepository)(nil).GetMainAccountByUserID), userID)
}
