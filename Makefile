GO = go

SHELL := /bin/bash
.DEFAULT_GOAL := precommit

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef

.PHONY: precommit
precommit: ## build pipeline
precommit: mod gen test

.PHONY: ci
ci: ## CI build pipeline
ci: precommit diff

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	$(GO) mod tidy

.PHONY: gen
gen: ## go generate
	$(call print-target)
	$(GO) generate ./...

.PHONY: test
test: ## go test
	$(call print-target)
	$(GO) test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...

.PHONY: diff
diff: ## git diff
	$(call print-target)
	if ! git diff --quiet; then \
	  echo; \
	  echo 'Working tree is not clean, did you forget to run "make precommit"?'; \
	  echo; \
	  git status; \
	  exit 1; \
	fi
