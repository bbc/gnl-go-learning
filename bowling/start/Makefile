APP_DIR=$(shell pwd)
PROJECT_DIR=${APP_DIR}/../..
BIN_DIR=${PROJECT_DIR}/bin
TMP_DIR=${PROJECT_DIR}/tmp
SHELL=/bin/bash

help:
	@echo 'make fmt     Reformat code and rebuild mocks'
	@echo 'make mocks   Build mocks for all interfaces, if present'
	@echo 'make clean   Reformat code and build mocks'
	@echo 'make test    Run tests'

fmt:
	cd $(APP_DIR) && $(BIN_DIR)/gotree-surgeon.sh -f .

mocks:
	cd $(APP_DIR) && $(BIN_DIR)/gotree-surgeon.sh -mv . $(BIN_DIR)/mockgen

clean: mocks fmt
	rm -rf $(TMP_DIR)

test:
	cd $(APP_DIR) && go test ./...

