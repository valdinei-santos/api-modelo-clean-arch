package getallcomtelefone_test

import (
	"errors"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/dto"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/usecases/get-com-telefone"
	dtoTelefone "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/dto"
	mockRepoTel "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/telefone/infra/repository/mocks"
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

	telefonesOK := make([]dto.Telefone, 2)
	telefonesOK[0] = dto.Telefone{CPF: "1", Numero: "48999448383"}
	telefonesOK[1] = dto.Telefone{CPF: "2", Numero: "4832453548"}

	t.Run("Caso de Sucesso", func(t *testing.T) {
		repoCli := mocks.NewMockIRepository(control)
		repoCli.EXPECT().FindById(gomock.Any()).Return(clienteOK, nil)

		repoTel := mockRepoTel.NewMockIRepository(control)
		repoTel.EXPECT().FindAll(gomock.Any()).Return(telefonesOK, nil)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(repoCli, repoTel, l)
		//err := uc.Execute("", tarifasOK_UC)
		err, resp := uc.Execute(clienteOK.CPF)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")

		repoCli := mocks.NewMockIRepository(control)
		repoCli.EXPECT().FindById(gomock.Any()).Return(nil, errExpect)

		// Quando da erro não chama o FindAll do Telefone
		repoTel := mockRepoTel.NewMockIRepository(control)
		//repoTel.EXPECT().FindAll(gomock.Any(), gomock.Any()).Return(nil, errExpect)

		l := mockLog.NewMockLogger(control)
		l.EXPECT().Info(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		uc := usecase.NewUseCase(repoCli, repoTel, l)
		//err := uc.Execute("", tarifasOK_UC)
		err, resp := uc.Execute(clienteOK.CPF)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

}
