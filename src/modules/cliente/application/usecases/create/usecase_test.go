package create_test

import (
	"errors"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"

	//"api-modelo-clean-arch/application/extrato/getdados/mock"
	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// Com Mock Testify

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

	clienteOK := &dto.Request{
		CPF:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mocks.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute("", clienteOK)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mocks.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errExpect)

		uc := usecase.NewUseCase(r)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute("", clienteOK)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

}
