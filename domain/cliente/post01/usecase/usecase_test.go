package usecase_test

import (
	"errors"

	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/post01/usecase/mocks"

	//"api-modelo-clean-arch/domain/extrato/getdados/mock"
	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/mock/gomock"
)

// Com Mock Testify

// PRESENTER
type PresenterMock struct {
	mock.Mock
}

func (p *PresenterMock) Show(stamp string, out *usecase.Response) error {
	args := p.Called(stamp, out)
	return args.Error(0)
}

func (p *PresenterMock) ShowError(stamp string, msgErro string) error {
	args := p.Called(stamp, msgErro)
	return args.Error(0)
}

// REPOSITORY
/* type RepositoryMock struct {
	mock.Mock
}

func (p *RepositoryMock) QueryDados(stamp string, nrMatricula int) (*entities.Dados, error) {
	args := p.Called(stamp, nrMatricula)
	return args.Get(0).(*entities.Dados), args.Error(1)
} */

func Test_Execute(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	clienteOK := &usecase.Request{
		CPF:       "1",
		Nome:      "Cliente 1",
		DtNasc:    "02/07/1975",
		Telefones: []string{"4832453548", "48999884455"},
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mocks.NewMockIRepository(control)
		r.EXPECT().InsertCliente(gomock.Any(), gomock.Any()).Return(nil)
		//v := mocks.NewMockIView(control)
		//v.EXPECT().Show("", "").Return(nil)
		p := mocks.NewMockIPresenter(control)
		p.EXPECT().Show(gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, p)
		//err := uc.Execute("", tarifasOK_UC)
		err := uc.Execute("", clienteOK)
		assert.Nil(t, err)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mocks.NewMockIRepository(control)
		r.EXPECT().InsertCliente(gomock.Any(), gomock.Any()).Return(errExpect)
		//v := mocks.NewMockIView(control)
		//v.EXPECT().Show("", "").Return(nil)
		p := mocks.NewMockIPresenter(control)
		p.EXPECT().ShowError(gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, p)
		//err := uc.Execute("", tarifasOK_UC)
		err := uc.Execute("", clienteOK)
		assert.NotNil(t, err)
	})

}
