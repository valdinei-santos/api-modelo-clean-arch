// Code generated by MockGen. DO NOT EDIT.
// Source: src/modules/cliente/application/get-all/usecase/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=src/modules/cliente/application/get-all/usecase/interfaces.go -destination=src/modules/cliente/application/get-all/usecase/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	entities "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	dto "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
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
func (m *MockIUsecase) Execute(stamp string) (*dto.ResponseClientes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", stamp)
	ret0, _ := ret[0].(*dto.ResponseClientes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIUsecaseMockRecorder) Execute(stamp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIUsecase)(nil).Execute), stamp)
}

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockIRepository) FindAll(stamp string) (*[]entities.ClienteComTel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", stamp)
	ret0, _ := ret[0].(*[]entities.ClienteComTel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIRepositoryMockRecorder) FindAll(stamp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIRepository)(nil).FindAll), stamp)
}

// FindAllTelefone mocks base method.
func (m *MockIRepository) FindAllTelefone(stamp, cpf string) ([]entities.Telefone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllTelefone", stamp, cpf)
	ret0, _ := ret[0].([]entities.Telefone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTelefone indicates an expected call of FindAllTelefone.
func (mr *MockIRepositoryMockRecorder) FindAllTelefone(stamp, cpf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTelefone", reflect.TypeOf((*MockIRepository)(nil).FindAllTelefone), stamp, cpf)
}
