// Code generated by MockGen. DO NOT EDIT.
// Source: src/modules/cliente/usecases/get-all/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=src/modules/cliente/usecases/get-all/interfaces.go -destination=src/modules/cliente/usecases/get-all/mocks/mocks.go -package=mocks
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

// QueryLoadAllClientes mocks base method.
func (m *MockIRepository) QueryLoadAllClientes(stamp string) (*[]entities.ClienteComTel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLoadAllClientes", stamp)
	ret0, _ := ret[0].(*[]entities.ClienteComTel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLoadAllClientes indicates an expected call of QueryLoadAllClientes.
func (mr *MockIRepositoryMockRecorder) QueryLoadAllClientes(stamp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLoadAllClientes", reflect.TypeOf((*MockIRepository)(nil).QueryLoadAllClientes), stamp)
}

// QueryLoadDataTelefone mocks base method.
func (m *MockIRepository) QueryLoadDataTelefone(stamp, cpf string) ([]entities.Telefone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLoadDataTelefone", stamp, cpf)
	ret0, _ := ret[0].([]entities.Telefone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLoadDataTelefone indicates an expected call of QueryLoadDataTelefone.
func (mr *MockIRepositoryMockRecorder) QueryLoadDataTelefone(stamp, cpf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLoadDataTelefone", reflect.TypeOf((*MockIRepository)(nil).QueryLoadDataTelefone), stamp, cpf)
}
