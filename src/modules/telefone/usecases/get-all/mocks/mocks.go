// Code generated by MockGen. DO NOT EDIT.
// Source: src/modules/telefone/usecases/get-all/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=src/modules/telefone/usecases/get-all/interfaces.go -destination=src/modules/telefone/usecases/get-all/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dto "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockIUsecase is a mock of IUsecase interface.
type MockIUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUsecaseMockRecorder
}

// MockIUsecaseMockRecorder is the mock recorder for MockIUsecase.
type MockIUsecaseMockRecorder struct {
	mock *MockIUsecase
}

// NewMockIUsecase creates a new mock instance.
func NewMockIUsecase(ctrl *gomock.Controller) *MockIUsecase {
	mock := &MockIUsecase{ctrl: ctrl}
	mock.recorder = &MockIUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUsecase) EXPECT() *MockIUsecaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIUsecase) Execute(cpf string) (*dto.ResponseAll, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", cpf)
	ret0, _ := ret[0].(*dto.ResponseAll)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIUsecaseMockRecorder) Execute(cpf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIUsecase)(nil).Execute), cpf)
}
