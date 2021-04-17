package todosvc

import (
	"context"

	todov1 "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct{}

var _ todov1.TodoServiceServer = &Server{}

func (s *Server) CreateTodo(context.Context, *todov1.CreateTodoRequest) (*todov1.Todo, error) {
	// TODO: Implement me.
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}

func (s *Server) GetTodo(context.Context, *todov1.GetTodoRequest) (*todov1.Todo, error) {
	// TODO: Implement me.
	return nil, status.Errorf(codes.Unimplemented, "method GetTodo not implemented")
}

func (s *Server) UpdateTodo(context.Context, *todov1.UpdateTodoRequest) (*todov1.Todo, error) {
	// TODO: Implement me.
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}

func (s *Server) DeleteTodo(context.Context, *todov1.DeleteTodoRequest) (*emptypb.Empty, error) {
	// TODO: Implement me.
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}

func (s *Server) ListTodos(context.Context, *todov1.ListTodosRequest) (*todov1.ListTodosResponse, error) {
	// TODO: Implement me.
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}
