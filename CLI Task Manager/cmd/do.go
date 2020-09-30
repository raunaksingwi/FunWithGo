package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid use of do")
			fmt.Println("Usage: todo do <task no>")
			return
		}
		databse := getBboltDatabase()
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Could not convert task id to int")
			panic(err)
		}

		databse.doTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
