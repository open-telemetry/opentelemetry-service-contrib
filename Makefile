include ./Makefile.Common

RUN_CONFIG?=local/config.yaml
CMD?=
OTEL_VERSION=main

BUILD_INFO_IMPORT_PATH=github.com/open-telemetry/opentelemetry-collector-contrib/internal/version
VERSION=$(shell git describe --always --match "v[0-9]*" HEAD)
BUILD_INFO=-ldflags "-X $(BUILD_INFO_IMPORT_PATH).Version=$(VERSION)"

COMP_REL_PATH=internal/components/components.go
MOD_NAME=github.com/open-telemetry/opentelemetry-collector-contrib

# ALL_MODULES includes ./* dirs (excludes . dir and example with go code)
ALL_MODULES := $(shell find . -type f -name "go.mod" -not -path './internal/core/*' -exec dirname {} \; | sort | egrep '^./' )
# Modules to run integration tests on.
# XXX: Find a way to automatically populate this. Too slow to run across all modules when there are just a few.
INTEGRATION_TEST_MODULES := \
	receiver/dockerstatsreceiver \
	receiver/jmxreceiver/ \
	receiver/redisreceiver \
	receiver/zookeeperreceiver \
	receiver/kafkametricsreceiver \
	receiver/nginxreceiver \
	internal/common

.DEFAULT_GOAL := all

all-modules:
	@echo $(ALL_MODULES) | tr ' ' '\n' | sort

.PHONY: all
all: common gotest otelcontribcol otelcontribcol-unstable

.PHONY: e2e-test
e2e-test: otelcontribcol otelcontribcol-unstable
	$(MAKE) -C testbed run-tests
	$(MAKE) -C testbed run-tests TESTS_DIR=tests_unstable_exe

.PHONY: unit-tests-with-cover
unit-tests-with-cover:
	@echo Verifying that all packages have test files to count in coverage
	@internal/buildscripts/check-test-files.sh $(subst github.com/open-telemetry/opentelemetry-collector-contrib/,./,$(ALL_PKGS))
	@$(MAKE) for-all-target TARGET="do-unit-tests-with-cover"

.PHONY: integration-tests-with-cover
integration-tests-with-cover:
	@echo $(INTEGRATION_TEST_MODULES)
	@$(MAKE) for-all-target TARGET="do-integration-tests-with-cover" ALL_MODULES="$(INTEGRATION_TEST_MODULES)"

# Long-running e2e tests
.PHONY: stability-tests
stability-tests: otelcontribcol
	@echo Stability tests are disabled until we have a stable performance environment.
	@echo To enable the tests replace this echo by $(MAKE) -C testbed run-stability-tests

.PHONY: gotidy
gotidy:
	$(MAKE) for-all CMD="rm -fr go.sum"
	$(MAKE) for-all CMD="go mod tidy"

.PHONY: gomoddownload
gomoddownload:
	@$(MAKE) for-all CMD="go mod download"

.PHONY: gotest
gotest:
	$(MAKE) for-all-target TARGET="test"

.PHONY: gofmt
gofmt:
	$(MAKE) for-all-target TARGET="fmt"

.PHONY: golint
golint:
	$(MAKE) for-all-target TARGET="lint"

.PHONY: for-all
for-all:
	@echo "running $${CMD} in root"
	@$${CMD}
	@set -e; for dir in $(ALL_MODULES); do \
	  (cd "$${dir}" && \
	  	echo "running $${CMD} in $${dir}" && \
	 	$${CMD} ); \
	done

.PHONY: add-tag
add-tag:
	@[ "${TAG}" ] || ( echo ">> env var TAG is not set"; exit 1 )
	@echo "Adding tag ${TAG}"
	@git tag -a ${TAG} -s -m "Version ${TAG}"
	@set -e; for dir in $(ALL_MODULES); do \
	  (echo Adding tag "$${dir:2}/$${TAG}" && \
	 	git tag -a "$${dir:2}/$${TAG}" -s -m "Version ${dir:2}/${TAG}" ); \
	done

.PHONY: push-tag
push-tag:
	@[ "${TAG}" ] || ( echo ">> env var TAG is not set"; exit 1 )
	@echo "Pushing tag ${TAG}"
	@git push upstream ${TAG}
	@set -e; for dir in $(ALL_MODULES); do \
	  (echo Pushing tag "$${dir:2}/$${TAG}" && \
	 	git push upstream "$${dir:2}/$${TAG}"); \
	done

