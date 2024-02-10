MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

BIN_TARGET_DIR := $(abspath $(lastword $(MAKEFILE_DIR)/../../bin/))

all:
	@set -e
	@set -o pipefail
	@cd $(MAKEFILE_DIR)
	@go build -o configspy $(MAKEFILE_DIR)main.go
	@mv configspy $(MAKEFILE_DIR)../../bin/
	@printf '[BUILT] %s\n' "`realpath $(MAKEFILE_DIR)../../bin/configspy | xargs file`"
