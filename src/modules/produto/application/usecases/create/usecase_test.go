package create_test

import (
	"errors"
	"testing"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/application/usecases/create"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/produto/infra/repository/mocks"

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
		r := mocks.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(r)
		resp, err := uc.Execute("", produtoOK)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Produto inserido com sucesso", resp.Message)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mocks.NewMockIRepository(control)
		r.EXPECT().Save(gomock.Any(), gomock.Any()).Return(errExpect)

		uc := usecase.NewUseCase(r)
		resp, err := uc.Execute("", produtoOK)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}
