#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

CORELIBS_MK_VERSION := v0.1.4

SHELL = /bin/bash
LOCAL_CORELIBS_PATH ?= ..

.PHONY: help version
.PHONY: local unlocal be-update tidy
.PHONY: build clean
.PHONY: test coverage goconvey reportcard

#
#: Custom functions
#

define __list_corelibs
$(shell grep -h -v '^module' go.mod \
		| grep 'go-corelibs/' \
		| perl -pe 's!^.*(github\.com/go-corelibs/\S+).*$$!$${1}!' \
		| sort -u -V)
endef

define __gofmt_check
 $(shell find * -name "*.go" | while read SRC; do \
		gofmt -s "$${SRC}" > "$${SRC}.fmts"; \
		if ! cmp "$${SRC}" "$${SRC}.fmts" 2> /dev/null; then \
			echo "can simplify: $${SRC}\\n"; \
		fi; \
		rm -f "$${SRC}.fmts"; \
done)
endef

#
#: Actual targets
#

help:
	@echo "# usage: make <help|version>"
	@echo "#        make <local|unlocal|be-update|tidy>"
	@echo "#        make <build|clean|fmt>"
	@echo "#        make <test|coverage|goconvey|reportcard>"

version:
	@echo "# Go-CoreLibs Makefile: ${CORELIBS_MK_VERSION}"

local: export FOUND_LIBS=$(call __list_corelibs)
local:
	@if [ -n "$${FOUND_LIBS}" ]; then \
		for found_lib in $${FOUND_LIBS}; do \
			name=`basename $${found_lib}`; \
			echo "# go mod local ${LOCAL_CORELIBS_PATH}/$${name}"; \
			go mod edit -replace=$${found_lib}=${LOCAL_CORELIBS_PATH}/$${name}; \
		done; \
	else \
		echo "# nothing to do"; \
	fi

unlocal: export FOUND_LIBS=$(call __list_corelibs)
unlocal:
	@if [ -n "$${FOUND_LIBS}" ]; then \
		for found_lib in $${FOUND_LIBS}; do \
			name=`basename $${found_lib}`; \
			echo "# go mod unlocal go-corelibs/$${name}"; \
			go mod edit -dropreplace=$${found_lib}; \
		done; \
	else \
		echo "# nothing to do"; \
	fi

be-update: export GOPROXY=direct
be-update: export FOUND_LIBS=$(call __list_corelibs)
be-update:
	@if [ -n "$${FOUND_LIBS}" ]; then \
		for found_lib in $${FOUND_LIBS}; do \
			name=`basename $${found_lib}`; \
			echo "# go get go-corelibs/$${name}"; \
			go get $${found_lib}@latest; \
		done; \
	else \
		echo "# nothing to do"; \
	fi

tidy:
	@go mod tidy

deps:
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
	@gofmt -w -s `find * -name "*.go"`

test:
	@go test -v ./...

coverage:
	@go test -race -coverprofile=coverage.out -covermode=atomic -coverpkg=./... -v ./...
	@go tool cover -html=coverage.out -o=coverage.html

goconvey:
	@echo "# running goconvey... (press <CTRL+c> to stop)"
	@goconvey 2>&1 > /dev/null

reportcard:
	@echo "# running reportcard utilities (no output from each is good)"
	@echo "#: go vet"
	@go vet ./...
	@echo "#: gocyclo"
	@gocyclo -over 15 `find * -name "*.go"`
	@echo "#: ineffassign"
	@ineffassign ./...
	@echo "#: misspell"
	@misspell ./...
	@echo "#: gofmt -s"
	@echo -e -n "$(call __gofmt_check)"
