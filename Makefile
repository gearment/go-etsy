.PHONY: fmt generate test lint run help grow dockerize-dist mock test-server
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOFMT := "goimports"
ETSY_SPEC?=$$(ls etsy_*.json | sort -dr | head -n 1)
ETSY_FILES?=$$(find client -name '*.go')

fmt: ## Run gofmt for all .go files
	@$(GOFMT) -w $(GOFMT_FILES)

test: ## Run go test for whole project
	@go test -v ./...

lint: ## Run linter
	@golangci-lint run ./...

mock:
	clear
	@mockery

tidy:
	@go mod tidy

test-server:
	clear
	@goconvey

etsy-gen:
	@rm -rf client
	@openapi-generator-cli generate \
		-i $(ETSY_SPEC) \
		-g go \
		-o client \
		--type-mappings integer=int64,number=float64 \
		--additional-properties=disallowAdditionalPropertiesIfNotPresent=false,enumClassPrefix=true,generateInterfaces=true,generateMarshalJSON=true,generateUnmarshalJSON,hideGenerationTimestamp=false,isGoSubmodule=false,packageName=goEtsy,packageVersion=1.0.0,structPrefix=true,useDefaultValuesForRequiredVars=false,withGoMod=false,withXml=false \
		--git-repo-id=go-etsy/client \
		--git-user-id=gearment \
		--inline-schema-options RESOLVE_INLINE_ENUMS=true \
		--skip-validate-spec
	@for file in $(shell find client -name '*.go'); do \
		sed -i 's/auth\["api_key"\]/auth\["x-api-key"\]/g' $$file; \
		$(GOFMT) -w $$file; \
	done

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
