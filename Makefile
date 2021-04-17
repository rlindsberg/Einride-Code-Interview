.PHONY: all
all: \
	protoc-generate \
	go-lint \
	go-test \
	go-mod-tidy

include tools/gapic-generator-go/rules.mk
include tools/golangci-lint/rules.mk
include tools/grpc-go/rules.mk
include tools/protobuf-go/rules.mk
include tools/protoc/rules.mk

.PHONY: clean
clean:
	rm -rf tools/*/*/

.PHONY: proto/api-common-protos
proto/api-common-protos:
	@git submodule update --init $@

protoc_plugins := \
	$(protoc_gen_go) \
	$(protoc_gen_go_grpc) \
	$(protoc_gen_go_gapic) \
	$(protoc_gen_go_cli)

.PHONY: protoc-generate
protoc-generate: ./scripts/protoc-generate.bash $(protoc) $(protoc_plugins) proto/api-common-protos
	$(info [$@] generating stubs...)
	@$<

.PHONY: go-test
go-test:
	$(info [$@] running Go test suite...)
	@go test -race -covermode=atomic ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v

.PHONY: run-server
run-server:
	$(info [$@] starting...)
	@go run ./cmd/server

.PHONY: integration-test
integration-test: ./scripts/integration-test.bash
	$(info [$@] running integration tests...)
	@$<
