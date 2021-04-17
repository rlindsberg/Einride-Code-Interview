#!/usr/bin/env bash
#
# Minimal integration test that verifies that the server works.

set -euo pipefail

shopt -s expand_aliases
export INTERVIEWCTL_TODO_ADDRESS=localhost:8080
export INTERVIEWCTL_TODO_INSECURE=true
alias interviewctl='go run ./proto/gen/cli'

interviewctl todo create-todo --todo_id 1 --todo.title 'Write code'
interviewctl todo get-todo --name 'todos/1'
