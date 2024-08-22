.PHONY: run-server
.DEFAULT_GOAL := help

KEY_FILE := credentials/develop/server.key
CRT_FILE := credentials/develop/server.crt

serve: export KEY_FILE := $(KEY_FILE)
serve: export CRT_FILE := $(CRT_FILE)
serve: ## Run server
	go run cmd/server/main.go

client: ## Run client
	go run cmd/client/main.go

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
