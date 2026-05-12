MAIN_PATH = './main.go'

run:
	go run $(MAIN_PATH)

generate-docs:
	asyncapi generate fromTemplate asyncapi.yml @asyncapi/html-template --output docs

help:
	@echo "Generate AsyncApi docs. Dependencies needs: @asyncapi/cli @asyncapi/html-template"