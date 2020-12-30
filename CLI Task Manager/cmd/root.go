package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// TODO: Write better help message for subcommands
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A TODO list",
	Long:  `A TODO list which is helpful to track action items`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
