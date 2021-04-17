// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	field_maskpb "google.golang.org/genproto/protobuf/field_mask"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
)

var UpdateTodoInput todov1pb.UpdateTodoRequest

var UpdateTodoFromFile string

func init() {
	TodoServiceCmd.AddCommand(UpdateTodoCmd)

	UpdateTodoInput.Todo = new(todov1pb.Todo)

	UpdateTodoInput.UpdateMask = new(field_maskpb.FieldMask)

	UpdateTodoCmd.Flags().StringVar(&UpdateTodoInput.Todo.Name, "todo.name", "", "Resource name of the todo.")

	UpdateTodoCmd.Flags().StringVar(&UpdateTodoInput.Todo.Title, "todo.title", "", "Required. The title of the todo.")

	UpdateTodoCmd.Flags().BoolVar(&UpdateTodoInput.Todo.Completed, "todo.completed", false, "Flag indicating if the todo has been completed.")

	UpdateTodoCmd.Flags().StringSliceVar(&UpdateTodoInput.UpdateMask.Paths, "update_mask.paths", []string{}, "The set of field mask paths.")

	UpdateTodoCmd.Flags().StringVar(&UpdateTodoFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var UpdateTodoCmd = &cobra.Command{
	Use:   "update-todo",
	Short: "Update a todo.",
	Long:  "Update a todo.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if UpdateTodoFromFile == "" {

			cmd.MarkFlagRequired("todo.title")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if UpdateTodoFromFile != "" {
			in, err = os.Open(UpdateTodoFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &UpdateTodoInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Todo", "UpdateTodo", &UpdateTodoInput)
		}
		resp, err := TodoClient.UpdateTodo(ctx, &UpdateTodoInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
