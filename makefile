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
	mockgen -source=src/modules/cliente/infra/repository/interfaces.go -destination=src/modules/cliente/infra/repository/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/application/usecases/create/interfaces.go -destination=src/modules/cliente/application/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/application/usecases/get/interfaces.go -destination=src/modules/cliente/application/usecases/get/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/cliente/application/usecases/get-all/interfaces.go -destination=src/modules/cliente/application/usecases/get-all/mocks/mocks.go -package=mocks
	
	mockgen -source=src/modules/telefone/infra/repository/interfaces.go -destination=src/modules/telefone/infra/repository/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/application/usecases/create/interfaces.go -destination=src/modules/telefone/application/usecases/create/mocks/mocks.go -package=mocks
	mockgen -source=src/modules/telefone/application/usecases/get-all/interfaces.go -destination=src/modules/telefone/application/usecases/get-all/mocks/mocks.go -package=mocks
	go mod tidy
test:
	@go test ./ ...
