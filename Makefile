##########################################################################################
# Metadata
##########################################################################################

BIN_NAME={{app_name}}

##########################################################################################
# Commands
##########################################################################################

GO ?= go
GET ?= $(GO) get
GORUN ?= $(GO) run
GOTEST ?= go test
GOTOOL ?= $(GO) tool

##########################################################################################
# Containers
##########################################################################################

.PHONY: run/container
run/container:
	docker-compose up -d

##########################################################################################
# Setup
##########################################################################################

.PHONY: setup/local
setup/local:
	@$(GO) get -u github.com/rakyll/gotest &&\
	$(GO) get -u github.com/axw/gocov/gocov &&\
    $(GO) mod tidy

.PHONY: setup/cover
setup/cover:
	@rm -rf ${COVER_DIR}
	@mkdir -p ${COVER_DIR}


##########################################################################################
# Simplifiers
##########################################################################################

.PHONY: run/api
run/api: 
	@$(GORUN) main.go

.PHONY: deps-down
deps-down: deps/local/down

.PHONY: test-unit
test-unit: test/local/unit

.PHONY: test-integration
test-integration: test/local/live-integration