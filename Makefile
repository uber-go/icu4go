export ICU_LIB = /usr/local/opt/icu4c
export CGO_CFLAGS += -I${ICU_LIB}/include
export CGO_LDFLAGS += -L${ICU_LIB}/lib -licui18n -licuuc -licudata

ALL_SRC := $(shell find . -name "*.go" | grep -v -e vendor \
           	-e ".*/\..*" \
           	-e ".*/_.*")

include .build/lint.mk

.DEFAULT_GOAL := all

.PHONY: all
all: lint test

.PHONY: test
test:
	go test -cover `glide novendor`

