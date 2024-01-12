#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

CORELIB_PKG := go-corelibs/slices
VERSION_TAGS += MAIN
MAIN_MK_SUMMARY := ${CORELIB_PKG}
MAIN_MK_VERSION := v1.2.0

include CoreLibs.mk