.PHONY: delete-tag
delete-tag:
	@[ "${TAG}" ] || ( echo ">> env var TAG is not set"; exit 1 )
	@echo "Deleting tag ${TAG}"
	@git tag -d ${TAG}
	@set -e; for dir in $(ALL_MODULES); do \
	  (echo Deleting tag "$${dir:2}/$${TAG}" && \
	 	git tag -d "$${dir:2}/$${TAG}" ); \
	done

DEPENDABOT_PATH=".github/dependabot.yml"
.PHONY: gendependabot
gendependabot:
	@echo "Recreate dependabot.yml file"
	@echo "# File generated by \"make gendependabot\"; DO NOT EDIT." > ${DEPENDABOT_PATH}
	@echo "version: 2" >> ${DEPENDABOT_PATH}
	@echo "updates:" >> ${DEPENDABOT_PATH}
	@echo "Add entry for \"/\""
	@echo "  - package-ecosystem: \"github-actions\"" >> ${DEPENDABOT_PATH}
	@echo "    directory: \"/\"" >> ${DEPENDABOT_PATH}
	@echo "    schedule:" >> ${DEPENDABOT_PATH}
	@echo "      interval: \"weekly\"" >> ${DEPENDABOT_PATH}
	@echo "Add entry for \"/\""
	@echo "  - package-ecosystem: \"gomod\"" >> ${DEPENDABOT_PATH}
	@echo "    directory: \"/\"" >> ${DEPENDABOT_PATH}
	@echo "    schedule:" >> ${DEPENDABOT_PATH}
	@echo "      interval: \"weekly\"" >> ${DEPENDABOT_PATH}
	@set -e; for dir in $(ALL_MODULES); do \
		echo "Add entry for \"$${dir:1}\""; \
		echo "  - package-ecosystem: \"gomod\"" >> ${DEPENDABOT_PATH}; \
		echo "    directory: \"$${dir:1}\"" >> ${DEPENDABOT_PATH}; \
		echo "    schedule:" >> ${DEPENDABOT_PATH}; \
		echo "      interval: \"weekly\"" >> ${DEPENDABOT_PATH}; \
	done

GOMODULES = $(ALL_MODULES) $(PWD)
.PHONY: $(GOMODULES)
MODULEDIRS = $(GOMODULES:%=for-all-target-%)
for-all-target: $(MODULEDIRS)
$(MODULEDIRS):
	$(MAKE) -C $(@:for-all-target-%=%) $(TARGET)
.PHONY: for-all-target

TOOLS_MOD_DIR := ./internal/tools
.PHONY: install-tools
install-tools:
	cd $(TOOLS_MOD_DIR) && go install github.com/client9/misspell/cmd/misspell
	cd $(TOOLS_MOD_DIR) && go install github.com/golangci/golangci-lint/cmd/golangci-lint
	cd $(TOOLS_MOD_DIR) && go install github.com/google/addlicense
	cd $(TOOLS_MOD_DIR) && go install github.com/jstemmer/go-junit-report
	cd $(TOOLS_MOD_DIR) && go install github.com/pavius/impi/cmd/impi
	cd $(TOOLS_MOD_DIR) && go install github.com/tcnksm/ghr
	cd $(TOOLS_MOD_DIR) && go install go.opentelemetry.io/collector/cmd/checkdoc
	cd $(TOOLS_MOD_DIR) && go install go.opentelemetry.io/collector/cmd/issuegenerator
	cd $(TOOLS_MOD_DIR) && go install go.opentelemetry.io/collector/cmd/mdatagen

.PHONY: run
run:
	GO111MODULE=on go run --race ./cmd/otelcontribcol/... --config ${RUN_CONFIG} ${RUN_ARGS}

.PHONY: docker-component # Not intended to be used directly
docker-component: check-component
	GOOS=linux GOARCH=amd64 $(MAKE) $(COMPONENT)
	cp ./bin/$(COMPONENT)_linux_amd64 ./cmd/$(COMPONENT)/$(COMPONENT)
	docker build -t $(COMPONENT) ./cmd/$(COMPONENT)/
	rm ./cmd/$(COMPONENT)/$(COMPONENT)

.PHONY: check-component
check-component:
ifndef COMPONENT
	$(error COMPONENT variable was not defined)
endif

.PHONY: docker-otelcontribcol
docker-otelcontribcol:
	COMPONENT=otelcontribcol $(MAKE) docker-component

.PHONY: generate
generate:
	$(MAKE) for-all CMD="go generate ./..."

# Build the Collector executable.
.PHONY: otelcontribcol
otelcontribcol:
	GO111MODULE=on CGO_ENABLED=0 go build -trimpath -o ./bin/otelcontribcol_$(GOOS)_$(GOARCH)$(EXTENSION) \
		$(BUILD_INFO) ./cmd/otelcontribcol

