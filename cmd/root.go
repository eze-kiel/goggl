package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// TimeFormat is the format I use all across the program
var TimeFormat = "2006-01-02_15:04:05"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goggl",
	Short: "Manage your work sessions from the terminal",
	Long: `Goggl helps you manage your work sessions from the terminal easily.
When a session is created, a JSON file is created under $HOME/.goggl/running.
When the session is stopped, the file is moved unde $HOME/.goggl/history for later
usage.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
