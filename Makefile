APP_NAME := thd-project
PKG := ./...

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: build
build:
	go build -o $(APP_NAME) cmd/server/main.go

.PHONY: test
test:
	go test -v -cover $(PKG)

.PHONY: fmt
fmt:
	go fmt $(PKG)

.PHONY: swagger
swagger:
	# TODO: Generate docs into docs/openapi.yaml
	@echo "Swagger generation not yet implemented 🚧"
