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
	mockgen -source=src/modules/cliente/application/get/usecase/interfaces.go -destination=src/modules/cliente/application/get/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/application/get-all/usecase/interfaces.go -destination=src/modules/cliente/application/get-all/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/application/create/usecase/interfaces.go -destination=src/modules/cliente/application/create/usecase/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/infra/repository/interfaces.go -destination=src/modules/telefone/infra/repository/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/application/usecases/create/interfaces.go -destination=src/modules/telefone/application/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/application/usecases/get-all/interfaces.go -destination=src/modules/telefone/application/usecases/get-all/mocks/mocks.go -package=mocks
	go mod tidy
test:
	@go test ./ ...
