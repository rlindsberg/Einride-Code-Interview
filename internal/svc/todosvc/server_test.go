package todosvc

import (
	"context"
	"testing"

	todov1 "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gotest.tools/v3/assert"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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
	// Done: Implement me.
	t.Run("Get before create", func(t *testing.T) {
		var server Server
		got, err := server.GetTodo(context.Background(), &todov1.GetTodoRequest{
			Name: "todos/1",
		})
		assert.Assert(t, got == nil)
		assert.Assert(t, err != nil)
		assert.Equal(t, codes.NotFound, status.Code(err))
	})

	t.Run("Create and get", func(t *testing.T) {
		var server Server
		// create todo
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
		got, err = server.GetTodo(context.Background(), &todov1.GetTodoRequest{
			Name: "todos/1",
		})
		assert.Assert(t, got != nil)
		assert.Equal(t, got.Name, myToDo.Name)
		assert.Equal(t, got.Title, myToDo.Title)
		assert.Assert(t, err == nil)
		assert.Equal(t, codes.OK, status.Code(err))
	})
}

func TestServer_UpdateTodo(t *testing.T) {
	// TODO: Implement me.
	var server Server

	// create a to-do
	myToDo := todov1.Todo{
		Name:       "todos/1",
		CreateTime: nil,
		UpdateTime: nil,
		Title:      "Write shopping list",
		Completed:  false,
	}
	got, err := server.CreateTodo(context.Background(), &todov1.CreateTodoRequest{
		Todo:   &myToDo,
		TodoId: "1",
	})

	// update title but not completed
	newToDo := todov1.Todo{
		Name:       "todos/1",
		CreateTime: nil,
		UpdateTime: nil,
		Title:      "Write washing list",
		Completed:  true,
	}
	mask := field_mask.FieldMask{
		Paths: []string{"Title"},
	}
	got, err = server.UpdateTodo(context.Background(), &todov1.UpdateTodoRequest{
		Todo:       &newToDo,
		UpdateMask: &mask,
	})
	assert.Assert(t, got.Title == newToDo.Title)
	assert.Assert(t, got.Completed != newToDo.Completed)
	assert.Assert(t, err == nil)
	assert.Equal(t, codes.OK, status.Code(err))
}

func TestServer_DeleteTodo(t *testing.T) {
	// TODO: Implement me.
	var server Server
	// create a to-do
	myToDo := todov1.Todo{
		Name:       "todos/1",
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

	// delete the to-do
	gotAfterDelete, errAfterDelete := server.DeleteTodo(context.Background(), &todov1.DeleteTodoRequest{
		Name: "todos/1",
	})
	assert.Assert(t, gotAfterDelete == nil)
	assert.Assert(t, errAfterDelete == nil)
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
