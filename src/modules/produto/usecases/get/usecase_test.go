package get_test

import (
	"errors"

	"time"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/domain/entities"
	mockRepo "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/usecases/get"
	"go.uber.org/mock/gomock"

	//"api-modelo-clean-arch/application/extrato/getdados/mock"
	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	produtoOK := &entities.Produto{
		ID:              1,
		Nome:            "Produto 1",
		Descricao:       "Produto 1",
		Preco:           10.0,
		QtdEstoque:      100,
		DataCriacao:     time.Now(),
		DataAtualizacao: time.Now(),
		FlAtivo:         "true",
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().FindById(gomock.Any()).Return(produtoOK, nil)
		//r.EXPECT().FindByIdTelefone(gomock.Any(), gomock.Any()).Return(telefonesOK, nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, l)
		//err := uc.Execute("", tarifasOK_UC)
		err, resp := uc.Execute(produtoOK.ID)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().FindById(gomock.Any()).Return(nil, errExpect)
		// Quando da erro não chama o QueryLoadDataTelefone
		//r.EXPECT().QueryLoadDataTelefone(gomock.Any(), gomock.Any()).Return(telefonesOK, errExpect)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, l)
		//err := uc.Execute("", tarifasOK_UC)
		err, resp := uc.Execute(produtoOK.ID)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

}
