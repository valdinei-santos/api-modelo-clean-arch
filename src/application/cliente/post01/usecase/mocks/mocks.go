// Code generated by MockGen. DO NOT EDIT.
// Source: application/cliente/post01/usecase/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=application/cliente/post01/usecase/interfaces.go -destination=application/cliente/post01/usecase/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/application/cliente/post01/usecase"
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
func (m *MockIUsecase) Execute(stamp string, p *usecase.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", stamp, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockIUsecaseMockRecorder) Execute(stamp, p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIUsecase)(nil).Execute), stamp, p)
}

// MockIPresenter is a mock of IPresenter interface.
type MockIPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockIPresenterMockRecorder
}

// MockIPresenterMockRecorder is the mock recorder for MockIPresenter.
type MockIPresenterMockRecorder struct {
	mock *MockIPresenter
}

// NewMockIPresenter creates a new mock instance.
func NewMockIPresenter(ctrl *gomock.Controller) *MockIPresenter {
	mock := &MockIPresenter{ctrl: ctrl}
	mock.recorder = &MockIPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPresenter) EXPECT() *MockIPresenterMockRecorder {
	return m.recorder
}

// Show mocks base method.
func (m *MockIPresenter) Show(stamp string, t *usecase.Response) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", stamp, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// Show indicates an expected call of Show.
func (mr *MockIPresenterMockRecorder) Show(stamp, t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockIPresenter)(nil).Show), stamp, t)
}

// ShowError mocks base method.
func (m *MockIPresenter) ShowError(stamp, msgErro string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowError", stamp, msgErro)
	ret0, _ := ret[0].(error)
	return ret0
}

// ShowError indicates an expected call of ShowError.
func (mr *MockIPresenterMockRecorder) ShowError(stamp, msgErro any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowError", reflect.TypeOf((*MockIPresenter)(nil).ShowError), stamp, msgErro)
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

// InsertCliente mocks base method.
func (m *MockIRepository) InsertCliente(stamp string, p *usecase.Cliente) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCliente", stamp, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCliente indicates an expected call of InsertCliente.
func (mr *MockIRepositoryMockRecorder) InsertCliente(stamp, p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCliente", reflect.TypeOf((*MockIRepository)(nil).InsertCliente), stamp, p)
}

// InsertTelefone mocks base method.
func (m *MockIRepository) InsertTelefone(stamp string, t *usecase.Telefone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTelefone", stamp, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertTelefone indicates an expected call of InsertTelefone.
func (mr *MockIRepositoryMockRecorder) InsertTelefone(stamp, t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTelefone", reflect.TypeOf((*MockIRepository)(nil).InsertTelefone), stamp, t)
}
