package create_test

import (
	"errors"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/application/usecases/create"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	mockRepo "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository/mocks"

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

	telefoneOK := &dto.Request{
		CPF:    "1",
		Numero: "48 999448383",
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any()).Return(nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := create.NewUseCase(r, l)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute(telefoneOK)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any()).Return(errExpect)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := create.NewUseCase(r, l)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute(telefoneOK)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

}
