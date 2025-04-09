#.PHONY: default run build mock test clean
# Variables
APP_NAME=api-modelo-clean-arch

# Tasks
default: run

run:
	@go run api/server.go
build:
	@go build -o $(APP_NAME) api/server.go
mock:
	mockgen -source=src/modules/cliente/get01/usecase/interfaces.go -destination=src/modules/cliente/get01/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/get02/usecase/interfaces.go -destination=src/modules/cliente/get02/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/post01/usecase/interfaces.go -destination=src/modules/cliente/post01/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente02/get-cliente/usecase/interfaces.go -destination=src/modules/cliente02/get-cliente/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente02/get-clientes/usecase/interfaces.go -destination=src/modules/cliente02/get-clientes/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente02/create-cliente/usecase/interfaces.go -destination=src/modules/cliente02/create-cliente/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/get/interfaces.go -destination=src/modules/cliente/usecases/get/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/get-all/interfaces.go -destination=src/modules/cliente/usecases/get-all/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/create/interfaces.go -destination=src/modules/cliente/usecases/create//mocks/mocks.go -package=mocks
	go mod tidy
test:
	@go test ./ ...
