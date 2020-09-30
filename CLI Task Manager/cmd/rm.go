package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a task",
	Long: `Removes a task from the list of all tasks. This does not mark the
task as completed and removes all of it's data`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid use of rm")
			fmt.Println("Usage: todo rm <task no>")
			return
		}
		databse := getBboltDatabase()
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Could not convert task id to int")
			panic(err)
		}

		databse.rmTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
