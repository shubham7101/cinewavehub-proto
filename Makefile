# Makefile for generating Go code from protobuf files

# Define the output directory for generated files
OUT_DIR = gen

# Define the path for proto files
PROTO_DIR = ./proto

# Find all proto files
PROTO_FILES = $(shell find $(PROTO_DIR) -name "*.proto")

# Rule to generate Go files from proto definitions
.PHONY: gen clean help

gen:
	mkdir -p $(OUT_DIR)
	protoc --go_out=paths=source_relative:$(OUT_DIR) \
	       --go-grpc_out=paths=source_relative:$(OUT_DIR) \
	       --proto_path=$(PROTO_DIR) \
	       $(PROTO_FILES)

# Rule to clean up generated files
clean:
	rm -rf $(OUT_DIR)

# Help target
help:
	@echo "Available targets:"
	@echo "  gen    - Generate Go files from proto definitions"
	@echo "  clean  - Remove generated files"
