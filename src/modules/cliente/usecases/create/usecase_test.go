package create_test

import (
	"errors"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	mockRepo "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/create"

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

	repo := mockRepo.NewMockIRepository(control)
	log := mockLog.NewMockLogger(control)

	clienteOK := &dto.Request{
		CPF:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		repo.EXPECT().Save(gomock.Any()).Return(nil)
		log.EXPECT().Debug(gomock.Any(), gomock.Any())

		uc := usecase.NewUseCase(repo, log)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute(clienteOK)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		log.EXPECT().Debug(gomock.Any(), gomock.Any())
		repo.EXPECT().Save(gomock.Any()).Return(errExpect)
		log.EXPECT().Error(gomock.Any(), gomock.Any())

		uc := usecase.NewUseCase(repo, log)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute(clienteOK)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

}
