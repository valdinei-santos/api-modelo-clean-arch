package createcomtelefone_test

import (
	"errors"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/create-com-telefone"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"
	dtoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	mockRepoTel "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository/mocks"

	//repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"

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

	clienteOK := &dto.Cliente{
		CPF:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
	}

	// Cria um slice de DTO Telefone
	var telefones []*dtoTelefone.Telefone
	tel := &dtoTelefone.Telefone{
		CPF:    "1",
		Numero: "48999448282",
	}
	telefones = append(telefones, tel)

	clienteComTel := &dto.RequestComTelefone{
		CPF:       clienteOK.CPF,
		Nome:      clienteOK.Nome,
		DtNasc:    clienteOK.DtNasc,
		Telefones: []string{"48999448282"},
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		repoCli := mocks.NewMockIRepository(control)
		repoCli.EXPECT().Save(gomock.Any()).Return(nil)

		repoTel := mockRepoTel.NewMockIRepository(control)
		repoTel.EXPECT().Save(gomock.Any()).Return(nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(repoCli, repoTel, l)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute(clienteComTel)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		repoCli := mocks.NewMockIRepository(control)
		repoCli.EXPECT().Save(gomock.Any()).Return(errExpect)

		repoTel := mockRepoTel.NewMockIRepository(control)
		repoTel.EXPECT().Save(gomock.Any()).Return(nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(repoCli, repoTel, l)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute(clienteComTel)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

}
