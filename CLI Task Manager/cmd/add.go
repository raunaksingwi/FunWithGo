package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		databse := getBboltDatabase()
		taskName := ""
		for _, arg := range args {
			taskName = taskName + arg + " "
		}
		databse.addTask(taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
