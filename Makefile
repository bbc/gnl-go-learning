PROJECT_DIR=$(shell pwd)
BIN_DIR=${PROJECT_DIR}/bin
TMP_DIR=${PROJECT_DIR}/tmp
SHELL=/bin/bash

help:
	@echo 'make init        Install and prepare dependencies and support tools'

init:
	rm -rf $(BIN_DIR) $(TMP_DIR)
	mkdir -p $(BIN_DIR) $(TMP_DIR)
	rm -rf $(TMP_DIR)/gnl-tools-go
	git clone -q git@github.com:bbc/gnl-tools-go.git $(TMP_DIR)/gnl-tools-go
	cp -af $(TMP_DIR)/gnl-tools-go/scripts/build-go-app.sh $(BIN_DIR)/
	cp -af $(TMP_DIR)/gnl-tools-go/scripts/gotree-surgeon.sh $(BIN_DIR)/
	GOBIN=$(BIN_DIR)/ go install github.com/golang/mock/mockgen@v1.6.0
	rm -rf $(TMP_DIR)/gnl-tools-go
