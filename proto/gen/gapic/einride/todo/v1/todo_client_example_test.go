// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package todo_test

import (
	"context"

	todo "github.com/einride-interviews/backend-software-engineer/proto/gen/gapic/einride/todo/v1"
	todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	"google.golang.org/api/iterator"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := todo.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateTodo() {
	// import todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"

	ctx := context.Background()
	c, err := todo.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &todov1pb.CreateTodoRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateTodo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetTodo() {
	// import todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"

	ctx := context.Background()
	c, err := todo.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &todov1pb.GetTodoRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetTodo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTodo() {
	// import todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"

	ctx := context.Background()
	c, err := todo.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &todov1pb.UpdateTodoRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateTodo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteTodo() {
	ctx := context.Background()
	c, err := todo.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &todov1pb.DeleteTodoRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTodo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListTodos() {
	// import todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := todo.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &todov1pb.ListTodosRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListTodos(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}