# Build the Collector executable, including unstable functionality.
.PHONY: otelcontribcol-unstable
otelcontribcol-unstable:
	GO111MODULE=on CGO_ENABLED=0 go build -trimpath -o ./bin/otelcontribcol_unstable_$(GOOS)_$(GOARCH)$(EXTENSION) \
		$(BUILD_INFO) -tags enable_unstable ./cmd/otelcontribcol

.PHONY: otelcontribcol-all-sys
otelcontribcol-all-sys: otelcontribcol-darwin_amd64 otelcontribcol-darwin_arm64 otelcontribcol-linux_amd64 otelcontribcol-linux_arm64 otelcontribcol-windows_amd64

.PHONY: otelcontribcol-darwin_amd64
otelcontribcol-darwin_amd64:
	GOOS=darwin  GOARCH=amd64 $(MAKE) otelcontribcol

.PHONY: otelcontribcol-darwin_arm64
otelcontribcol-darwin_arm64:
	GOOS=darwin  GOARCH=arm64 $(MAKE) otelcontribcol

.PHONY: otelcontribcol-linux_amd64
otelcontribcol-linux_amd64:
	GOOS=linux   GOARCH=amd64 $(MAKE) otelcontribcol

.PHONY: otelcontribcol-linux_arm64
otelcontribcol-linux_arm64:
	GOOS=linux   GOARCH=arm64 $(MAKE) otelcontribcol

.PHONY: otelcontribcol-windows_amd64
otelcontribcol-windows_amd64:
	GOOS=windows GOARCH=amd64 EXTENSION=.exe $(MAKE) otelcontribcol

.PHONY: update-dep
update-dep:
	$(MAKE) for-all CMD="$(PWD)/internal/buildscripts/update-dep"
	$(MAKE) gotidy
	$(MAKE) otelcontribcol

.PHONY: update-otel
update-otel:
	cd $(TOOLS_MOD_DIR) && go get go.opentelemetry.io/collector/cmd/checkdoc@$(OTEL_VERSION)
	cd $(TOOLS_MOD_DIR) && go get go.opentelemetry.io/collector/cmd/issuegenerator@$(OTEL_VERSION)
	cd $(TOOLS_MOD_DIR) && go get go.opentelemetry.io/collector/cmd/mdatagen@$(OTEL_VERSION)
	$(MAKE) update-dep MODULE=go.opentelemetry.io/collector VERSION=$(OTEL_VERSION)

.PHONY: otel-from-tree
otel-from-tree:
	# This command allows you to make changes to your local checkout of otel core and build
	# contrib against those changes without having to push to github and update a bunch of
	# references. The workflow is:
	#
	# 1. Hack on changes in core (assumed to be checked out in ../opentelemetry-collector from this directory)
	# 2. Run `make otel-from-tree` (only need to run it once to remap go modules)
	# 3. You can now build contrib and it will use your local otel core changes.
	# 4. Before committing/pushing your contrib changes, undo by running `make otel-from-lib`.
	$(MAKE) for-all CMD="go mod edit -replace go.opentelemetry.io/collector=$(SRC_ROOT)/../opentelemetry-collector"

.PHONY: otel-from-lib
otel-from-lib:
	# Sets opentelemetry core to be not be pulled from local source tree. (Undoes otel-from-tree.)
	$(MAKE) for-all CMD="go mod edit -dropreplace go.opentelemetry.io/collector"

.PHONY: build-examples
build-examples:
	docker-compose -f examples/tracing/docker-compose.yml build
	docker-compose -f exporter/splunkhecexporter/example/docker-compose.yml build

.PHONY: deb-rpm-package
%-package: ARCH ?= amd64
%-package:
	$(MAKE) otelcontribcol-linux_$(ARCH)
	docker build -t otelcontribcol-fpm internal/buildscripts/packaging/fpm
	docker run --rm -v $(CURDIR):/repo -e PACKAGE=$* -e VERSION=$(VERSION) -e ARCH=$(ARCH) otelcontribcol-fpm

# Verify existence of READMEs for components specified as default components in the collector.
.PHONY: checkdoc
checkdoc:
	checkdoc --project-path $(CURDIR) --component-rel-path $(COMP_REL_PATH) --module-name $(MOD_NAME)

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# List of directories where certificates are stored for unit tests.
CERT_DIRS := receiver/sapmreceiver/testdata \
             receiver/signalfxreceiver/testdata \
             receiver/splunkhecreceiver/testdata

# Generate certificates for unit tests relying on certificates.
.PHONY: certs
certs:
	$(foreach dir, $(CERT_DIRS), $(call exec-command, @internal/buildscripts/gen-certs.sh -o $(dir)))
