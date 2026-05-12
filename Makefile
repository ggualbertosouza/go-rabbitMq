MAIN_PATH = ./main.go
DOCS_PATH = ./docs
ASYNCAPI_FILE = asyncapi.yml

.PHONY: run docs clean-docs help

run:
	@go run $(MAIN_PATH)

docs: clean-docs
	@echo "Generating AsyncAPI docs..."

	@npx @asyncapi/cli@latest generate fromTemplate \
		$(ASYNCAPI_FILE) \
		@asyncapi/html-template \
		--output $(DOCS_PATH) \
		--force-write

	@echo "Docs generated successfully"

clean-docs:
	@echo "Removing old docs..."
	@rm -rf $(DOCS_PATH)

help:
	@echo ""
	@echo "Available commands:"
	@echo ""
	@echo "make run         - Run application locally"
	@echo "make docs        - Generate AsyncAPI documentation"
	@echo "make clean-docs  - Remove generated documentation"
	@echo ""