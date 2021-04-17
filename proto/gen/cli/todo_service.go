// Code generated. DO NOT EDIT.

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	gapic "github.com/einride-interviews/backend-software-engineer/proto/gen/gapic/einride/todo/v1"
)

var TodoConfig *viper.Viper
var TodoClient *gapic.Client
var TodoSubCommands []string = []string{
	"create-todo",
	"get-todo",
	"update-todo",
	"delete-todo",
	"list-todos",
}

func init() {
	rootCmd.AddCommand(TodoServiceCmd)

	TodoConfig = viper.New()
	TodoConfig.SetEnvPrefix("INTERVIEWCTL_TODO")
	TodoConfig.AutomaticEnv()

	TodoServiceCmd.PersistentFlags().Bool("insecure", false, "Make insecure client connection. Or use INTERVIEWCTL_TODO_INSECURE. Must be used with \"address\" option")
	TodoConfig.BindPFlag("insecure", TodoServiceCmd.PersistentFlags().Lookup("insecure"))
	TodoConfig.BindEnv("insecure")

	TodoServiceCmd.PersistentFlags().String("address", "", "Set API address used by client. Or use INTERVIEWCTL_TODO_ADDRESS.")
	TodoConfig.BindPFlag("address", TodoServiceCmd.PersistentFlags().Lookup("address"))
	TodoConfig.BindEnv("address")

	TodoServiceCmd.PersistentFlags().String("token", "", "Set Bearer token used by the client. Or use INTERVIEWCTL_TODO_TOKEN.")
	TodoConfig.BindPFlag("token", TodoServiceCmd.PersistentFlags().Lookup("token"))
	TodoConfig.BindEnv("token")

	TodoServiceCmd.PersistentFlags().String("api_key", "", "Set API Key used by the client. Or use INTERVIEWCTL_TODO_API_KEY.")
	TodoConfig.BindPFlag("api_key", TodoServiceCmd.PersistentFlags().Lookup("api_key"))
	TodoConfig.BindEnv("api_key")
}

var TodoServiceCmd = &cobra.Command{
	Use:       "todo",
	Short:     "Todo service.",
	Long:      "Todo service.",
	ValidArgs: TodoSubCommands,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		var opts []option.ClientOption

		address := TodoConfig.GetString("address")
		if address != "" {
			opts = append(opts, option.WithEndpoint(address))
		}

		if TodoConfig.GetBool("insecure") {
			if address == "" {
				return fmt.Errorf("Missing address to use with insecure connection")
			}

			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				return err
			}
			opts = append(opts, option.WithGRPCConn(conn))
		}

		if token := TodoConfig.GetString("token"); token != "" {
			opts = append(opts, option.WithTokenSource(oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
					TokenType:   "Bearer",
				})))
		}

		if key := TodoConfig.GetString("api_key"); key != "" {
			opts = append(opts, option.WithAPIKey(key))
		}

		TodoClient, err = gapic.NewClient(ctx, opts...)
		return
	},
}
