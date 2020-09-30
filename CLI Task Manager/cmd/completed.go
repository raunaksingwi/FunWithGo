package cmd

import (
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		databse := getBboltDatabase()
		databse.listCompletedTasks()
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
