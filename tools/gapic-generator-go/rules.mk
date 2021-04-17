gapic_generator_go_cwd := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
protoc_gen_go_gapic := $(gapic_generator_go_cwd)/bin/protoc-gen-go_gapic
protoc_gen_go_cli := $(gapic_generator_go_cwd)/bin/protoc-gen-go_cli
PATH := $(dir $(protoc_gen_go_gapic)):$(PATH)
PATH := $(dir $(protoc_gen_go_cli)):$(PATH)

$(protoc_gen_go_gapic): $(gapic_generator_go_cwd)/go.mod
	$(info [gapic-generator-go] building $@...)
	@cd $(gapic_generator_go_cwd) && go build -o $@ github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_gapic
	@cd $(gapic_generator_go_cwd) && go mod tidy -v

$(protoc_gen_go_cli): $(gapic_generator_go_cwd)/go.mod
	$(info [gapic-generator-go] building $@...)
	@cd $(gapic_generator_go_cwd) && go build -o $@ github.com/googleapis/gapic-generator-go/cmd/protoc-gen-go_cli
	@cd $(gapic_generator_go_cwd) && go mod tidy -v
