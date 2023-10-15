GO := go

.PHONY: go.build.%
go.build.%:
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(info ${COMMAND})
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(info ${PLATFORM})
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(info ${OS})
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	$(info ${ARCH})
	@echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)"
	@mkdir -p $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) -o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(ROOT_PACKAGE)/cmd/$(COMMAND)
