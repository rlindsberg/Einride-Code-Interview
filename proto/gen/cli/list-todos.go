// Code generated. DO NOT EDIT.

package main

import (
	"github.com/spf13/cobra"

	"fmt"

	"github.com/golang/protobuf/jsonpb"

	"google.golang.org/api/iterator"

	"os"

	todov1pb "github.com/einride-interviews/backend-software-engineer/proto/gen/go/einride/todo/v1"
)

var ListTodosInput todov1pb.ListTodosRequest

var ListTodosFromFile string

func init() {
	TodoServiceCmd.AddCommand(ListTodosCmd)

	ListTodosCmd.Flags().Int32Var(&ListTodosInput.PageSize, "page_size", 10, "Default is 10. The maximum number of results to return. The...")

	ListTodosCmd.Flags().StringVar(&ListTodosInput.PageToken, "page_token", "", "A page token, received from a previous call....")

	ListTodosCmd.Flags().StringVar(&ListTodosFromFile, "from_file", "", "Absolute path to JSON file containing request payload")

}

var ListTodosCmd = &cobra.Command{
	Use:   "list-todos",
	Short: "List todos.",
	Long:  "List todos.",
	PreRun: func(cmd *cobra.Command, args []string) {

		if ListTodosFromFile == "" {

		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		in := os.Stdin
		if ListTodosFromFile != "" {
			in, err = os.Open(ListTodosFromFile)
			if err != nil {
				return err
			}
			defer in.Close()

			err = jsonpb.Unmarshal(in, &ListTodosInput)
			if err != nil {
				return err
			}

		}

		if Verbose {
			printVerboseInput("Todo", "ListTodos", &ListTodosInput)
		}
		iter := TodoClient.ListTodos(ctx, &ListTodosInput)

		// populate iterator with a page
		_, err = iter.Next()
		if err != nil && err != iterator.Done {
			return err
		}

		if Verbose {
			fmt.Print("Output: ")
		}
		printMessage(iter.Response)

		return err
	},
}
