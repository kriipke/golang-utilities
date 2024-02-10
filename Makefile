MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

all:
	@make -f $(MAKEFILE_DIR)src/configspy/Makefile
