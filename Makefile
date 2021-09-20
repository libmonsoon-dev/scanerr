APP_NAME = scanerr
LINTER_HOME = /tmp/go/lint/${APP_NAME}
# https://github.com/golangci/golangci-lint/releases
LINTER_VERSION = v1.42.1

GO = go

dependency:
	$(GO) mod tidy
	$(GO) mod vendor

generate:
	$(GO) generate -x ./...

git-add-generated:
	git add "**/*_string.go" "*_gen.go"

type-check:
	$(GO) build -v ./...

lint:
	@mkdir -p ${LINTER_HOME}
	@docker run --rm \
		-v $(shell pwd):/app \
		-v ${LINTER_HOME}:/root \
		-e GOLANGCI_LINT_CACHE=/root/lint/cache \
		-w /app \
	golangci/golangci-lint:${LINTER_VERSION} golangci-lint run -v

test:
	$(GO) test -race ./...

pre-commit: generate type-check git-add-generated lint test
