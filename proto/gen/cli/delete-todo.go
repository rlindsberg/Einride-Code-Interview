// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"github.com/golang/protobuf/jsonpb"

	"os"

	todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
)

var DeleteTodoInput todov1pb.DeleteTodoRequest

var DeleteTodoFromFile string

func init() {
	TodoServiceCmd.AddCommand(DeleteTodoCmd)

	DeleteTodoCmd.Flags().StringVar(&DeleteTodoInput.Name, "name", "", "Required. Resource name of the todo to delete.")

	DeleteTodoCmd.Flags().StringVar(&DeleteTodoFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var DeleteTodoCmd = &cobra.Command{
	Use:   "delete-todo",
	Short: "Delete a todo.",
	Long:  "Delete a todo.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if DeleteTodoFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if DeleteTodoFromFile != "" {
			in, err = os.Open(DeleteTodoFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &DeleteTodoInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Todo", "DeleteTodo", &DeleteTodoInput)
		}
		err = TodoClient.DeleteTodo(ctx, &DeleteTodoInput)

		return err
	},
}
