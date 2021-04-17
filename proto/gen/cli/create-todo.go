// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
)

var CreateTodoInput todov1pb.CreateTodoRequest

var CreateTodoFromFile string

func init() {
	TodoServiceCmd.AddCommand(CreateTodoCmd)

	CreateTodoInput.Todo = new(todov1pb.Todo)

	CreateTodoCmd.Flags().StringVar(&CreateTodoInput.Todo.Name, "todo.name", "", "Resource name of the todo.")

	CreateTodoCmd.Flags().StringVar(&CreateTodoInput.Todo.Title, "todo.title", "", "Required. The title of the todo.")

	CreateTodoCmd.Flags().BoolVar(&CreateTodoInput.Todo.Completed, "todo.completed", false, "Flag indicating if the todo has been completed.")

	CreateTodoCmd.Flags().StringVar(&CreateTodoInput.TodoId, "todo_id", "", "The ID to use for the todo, which will become the...")

	CreateTodoCmd.Flags().StringVar(&CreateTodoFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var CreateTodoCmd = &cobra.Command{
	Use:   "create-todo",
	Short: "Create a todo.",
	Long:  "Create a todo.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if CreateTodoFromFile == "" {

			cmd.MarkFlagRequired("todo.title")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if CreateTodoFromFile != "" {
			in, err = os.Open(CreateTodoFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &CreateTodoInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Todo", "CreateTodo", &CreateTodoInput)
		}
		resp, err := TodoClient.CreateTodo(ctx, &CreateTodoInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
