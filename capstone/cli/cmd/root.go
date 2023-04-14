package cmd

import (
	"fmt"
	"os"

	"github.com/aguazul-marco/hackbright/capstone/cli/client"
	"github.com/spf13/cobra"
)

var DeleteFlag bool
var rootCmd = &cobra.Command{
	Use:     "goals",
	Example: "goals `firstname` `lastname`",
	Short:   "a simple CLI to help keep track a users daily, weekly, and even monthly goals",
	Args:    cobra.RangeArgs(0, 2),
	Long: ` "goals" is a simple CLI to help keep track a users daily, weekly, and even monthly goals. Begin by creating a user account. Then with the user_ID and subcommands the users are able to create, 
	delete and update their goals. Flags will assist with retrieving complete and incomplete goals.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := client.NewClient()
		client.CreateUser(args[0], args[1])

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error occured while executing the CLI: %v", err)
		os.Exit(1)
	}
}

func addSubCommand() {
	rootCmd.AddCommand(CreateCmd)
	rootCmd.AddCommand(RetrieveCmd)
	rootCmd.AddCommand(UpdateCmd)
	rootCmd.AddCommand(DeleteCmd)
}

func init() {
	addSubCommand()
}
