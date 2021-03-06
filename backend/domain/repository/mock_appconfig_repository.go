// Code generated by MockGen. DO NOT EDIT.
// Source: appconfig_repository.go

// Package repository is a generated GoMock package.
package repository

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAppConfigRepository is a mock of AppConfigRepository interface
type MockAppConfigRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAppConfigRepositoryMockRecorder
}

// MockAppConfigRepositoryMockRecorder is the mock recorder for MockAppConfigRepository
type MockAppConfigRepositoryMockRecorder struct {
	mock *MockAppConfigRepository
}

// NewMockAppConfigRepository creates a new mock instance
func NewMockAppConfigRepository(ctrl *gomock.Controller) *MockAppConfigRepository {
	mock := &MockAppConfigRepository{ctrl: ctrl}
	mock.recorder = &MockAppConfigRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAppConfigRepository) EXPECT() *MockAppConfigRepositoryMockRecorder {
	return m.recorder
}

// GetPaymentPeriod mocks base method
func (m *MockAppConfigRepository) GetPaymentPeriod() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentPeriod")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentPeriod indicates an expected call of GetPaymentPeriod
func (mr *MockAppConfigRepositoryMockRecorder) GetPaymentPeriod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentPeriod", reflect.TypeOf((*MockAppConfigRepository)(nil).GetPaymentPeriod))
}

// SetPaymentPeriod mocks base method
func (m *MockAppConfigRepository) SetPaymentPeriod(period int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPaymentPeriod", period)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPaymentPeriod indicates an expected call of SetPaymentPeriod
func (mr *MockAppConfigRepositoryMockRecorder) SetPaymentPeriod(period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPaymentPeriod", reflect.TypeOf((*MockAppConfigRepository)(nil).SetPaymentPeriod), period)
}

// GetCurrentPeriod mocks base method
func (m *MockAppConfigRepository) GetCurrentPeriod() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentPeriod")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentPeriod indicates an expected call of GetCurrentPeriod
func (mr *MockAppConfigRepositoryMockRecorder) GetCurrentPeriod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentPeriod", reflect.TypeOf((*MockAppConfigRepository)(nil).GetCurrentPeriod))
}

// SetCurrentPeriod mocks base method
func (m *MockAppConfigRepository) SetCurrentPeriod(period int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrentPeriod", period)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCurrentPeriod indicates an expected call of SetCurrentPeriod
func (mr *MockAppConfigRepositoryMockRecorder) SetCurrentPeriod(period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentPeriod", reflect.TypeOf((*MockAppConfigRepository)(nil).SetCurrentPeriod), period)
}
