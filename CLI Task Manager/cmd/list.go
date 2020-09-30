package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks in TODO.",
	Long: `Lists all the tasks that are currently in the TODO list
which are not completed`,
	Run: func(cmd *cobra.Command, args []string) {
		databse := getBboltDatabase()
		databse.listPendingTasks()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
