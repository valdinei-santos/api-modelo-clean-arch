package usecase_test

import (
	"time"

	"github.com/valdinei-santos/api-modelo-clean-arch/config"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/infra/repository"
	"github.com/valdinei-santos/api-modelo-clean-arch/domain/cliente/get01/usecase"

	//"api-modelo-clean-arch/domain/extrato/getdados/mock"
	"testing"

	_ "github.com/godror/godror"

	//"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/mock"
)

// Com Mock Testify

// PRESENTER
type PresenterMock struct {
	mock.Mock
}

func (p *PresenterMock) Show(stamp string, out *usecase.Response) error {
	args := p.Called(stamp, out)
	return args.Error(0)
}

func (p *PresenterMock) ShowError(stamp string, msgErro string) error {
	args := p.Called(stamp, msgErro)
	return args.Error(0)
}

// REPOSITORY
/* type RepositoryMock struct {
	mock.Mock
}

func (p *RepositoryMock) QueryDados(stamp string, nrMatricula int) (*entity.Dados, error) {
	args := p.Called(stamp, nrMatricula)
	return args.Get(0).(*entity.Dados), args.Error(1)
} */

func Test_Execute(t *testing.T) {
	dbTest := config.InitDBTest()
	defer dbTest.Close()
	repo := repository.NewRepoOracle(dbTest)
	stamp := time.Now().Format(("20060102150405"))
	cpf := "01234567890"
	nome := "João da Silva"
	dtNasc := "02/07/1975"

	// Com gomock
	/* controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockIUsecase(controller)
	service.EXPECT().
		Execute(stamp, nrMatriculaInt).
		Return(nil)

	err := service.Execute(stamp, nrMatriculaInt)
	assert := assert.New(t)
	assert.Nil(err)
	//assert.Equal(d.NrMatricula, result.NrMatricula)
	*/

	// Com Testify
	p := &usecase.Cliente{
		Nome:   nome,
		DtNasc: dtNasc,
		CPF:    cpf,
	}

	t1 := &usecase.Telefone{
		Numero: "48912345678",
	}
	t2 := &usecase.Telefone{
		Numero: "48987654321",
	}
	var listaTel []*usecase.Telefone
	listaTel = append(listaTel, t1)
	listaTel = append(listaTel, t2)

	res := usecase.Response{
		Cliente:   p,
		Telefones: listaTel,
	}
	// create an instance of our test object
	//thePresenterMock := PresenterMock{} // Cria o Mock
	thePresenterMock := new(PresenterMock) // Cria o Mock

	// setup expectations
	// Quando os dados que serão passados como parametros não podem ser previstos
	//thePresenterMock.On("Show", mock.Anything, mock.Anything).Return(nil) // Se Show(mock.Anything, mock.Anything) é chamado, então retorna nil

	// Quando os dados que serão passados como parametros podem ser fixados
	thePresenterMock.On("Show", stamp, &res).Return(nil)

	// g := greeter{&theDBMock, "sg"}
	// call the code we are testing
	u := usecase.UseCase{repo, thePresenterMock}
	u.Execute(stamp, cpf)

	// assert that the expectations were met
	thePresenterMock.AssertExpectations(t)

	//assert.Equal(t, "Message is: nil", u.Execute(stamp, nrMatriculaInt))
	//thePresenterMock.AssertExpectations(t)

}
