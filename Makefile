DOCS_PATH = ./docs
ASYNCAPI_FILE = asyncapi.yml

SERVER_PATH = ./cmd/server/main.go
CONSUMERS_PATH = ./cmd/consumers/main.go

.PHONY: run-server run-consumers docs clean-docs infra-up infra-down infra-logs help

run-server:
	@go run $(SERVER_PATH)

run-consumers:
	@go run $(CONSUMERS_PATH)

docs: clean-docs
	@echo "Removing old docs..."
	@rm -rf $(DOCS_PATH)

	@echo "Generating AsyncAPI docs..."
	@npx @asyncapi/cli@latest generate fromTemplate \
		$(ASYNCAPI_FILE) \
		@asyncapi/html-template \
		--output $(DOCS_PATH) \
		--force-write

	@echo "Docs generated successfully"

infra-up:
	@echo "Starting infrastructure..."
	@docker compose up --build -d

infra-down:
	@echo "Stopping infrastructure..."
	@docker compose down

infra-logs:
	@docker compose logs -f

help:
	@echo ""
	@echo "Available commands:"
	@echo ""
	@echo "make run          - Run application locally"
	@echo "make docs         - Generate AsyncAPI documentation"
	@echo "make infra-up     - Start RabbitMQ infrastructure"
	@echo "make infra-down   - Stop infrastructure"
	@echo "make infra-logs   - Follow infrastructure logs"
	@echo ""