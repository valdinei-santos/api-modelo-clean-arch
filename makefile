.PHONY: default run run-with-docs build test docs clean
# Variables
APP_NAME=api-modelo-clean-arch

# Tasks
default: run-with-docs

run:
	@go run api/main.go
run-with-docs:
	@cd api && swag init -g main.go -o docs
	@go run api/main.go
build:
	@go build -o $(APP_NAME) api/main.go
test:
	@go test ./ ...
docs:
	@cd api && swag init -g main.go -o docs
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs