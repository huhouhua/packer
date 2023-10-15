.RHONY: all
all: tidy

include scripts/make-rules/common.mk
include scripts/make-rules/golang.mk
include scripts/make-rules/image.mk

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build

.PHONY: start_docker
start_docker:
	@$(DOCKER) compose up

.PHONY: down_docker
down_docker:
	@$(DOCKER)  compose down

.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: install_easyjson
install_easyjson:
	@$(GO) install github.com/mailru/easyjson/...@latest