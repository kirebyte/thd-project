APP_NAME := thd-project
PKG := ./...

define feedback
  @if [ "$$LOG_LEVEL" = "DEBUG" ]; then \
    echo "[DEBUG] $(1)" >&2; \
  elif [ "$$LOG_LEVEL" = "INFO" ]; then \
    echo "[INFO] $(1)" >&2; \
  fi
endef

.PHONY: run
run:
	go run cmd/server/main.go
	$(call feedback,App started successfully)

.PHONY: build
build:
	go build -o bin/$(APP_NAME) cmd/server/main.go
	$(call feedback,App built to bin/$(APP_NAME))

.PHONY: test
test:
	go test -cover $(PKG)
	$(call feedback,Test coverage complete)

.PHONY: coverage
coverage:
	go test -coverprofile=coverage.out $(PKG)
	go tool cover -html=coverage.out -o docs/coverage.html
	$(call feedback,Coverage report generated at coverage.html)

# Lint: basic formatting + vetting (fast, safe)
.PHONY: lint
lint:
	go fmt $(PKG)
	go vet $(PKG)
	$(call feedback,Basic linting complete)

# Lint-strict: includes staticcheck for deep analysis (slow, optional)
.PHONY: lint-strict
lint-strict:
	go fmt $(PKG)
	go vet $(PKG)
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...
	$(call feedback,Code linted and validated strictly)

.PHONY: check
check: lint test
	$(call feedback,Codebase checked and validated)

.PHONY: swagger
swagger:
	# TODO: Generate doc from docs/openapi.yaml
	@echo "[WARN] Swagger generation not yet implemented 🚧" >&2

.PHONY: swagger-validate
swagger-validate:
	go run github.com/getkin/kin-openapi/cmd/validate@latest docs/openapi.yaml
	$(call feedback,Swagger spec is valid ✅)

.PHONY: db-init
db-init:
	go run cmd/dbinit/main.go
	$(call feedback,Checked and initialized the SQLite database)

