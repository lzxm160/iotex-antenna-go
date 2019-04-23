########################################################################################################################
# Copyright (c) 2019 IoTeX
# This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
# warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
# permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
# License 2.0 that can be found in the LICENSE file.
########################################################################################################################

# Go parameters
GOCMD=go
GOLINT=golint
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_TARGET_SERVER=Antenna

# Pkgs
ALL_PKGS := $(shell go list ./... )
PKGS := $(shell go list ./... | grep -v /test/ )
ROOT_PKG := "github.com/iotexproject/iotex-antenna-go"

# Docker parameters
DOCKERCMD=docker

# Package Info
PACKAGE_VERSION := $(shell git describe --tags --always)
PACKAGE_COMMIT_ID := $(shell git rev-parse HEAD)
GIT_STATUS := $(shell git status --porcelain)
ifdef GIT_STATUS
	GIT_STATUS := "dirty"
else
	GIT_STATUS := "clean"
endif
GO_VERSION := $(shell go version)
BUILD_TIME=$(shell date +%F-%Z/%T)
VersionImportPath := github.com/iotexproject/iotex-antenna/version
PackageFlags += -X '$(VersionImportPath).PackageVersion=$(PACKAGE_VERSION)'
PackageFlags += -X '$(VersionImportPath).PackageCommitID=$(PACKAGE_COMMIT_ID)'
PackageFlags += -X '$(VersionImportPath).GitStatus=$(GIT_STATUS)'
PackageFlags += -X '$(VersionImportPath).GoVersion=$(GO_VERSION)'
PackageFlags += -X '$(VersionImportPath).BuildTime=$(BUILD_TIME)'
PackageFlags += -s -w

TEST_IGNORE= ".git,vendor"
COV_OUT := profile.coverprofile
COV_REPORT := overalls.coverprofile
COV_HTML := coverage.html

LINT_LOG := lint.log

V ?= 0
ifeq ($(V),0)
	ECHO_V = @
else
	VERBOSITY_FLAG = -v
	DEBUG_FLAG = -debug
endif

all: build test clean

.PHONY: build
build:
	$(GOBUILD) -ldflags "$(PackageFlags)" -o ./$(BUILD_TARGET_SERVER) -v ./$(BUILD_TARGET_SERVER)

.PHONY: test
test: lint
	go list ./... | grep -v /vendor/ | xargs $(GOTEST)

.PHONY: lint
lint:
	@echo "Installing golint..."
	go get golang.org/x/lint/golint
	go list ./... | grep -v /vendor/ | xargs $(GOLINT)

.PHONY: clean
clean:
	@echo "Cleaning..."
	$(ECHO_V)rm -rf ./bin/$(BUILD_TARGET_SERVER)
	$(ECHO_V)rm -rf $(COV_REPORT) $(COV_HTML) $(LINT_LOG)
	$(ECHO_V)find . -name $(COV_OUT) -delete
	$(ECHO_V)find . -name $(TESTBED_COV_OUT) -delete
	$(ECHO_V)$(GOCLEAN) -i $(PKGS)