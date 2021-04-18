package todosvc

import (
	"context"
	"testing"

	todov1 "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gotest.tools/v3/assert"
)

func TestServer_CreateTodo(t *testing.T) {
	// Done: Implement me.
	var server Server
	myToDo := todov1.Todo{
		Name:       "Task 1",
		CreateTime: nil,
		UpdateTime: nil,
		Title:      "Write shopping list",
		Completed:  false,
	}
	got, err := server.CreateTodo(context.Background(), &todov1.CreateTodoRequest{
		Todo:   &myToDo,
		TodoId: "1",
	})
	assert.Assert(t, got != nil)
	assert.Assert(t, err == nil)
	assert.Equal(t, codes.OK, status.Code(err))
}

func TestServer_GetTodo(t *testing.T) {
	// TODO: Implement me.
	var server Server
	got, err := server.GetTodo(context.Background(), &todov1.GetTodoRequest{})
	assert.Assert(t, got == nil)
	assert.Assert(t, err != nil)
	assert.Equal(t, codes.OK, status.Code(err))
}

func TestServer_UpdateTodo(t *testing.T) {
	// TODO: Implement me.
	var server Server
	got, err := server.UpdateTodo(context.Background(), &todov1.UpdateTodoRequest{})
	assert.Assert(t, got == nil)
	assert.Assert(t, err != nil)
	assert.Equal(t, codes.OK, status.Code(err))
}

func TestServer_DeleteTodo(t *testing.T) {
	// TODO: Implement me.
	var server Server
	got, err := server.DeleteTodo(context.Background(), &todov1.DeleteTodoRequest{})
	assert.Assert(t, got == nil)
	assert.Assert(t, err != nil)
	assert.Equal(t, codes.OK, status.Code(err))
}

func TestServer_ListTodos(t *testing.T) {
	// TODO: Implement me.
	var server Server
	got, err := server.ListTodos(context.Background(), &todov1.ListTodosRequest{})
	assert.Assert(t, got == nil)
	assert.Assert(t, err != nil)
	assert.Equal(t, codes.OK, status.Code(err))
}
