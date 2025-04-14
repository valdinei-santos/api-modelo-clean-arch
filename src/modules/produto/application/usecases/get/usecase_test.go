package get_test

import (
	"errors"

	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"
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

	clienteOK := &entities.Cliente{
		Cpf:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
	}

	/* telefonesOK := make([]entities.Telefone, 2)
	telefonesOK[0] = entities.Telefone{Cpf: "1", Numero: "48999448383"}
	telefonesOK[1] = entities.Telefone{Cpf: "2", Numero: "4832453548"} */

	/* telefoneOK := &entities.Telefone{
		Cpf:    "1",
		Numero: "48999448383",
	}

	clienteTelOK := &usecase.Cliente{
		Nome:      "Cliente 1",
		DtNasc:    "02/07/1975",
		CPF:       "1",
		Telefones: []string{"48999448383"},
	} */

	t.Run("Caso de Sucesso", func(t *testing.T) {
		r := mocks.NewMockIRepository(control)
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(clienteOK, nil)
		//r.EXPECT().FindByIdTelefone(gomock.Any(), gomock.Any()).Return(telefonesOK, nil)

		uc := usecase.NewUseCase(r)
		//err := uc.Execute("", tarifasOK_UC)
		err, resp := uc.Execute("", clienteOK.Cpf)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Caso de Erro", func(t *testing.T) {
		errExpect := errors.New("dummy error")
		r := mocks.NewMockIRepository(control)
		r.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, errExpect)
		// Quando da erro não chama o QueryLoadDataTelefone
		//r.EXPECT().QueryLoadDataTelefone(gomock.Any(), gomock.Any()).Return(telefonesOK, errExpect)

		uc := usecase.NewUseCase(r)
		//err := uc.Execute("", tarifasOK_UC)
		err, resp := uc.Execute("", clienteOK.Cpf)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

}
