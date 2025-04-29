// pkg/errors/errors.go

package errors

import "errors"

// Códigos de erro para a API
const (
	UserNotFoundCode       = 1001
	InvalidUserInputCode   = 1002
	EmailAlreadyExistsCode = 1003
	FailedToCreateUserCode = 1004
	FailedToUpdateUserCode = 1005
	UnauthorizedErrorCode  = 2001
	ForbiddenErrorCode     = 2002
	DatabaseConnectionCode = 3001
	RecordNotFoundCode     = 3002
	// ... outros códigos de erro
)

// Variáveis de erro pré-definidas com mensagens (opcional, mas útil)
var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidUserInput   = errors.New("invalid user input")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrFailedToCreateUser = errors.New("failed to create user")
	ErrFailedToUpdateUser = errors.New("failed to update user")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")
	ErrDatabaseConnection = errors.New("database connection error")
	ErrRecordNotFound     = errors.New("record not found")
	// ... outras variáveis de erro
)

// Tipo para erros com código (se você quiser incluir o código diretamente no erro)
type CodedError struct {
	Code    int
	Message string
}

func (e CodedError) Error() string {
	return e.Message
}

// Função para criar um erro codificado (opcional)
func NewCodedError(code int, message string) error {
	return CodedError{Code: code, Message: message}
}
