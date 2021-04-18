package todosvc

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"sync"

	todov1 "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	"github.com/mennanov/fieldmask-utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct{
	sync.Mutex
	todos map[string]*todov1.Todo
}

var _ todov1.TodoServiceServer = &Server{}

func (s *Server) CreateTodo(c context.Context,  r *todov1.CreateTodoRequest) (*todov1.Todo, error) {
	// Done: Implement me.
	if s.todos == nil{
		s.todos = make(map[string]*todov1.Todo)
	}
	s.Lock()
	todo := r.Todo
	key := r.TodoId

	// check if the todoId exists
	if _, ok := s.todos[key]; ok{
		s.Unlock()
		return nil, status.Errorf(codes.AlreadyExists, "Todo cannot be created. Duplicate found.")
	}

	// create todo
	s.todos[key] = todo
	s.Unlock()
	return todo, status.Errorf(codes.OK, "Todo created successfully")
}

func (s *Server) GetTodo(c context.Context, r *todov1.GetTodoRequest) (*todov1.Todo, error) {
	// Done: Implement me.
	if s.todos == nil{
		return nil, status.Errorf(codes.NotFound, "The requested todo is not found")
	}
	s.Lock()
	key := r.Name[6:]

	if res, ok := s.todos[key]; ok{
		s.Unlock()
		return res, status.Errorf(codes.OK, "Todo get successfully")
	} else{
		s.Unlock()
		return nil, status.Errorf(codes.NotFound, "The requested todo is not found")
	}
}

func (s *Server) UpdateTodo(c context.Context, r *todov1.UpdateTodoRequest) (*todov1.Todo, error) {
	// Done: Implement me.
	if s.todos == nil{
		return nil, status.Errorf(codes.NotFound, "The requested todo is not found")
	}
	s.Lock()
	todo := r.Todo

	// The to-do's `name` field is used to identify the to-do to be updated.
	key := r.Todo.Name[6:]

	// created a mask
	todoDst := &todov1.Todo{}
	paths, err := fieldmask_utils.MaskFromPaths(r.UpdateMask.Paths, generator.CamelCase)
	fmt.Println(err)

	// update to-do with mask
	err = fieldmask_utils.StructToStruct(paths, todo, todoDst)
	fmt.Println(err)

	s.todos[key] = todoDst
	todoUpdated := s.todos[key]

	s.Unlock()
	return todoUpdated, status.Errorf(codes.OK, "Todo updated successfully")
}

func (s *Server) DeleteTodo(c context.Context, r *todov1.DeleteTodoRequest) (*emptypb.Empty, error) {
	// TODO: Implement me.
	if s.todos == nil{
		return nil, status.Errorf(codes.NotFound, "The requested todo is not found")
	}
	todoToDelete := r.Name[6:]
	if _, ok := s.todos[todoToDelete]; ok{
		delete(s.todos, todoToDelete)
	}
	return nil, status.Errorf(codes.OK, "Todo deleted successfully")
}

func (s *Server) ListTodos(context.Context, *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	// TODO: Implement me.
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
