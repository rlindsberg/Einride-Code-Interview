package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/einride-interviews/backend-software-engineer/internal/svc/todosvc"
	todov1 "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(0)
	var todoServiceServer todosvc.Server
	grpcServer := grpc.NewServer()
	todov1.RegisterTodoServiceServer(grpcServer, &todoServiceServer)
	port := 8080
	if envPortStr, ok := os.LookupEnv("PORT"); ok {
		envPort, err := strconv.Atoi(envPortStr)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("using port from environment: %d", envPort)
		port = envPort
	}
	address := fmt.Sprintf("localhost:%d", port)
	log.Printf("listening on %s...", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err := lis.Close(); err != nil {
			log.Panic(err)
		}
	}()
	if err := grpcServer.Serve(lis); err != nil {
		log.Panic(err)
	}
}
