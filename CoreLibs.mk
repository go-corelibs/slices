#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

SHELL := /bin/bash

VERSION_TAGS        += CORELIBS
CORELIBS_MK_SUMMARY := Go-CoreLibs.mk
CORELIBS_MK_VERSION := v0.1.13

GOPKG_KEYS          ?=
GOPKG_AUTO_CORELIBS ?= true
LOCAL_CORELIBS_PATH ?= ..

.PHONY: help version
.PHONY: local unlocal be-update tidy
.PHONE: corelibs packages
.PHONY: deps build clean fmt
.PHONY: test coverage goconvey reportcard

#
#: Custom functions
#

define __list_gopkgs
$(if ${GOPKG_KEYS},$(foreach key,${GOPKG_KEYS},$(shell \
		PKG="$($(key)_GO_PACKAGE)"; \
		if [ -n "$${PKG}" -a "$${PKG}" != "nil" ]; then \
			echo "$${PKG}$(1)"; \
		fi; \
	)))
endef

define __list_gopkgs_latest
$(call __list_gopkgs,@latest)
endef

define __list_corelibs
$(shell grep -h -v '^module' go.mod \
		| grep -P '^(require)?\s*github.com/go-corelibs/' \
		| grep -v "github.com/${CORELIB_PKG} v" \
		| grep -v "// indirect" \
		| perl -pe 's!^(require)?\s*!!;s!\s+v\d+(.\d)*.*$$!!;' \
		| sort -u -V \
		| while read MODULE; do \
			NAME=$$(basename "$${MODULE}"); \
			if [ -d "${LOCAL_CORELIBS_PATH}/$${NAME}" ]; then \
				echo "$${MODULE}$(1)"; \
			fi; \
	done)
endef

define __list_corelibs_latest
$(call __list_corelibs,@latest)
endef

#
#: Actual targets
#

help: export FOUND_PKGS=$(call __list_gopkgs)
help: export FOUND_LIBS=$(call __list_corelibs)
help:
	@echo "# usage: make <help|version>"
	@echo "#        make <local|unlocal|be-update|tidy>"
	@echo "#        make <corelibs|packages>"
	@echo "#        make <deps|build|clean|fmt>"
	@echo "#        make <test|coverage|goconvey|reportcard>"
	@echo "#"
	@echo "# targets:"
	@echo "#"
	@echo "#  help           - this help screen"
	@echo "#  version        - build system versions"
	@echo "#"
	@echo "#  local          - go mod edit -replace"
	@echo "#  unlocal        - go mod edit -dropreplace"
	@echo "#  be-update      - go get @latest"
	@echo "#  tidy           - go mod tidy"
	@echo "#"
	@echo "#  corelibs       - list detected go-corelibs"
	@echo "#  packages       - list configured GOPKGS"
	@echo "#"
	@echo "#  deps           - install dependencies"
	@echo "#  build          - go build -v ./..."
	@echo "#  clean          - cleanup artifacts"
	@echo "#  fmt            - gofmt -s, goimports"
	@echo "#"
	@echo "#  test           - go test -race -v ./..."
	@echo "#  coverage       - go test cover -v ./..."
	@echo "#  goconvey       - goconvey -host=0.0.0.0"
	@echo "#  reportcard     - code sanity and style report"
	@if [ -n "$${FOUND_PKGS}" -o -n "$${FOUND_LIBS}" ]; then \
		if [ -n "$${FOUND_PKGS}" ]; then \
			echo "#"; \
			echo "# configured packages:"; \
			echo "#"; \
			for pkg in $${FOUND_PKGS}; do \
				echo "#  - $${pkg}"; \
			done; \
		fi; \
		if [ -n "$${FOUND_LIBS}" ]; then \
			echo "#"; \
			echo "# detected go-corelibs:"; \
			echo "#"; \
			for pkg in $${FOUND_LIBS}; do \
				echo "#  - $${pkg}"; \
			done; \
		fi; \
	fi

corelibs: export FOUND_LIBS=$(call __list_corelibs)
corelibs:
	@if [ -n "$${FOUND_LIBS}" ]; then \
		for FOUND in $${FOUND_LIBS}; do \
			echo "# $${FOUND}"; \
		done; \
	else \
		echo "# no go-corelibs detected"; \
	fi

packages: export FOUND_PKGS=$(call __list_gopkgs)
packages:
	@if [ -n "$${FOUND_PKGS}" ]; then \
		for FOUND in $${FOUND_PKGS}; do \
			echo "# $${FOUND}"; \
		done; \
	else \
		echo "# no GOPKGS configured"; \
	fi

