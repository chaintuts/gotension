# This file contains a make script for the GoTension application
#
# Author: Josh McIntyre
#

# This block defines makefile variables
SRC_FILES=src/gotension/*.go
PACKAGE_NAME=./src/gotension
RES_FILES=res/*.json
BUILD_DIR=bin/gotension

# This rule builds the application
# Here we are simply wrapping Go's build tool
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR) $(SRC_FILES)
	cp $(RES_FILES) $(BUILD_DIR)

# Build unit tests
test:
	mkdir -p $(BUILD_DIR)
	go test $(PACKAGE_NAME) -o $(BUILD_DIR) -v

# This rule cleans the build directory
clean: $(BUILD_DIR)
	rm -r $(BUILD_DIR)/*
	rmdir $(BUILD_DIR)
