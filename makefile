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
	mockgen -source=src/infra/logger/interfaces.go -destination=src/infra/logger/mocks/mocks.go -package=mocks

	mockgen -source=src/modules/cliente/infra/repository/interfaces.go -destination=src/modules/cliente/infra/repository/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/create/interfaces.go -destination=src/modules/cliente/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/create-com-telefone/interfaces.go -destination=src/modules/cliente/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/get/interfaces.go -destination=src/modules/cliente/usecases/get/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/get-all/interfaces.go -destination=src/modules/cliente/usecases/get-all/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/usecases/get-all/interfaces.go -destination=src/modules/cliente/usecases/get-all/mocks/mocks.go -package=mocks
	
	mockgen -source=src/modules/telefone/infra/repository/interfaces.go -destination=src/modules/telefone/infra/repository/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/usecases/create/interfaces.go -destination=src/modules/telefone/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/usecases/get-all/interfaces.go -destination=src/modules/telefone/usecases/get-all/mocks/mocks.go -package=mocks
	
	mockgen -source=src/modules/produto/infra/repository/interfaces.go -destination=src/modules/produto/infra/repository/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/produto/usecases/create/interfaces.go -destination=src/modules/produto/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/produto/usecases/get/interfaces.go -destination=src/modules/produto/usecases/get/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/produto/usecases/get-all/interfaces.go -destination=src/modules/produto/usecases/get-all/mocks/mocks.go -package=mocks
	go mod tidy
test:
	@go test ./ ...
