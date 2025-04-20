PROTO_DIR = proto
OUT_DIR = proto/gen
PROTO_FILES = $(wildcard $(PROTO_DIR)/*.proto)

.PHONY: proto
proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
