package create_test

import (
	"errors"
	"testing"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	mockRepo "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/usecases/create"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Execute(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	produtoOK := &dto.Request{
		Nome:       "Produto 1",
		Descricao:  "Descrição do Produto 1",
		Preco:      100.50,
		QtdEstoque: 10,
		Categoria:  "Categoria 1",
		FlAtivo:    "S",
	}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any()).Return(nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, l)
		resp, err := uc.Execute(produtoOK)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Produto inserido com sucesso", resp.Message)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any()).Return(errExpect)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r, l)
		resp, err := uc.Execute(produtoOK)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}