version: LIST=$(foreach key,${VERSION_TAGS},\\n# $($(key)_MK_SUMMARY) $($(key)_MK_VERSION))
version:
	@echo -e -n "${LIST}" | column -t -N '#,SYSTEM,VERSION'

local: export FOUND_PKGS=$(call __list_gopkgs)
local: export FOUND_LIBS=$(call __list_corelibs)
local:
	@if [ -n "$${FOUND_PKGS}" -o -n "$${FOUND_LIBS}" ]; then \
		for found in $${FOUND_PKGS} $${FOUND_LIBS}; do \
			name=`basename $${found}`; \
			echo "# go mod local go-corelibs/$${name}"; \
			go mod edit -replace=$${found}=${LOCAL_CORELIBS_PATH}/$${name}; \
		done; \
	else \
		echo "# nothing to do"; \
	fi

unlocal: export FOUND_PKGS=$(call __list_gopkgs)
unlocal: export FOUND_LIBS=$(call __list_corelibs)
unlocal:
	@if [ -n "$${FOUND_PKGS}" -o -n "$${FOUND_LIBS}" ]; then \
		for found in $${FOUND_PKGS} $${FOUND_LIBS}; do \
			name=`basename $${found}`; \
			echo "# go mod unlocal go-corelibs/$${name}"; \
			go mod edit -dropreplace=$${found}; \
		done; \
	else \
		echo "# nothing to do"; \
	fi

be-update: export GOPROXY=direct
be-update: export FOUND_PKGS=$(call __list_gopkgs_latest)
be-update: export FOUND_LIBS=$(call __list_corelibs_latest)
be-update:
	@if [ "${GOPKG_AUTO_CORELIBS}" == "true" -a -n "$${FOUND_LIBS}" ]; then \
		if [ -n "$${FOUND_PKGS}" ]; then \
			echo "# go get $${FOUND_PKGS} $${FOUND_LIBS}"; \
			go get $${FOUND_PKGS} $${FOUND_LIBS}; \
		else \
			echo "# go get $${FOUND_LIBS}"; \
			go get $${FOUND_LIBS}; \
		fi; \
	elif [ -n "$${FOUND_PKGS}" ]; then \
		echo "# go get $${FOUND_PKGS}"; \
		go get $${FOUND_PKGS}; \
	else \
		echo "# nothing to do"; \
	fi

tidy:
	@go mod tidy

deps:
	@echo "# go install goconvey"
	@go install github.com/smartystreets/goconvey@latest
	@echo "# go install govulncheck"
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "# go install gocyclo"
	@go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@echo "# go install ineffassign"
	@go install github.com/gordonklaus/ineffassign@latest
	@echo "# go install misspell"
	@go install github.com/client9/misspell/cmd/misspell@latest
	@echo "# go get ./..."
	@go get ./...

build:
	@go build -v ./...

clean:
	@rm -fv coverage.{out,html}

fmt:
	@echo "# gofmt -s..."
	@gofmt -w -s `find * -name "*.go"`
	@echo "# goimports..."
	@goimports -w \
		-local "github.com/go-corelibs" \
		`find * -name "*.go"`

test:
	@go test -race -v ./...

coverage:
	@go test -race -coverprofile=coverage.out -covermode=atomic -coverpkg=./... -v ./...
	@go tool cover -html=coverage.out -o=coverage.html
	@go tool cover -func=coverage.out

goconvey:
	@echo "# running goconvey... (press <CTRL+c> to stop)"
	@goconvey -host=0.0.0.0 -launchBrowser=false -depth=-1

reportcard:
	@echo "# code sanity and style report"
	@echo "#: go vet"
	@go vet ./...
	@echo "#: gocyclo"
	@gocyclo -over 15 `find * -name "*.go"` || true
	@echo "#: ineffassign"
	@ineffassign ./...
	@echo "#: misspell"
	@misspell ./...
	@echo "#: gofmt -s"
	@echo -e -n `find * -name "*.go" | while read SRC; do \
		gofmt -s "$${SRC}" > "$${SRC}.fmts"; \
		if ! cmp "$${SRC}" "$${SRC}.fmts" 2> /dev/null; then \
			echo "can simplify: $${SRC}\\n"; \
		fi; \
		rm -f "$${SRC}.fmts"; \
	done`
	@echo "#: govulncheck"
	@echo -e -n `govulncheck ./... \
		| egrep '^Vulnerability #' \
		| sort -u -V \
		| while read LINE; do \
			echo "$${LINE}\n"; \
		done`
