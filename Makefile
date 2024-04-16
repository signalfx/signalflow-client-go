TOOLS = $(CURDIR)/.tools
GO = go

SHELL := /bin/bash
.DEFAULT_GOAL := precommit
TOOLS_MOD_DIR := ./internal/tools

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef

$(TOOLS):
	$(call print-target)
	mkdir -p $@

$(TOOLS)/%: $(TOOLS) $(TOOLS_MOD_DIR)/go.mod
	$(call print-target)
	cd $(TOOLS_MOD_DIR) && $(GO) build -o $@ $(PACKAGE)

GOLANGCI_LINT = $(TOOLS)/golangci-lint
$(TOOLS)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/cmd/golangci-lint

MISSPELL = $(TOOLS)/misspell
$(TOOLS)/misspell: PACKAGE=github.com/client9/misspell/cmd/misspell

.PHONY: precommit
precommit: ## build pipeline
precommit: spell mod gen lint test

.PHONY: ci
ci: ## CI build pipeline
ci: precommit diff

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	$(GO) mod tidy
	cd $(TOOLS_MOD_DIR) && $(GO) mod tidy

.PHONY: spell
spell: ## misspell
spell: $(MISSPELL)
	$(call print-target)
	$(MISSPELL) -error -locale=US -w **.md

.PHONY: lint
lint: ## golangci-lint
lint: $(GOLANGCI_LINT)
	$(call print-target)
	$(GOLANGCI_LINT) run --fix

.PHONY: gen
gen: ## go generate
	$(call print-target)
	$(GO) generate ./...

.PHONY: test
test: ## go test
	$(call print-target)
	$(GO) test -race ./...

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
