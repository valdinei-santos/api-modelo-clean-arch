package getall_test

import (
	"errors"

	mockLog "github.com/valdinei-santos/api-modelo-clean-arch/src/infra/logger/mocks"
	usecase "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/application/usecases/get-all"
	"github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/domain/entities"
	mockRepo "github.com/valdinei-santos/api-modelo-clean-arch/src/modules/cliente/infra/repository/mocks"
	"go.uber.org/mock/gomock"

	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func Test_Execute(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

	cliente1 := &entities.Cliente{
		Cpf:    "1",
		Nome:   "Cliente 1",
		DtNasc: "02/07/1975",
	}
	cliente2 := &entities.Cliente{
		Cpf:    "2",
		Nome:   "Cliente 2",
		DtNasc: "02/07/1975",
	}

	clientesOK := make([]*entities.Cliente, 2)
	clientesOK[0] = cliente1
	clientesOK[1] = cliente2

	//telefonesOK := make([]entities.Telefone, 2)
	//telefonesOK[0] = entities.Telefone{Cpf: "1", Numero: "48999448383"}
	//telefonesOK[1] = entities.Telefone{Cpf: "2", Numero: "4832453548"}

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
		r := mockRepo.NewMockIRepository(control)
		r.EXPECT().FindAll().Return(clientesOK, nil)

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
