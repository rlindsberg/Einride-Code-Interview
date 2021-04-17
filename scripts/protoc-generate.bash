#!/usr/bin/env bash
#
# Generate proto/gRPC stubs and GAPIC CLI client.

set -euo pipefail

go_module=$(head -1 go.mod | cut -d' ' -f2)
out=proto/gen
go_out="${out}/go"
gapic_out="${out}/gapic"
cli_out="${out}/cli"
cli_root="interviewctl"

rm -rf "$out"
mkdir -p "$go_out" "$gapic_out" "$cli_out"

protoc \
  -I proto/api-common-protos \
  -I proto/src \
  --go_out="$go_out" \
  --go_opt=module="${go_module}/${go_out}" \
  --go-grpc_out="$go_out" \
  --go-grpc_opt=module="${go_module}/${go_out},gen_unstable_server_interfaces=true" \
  $(find proto/src -type f -name '*.proto')

files_with_gapic_clients=$(grep -lr default_host proto/src)
for file in $files_with_gapic_clients; do
  go_package=$(grep 'option go_package =' "$file" | cut -d'"' -f2)
  go_package_path=$(cut -d';' -f1 <<<"$go_package")
  go_package_name=$(cut -d';' -f2 <<<"$go_package")
  go_gapic_package_path=$(sed "s#${go_out}#${gapic_out}#" <<<"$go_package_path")
  go_gapic_package_name=$(basename $(dirname "$go_gapic_package_path"))
  go_gapic_package="${go_gapic_package_path};${go_gapic_package_name}"
  protoc \
    -I proto/api-common-protos \
    -I proto/src \
    --go_gapic_out="$gapic_out" \
    --go_gapic_opt="module=${go_module}/${gapic_out},go-gapic-package=${go_gapic_package}" \
    --go_cli_out "$cli_out" \
    --go_cli_opt "root=${cli_root},gapic=${go_gapic_package}" \
    "$file"
done
