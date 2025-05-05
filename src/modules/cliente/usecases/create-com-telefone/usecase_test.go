package createcomtelefone_test

import (
	"errors"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	mockRepoCli "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/create-com-telefone"
	mockRepoTel "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository/mocks"

	//repoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository"

	//"api-modelo-clean-arch/application/extrato/getdados/mock"
	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Execute(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	repoCli := mockRepoCli.NewMockIRepository(control)
	repoTel := mockRepoTel.NewMockIRepository(control)
	log := mockLog.NewMockLogger(control)
	mockTx := mockRepoCli.NewMockITransaction(control)
	//mockTx := &mockTransaction.MockTransaction{}

	requestOK := &dto.RequestComTelefone{
		CPF:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
		Telefones: []string{
			"11999999999",
			"11888888888",
		},
	}

	requestNOK := &dto.RequestComTelefone{
		CPF:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
	}

	/* clienteOK := &dto.Cliente{
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
	} */

	t.Run("Caso de Sucesso", func(t *testing.T) {
		// repoCli := mockRepoCli.NewMockIRepository(control)
		// repoTel := mockRepoTel.NewMockIRepository(control)
		// l := mockLog.NewMockLogger(control)
		// mockTx := &mockTransaction.MockTransaction{}

		log.EXPECT().Debug(gomock.Any())
		repoCli.EXPECT().BeginTransaction().Return(mockTx, nil).Times(1)
		repoCli.EXPECT().Save(gomock.Any()).Return(nil).Times(1)
		repoTel.EXPECT().SaveAll(gomock.Any()).Return(nil).Times(1)
		mockTx.EXPECT().Commit().Return(nil).Times(1)
		//mockTx.EXPECT().Rollback().Return(nil).Times(1)

		uc := usecase.NewUseCase(repoCli, repoTel, log)
		//err := uc.Execute("", tarifasOK_UC)
		//resp, err := uc.Execute(clienteComTel)
		resp, err := uc.Execute(requestOK)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, 1, resp.StatusCode)
		assert.Equal(t, "Insert OK", resp.Message)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		log.EXPECT().Debug(gomock.Any())
		repoCli.EXPECT().BeginTransaction().Return(mockTx, nil).Times(1)
		repoCli.EXPECT().Save(gomock.Any()).Return(errExpect).Times(1)
		mockTx.EXPECT().Rollback().Times(1)
		log.EXPECT().Error(gomock.Any(), gomock.Any(), gomock.Any())

		uc := usecase.NewUseCase(repoCli, repoTel, log)
		resp, err := uc.Execute(requestNOK)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})

}
