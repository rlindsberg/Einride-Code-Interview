// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"os"

	todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
)

var GetTodoInput todov1pb.GetTodoRequest

var GetTodoFromFile string

func init() {
	TodoServiceCmd.AddCommand(GetTodoCmd)

	GetTodoCmd.Flags().StringVar(&GetTodoInput.Name, "name", "", "Required. Resource name of the todo to retrieve.")

	GetTodoCmd.Flags().StringVar(&GetTodoFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var GetTodoCmd = &cobra.Command{
	Use:   "get-todo",
	Short: "Get a todo.",
	Long:  "Get a todo.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if GetTodoFromFile == "" {

			cmd.MarkFlagRequired("name")

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if GetTodoFromFile != "" {
			in, err = os.Open(GetTodoFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &GetTodoInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Todo", "GetTodo", &GetTodoInput)
		}
		resp, err := TodoClient.GetTodo(ctx, &GetTodoInput)

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(resp)

		return err
	},
}
