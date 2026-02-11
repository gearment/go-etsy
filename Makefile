.PHONY: fmt generate test lint run help grow dockerize-dist mock test-server
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)
GOFMT := "goimports"
ETSY_SPEC?=$$(ls etsy_*.json | sort -dr | head -n 1)
ETSY_FILES?=$$(find client -name '*.go')
GO_POST_PROCESS_FILE=$$(which gofmt)
export GO_POST_PROCESS_FILE

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

preprocess-spec:
	@jq . $(ETSY_SPEC) > $(ETSY_SPEC).tmp
	@bash scripts/fix_includes.sh $(ETSY_SPEC).tmp $(ETSY_SPEC)
	@mv $(ETSY_SPEC).tmp $(ETSY_SPEC)
	@sed -i '' 's|"application/x-www-form-urlencoded"|"application/json"|g' $(ETSY_SPEC)
	@rm -f $(ETSY_SPEC).tmp

generate: preprocess-spec
	@rm -rf client
	@openapi-generator generate \
		-i $(ETSY_SPEC) \
		-g go \
		-t custom-templates/go \
		-o client \
		--type-mappings integer=int64,number=float64 \
		--additional-properties=disallowAdditionalPropertiesIfNotPresent=false,enumClassPrefix=true,generateInterfaces=true,generateMarshalJSON=true,generateUnmarshalJSON,hideGenerationTimestamp=false,isGoSubmodule=false,packageName=goEtsy,packageVersion=1.0.0,structPrefix=true,useDefaultValuesForRequiredVars=false,withGoMod=false,withXml=false \
		--git-repo-id=go-etsy/client \
		--git-user-id=gearment \
		--inline-schema-options RESOLVE_INLINE_ENUMS=true \
		--skip-validate-spec \
		--skip-overwrite \
		--ignore-file-override=custom-templates/go/.openapi-generator-ignore
	@for file in $(shell find client -name '*.go'); do \
		sed -i '' 's/auth\["api_key"\]/auth\["x-api-key"\]/g' $$file; \
		$(GOFMT) -w $$file; \
	done

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
