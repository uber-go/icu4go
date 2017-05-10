export ICU_LIB = /usr/local/opt/icu4c
export CGO_CFLAGS += -I${ICU_LIB}/include
export CGO_LDFLAGS += -L${ICU_LIB}/lib -licui18n -licuuc -licudata

ALL_SRC := $(shell find . -name "*.go" | grep -v -e vendor \
           	-e ".*/\..*" \
           	-e ".*/_.*")

PACKAGES := $(shell glide nv)

include .build/lint.mk

.DEFAULT_GOAL := all

.PHONY: build
build:
	go build -i $(PACKAGES)

.PHONY: install
install:
	glide --version || go get github.com/Masterminds/glide
	glide install


.PHONY: all
all: lint test


.PHONY: test
test:
	go test -cover $(PACKAGES)


.PHONY: install_ci
install_ci: install
	go get github.com/wadey/gocovmerge
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cmd/cover
	go get github.com/golang/lint/golint


.PHONY: test_ci
test_ci: install_ci build test
	./.build/cover.sh $(shell go list $(PACKAGES))
