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
	mockgen -source=application/cliente/get01/usecase/interfaces.go -destination=application/cliente/get01/usecase/mocks/mocks.go -package=mocks
	mockgen -source=application/cliente/get02/usecase/interfaces.go -destination=application/cliente/get02/usecase/mocks/mocks.go -package=mocks
	mockgen -source=application/cliente/post01/usecase/interfaces.go -destination=application/cliente/post01/usecase/mocks/mocks.go -package=mocks
	go mod tidy
test:
	@go test ./ ...
