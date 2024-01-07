#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += SLICES
SLICES_MK_SUMMARY := go-corelibs/slices
SLICES_MK_VERSION := v1.0.1

include CoreLibs.mk
