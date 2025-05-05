package getall_test

import (
	"errors"
	"time"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/domain/entities"
	mockRepo "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/usecases/get-all"
	"go.uber.org/mock/gomock"

	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func Test_Execute(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	produto1 := &entities.Produto{
		ID:              1,
		Nome:            "Produto 1",
		Descricao:       "Produto 1",
		Preco:           10.0,
		QtdEstoque:      100,
		DataCriacao:     time.Now(),
		DataAtualizacao: time.Now(),
		FlAtivo:         "true",
	}
	produto2 := &entities.Produto{
		ID:              2,
		Nome:            "Produto 2",
		Descricao:       "Produto 2",
		Preco:           9.0,
		QtdEstoque:      10,
		DataCriacao:     time.Now(),
		DataAtualizacao: time.Now(),
		FlAtivo:         "true",
	}

	produtosOK := make([]*entities.Produto, 2)
	produtosOK[0] = produto1
	produtosOK[1] = produto2

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().FindAll().Return(produtosOK, nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, l)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().FindAll().Return(nil, errExpect)
		// Quando da erro não chama o QueryLoadDataTelefone
		//r.EXPECT().QueryLoadDataTelefone(gomock.Any(), gomock.Any()).Return(telefonesOK, errExpect)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, l)
		//err := uc.Execute("", tarifasOK_UC)
		resp, err := uc.Execute()
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

}
