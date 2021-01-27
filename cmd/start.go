package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new working session",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := start(); err != nil {
			logrus.Fatalf("cannot start a new work session: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func start() error {
	// Get user home directory
	homedir, err := homedir.Dir()
	if err != nil {
		return nil
	}

	// Set running path
	runDir := homedir + "/.goggl/running/"
	return nil
}
